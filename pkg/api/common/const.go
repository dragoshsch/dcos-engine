package common

// the orchestrators supported
const (
	// DCOS is the string constant for DCOS orchestrator type and defaults to DCOS188
	DCOS string = "DCOS"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents per agent pool
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents per agent pool
	MaxAgentCount = 100
	// MinPort specifies the minimum tcp port to open
	MinPort = 1
	// MaxPort specifies the maximum tcp port to open
	MaxPort = 65535
	// MaxDisks specifies the maximum attached disks to add to the cluster
	MaxDisks = 4
	// MinDiskSizeGB specifies the minimum attached disk size
	MinDiskSizeGB = 1
	// MaxDiskSizeGB specifies the maximum attached disk size
	MaxDiskSizeGB = 1023
	// MinIPAddressCount specifies the minimum number of IP addresses per network interface
	MinIPAddressCount = 1
	// MaxIPAddressCount specifies the maximum number of IP addresses per network interface
	MaxIPAddressCount = 256
)

// Availability profiles
const (
	// AvailabilitySet means that the vms are in an availability set
	AvailabilitySet = "AvailabilitySet"
	// VirtualMachineScaleSets means that the vms are in a virtual machine scaleset
	VirtualMachineScaleSets = "VirtualMachineScaleSets"
)

// storage profiles
const (
	// StorageAccount means that the nodes use raw storage accounts for their os and attached volumes
	StorageAccount = "StorageAccount"
	// ManagedDisks means that the nodes use managed disks for their os and attached volumes
	ManagedDisks = "ManagedDisks"
)

const (
	// DCOSVersion1Dot12Dot0 is the major.minor.patch string for 1.12.0 versions of DCOS
	DCOSVersion1Dot12Dot0 string = "1.12.0"
	// DCOSDefaultVersion is a default DCOS version
	DCOSDefaultVersion string = DCOSVersion1Dot12Dot0
)

// AllDCOSSupportedVersions maintain a list of available dcos versions in acs-engine
var AllDCOSSupportedVersions = []string{
	DCOSVersion1Dot12Dot0,
}

// GetAllSupportedDCOSVersions returns a slice of all supported DCOS versions.
func GetAllSupportedDCOSVersions() []string {
	return AllDCOSSupportedVersions
}
