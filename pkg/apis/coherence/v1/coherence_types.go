/*
 * Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.
 * Licensed under the Universal Permissive License v 1.0 as shown at
 * http://oss.oracle.com/licenses/upl.
 */

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
	"time"
)

// Common Coherence API structs

// NOTE: This file is used to generate the CRDs use by the Operator. The CRD files should not be manually edited
// NOTE: json tags are required. Any new fields you add must have json tags for the fields to be serialized.

// ----- constants ----------------------------------------------------------

const (
	// The default number of replicas that will be created for a role if no value is specified in the spec
	DefaultReplicas int32 = 3

	// The defaultrole name that will be used for a role if no value is specified in the spec
	DefaultRoleName = "storage"

	// The suffix appended to a cluster name to give the WKA service name
	WKAServiceNameSuffix = "-wka"

	// The key of the label used to hold the Coherence cluster name
	CoherenceClusterLabel string = "coherenceCluster"

	// The key of the label used to hold the Coherence role name
	CoherenceRoleLabel string = "coherenceRole"

	// The key of the label used to hold the component name
	CoherenceComponentLabel string = "component"
)

// ----- ApplicationSpec struct ---------------------------------------------

// The specification of the application deployed into the Coherence
// role members.
// +k8s:openapi-gen=true
type ApplicationSpec struct {
	// The application type to execute.
	// This field would be set if using the Coherence Graal image and running a none-Java
	// application. For example if the application was a Node application this field
	// would be set to "node". The default is to run a plain Java application.
	// +optional
	Type *string `json:"type,omitempty"`
	// Class is the Coherence container main class.  The default value is
	// com.tangosol.net.DefaultCacheServer.
	// +optional
	MainClass *string `json:"mainClass,omitempty"`
	// Args is the optional arguments to pass to the main class.
	// +optional
	Args []string `json:"args,omitempty"`
	// The inlined application image definition
	ImageSpec `json:",inline"`
	// The folder in the custom artifacts Docker image containing jar
	// files to be added to the classpath of the Coherence container.
	// If not set the libDir is "/files/lib".
	// +optional
	LibDir *string `json:"libDir,omitempty"`
	// The folder in the custom artifacts Docker image containing
	// configuration files to be added to the classpath of the Coherence container.
	// If not set the configDir is "/files/conf".
	// +optional
	ConfigDir *string `json:"configDir,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this ApplicationSpec struct with any nil or not set
// values set by the corresponding value in the defaults Images struct.
func (in *ApplicationSpec) DeepCopyWithDefaults(defaults *ApplicationSpec) *ApplicationSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := ApplicationSpec{}
	clone.ImageSpec = *in.ImageSpec.DeepCopyWithDefaults(&defaults.ImageSpec)

	if in.Type != nil {
		clone.Type = in.Type
	} else {
		clone.Type = defaults.Type
	}

	if in.MainClass != nil {
		clone.MainClass = in.MainClass
	} else {
		clone.MainClass = defaults.MainClass
	}

	if in.Args != nil {
		clone.Args = in.Args
	} else {
		clone.Args = defaults.Args
	}

	if in.LibDir != nil {
		clone.LibDir = in.LibDir
	} else {
		clone.LibDir = defaults.LibDir
	}

	if in.ConfigDir != nil {
		clone.ConfigDir = in.ConfigDir
	} else {
		clone.ConfigDir = defaults.ConfigDir
	}

	return &clone
}

// ----- CoherenceSpec struct -----------------------------------------------

// The Coherence specific configuration.
// +k8s:openapi-gen=true
type CoherenceSpec struct {
	// The Coherence images configuration.
	// +optional
	Images *CoherenceImagesSpec `json:"images,omitempty"`
	// A boolean flag indicating whether members of this role are storage enabled.
	// This value will set the corresponding coherence.distributed.localstorage System property.
	// If not specified the default value is true.
	// This flag is also used to configure the ScalingPolicy value if a value is not specified. If the
	// StorageEnabled field is not specified or is true the scaling will be safe, if StorageEnabled is
	// set to false scaling will be parallel.
	// +optional
	StorageEnabled *bool `json:"storageEnabled,omitempty"`
	// ScalingPolicy describes how the replicas of the cluster role will be scaled.
	// The default if not specified is based upon the value of the StorageEnabled field.
	// If StorageEnabled field is not specified or is true the default scaling will be safe, if StorageEnabled is
	// set to false the default scaling will be parallel.
	// +optional
	ScalingPolicy *ScalingPolicy `json:"scalingPolicy,omitempty"`
	// CacheConfig is the name of the cache configuration file to use
	// +optional
	CacheConfig *string `json:"cacheConfig,omitempty"`
	// PofConfig is the name of the POF configuration file to use when using POF serializer
	// +optional
	PofConfig *string `json:"pofConfig,omitempty"`
	// OverrideConfig is name of the Coherence operational configuration override file,
	// the default is tangosol-coherence-override.xml
	// +optional
	OverrideConfig *string `json:"overrideConfig,omitempty"`
	// Persistence values configure the on-disc data persistence settings.
	// The bool Enabled enables or disabled on disc persistence of data.
	// +optional
	Persistence *PersistentStorageSpec `json:"persistence,omitempty"`
	// Snapshot values configure the on-disc persistence data snapshot (backup) settings.
	// The bool Enabled enables or disabled a different location for
	// persistence snapshot data. If set to false then snapshot files will be written
	// to the same volume configured for persistence data in the Persistence section.
	// +optional
	Snapshot *PersistentStorageSpec `json:"snapshot,omitempty"`
	// Management configures Coherence management over REST
	//   Note: Coherence management over REST will be available in 12.2.1.4.
	// +optional
	Management *PortSpecWithSSL `json:"management,omitempty"`
	// Metrics configures Coherence metrics publishing
	//   Note: Coherence metrics publishing will be available in 12.2.1.4.
	// +optional
	Metrics *PortSpecWithSSL `json:"metrics,omitempty"`
	// The handler to use to determine whether a role is Status HA.
	// If not set the default handler will be used.
	// In most use-cases the default handler would suffice but in
	// advanced use-cases where the application code has a different
	// concept of Status HA to just checking Coherence services then
	// a different handler may be specified.
	// +optional
	StatusHA *StatusHAHandler `json:"statusHA,omitempty"`
	// The timeout in seconds used by curl when requesting site and rack info.
	// +optional
	CurlTimeout *int `json:"curlTimeout,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this CoherenceSpec struct with any nil or not set
// values set by the corresponding value in the defaults CoherenceSpec struct.
func (in *CoherenceSpec) DeepCopyWithDefaults(defaults *CoherenceSpec) *CoherenceSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := CoherenceSpec{}
	clone.Images = in.Images.DeepCopyWithDefaults(defaults.Images)
	clone.Persistence = in.Persistence.DeepCopyWithDefaults(defaults.Persistence)
	clone.Snapshot = in.Snapshot.DeepCopyWithDefaults(defaults.Snapshot)
	clone.Management = in.Management.DeepCopyWithDefaults(defaults.Management)
	clone.Metrics = in.Metrics.DeepCopyWithDefaults(defaults.Metrics)

	if in.StorageEnabled != nil {
		clone.StorageEnabled = in.StorageEnabled
	} else {
		clone.StorageEnabled = defaults.StorageEnabled
	}

	if in.ScalingPolicy != nil {
		clone.ScalingPolicy = in.ScalingPolicy
	} else {
		clone.ScalingPolicy = defaults.ScalingPolicy
	}

	if in.CacheConfig != nil {
		clone.CacheConfig = in.CacheConfig
	} else {
		clone.CacheConfig = defaults.CacheConfig
	}

	if in.PofConfig != nil {
		clone.PofConfig = in.PofConfig
	} else {
		clone.PofConfig = defaults.PofConfig
	}

	if in.OverrideConfig != nil {
		clone.OverrideConfig = in.OverrideConfig
	} else {
		clone.OverrideConfig = defaults.OverrideConfig
	}

	if in.StatusHA != nil {
		clone.StatusHA = in.StatusHA
	} else {
		clone.StatusHA = defaults.StatusHA
	}

	if in.CurlTimeout != nil {
		clone.CurlTimeout = in.CurlTimeout
	} else {
		clone.CurlTimeout = defaults.CurlTimeout
	}

	return &clone
}

// ----- CoherenceImagesSpec struct -----------------------------------------

// The Coherence specific images configuration.
// +k8s:openapi-gen=true
type CoherenceImagesSpec struct {
	Coherence      *ImageSpec `json:"coherence,omitempty"`
	CoherenceUtils *ImageSpec `json:"coherenceUtils,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this CoherenceImagesSpec struct with any nil or not set
// values set by the corresponding value in the defaults CoherenceImagesSpec struct.
func (in *CoherenceImagesSpec) DeepCopyWithDefaults(defaults *CoherenceImagesSpec) *CoherenceImagesSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := CoherenceImagesSpec{}
	clone.Coherence = in.Coherence.DeepCopyWithDefaults(defaults.Coherence)
	clone.CoherenceUtils = in.CoherenceUtils.DeepCopyWithDefaults(defaults.CoherenceUtils)

	return &clone
}

// ----- JVMSpec struct -----------------------------------------------------

// The JVM specific configuration.
// +k8s:openapi-gen=true
type JVMSpec struct {
	// MaxHeap is the min/max heap value to pass to the JVM.
	// The format should be the same as that used for Java's -Xms and -Xmx JVM options.
	// If not set the JVM defaults are used.
	// +optional
	MaxHeap *string `json:"maxHeap,omitempty"`
	// The optional GC parameters. If not set defaults to enabling the G1 collector.
	// +optional
	GC *string `json:"gc,omitempty"`
	// Args specifies the options (System properties, -XX: args etc) to pass to the JVM.
	// +optional
	Args []string `json:"args,omitempty"`
	// The settings for enabling debug mode in the JVM.
	// +optional
	Debug *DebugSpec `json:"debug,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this JVMSpec struct with any nil or not set
// values set by the corresponding value in the defaults JVMSpec struct.
func (in *JVMSpec) DeepCopyWithDefaults(defaults *JVMSpec) *JVMSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := JVMSpec{}
	clone.Debug = in.Debug.DeepCopyWithDefaults(defaults.Debug)

	if in.MaxHeap != nil {
		clone.MaxHeap = in.MaxHeap
	} else {
		clone.MaxHeap = defaults.MaxHeap
	}

	if in.GC != nil {
		clone.GC = in.GC
	} else {
		clone.GC = defaults.GC
	}

	if in.Args != nil {
		// Merge Args
		clone.Args = []string{}
		clone.Args = append(clone.Args, in.Args...)
		clone.Args = append(clone.Args, defaults.Args...)
	} else {
		clone.Args = defaults.Args
	}

	return &clone
}

// ----- ImageSpec struct ---------------------------------------------------

// CoherenceInternalImageSpec defines the settings for a Docker image
// +k8s:openapi-gen=true
type ImageSpec struct {
	// Docker image name.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// +optional
	Image *string `json:"image,omitempty"`
	// Image pull policy.
	// One of Always, Never, IfNotPresent.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// +optional
	ImagePullPolicy *corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this ImageSpec struct with any nil or not set values set
// by the corresponding value in the defaults ImageSpec struct.
func (in *ImageSpec) DeepCopyWithDefaults(defaults *ImageSpec) *ImageSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := ImageSpec{}

	if in.Image != nil {
		clone.Image = in.Image
	} else {
		clone.Image = defaults.Image
	}

	if in.ImagePullPolicy != nil {
		clone.ImagePullPolicy = in.ImagePullPolicy
	} else {
		clone.ImagePullPolicy = defaults.ImagePullPolicy
	}

	return &clone
}

// ----- LoggingSpec struct -------------------------------------------------

// LoggingSpec defines the settings for the Coherence Pod logging
// +k8s:openapi-gen=true
type LoggingSpec struct {
	// The default being 5 (info level).
	// +optional
	Level *int32 `json:"level,omitempty"`
	// ConfigFile allows the location of the Java util logging configuration file to be overridden.
	//  If this value is not set the logging.properties file embedded in this chart will be used.
	//  If this value is set the configuration will be located by trying the following locations in order:
	//    1. If store.logging.configMapName is set then the config map will be mounted as a volume and the logging
	//         properties file will be located as a file location relative to the ConfigMap volume mount point.
	//    2. If userArtifacts.imageName is set then using this value as a file name relative to the location of the
	//         configuration files directory in the user artifacts image.
	//    3. Using this value as an absolute file name.
	// +optional
	ConfigFile *string `json:"configFile,omitempty"`
	// ConfigMapName allows a config map to be mounted as a volume containing the logging
	//  configuration file to use.
	// +optional
	ConfigMapName *string `json:"configMapName,omitempty"`
	// Configures whether Fluentd is enabled and the configuration
	// of the Fluentd side-car container
	// +optional
	Fluentd *FluentdSpec `json:"fluentd,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this LoggingSpec struct with any nil or not set values set
// by the corresponding value in the defaults LoggingSpec struct.
func (in *LoggingSpec) DeepCopyWithDefaults(defaults *LoggingSpec) *LoggingSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := LoggingSpec{}
	clone.Fluentd = in.Fluentd.DeepCopyWithDefaults(defaults.Fluentd)

	if in.Level != nil {
		clone.Level = in.Level
	} else {
		clone.Level = defaults.Level
	}

	if in.ConfigFile != nil {
		clone.ConfigFile = in.ConfigFile
	} else {
		clone.ConfigFile = defaults.ConfigFile
	}

	if in.ConfigMapName != nil {
		clone.ConfigMapName = in.ConfigMapName
	} else {
		clone.ConfigMapName = defaults.ConfigMapName
	}

	return &clone
}

// ----- PersistentStorageSpec struct ---------------------------------------

// PersistenceStorageSpec defines the persistence settings for the Coherence
// +k8s:openapi-gen=true
type PersistentStorageSpec struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
	// PersistentVolumeClaim allows the configuration of a normal k8s persistent volume claim
	// for persistence data.
	// +optional
	PersistentVolumeClaim *corev1.PersistentVolumeClaimSpec `json:"persistentVolumeClaim,omitempty"` // from k8s.io/api/core/v1
	// Volume allows the configuration of a normal k8s volume mapping
	// for persistence data instead of a persistent volume claim. If a value is defined
	// for store.persistence.volume then no PVC will be created and persistence data
	// will instead be written to this volume. It is up to the deployer to understand
	// the consequences of this and how the guarantees given when using PVCs differ
	// to the storage guarantees for the particular volume type configured here.
	// +optional
	Volume *corev1.VolumeSource `json:"volume,omitempty"` // from k8s.io/api/core/v1
}

// DeepCopyWithDefaults returns a copy of this PersistentStorageSpec struct with any nil or not set values set
// by the corresponding value in the defaults PersistentStorageSpec struct.
func (in *PersistentStorageSpec) DeepCopyWithDefaults(defaults *PersistentStorageSpec) *PersistentStorageSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := PersistentStorageSpec{}

	if in.Enabled != nil {
		clone.Enabled = in.Enabled
	} else {
		clone.Enabled = defaults.Enabled
	}

	if in.PersistentVolumeClaim != nil {
		clone.PersistentVolumeClaim = in.PersistentVolumeClaim
	} else {
		clone.PersistentVolumeClaim = defaults.PersistentVolumeClaim
	}

	if in.Volume != nil {
		clone.Volume = in.Volume
	} else {
		clone.Volume = defaults.Volume
	}

	return &clone
}

// ----- SSLSpec struct -----------------------------------------------------

// SSLSpec defines the SSL settings for a Coherence component over REST endpoint.
// +k8s:openapi-gen=true
type SSLSpec struct {
	// Enabled is a boolean flag indicating whether enables or disables SSL on the Coherence management
	// over REST endpoint, the default is false (disabled).
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
	// Secrets is the name of the k8s secrets containing the Java key stores and password files.
	//   This value MUST be provided if SSL is enabled on the Coherence management over ReST endpoint.
	// +optional
	Secrets *string `json:"secrets,omitempty"`
	// Keystore is the name of the Java key store file in the k8s secret to use as the SSL keystore
	//   when configuring component over REST to use SSL.
	// +optional
	KeyStore *string `json:"keyStore,omitemtpy"`
	// KeyStorePasswordFile is the name of the file in the k8s secret containing the keystore
	//   password when configuring component over REST to use SSL.
	// +optional
	KeyStorePasswordFile *string `json:"keyStorePasswordFile,omitempty"`
	// KeyStorePasswordFile is the name of the file in the k8s secret containing the key
	//   password when configuring component over REST to use SSL.
	// +optional
	KeyPasswordFile *string `json:"keyPasswordFile,omitempty"`
	// KeyStoreAlgorithm is the name of the keystore algorithm for the keystore in the k8s secret
	//   used when configuring component over REST to use SSL. If not set the default is SunX509
	// +optional
	KeyStoreAlgorithm *string `json:"keyStoreAlgorithm,omitempty"`
	// KeyStoreProvider is the name of the keystore provider for the keystore in the k8s secret
	//   used when configuring component over REST to use SSL.
	// +optional
	KeyStoreProvider *string `json:"keyStoreProvider,omitempty"`
	// KeyStoreType is the name of the Java keystore type for the keystore in the k8s secret used
	//   when configuring component over REST to use SSL. If not set the default is JKS.
	// +optional
	KeyStoreType *string `json:"keyStoreType,omitempty"`
	// TrustStore is the name of the Java trust store file in the k8s secret to use as the SSL
	//   trust store when configuring component over REST to use SSL.
	// +optional
	TrustStore *string `json:"trustStore,omitempty"`
	// TrustStorePasswordFile is the name of the file in the k8s secret containing the trust store
	//   password when configuring component over REST to use SSL.
	// +optional
	TrustStorePasswordFile *string `json:"trustStorePasswordFile,omitempty"`
	// TrustStoreAlgorithm is the name of the keystore algorithm for the trust store in the k8s
	//   secret used when configuring component over REST to use SSL.  If not set the default is SunX509.
	// +optional
	TrustStoreAlgorithm *string `json:"trustStoreAlgorithm,omitempty"`
	// TrustStoreProvider is the name of the keystore provider for the trust store in the k8s
	//   secret used when configuring component over REST to use SSL.
	// +optional
	TrustStoreProvider *string `json:"trustStoreProvider,omitempty"`
	// TrustStoreType is the name of the Java keystore type for the trust store in the k8s secret
	//   used when configuring component over REST to use SSL. If not set the default is JKS.
	// +optional
	TrustStoreType *string `json:"trustStoreType,omitempty"`
	// RequireClientCert is a boolean flag indicating whether the client certificate will be
	//   authenticated by the server (two-way SSL) when configuring component over REST to use SSL.
	//   If not set the default is false
	// +optional
	RequireClientCert *bool `json:"requireClientCert,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this SSLSpec struct with any nil or not set values set
// by the corresponding value in the defaults SSLSpec struct.
func (in *SSLSpec) DeepCopyWithDefaults(defaults *SSLSpec) *SSLSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := SSLSpec{}

	if in.Enabled != nil {
		clone.Enabled = in.Enabled
	} else {
		clone.Enabled = defaults.Enabled
	}

	if in.Secrets != nil {
		clone.Secrets = in.Secrets
	} else {
		clone.Secrets = defaults.Secrets
	}

	if in.KeyStore != nil {
		clone.KeyStore = in.KeyStore
	} else {
		clone.KeyStore = defaults.KeyStore
	}

	if in.KeyStorePasswordFile != nil {
		clone.KeyStorePasswordFile = in.KeyStorePasswordFile
	} else {
		clone.KeyStorePasswordFile = defaults.KeyStorePasswordFile
	}

	if in.KeyPasswordFile != nil {
		clone.KeyPasswordFile = in.KeyPasswordFile
	} else {
		clone.KeyPasswordFile = defaults.KeyPasswordFile
	}

	if in.KeyStoreAlgorithm != nil {
		clone.KeyStoreAlgorithm = in.KeyStoreAlgorithm
	} else {
		clone.KeyStoreAlgorithm = defaults.KeyStoreAlgorithm
	}

	if in.KeyStoreProvider != nil {
		clone.KeyStoreProvider = in.KeyStoreProvider
	} else {
		clone.KeyStoreProvider = defaults.KeyStoreProvider
	}

	if in.KeyStoreType != nil {
		clone.KeyStoreType = in.KeyStoreType
	} else {
		clone.KeyStoreType = defaults.KeyStoreType
	}

	if in.TrustStore != nil {
		clone.TrustStore = in.TrustStore
	} else {
		clone.TrustStore = defaults.TrustStore
	}

	if in.TrustStorePasswordFile != nil {
		clone.TrustStorePasswordFile = in.TrustStorePasswordFile
	} else {
		clone.TrustStorePasswordFile = defaults.TrustStorePasswordFile
	}

	if in.TrustStoreAlgorithm != nil {
		clone.TrustStoreAlgorithm = in.TrustStoreAlgorithm
	} else {
		clone.TrustStoreAlgorithm = defaults.TrustStoreAlgorithm
	}

	if in.TrustStoreProvider != nil {
		clone.TrustStoreProvider = in.TrustStoreProvider
	} else {
		clone.TrustStoreProvider = defaults.TrustStoreProvider
	}

	if in.TrustStoreType != nil {
		clone.TrustStoreType = in.TrustStoreType
	} else {
		clone.TrustStoreType = defaults.TrustStoreType
	}

	if in.RequireClientCert != nil {
		clone.RequireClientCert = in.RequireClientCert
	} else {
		clone.RequireClientCert = defaults.RequireClientCert
	}

	return &clone
}

// ----- PortSpec struct ----------------------------------------------------
// PortSpec defines the port settings for a Coherence component
// +k8s:openapi-gen=true
type PortSpec struct {
	// Port specifies the port used.
	// +optional
	Port int32 `json:"port,omitempty"`
	// Protocol for container port. Must be UDP or TCP. Defaults to "TCP"
	// +optional
	Protocol *string `json:"protocol,omitempty"`
	// Service specifies the service used to expose the port.
	// +optional
	Service *ServiceSpec `json:"service,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this PortSpec struct with any nil or not set values set
// by the corresponding value in the defaults PortSpec struct.
func (in *PortSpec) DeepCopyWithDefaults(defaults *PortSpec) *PortSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := PortSpec{}

	if in.Port != 0 {
		clone.Port = in.Port
	} else {
		clone.Port = defaults.Port
	}

	if in.Protocol != nil {
		clone.Protocol = in.Protocol
	} else {
		clone.Protocol = defaults.Protocol
	}

	if in.Service != nil {
		clone.Service = in.Service
	} else {
		clone.Service = defaults.Service
	}

	return &clone
}

// ----- NamedPortSpec struct ----------------------------------------------------
// NamedPortSpec defines a named port for a Coherence component
// +k8s:openapi-gen=true
type NamedPortSpec struct {
	// Name specifies the name of th port.
	// +optional
	Name     string `json:"name,omitempty"`
	PortSpec `json:",inline"`
}

// DeepCopyWithDefaults returns a copy of this NamedPortSpec struct with any nil or not set values set
// by the corresponding value in the defaults NamedPortSpec struct.
func (in *NamedPortSpec) DeepCopyWithDefaults(defaults *NamedPortSpec) *NamedPortSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := NamedPortSpec{}

	if in.Name != "" {
		clone.Name = in.Name
	} else {
		clone.Name = defaults.Name
	}

	if in.Port != 0 {
		clone.Port = in.Port
	} else {
		clone.Port = defaults.Port
	}

	if in.Protocol != nil {
		clone.Protocol = in.Protocol
	} else {
		clone.Protocol = defaults.Protocol
	}

	if in.Service != nil {
		clone.Service = in.Service
	} else {
		clone.Service = defaults.Service
	}

	return &clone
}

// Merge merges two arrays of NamedPortSpec structs.
// Any NamedPortSpec instances in both arrays that share the same name will be merged,
// the field set in the primary NamedPortSpec will take presedence over those in the
// secondary NamedPortSpec.
func MergeNamedPortSpecs(primary, secondary []NamedPortSpec) []NamedPortSpec {
	if primary == nil {
		return secondary
	}

	if secondary == nil {
		return primary
	}

	if len(primary) == 0 && len(secondary) == 0 {
		return []NamedPortSpec{}
	}

	var mr []NamedPortSpec
	mr = append(mr, primary...)

	for _, p := range secondary {
		found := false
		for i, pp := range primary {
			if pp.Name == p.Name {
				clone := pp.DeepCopyWithDefaults(&p)
				mr[i] = *clone
				found = true
				break
			}
		}

		if !found {
			mr = append(mr, p)
		}
	}

	return mr
}

// ----- DebugSpec struct ----------------------------------------------------------

// The JVM Debug specific configuration.
// +k8s:openapi-gen=true
type DebugSpec struct {
	// Enabled is a flag to enable or disable running the JVM in debug mode. Default is disabled.
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
	// Attach specifies the address of the debugger that the JVM should attempt to connect back to
	// instead of listening on port 5005.
	// +optional
	Attach *string `json:"attach,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this DebugSpec struct with any nil or not set values set
// by the corresponding value in the defaults DebugSpec struct.
func (in *DebugSpec) DeepCopyWithDefaults(defaults *DebugSpec) *DebugSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := DebugSpec{}

	if in.Enabled != nil {
		clone.Enabled = in.Enabled
	} else {
		clone.Enabled = defaults.Enabled
	}

	if in.Attach != nil {
		clone.Attach = in.Attach
	} else {
		clone.Attach = defaults.Attach
	}

	return &clone
}

// ----- PortSpecWithSSL struct ----------------------------------------------------

// PortSpecWithSSL defines a port with SSL settings for a Coherence component
// +k8s:openapi-gen=true
type PortSpecWithSSL struct {
	// Enable or disable flag.
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
	// SSL configures SSL settings for a Coherence component
	// +optional
	SSL *SSLSpec `json:"ssl,omitempty"`
}

// IsSSLEnabled returns true if this port is SSL enabled
func (in *PortSpecWithSSL) IsSSLEnabled() bool {
	return in != nil && in.Enabled != nil && *in.Enabled
}

// DeepCopyWithDefaults returns a copy of this PortSpecWithSSL struct with any nil or not set values set
// by the corresponding value in the defaults PortSpecWithSSL struct.
func (in *PortSpecWithSSL) DeepCopyWithDefaults(defaults *PortSpecWithSSL) *PortSpecWithSSL {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := PortSpecWithSSL{}

	if in.Enabled != nil {
		clone.Enabled = in.Enabled
	} else {
		clone.Enabled = defaults.Enabled
	}

	if in.SSL != nil {
		clone.SSL = in.SSL
	} else {
		clone.SSL = defaults.SSL
	}

	return &clone
}

// ----- ServiceSpec struct -------------------------------------------------
// ServiceSpec defines the settings for a Service
// +k8s:openapi-gen=true
type ServiceSpec struct {
	// Enabled controls whether to create the service yaml or not
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
	// An optional name to use to override the generated service name.
	// +optional
	Name *string `json:"name,omitempty"`
	// The service port value
	// +optional
	Port *int32 `json:"port,omitempty"`
	// Type is the K8s service type (typically ClusterIP or LoadBalancer)
	// The default is "ClusterIP".
	// +optional
	Type *corev1.ServiceType `json:"type,omitempty"`
	// LoadBalancerIP is the IP address of the load balancer
	// +optional
	LoadBalancerIP *string `json:"loadBalancerIP,omitempty"`
	// Annotations is free form yaml that will be added to the service annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// Supports "ClientIP" and "None". Used to maintain session affinity.
	// Enable client IP based session affinity.
	// Must be ClientIP or None.
	// Defaults to None.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +optional
	SessionAffinity *corev1.ServiceAffinity `json:"sessionAffinity,omitempty"`
	// If specified and supported by the platform, this will restrict traffic through the cloud-provider
	// load-balancer will be restricted to the specified client IPs. This field will be ignored if the
	// cloud-provider does not support the feature."
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	// +optional
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty"`
	// externalName is the external reference that kubedns or equivalent will
	// return as a CNAME record for this service. No proxying will be involved.
	// Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123)
	// and requires Type to be ExternalName.
	// +optional
	ExternalName *string `json:"externalName,omitempty"`
	// externalTrafficPolicy denotes if this Service desires to route external
	// traffic to node-local or cluster-wide endpoints. "Local" preserves the
	// client source IP and avoids a second hop for LoadBalancer and Nodeport
	// type services, but risks potentially imbalanced traffic spreading.
	// "Cluster" obscures the client source IP and may cause a second hop to
	// another node, but should have good overall load-spreading.
	// +optional
	ExternalTrafficPolicy *corev1.ServiceExternalTrafficPolicyType `json:"externalTrafficPolicy,omitempty"`
	// healthCheckNodePort specifies the healthcheck nodePort for the service.
	// If not specified, HealthCheckNodePort is created by the service api
	// backend with the allocated nodePort. Will use user-specified nodePort value
	// if specified by the client. Only effects when Type is set to LoadBalancer
	// and ExternalTrafficPolicy is set to Local.
	// +optional
	HealthCheckNodePort *int32 `json:"healthCheckNodePort,omitempty"`
	// publishNotReadyAddresses, when set to true, indicates that DNS implementations
	// must publish the notReadyAddresses of subsets for the Endpoints associated with
	// the Service. The default value is false.
	// The primary use case for setting this field is to use a StatefulSet's Headless Service
	// to propagate SRV records for its Pods without respect to their readiness for purpose
	// of peer discovery.
	// +optional
	PublishNotReadyAddresses *bool `json:"publishNotReadyAddresses,omitempty"`
	// sessionAffinityConfig contains the configurations of session affinity.
	// +optional
	SessionAffinityConfig *corev1.SessionAffinityConfig `json:"sessionAffinityConfig,omitempty"`
}

// Set the Type of the service.
func (in *ServiceSpec) SetServiceType(t corev1.ServiceType) {
	if in != nil {
		in.Type = &t
	}
}

// DeepCopyWithDefaults returns a copy of this ServiceSpec struct with any nil or not set values set
// by the corresponding value in the defaults PortSpecWithSSL struct.
func (in *ServiceSpec) DeepCopyWithDefaults(defaults *ServiceSpec) *ServiceSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := ServiceSpec{}

	if in.Enabled != nil {
		clone.Enabled = in.Enabled
	} else {
		clone.Enabled = defaults.Enabled
	}

	if in.Type != nil {
		clone.Type = in.Type
	} else {
		clone.Type = defaults.Type
	}

	if in.Name != nil {
		clone.Name = in.Name
	} else {
		clone.Name = defaults.Name
	}

	if in.Port != nil {
		clone.Port = in.Port
	} else {
		clone.Port = defaults.Port
	}

	if in.LoadBalancerIP != nil {
		clone.LoadBalancerIP = in.LoadBalancerIP
	} else {
		clone.LoadBalancerIP = defaults.LoadBalancerIP
	}

	if in.Annotations != nil {
		clone.Annotations = in.Annotations
	} else {
		clone.Annotations = defaults.Annotations
	}

	if in.Port != nil {
		clone.Port = in.Port
	} else {
		clone.Port = defaults.Port
	}

	if in.SessionAffinity != nil {
		clone.SessionAffinity = in.SessionAffinity
	} else {
		clone.SessionAffinity = defaults.SessionAffinity
	}

	if in.LoadBalancerSourceRanges != nil {
		clone.LoadBalancerSourceRanges = in.LoadBalancerSourceRanges
	} else {
		clone.LoadBalancerSourceRanges = defaults.LoadBalancerSourceRanges
	}

	if in.ExternalName != nil {
		clone.ExternalName = in.ExternalName
	} else {
		clone.ExternalName = defaults.ExternalName
	}

	if in.ExternalTrafficPolicy != nil {
		clone.ExternalTrafficPolicy = in.ExternalTrafficPolicy
	} else {
		clone.ExternalTrafficPolicy = defaults.ExternalTrafficPolicy
	}

	if in.HealthCheckNodePort != nil {
		clone.HealthCheckNodePort = in.HealthCheckNodePort
	} else {
		clone.HealthCheckNodePort = defaults.HealthCheckNodePort
	}

	if in.PublishNotReadyAddresses != nil {
		clone.PublishNotReadyAddresses = in.PublishNotReadyAddresses
	} else {
		clone.PublishNotReadyAddresses = defaults.PublishNotReadyAddresses
	}

	if in.SessionAffinityConfig != nil {
		clone.SessionAffinityConfig = in.SessionAffinityConfig
	} else {
		clone.SessionAffinityConfig = defaults.SessionAffinityConfig
	}

	return &clone
}

// ----- StatusHAHandler ----------------------------------------------------

// StatusHAHandler is the handler that will be used to determine how to check for StatusHA in a CoherenceRole.
// StatusHA checking is primarily used during scaling of a role, a role must be in a safe Status HA state
// before scaling takes place. If StatusHA handler is disabled for a role (by specifically setting Enabled
// to false then no check will take place and a role will be assumed to be safe).
// +k8s:openapi-gen=true
type StatusHAHandler struct {
	corev1.Handler `json:",inline"`
	// Number of seconds after which the handler times out (only applies to http and tcp handlers).
	// Defaults to 1 second. Minimum value is 1.
	// +optional
	TimeoutSeconds int `json:"timeoutSeconds,omitempty"`
	// An optional flag to enable or disable the StatusHA check.
	// The default value if not set is true.
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

var defaultStatusHA = StatusHAHandler{
	Enabled:        pointer.BoolPtr(true),
	TimeoutSeconds: 10,
	Handler: corev1.Handler{
		Exec:      &corev1.ExecAction{Command: []string{"/bin/sh", "-x", "/scripts/startCoherence.sh", "probe", "com.oracle.coherence.k8s.PodChecker", "statusha"}},
		HTTPGet:   nil,
		TCPSocket: nil,
	},
}

// Returns true if this handler is enabled.
func (in *StatusHAHandler) IsEnabled() bool {
	if in == nil {
		return false
	}

	return in.Enabled == nil || *in.Enabled
}

// Returns the timeout value in seconds.
func (in *StatusHAHandler) GetTimeout() time.Duration {
	if in == nil || in.TimeoutSeconds <= 0 {
		return time.Second
	}

	return time.Second * time.Duration(in.TimeoutSeconds)
}

// Obtain a default StatusHAHandler
func GetDefaultStatusHAHandler() *StatusHAHandler {
	return defaultStatusHA.DeepCopy()
}

// ----- ReadinessProbeSpec struct ------------------------------------------

// ReadinessProbeSpec defines the settings for the Coherence Pod readiness probe
// +k8s:openapi-gen=true
type ReadinessProbeSpec struct {
	// Number of seconds after the container has started before liveness probes are initiated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`
	// Number of seconds after which the probe times out.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
	// How often (in seconds) to perform the probe.
	// +optional
	PeriodSeconds *int32 `json:"periodSeconds,omitempty"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	// +optional
	SuccessThreshold *int32 `json:"successThreshold,omitempty"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	// +optional
	FailureThreshold *int32 `json:"failureThreshold,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this ReadinessProbeSpec struct with any nil or not set values set
// by the corresponding value in the defaults ReadinessProbeSpec struct.
func (in *ReadinessProbeSpec) DeepCopyWithDefaults(defaults *ReadinessProbeSpec) *ReadinessProbeSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := ReadinessProbeSpec{}

	if in.InitialDelaySeconds != nil {
		clone.InitialDelaySeconds = in.InitialDelaySeconds
	} else {
		clone.InitialDelaySeconds = defaults.InitialDelaySeconds
	}

	if in.TimeoutSeconds != nil {
		clone.TimeoutSeconds = in.TimeoutSeconds
	} else {
		clone.TimeoutSeconds = defaults.TimeoutSeconds
	}

	if in.PeriodSeconds != nil {
		clone.PeriodSeconds = in.PeriodSeconds
	} else {
		clone.PeriodSeconds = defaults.PeriodSeconds
	}

	if in.SuccessThreshold != nil {
		clone.SuccessThreshold = in.SuccessThreshold
	} else {
		clone.SuccessThreshold = defaults.SuccessThreshold
	}

	if in.FailureThreshold != nil {
		clone.FailureThreshold = in.FailureThreshold
	} else {
		clone.FailureThreshold = defaults.FailureThreshold
	}

	return &clone
}

// ----- FluentdSpec struct -------------------------------------------------

// FluentdSpec defines the settings for the fluentd image
// +k8s:openapi-gen=true
type FluentdSpec struct {
	ImageSpec `json:",inline"`
	// Controls whether or not log capture via a Fluentd sidecar container to an EFK stack is enabled.
	// If this flag i set to true it is expected that the coherence-monitoring-config secret exists in
	// the namespace that the cluster is being deployed to. This secret is either created by the
	// Coherence Operator Helm chart if it was installed with the correct parameters or it should
	// have already been created manually.
	Enabled *bool `json:"enabled,omitempty"`
	// The Fluentd configuration file configuring source for application log.
	// +optional
	ConfigFile *string `json:"configFile,omitempty"`
	// This value should be source.tag from fluentd.application.configFile.
	// +optional
	Tag *string `json:"tag,omitempty"`
}

// DeepCopyWithDefaults returns a copy of this FluentdSpec struct with any nil or not set values set
// by the corresponding value in the defaults FluentdSpec struct.
func (in *FluentdSpec) DeepCopyWithDefaults(defaults *FluentdSpec) *FluentdSpec {
	if in == nil {
		if defaults != nil {
			return defaults.DeepCopy()
		}
		return nil
	}

	if defaults == nil {
		return in.DeepCopy()
	}

	clone := FluentdSpec{}
	clone.ImageSpec = *in.ImageSpec.DeepCopyWithDefaults(&defaults.ImageSpec)

	if in.Enabled != nil {
		clone.Enabled = in.Enabled
	} else {
		clone.Enabled = defaults.Enabled
	}

	if in.ConfigFile != nil {
		clone.ConfigFile = in.ConfigFile
	} else {
		clone.ConfigFile = defaults.ConfigFile
	}

	if in.Tag != nil {
		clone.Tag = in.Tag
	} else {
		clone.Tag = defaults.Tag
	}

	return &clone
}

// ----- ScalingPolicy type -------------------------------------------------

// ScalingPolicy describes a policy for scaling a cluster role
type ScalingPolicy string

// Scaling policy constants
const (
	// Safe means that a role will be scaled up or down in a safe manner to ensure no data loss.
	SafeScaling ScalingPolicy = "Safe"
	// Parallel means that a role will be scaled up or down by adding or removing members in parallel.
	// If the members of the role are storage enabled then this could cause data loss
	ParallelScaling ScalingPolicy = "Parallel"
	// ParallelUpSafeDownScaling means that a role will be scaled up by adding or removing members in parallel
	// but will be scaled down in a safe manner to ensure no data loss.
	ParallelUpSafeDownScaling ScalingPolicy = "ParallelUpSafeDownScaling"
)
