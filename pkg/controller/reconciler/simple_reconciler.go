/*
 * Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.
 * Licensed under the Universal Permissive License v 1.0 as shown at
 * http://oss.oracle.com/licenses/upl.
 */

package reconciler

import (
	coh "github.com/oracle/coherence-operator/pkg/apis/coherence/v1"
	"github.com/oracle/coherence-operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// blank assignment to verify that SimpleReconciler implements reconcile.Reconciler.
// If the reconcile.Reconciler API was to change then we'd get a compile error here.
var _ reconcile.Reconciler = &SimpleReconciler{}

func NewConfigMapReconciler(mgr manager.Manager) SecondaryResourceReconciler {
	return NewSimpleReconciler(mgr, "configmap.controller", coh.ResourceTypeConfigMap, &corev1.ConfigMap{})
}

func NewSecretReconciler(mgr manager.Manager) SecondaryResourceReconciler {
	return NewSimpleReconciler(mgr, "secret.controller", coh.ResourceTypeSecret, &corev1.Secret{})
}

func NewServiceReconciler(mgr manager.Manager) SecondaryResourceReconciler {
	return NewSimpleReconciler(mgr, "service.controller", coh.ResourceTypeService, &corev1.Service{})
}

// NewSimpleReconciler returns a new SimpleReconciler.
func NewSimpleReconciler(mgr manager.Manager, name string, kind coh.ResourceType, template runtime.Object) SecondaryResourceReconciler {
	r := &SimpleReconciler{
		ReconcileSecondaryResource: ReconcileSecondaryResource{
			Kind:     kind,
			Template: template,
		},
	}

	r.SetCommonReconciler(name, mgr)
	return r
}

type SimpleReconciler struct {
	ReconcileSecondaryResource
}

func (in *SimpleReconciler) GetReconciler() reconcile.Reconciler { return in }

// Reconcile reads that state of the secondary resource for a deployment and makes changes based on the
// state read and the desired state based on the parent Coherence resource.
func (in *SimpleReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Attempt to lock the requested resource. If the resource is locked then another
	// request for the same resource is already in progress so requeue this one.
	if ok := in.Lock(request); !ok {
		return reconcile.Result{Requeue: true, RequeueAfter: 0}, nil
	}
	// Make sure that the request is unlocked when this method exits
	defer in.Unlock(request)

	// Obtain the parent Coherence resource
	deployment, err := in.FindDeployment(request)
	if err != nil {
		return reconcile.Result{}, err
	}

	storage, err := utils.NewStorage(request.NamespacedName, in.GetManager())
	if err != nil {
		return reconcile.Result{}, err
	}

	return in.ReconcileResources(request, deployment, storage)
}
