// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceCluster) DeepCopyInto(out *CoherenceCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceCluster.
func (in *CoherenceCluster) DeepCopy() *CoherenceCluster {
	if in == nil {
		return nil
	}
	out := new(CoherenceCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceClusterList) DeepCopyInto(out *CoherenceClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoherenceCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceClusterList.
func (in *CoherenceClusterList) DeepCopy() *CoherenceClusterList {
	if in == nil {
		return nil
	}
	out := new(CoherenceClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceClusterSpec) DeepCopyInto(out *CoherenceClusterSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.CoherenceRoleSpec.DeepCopyInto(&out.CoherenceRoleSpec)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]CoherenceRoleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceClusterSpec.
func (in *CoherenceClusterSpec) DeepCopy() *CoherenceClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceClusterStatus) DeepCopyInto(out *CoherenceClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceClusterStatus.
func (in *CoherenceClusterStatus) DeepCopy() *CoherenceClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CoherenceClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternal) DeepCopyInto(out *CoherenceInternal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternal.
func (in *CoherenceInternal) DeepCopy() *CoherenceInternal {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceInternal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalList) DeepCopyInto(out *CoherenceInternalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoherenceInternal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalList.
func (in *CoherenceInternalList) DeepCopy() *CoherenceInternalList {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceInternalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalSpec) DeepCopyInto(out *CoherenceInternalSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Coherence != nil {
		in, out := &in.Coherence, &out.Coherence
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CoherenceUtils != nil {
		in, out := &in.CoherenceUtils, &out.CoherenceUtils
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Store != nil {
		in, out := &in.Store, &out.Store
		*out = new(CoherenceInternalStoreSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Fluentd != nil {
		in, out := &in.Fluentd, &out.Fluentd
		*out = new(FluentdImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UserArtifacts != nil {
		in, out := &in.UserArtifacts, &out.UserArtifacts
		*out = new(UserArtifactsImageSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalSpec.
func (in *CoherenceInternalSpec) DeepCopy() *CoherenceInternalSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalStatus) DeepCopyInto(out *CoherenceInternalStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalStatus.
func (in *CoherenceInternalStatus) DeepCopy() *CoherenceInternalStatus {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceInternalStoreSpec) DeepCopyInto(out *CoherenceInternalStoreSpec) {
	*out = *in
	if in.StorageEnabled != nil {
		in, out := &in.StorageEnabled, &out.StorageEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ReadinessProbeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CacheConfig != nil {
		in, out := &in.CacheConfig, &out.CacheConfig
		*out = new(string)
		**out = **in
	}
	if in.PofConfig != nil {
		in, out := &in.PofConfig, &out.PofConfig
		*out = new(string)
		**out = **in
	}
	if in.OverrideConfig != nil {
		in, out := &in.OverrideConfig, &out.OverrideConfig
		*out = new(string)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LoggingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Main != nil {
		in, out := &in.Main, &out.Main
		*out = new(MainSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxHeap != nil {
		in, out := &in.MaxHeap, &out.MaxHeap
		*out = new(string)
		**out = **in
	}
	if in.JvmArgs != nil {
		in, out := &in.JvmArgs, &out.JvmArgs
		*out = new(string)
		**out = **in
	}
	if in.JavaOpts != nil {
		in, out := &in.JavaOpts, &out.JavaOpts
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodManagementPolicy != nil {
		in, out := &in.PodManagementPolicy, &out.PodManagementPolicy
		*out = new(appsv1.PodManagementPolicyType)
		**out = **in
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(PersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(PortSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(PortSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JMX != nil {
		in, out := &in.JMX, &out.JMX
		*out = new(JMXSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(CoherenceServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceInternalStoreSpec.
func (in *CoherenceInternalStoreSpec) DeepCopy() *CoherenceInternalStoreSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceInternalStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRole) DeepCopyInto(out *CoherenceRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRole.
func (in *CoherenceRole) DeepCopy() *CoherenceRole {
	if in == nil {
		return nil
	}
	out := new(CoherenceRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRoleList) DeepCopyInto(out *CoherenceRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoherenceRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRoleList.
func (in *CoherenceRoleList) DeepCopy() *CoherenceRoleList {
	if in == nil {
		return nil
	}
	out := new(CoherenceRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRoleSpec) DeepCopyInto(out *CoherenceRoleSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = new(Images)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageEnabled != nil {
		in, out := &in.StorageEnabled, &out.StorageEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ScalingPolicy != nil {
		in, out := &in.ScalingPolicy, &out.ScalingPolicy
		*out = new(ScalingPolicy)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ReadinessProbeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CacheConfig != nil {
		in, out := &in.CacheConfig, &out.CacheConfig
		*out = new(string)
		**out = **in
	}
	if in.PofConfig != nil {
		in, out := &in.PofConfig, &out.PofConfig
		*out = new(string)
		**out = **in
	}
	if in.OverrideConfig != nil {
		in, out := &in.OverrideConfig, &out.OverrideConfig
		*out = new(string)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LoggingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Main != nil {
		in, out := &in.Main, &out.Main
		*out = new(MainSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxHeap != nil {
		in, out := &in.MaxHeap, &out.MaxHeap
		*out = new(string)
		**out = **in
	}
	if in.JvmArgs != nil {
		in, out := &in.JvmArgs, &out.JvmArgs
		*out = new(string)
		**out = **in
	}
	if in.JavaOpts != nil {
		in, out := &in.JavaOpts, &out.JavaOpts
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodManagementPolicy != nil {
		in, out := &in.PodManagementPolicy, &out.PodManagementPolicy
		*out = new(appsv1.PodManagementPolicyType)
		**out = **in
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(PersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(PortSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(PortSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JMX != nil {
		in, out := &in.JMX, &out.JMX
		*out = new(JMXSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(CoherenceServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRoleSpec.
func (in *CoherenceRoleSpec) DeepCopy() *CoherenceRoleSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceRoleStatus) DeepCopyInto(out *CoherenceRoleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceRoleStatus.
func (in *CoherenceRoleStatus) DeepCopy() *CoherenceRoleStatus {
	if in == nil {
		return nil
	}
	out := new(CoherenceRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceServiceSpec) DeepCopyInto(out *CoherenceServiceSpec) {
	*out = *in
	in.ServiceSpec.DeepCopyInto(&out.ServiceSpec)
	if in.ManagementHttpPort != nil {
		in, out := &in.ManagementHttpPort, &out.ManagementHttpPort
		*out = new(int32)
		**out = **in
	}
	if in.MetricsHttpPort != nil {
		in, out := &in.MetricsHttpPort, &out.MetricsHttpPort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceServiceSpec.
func (in *CoherenceServiceSpec) DeepCopy() *CoherenceServiceSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdApplicationSpec) DeepCopyInto(out *FluentdApplicationSpec) {
	*out = *in
	if in.ConfigFile != nil {
		in, out := &in.ConfigFile, &out.ConfigFile
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdApplicationSpec.
func (in *FluentdApplicationSpec) DeepCopy() *FluentdApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(FluentdApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdImageSpec) DeepCopyInto(out *FluentdImageSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(FluentdApplicationSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdImageSpec.
func (in *FluentdImageSpec) DeepCopy() *FluentdImageSpec {
	if in == nil {
		return nil
	}
	out := new(FluentdImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images) DeepCopyInto(out *Images) {
	*out = *in
	if in.Coherence != nil {
		in, out := &in.Coherence, &out.Coherence
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CoherenceUtils != nil {
		in, out := &in.CoherenceUtils, &out.CoherenceUtils
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UserArtifacts != nil {
		in, out := &in.UserArtifacts, &out.UserArtifacts
		*out = new(UserArtifactsImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Fluentd != nil {
		in, out := &in.Fluentd, &out.Fluentd
		*out = new(FluentdImageSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images.
func (in *Images) DeepCopy() *Images {
	if in == nil {
		return nil
	}
	out := new(Images)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JMXSpec) DeepCopyInto(out *JMXSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxHeap != nil {
		in, out := &in.MaxHeap, &out.MaxHeap
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JMXSpec.
func (in *JMXSpec) DeepCopy() *JMXSpec {
	if in == nil {
		return nil
	}
	out := new(JMXSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingSpec) DeepCopyInto(out *LoggingSpec) {
	*out = *in
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(int32)
		**out = **in
	}
	if in.ConfigFile != nil {
		in, out := &in.ConfigFile, &out.ConfigFile
		*out = new(string)
		**out = **in
	}
	if in.ConfigMapName != nil {
		in, out := &in.ConfigMapName, &out.ConfigMapName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingSpec.
func (in *LoggingSpec) DeepCopy() *LoggingSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MainSpec) DeepCopyInto(out *MainSpec) {
	*out = *in
	if in.Class != nil {
		in, out := &in.Class, &out.Class
		*out = new(string)
		**out = **in
	}
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MainSpec.
func (in *MainSpec) DeepCopy() *MainSpec {
	if in == nil {
		return nil
	}
	out := new(MainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentStorageSpec) DeepCopyInto(out *PersistentStorageSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(corev1.Volume)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentStorageSpec.
func (in *PersistentStorageSpec) DeepCopy() *PersistentStorageSpec {
	if in == nil {
		return nil
	}
	out := new(PersistentStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSpec) DeepCopyInto(out *PortSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(SSLSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSpec.
func (in *PortSpec) DeepCopy() *PortSpec {
	if in == nil {
		return nil
	}
	out := new(PortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessProbeSpec) DeepCopyInto(out *ReadinessProbeSpec) {
	*out = *in
	if in.InitialDelaySeconds != nil {
		in, out := &in.InitialDelaySeconds, &out.InitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SuccessThreshold != nil {
		in, out := &in.SuccessThreshold, &out.SuccessThreshold
		*out = new(int32)
		**out = **in
	}
	if in.FailureThreshold != nil {
		in, out := &in.FailureThreshold, &out.FailureThreshold
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessProbeSpec.
func (in *ReadinessProbeSpec) DeepCopy() *ReadinessProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ReadinessProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLSpec) DeepCopyInto(out *SSLSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(string)
		**out = **in
	}
	if in.KeyStore != nil {
		in, out := &in.KeyStore, &out.KeyStore
		*out = new(string)
		**out = **in
	}
	if in.KeyStorePasswordFile != nil {
		in, out := &in.KeyStorePasswordFile, &out.KeyStorePasswordFile
		*out = new(string)
		**out = **in
	}
	if in.KeyPasswordFile != nil {
		in, out := &in.KeyPasswordFile, &out.KeyPasswordFile
		*out = new(string)
		**out = **in
	}
	if in.KeyStoreAlgorithm != nil {
		in, out := &in.KeyStoreAlgorithm, &out.KeyStoreAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.KeyStoreProvider != nil {
		in, out := &in.KeyStoreProvider, &out.KeyStoreProvider
		*out = new(string)
		**out = **in
	}
	if in.KeyStoreType != nil {
		in, out := &in.KeyStoreType, &out.KeyStoreType
		*out = new(string)
		**out = **in
	}
	if in.TrustStore != nil {
		in, out := &in.TrustStore, &out.TrustStore
		*out = new(string)
		**out = **in
	}
	if in.TrustStorePasswordFile != nil {
		in, out := &in.TrustStorePasswordFile, &out.TrustStorePasswordFile
		*out = new(string)
		**out = **in
	}
	if in.TrustStoreAlgorithm != nil {
		in, out := &in.TrustStoreAlgorithm, &out.TrustStoreAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.TrustStoreProvider != nil {
		in, out := &in.TrustStoreProvider, &out.TrustStoreProvider
		*out = new(string)
		**out = **in
	}
	if in.TrustStoreType != nil {
		in, out := &in.TrustStoreType, &out.TrustStoreType
		*out = new(string)
		**out = **in
	}
	if in.RequireClientCert != nil {
		in, out := &in.RequireClientCert, &out.RequireClientCert
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLSpec.
func (in *SSLSpec) DeepCopy() *SSLSpec {
	if in == nil {
		return nil
	}
	out := new(SSLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerIP != nil {
		in, out := &in.LoadBalancerIP, &out.LoadBalancerIP
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExternalPort != nil {
		in, out := &in.ExternalPort, &out.ExternalPort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserArtifactsImageSpec) DeepCopyInto(out *UserArtifactsImageSpec) {
	*out = *in
	in.ImageSpec.DeepCopyInto(&out.ImageSpec)
	if in.LibDir != nil {
		in, out := &in.LibDir, &out.LibDir
		*out = new(string)
		**out = **in
	}
	if in.ConfigDir != nil {
		in, out := &in.ConfigDir, &out.ConfigDir
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserArtifactsImageSpec.
func (in *UserArtifactsImageSpec) DeepCopy() *UserArtifactsImageSpec {
	if in == nil {
		return nil
	}
	out := new(UserArtifactsImageSpec)
	in.DeepCopyInto(out)
	return out
}
