package api

import (
	"time"
)

// VolumeID driver specific system wide unique volume identifier.
type VolumeID string

// BadVolumeID invalid volume ID, usually accompanied by an error.
const BadVolumeID = VolumeID("")

// SnapID driver specific system wide unique snap identifier.
type SnapID string

// BadSnapID invalid snap ID, usually accompanied by an error.
const BadSnapID = SnapID("")

// VolumeCos a number representing class of servcie.
type VolumeCos int

const (
	// VolumeCosNone minmum level of CoS
	VolumeCosNone = VolumeCos(0)
	// VolumeCosMedium in-between level of Cos
	VolumeCosMedium = VolumeCos(5)
	// VolumeCosMax maximum level of CoS
	VolumeCosMax = VolumeCos(9)
)

// VolumeStatus a health status.
type VolumeStatus string

const (
	// NotPresent This volume is not present.
	NotPresent = VolumeStatus("NotPresent")
	// Up status healthy
	Up = VolumeStatus("Up")
	// Down status failure.
	Down = VolumeStatus("Down")
	// Degraded status up but with degraded performance. In a RAID group, this may indicate a problem with one or more drives
	Degraded = VolumeStatus("Degraded")
)

// VolumeState is one of the below enumerations and reflects the state
// of a volume.
type VolumeState int

const (
	// VolumePending volume is transitioning to new state
	VolumePending VolumeState = 1 << iota
	// VolumeAvailable volume is ready to be assigned to a container
	VolumeAvailable
	// VolumeAttached is attached to container
	VolumeAttached
	// VolumeDetached is detached but associated with a container.
	VolumeDetached
	// VolumeError is in Error State
	VolumeError
	// VolumeDeleted is deleted, it will remain in this state while resources are
	// asynchronously reclaimed.
	VolumeDeleted
)

// VolumeStateAny a filter that selects all volumes
const VolumeStateAny = VolumePending | VolumeAvailable | VolumeAttached | VolumeDetached | VolumeError | VolumeDeleted

// Labels a name-value map
type Labels map[string]string

// VolumeLocator is a structure that is attached to a volume and is used to
// carry opaque metadata.
type VolumeLocator struct {
	// Name user friendly identifier
	Name string
	// VolumeLabels set of name-value pairs that acts as search filters.
	VolumeLabels Labels
}

// CreateOptions are passed in with a CreateRequest
type CreateOptions struct {
	// FailIfExists fail create request if a volume with matching Locator already exists.
	FailIfExists bool
	// CreateFromSnap will create a volume with specified SnapID
	CreateFromSnap SnapID
}

// Filesystem supported filesystems
type Filesystem string

const (
	FsNone Filesystem = "none"
	FsExt4 Filesystem = "ext4"
	FsXfs  Filesystem = "xfs"
	FsZfs  Filesystem = "zfs"
	FsNfs  Filesystem = "nfs"
)

// VolumeSpec has the properties needed to create a volume.
type VolumeSpec struct {
	// Ephemeral storage
	Ephemeral bool
	// Thin provisioned volume size in bytes
	Size uint64
	// Format disk with this FileSystem
	Format Filesystem
	// BlockSize for file system
	BlockSize int
	// HA Level specifies the number of nodes that are
	// allowed to fail, and yet data is availabel.
	// A value of 0 implies that data is not erasure coded,
	// a failure of a node will lead to data loss.
	HALevel int
	// This disk's CoS
	Cos VolumeCos
	// Perform dedupe on this disk
	Dedupe bool
	// SnapshotInterval in minutes, set to 0 to disable Snapshots
	SnapshotInterval int
	// Volume configuration labels
	ConfigLabels Labels
}

// MachineID is a node instance identifier for clustered systems.
type MachineID string

const MachineNone MachineID = ""

// Volume represents a live, created volume.
type Volume struct {
	// ID Self referential VolumeID
	ID VolumeID
	// Locator User specified locator
	Locator VolumeLocator
	// Ctime Volume creation time
	Ctime time.Time
	// Spec User specified VolumeSpec
	Spec *VolumeSpec
	// Usage Volume usage
	Usage uint64
	// LastScan time when an integrity check for run
	LastScan time.Time
	// Format Filesystem type if any
	Format Filesystem
	// Status see VolumeStatus
	Status VolumeStatus
	// State see VolumeState
	State VolumeState
	// AttachedOn - Node on which this volume is attached.
	AttachedOn MachineID
	// DevicePath
	DevicePath string
	// AttachPath
	AttachPath string
	// ReplicaSet Set of nodes no which this Volume is erasure coded - for clustered storage arrays
	ReplicaSet []MachineID
	// Error Last recorded error
	Error string
}

// VolumeSnap identifies a volume snapshot.
type VolumeSnap struct {
	// SnapID system generated ID
	ID SnapID
	// VolumeID Volume identifier.
	VolumeID VolumeID
	// Ctime Snap creation time.
	Ctime time.Time
	// SnapLabel arbitrary name value pairs
	SnapLabels Labels
	// Usage
	Usage uint64
}

// VolumeStats
type VolumeStats struct {
}

// VolumeAlerts
type VolumeAlerts struct {
}
