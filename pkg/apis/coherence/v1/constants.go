/*
 * Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.
 * Licensed under the Universal Permissive License v 1.0 as shown at
 * http://oss.oracle.com/licenses/upl.
 */

package v1

const (
	// The default number of replicas that will be created for a deployment if no value is specified in the spec
	DefaultReplicas int32 = 3
	// The suffix appended to a deployment name to give the WKA service name
	WKAServiceNameSuffix = "-wka"
	// The suffix appended to a deployment name to give the StatefulSet headless-service name
	HeadlessServiceNameSuffix = "-sts"

	// Label keys used to label k8s resources
	LabelCoherenceDeployment = "coherenceDeployment"
	LabelCoherenceCluster    = "coherenceCluster"
	LabelCoherenceRole       = "coherenceRole"
	LabelComponent           = "coherenceComponent"
	LabelPort                = "coherencePort"
	LabelCoherenceWKAMember  = "coherenceWKAMember"

	// Values used for the component label in k8s resources
	LabelComponentCoherenceStatefulSet = "coherence"
	LabelComponentCoherencePod         = "coherencePod"
	LabelComponentCoherenceHeadless    = "coherence-headless"
	LabelComponentPVC                  = "coherence-volume"
	LabelComponentPortService          = "coherence-service"
	LabelComponentPortServiceMonitor   = "coherence-service-monitor"
	LabelComponentWKA                  = "coherenceWkaService"

	StatusSelectorTemplate = LabelCoherenceCluster + "=%s," + LabelCoherenceDeployment + "=%s"

	// The default k8s service account name.
	DefaultServiceAccount = "default"

	// The affinity topology key for fault domains.
	AffinityTopologyKey = "failure-domain.beta.kubernetes.io/zone"

	// Container Names
	ContainerNameCoherence = "coherence"
	ContainerNameUtils     = "coherence-k8s-utils"

	// Volume names
	VolumeNamePersistence   = "persistence-volume"
	VolumeNameSnapshots     = "snapshot-volume"
	VolumeNameLogs          = "logs"
	VolumeNameUtils         = "coh-utils"
	VolumeNameJVM           = "jvm"
	VolumeNameManagementSSL = "management-ssl-config"
	VolumeNameMetricsSSL    = "metrics-ssl-config"

	// Volume mount paths
	VolumeMountRoot                = "/coherence-operator"
	VolumeMountPathPersistence     = VolumeMountRoot + "/persistence"
	VolumeMountPathSnapshots       = VolumeMountRoot + "/snapshot"
	VolumeMountPathUtils           = VolumeMountRoot + "/utils"
	VolumeMountPathJVM             = VolumeMountRoot + "/jvm"
	VolumeMountPathManagementCerts = VolumeMountRoot + "/coherence/certs/management"
	VolumeMountPathMetricsCerts    = VolumeMountRoot + "/coherence/certs/metrics"

	// Start command for the runner
	RunnerCommand = VolumeMountPathUtils + "/runner"
	// Start command for the test executable
	TestCommand = VolumeMountPathUtils + "/op-test"
	// Start command for the utils init container
	UtilsInitCommand = "/files/utils-init"

	ServiceMonitorKind         = "ServiceMonitor"
	ServiceMonitorGroup        = "monitoring.coreos.com"
	ServiceMonitorVersion      = "v1"
	ServiceMonitorGroupVersion = ServiceMonitorGroup + "/" + ServiceMonitorVersion

	// Port names
	PortNameCoherence  = "coherence"
	PortNameDebug      = "debug-port"
	PortNameHealth     = "health"
	PortNameManagement = "management"
	PortNameMetrics    = "metrics"

	DefaultDebugPort      int32 = 5005
	DefaultManagementPort int32 = 30000
	DefaultMetricsPort    int32 = 9612
	DefaultJmxmpPort      int32 = 9099
	DefaultHealthPort     int32 = 6676

	OperatorConfigName = "coherence-operator-config"

	OperatorConfigKeyHost = "operatorhost"

	DefaultReadinessPath = "/ready"
	DefaultLivenessPath  = "/healthz"

	EnvVarAppType                     = "APP_TYPE"
	EnvVarAppMainClass                = "COH_MAIN_CLASS"
	EnvVarAppMainArgs                 = "COH_MAIN_ARGS"
	EnvVarOperatorHost                = "OPERATOR_HOST"
	EnvVarOperatorTimeout             = "OPERATOR_REQUEST_TIMEOUT"
	EnvVarCoherenceHome               = "COHERENCE_HOME"
	EnvVarCohDependencyModules        = "DEPENDENCY_MODULES"
	EnvVarCohSkipVersionCheck         = "COH_SKIP_VERSION_CHECK"
	EnvVarCohClusterName              = "COH_CLUSTER_NAME"
	EnvVarCohWka                      = "COH_WKA"
	EnvVarCohAppDir                   = "COH_APP_DIR"
	EnvVarCohMachineName              = "COH_MACHINE_NAME"
	EnvVarCohMemberName               = "COH_MEMBER_NAME"
	EnvVarCohPodUID                   = "COH_POD_UID"
	EnvVarCohSkipSite                 = "COH_SKIP_SITE"
	EnvVarCohSite                     = "COH_SITE_INFO_LOCATION"
	EnvVarCohRack                     = "COH_RACK_INFO_LOCATION"
	EnvVarCohRole                     = "COH_ROLE"
	EnvVarCohUtilDir                  = "COH_UTIL_DIR"
	EnvVarCohHealthPort               = "COH_HEALTH_PORT"
	EnvVarCohCacheConfig              = "COH_CACHE_CONFIG"
	EnvVarCohOverride                 = "COH_OVERRIDE_CONFIG"
	EnvVarCohLogLevel                 = "COH_LOG_LEVEL"
	EnvVarCohStorage                  = "COH_STORAGE_ENABLED"
	EnvVarCohPersistenceMode          = "COH_PERSISTENCE_MODE"
	EnvVarCohPersistenceDir           = "COH_PERSISTENCE_DIR"
	EnvVarCohSnapshotDir              = "COH_SNAPSHOT_DIR"
	EnvVarCohTracingRatio             = "COH_TRACING_RATIO"
	EnvVarCohAllowEndangered          = "COH_ALLOW_ENDANGERED"
	EnvVarCohMgmtPrefix               = "COH_MGMT"
	EnvVarCohMetricsPrefix            = "COH_METRICS"
	EnvVarCohEnabledSuffix            = "_ENABLED"
	EnvVarCohPortSuffix               = "_PORT"
	EnvVarSuffixSSLEnabled            = "_SSL_ENABLED"
	EnvVarSuffixSSLCerts              = "_SSL_CERTS"
	EnvVarSuffixSSLKeyStore           = "_SSL_KEYSTORE"
	EnvVarSuffixSSLKeyStoreCredFile   = "_SSL_KEYSTORE_PASSWORD_FILE"
	EnvVarSuffixSSLKeyCredFile        = "_SSL_KEY_PASSWORD_FILE"
	EnvVarSuffixSSLKeyStoreAlgo       = "_SSL_KEYSTORE_ALGORITHM"
	EnvVarSuffixSSLKeyStoreProvider   = "_SSL_KEYSTORE_PROVIDER"
	EnvVarSuffixSSLKeyStoreType       = "_SSL_KEYSTORE_TYPE"
	EnvVarSuffixSSLTrustStore         = "_SSL_TRUSTSTORE"
	EnvVarSuffixSSLTrustStoreCredFile = "_SSL_TRUSTSTORE_PASSWORD_FILE"
	EnvVarSuffixSSLTrustStoreAlgo     = "_SSL_TRUSTSTORE_ALGORITHM"
	EnvVarSuffixSSLTrustStoreProvider = "_SSL_TRUSTSTORE_PROVIDER"
	EnvVarSuffixSSLTrustStoreType     = "_SSL_TRUSTSTORE_TYPE"
	EnvVarSuffixSSLRequireClientCert  = "_SSL_REQUIRE_CLIENT_CERT"
	EnvVarJavaHome                    = "JAVA_HOME"
	EnvVarJavaClasspath               = "CLASSPATH"
	EnvVarJvmClasspathJib             = "JVM_USE_JIB_CLASSPATH"
	EnvVarJvmExtraClasspath           = "JVM_EXTRA_CLASSPATH"
	EnvVarJvmExtraModulepath          = "JVM_EXTRA_MODULEPATH"
	EnvVarJvmArgs                     = "JVM_ARGS"
	EnvVarJvmUseContainerLimits       = "JVM_USE_CONTAINER_LIMITS"
	EnvVarJvmDebugEnabled             = "JVM_DEBUG_ENABLED"
	EnvVarJvmDebugPort                = "JVM_DEBUG_PORT"
	EnvVarJvmDebugSuspended           = "JVM_DEBUG_SUSPEND"
	EnvVarJvmDebugAttach              = "JVM_DEBUG_ATTACH"
	EnvVarJvmGcArgs                   = "JVM_GC_ARGS"
	EnvVarJvmGcCollector              = "JVM_GC_COLLECTOR"
	EnvVarJvmGcLogging                = "JVM_GC_LOGGING"
	EnvVarJvmMemoryHeap               = "JVM_HEAP_SIZE"
	EnvVarJvmMemoryInitialHeap        = "JVM_INITIAL_HEAP_SIZE"
	EnvVarJvmMemoryMaxHeap            = "JVM_MAX_HEAP_SIZE"
	EnvVarJvmMaxRAM                   = "JVM_MAX_RAM"
	EnvVarJvmRAMPercentage            = "JVM_RAM_PERCENTAGE"
	EnvVarJvmInitialRAMPercentage     = "JVM_INITIAL_RAM_PERCENTAGE"
	EnvVarJvmMaxRAMPercentage         = "JVM_MAX_RAM_PERCENTAGE"
	EnvVarJvmMinRAMPercentage         = "JVM_MIN_RAM_PERCENTAGE"
	EnvVarJvmMemoryDirect             = "JVM_DIRECT_MEMORY_SIZE"
	EnvVarJvmMemoryStack              = "JVM_STACK_SIZE"
	EnvVarJvmMemoryMeta               = "JVM_METASPACE_SIZE"
	EnvVarJvmMemoryNativeTracking     = "JVM_NATIVE_MEMORY_TRACKING"
	EnvVarJvmOomExit                  = "JVM_OOM_EXIT"
	EnvVarJvmOomHeapDump              = "JVM_OOM_HEAP_DUMP"
	EnvVarJvmJmxmpEnabled             = "JVM_JMXMP_ENABLED"
	EnvVarJvmJmxmpPort                = "JVM_JMXMP_PORT"
)
