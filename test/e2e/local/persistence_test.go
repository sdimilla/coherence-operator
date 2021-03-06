/*
 * Copyright (c) 2019, 2020 Oracle and/or its affiliates. All rights reserved.
 * Licensed under the Universal Permissive License v 1.0 as shown at
 * http://oss.oracle.com/licenses/upl.
 */

package local

import (
	goctx "context"
	"encoding/json"
	"flag"
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	v1 "github.com/oracle/coherence-operator/pkg/apis/coherence/v1"
	"github.com/oracle/coherence-operator/test/e2e/helper"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"net/http"
	"os"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"strconv"
	"testing"
	"time"
)

type snapshotActionType int

const (
	canaryServiceName = "CanaryService"

	create snapshotActionType = iota
	recover
	delete
)

// Deploy a Coherence resource with persistence enabled (this should enable active persistence).
// A PVC should be created for the StatefulSet. Create data in some caches, delete the deployment,
// re-deploy the deployment and assert that the data is recovered.
func TestActivePersistence(t *testing.T) {
	assertPersistence("persistence-active.yaml", "persistence-volume", false, false, true, t)
}

// Deploy a Coherence resource with the minimal default configuration. Persistence will be on-demand.
// Put data in a cache, take a snapshot, delete the data, recover the snapshot,
// assert that the data is recovered.
func TestOnDemandPersistence(t *testing.T) {
	assertPersistence("persistence-on-demand.yaml", "", true, true, false, t)
}

// Deploy a Coherence resource with snapshot enabled. Persistence will be on-demand,
// a PVC will be created for the StatefulSet to use for snapshots. Put data in a cache, take a snapshot,
// delete the deployment, re-deploy the deployment, recover the snapshot, assert that the data is recovered.
func TestSnapshotPersistence(t *testing.T) {
	assertPersistence("persistence-snapshot.yaml", "snapshot-volume", true, false, true, t)
}

func assertPersistence(yamlFile, pVolName string, isSnapshot, isClearCanary, isRestart bool, t *testing.T) {
	g := NewGomegaWithT(t)

	logf.SetLogger(zap.Logger())

	flags := &flag.FlagSet{}
	klog.InitFlags(flags)
	_ = flags.Set("v", "4")

	f := framework.Global
	ctx := helper.CreateTestContext(t)
	defer helper.DumpOperatorLogsAndCleanup(t, ctx)

	ns, err := ctx.GetWatchNamespace()
	g.Expect(err).NotTo(HaveOccurred())

	deployment, pods := ensurePods(g, ctx, yamlFile, ns, t)

	// check the pvc is created for the given volume
	if pVolName != "" {
		pvcName := ""
		for _, vol := range pods[0].Spec.Volumes {
			if vol.Name == pVolName {
				if vol.PersistentVolumeClaim != nil {
					pvcName = vol.PersistentVolumeClaim.ClaimName
				}
				break
			}
		}

		// check the pvc is created
		g.Expect(pvcName).NotTo(Equal(""))
		pvc := f.KubeClient.CoreV1().PersistentVolumeClaims(pvcName)
		g.Expect(pvc).ShouldNot(BeNil())
	}

	// create data in some caches
	err = helper.StartCanary(ns, deployment.GetName())
	g.Expect(err).NotTo(HaveOccurred())

	if isSnapshot {
		// take a snapshot
		err = processSnapshotRequest(pods[0], create)
		g.Expect(err).NotTo(HaveOccurred())

		defer processSnapshotRequest(pods[0], delete)
	}

	if isClearCanary {
		// delete the data
		err = helper.ClearCanary(ns, deployment.GetName())
		g.Expect(err).NotTo(HaveOccurred())
	}

	localStorageRestartEnv := os.Getenv("LOCAL_STORAGE_RESTART")
	localStorageRestart, err := strconv.ParseBool(localStorageRestartEnv)
	if err != nil {
		localStorageRestart = false
	}
	// restart Coherence may be on a different instance, local storage will not work
	if isRestart && !localStorageRestart {
		// dump the pod logs before deleting
		helper.DumpPodsForTest(t, ctx)
		// delete the deployment
		err = helper.WaitForCoherenceCleanup(f, ns)
		g.Expect(err).NotTo(HaveOccurred())

		// re-deploy the deployment
		deployment, pods = ensurePods(g, ctx, yamlFile, ns, t)
	}

	if isSnapshot {
		// recover the snapshot
		err = processSnapshotRequest(pods[0], recover)
		g.Expect(err).NotTo(HaveOccurred())
	}

	// assert that the data is recovered
	err = helper.CheckCanary(ns, deployment.GetName())
	g.Expect(err).NotTo(HaveOccurred())

	// cleanup the data
	_ = helper.ClearCanary(ns, deployment.GetName())
}

func ensurePods(g *GomegaWithT, ctx *framework.Context, yamlFile, ns string, t *testing.T) (v1.Coherence, []corev1.Pod) {
	f := framework.Global

	deployment, err := helper.NewSingleCoherenceFromYaml(ns, yamlFile)
	g.Expect(err).NotTo(HaveOccurred())

	d, _ := json.Marshal(deployment)
	fmt.Printf("Persistence Test installing deployment:\n%s\n", string(d))

	err = f.Client.Create(goctx.TODO(), &deployment, helper.DefaultCleanup(ctx))
	g.Expect(err).NotTo(HaveOccurred())

	_, err = helper.WaitForStatefulSetForDeployment(f.KubeClient, ns, &deployment, helper.RetryInterval, helper.Timeout, t)
	g.Expect(err).NotTo(HaveOccurred())

	// Get the list of Pods
	pods, err := helper.ListCoherencePodsForDeployment(f.KubeClient, ns, deployment.GetName())
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(len(pods)).Should(BeNumerically(">", 0))

	return deployment, pods
}

func processSnapshotRequest(pod corev1.Pod, actionType snapshotActionType) error {
	pf, ports, err := helper.StartPortForwarderForPod(&pod)
	if err != nil {
		return err
	}

	defer pf.Close()

	url := fmt.Sprintf("http://127.0.0.1:%d/management/coherence/cluster/services/%s/persistence/snapshots/snapshotOne",
		ports[v1.PortNameManagement], canaryServiceName)
	httpMethod := "POST"
	if actionType == delete {
		httpMethod = "DELETE"
	}
	if actionType == recover {
		url = url + "/recover"
	}

	client := &http.Client{}
	var resp *http.Response
	var req *http.Request
	// try a max of 5 times
	for i := 0; i < 5; i++ {
		req, err = http.NewRequest(httpMethod, url, nil)
		if err == nil {
			resp, err = client.Do(req)
			if err == nil {
				break
			}
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("snapshot request returned non-200 status %d", resp.StatusCode)
	}

	// wait for idle
	err = wait.Poll(helper.RetryInterval, helper.Timeout, func() (done bool, err error) {
		url = fmt.Sprintf("http://127.0.0.1:%d/management/coherence/cluster/services/%s/persistence?fields=operationStatus",
			ports[v1.PortNameManagement], canaryServiceName)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Cannot create idle check request: %v\n", url)
			return false, err
		}
		resp, err = client.Do(req)
		if err != nil {
			fmt.Printf("Error in send idle check request: %v\n", url)
			return false, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Idle check request with incorrect status code: %v\n", resp.StatusCode)
			return false, err
		}

		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print("Error in reading idle check response")
			return false, err
		}

		var data map[string]interface{}
		if err = json.Unmarshal(bs, &data); err != nil {
			fmt.Print("Error in unmarshalling idle check response")
			return false, nil
		}
		opStatus := data["operationStatus"]
		fmt.Printf("Persistence opStatus: %v\n", opStatus)
		if opStatus == "Idle" {
			return true, nil
		}
		return false, nil
	})

	return err
}
