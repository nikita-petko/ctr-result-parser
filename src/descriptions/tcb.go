package descriptions

import "fmt"

type tcbDescription int32

const (
	tcbDescriptionUsingRegion tcbDescription = iota + 1
	tcbDescriptionFailedToAllocateCodeset
	tcbDescriptionFailedToAllocateProcess
	tcbDescriptionFailedToAllocateResourceLimit
	tcbDescriptionExceedCodesetLimit
	tcbDescriptionExceedProcessLimit
	tcbDescriptionExceedResourceLimit
	tcbDescriptionExceedMemoryLimit
	tcbDescriptionExceedThreadLimit
	tcbDescriptionFailedToAllocateMemory
	tcbDescriptionInaccessiblePage
	tcbDescriptionMaxHandle
	tcbDescriptionFailedToAllocateThread
	tcbDescriptionInterruptNumberAlreadyPermitted
	tcbDescriptionSvcNumberAlreadyPermitted
	tcbDescriptionStaticPageAlreadyMapped
	tcbDescriptionIoPageAlreadyMapped
	tcbDescriptionWrongCapabilityFlag
	tcbDescriptionRequireNewSystem
	tcbDescriptionApplicationAssigned
)

var tcbDescriptionToString = map[tcbDescription]string{
	tcbDescriptionUsingRegion:                     "UsingRegion",
	tcbDescriptionFailedToAllocateCodeset:         "FailedToAllocateCodeset",
	tcbDescriptionFailedToAllocateProcess:         "FailedToAllocateProcess",
	tcbDescriptionFailedToAllocateResourceLimit:   "FailedToAllocateResourceLimit",
	tcbDescriptionExceedCodesetLimit:              "ExceedCodesetLimit",
	tcbDescriptionExceedProcessLimit:              "ExceedProcessLimit",
	tcbDescriptionExceedResourceLimit:             "ExceedResourceLimit",
	tcbDescriptionExceedMemoryLimit:               "ExceedMemoryLimit",
	tcbDescriptionExceedThreadLimit:               "ExceedThreadLimit",
	tcbDescriptionFailedToAllocateMemory:          "FailedToAllocateMemory",
	tcbDescriptionInaccessiblePage:                "InaccessiblePage",
	tcbDescriptionMaxHandle:                       "MaxHandle",
	tcbDescriptionFailedToAllocateThread:          "FailedToAllocateThread",
	tcbDescriptionInterruptNumberAlreadyPermitted: "InterruptNumberAlreadyPermitted",
	tcbDescriptionSvcNumberAlreadyPermitted:       "SvcNumberAlreadyPermitted",
	tcbDescriptionStaticPageAlreadyMapped:         "StaticPageAlreadyMapped",
	tcbDescriptionIoPageAlreadyMapped:             "IoPageAlreadyMapped",
	tcbDescriptionWrongCapabilityFlag:             "WrongCapabilityFlag",
	tcbDescriptionRequireNewSystem:                "RequireNewSystem",
	tcbDescriptionApplicationAssigned:             "ApplicationAssigned",
}

var tcbDescriptionDescription = map[tcbDescription]string{
	tcbDescriptionUsingRegion:                     "UsingRegion",
	tcbDescriptionFailedToAllocateCodeset:         "Failed to allocate codeset",
	tcbDescriptionFailedToAllocateProcess:         "Failed to allocate process",
	tcbDescriptionFailedToAllocateResourceLimit:   "Failed to allocate resource limit",
	tcbDescriptionExceedCodesetLimit:              "Codeset quota limit reached",
	tcbDescriptionExceedProcessLimit:              "Process quota reached",
	tcbDescriptionExceedResourceLimit:             "Resource limit quota reached",
	tcbDescriptionExceedMemoryLimit:               "Physical memory allocation limit reached",
	tcbDescriptionExceedThreadLimit:               "Thread allocation limit reached",
	tcbDescriptionFailedToAllocateMemory:          "Failed to allocate physical memory",
	tcbDescriptionInaccessiblePage:                "Contains inaccessible pages",
	tcbDescriptionMaxHandle:                       "Maximum number of handles reached",
	tcbDescriptionFailedToAllocateThread:          "Failed to allocate thread",
	tcbDescriptionInterruptNumberAlreadyPermitted: "Interrupt number accepted",
	tcbDescriptionSvcNumberAlreadyPermitted:       "SVC number accepted",
	tcbDescriptionStaticPageAlreadyMapped:         "Static page is mapped",
	tcbDescriptionIoPageAlreadyMapped:             "I/O page is already mapped",
	tcbDescriptionWrongCapabilityFlag:             "Wrong capability flag",
	tcbDescriptionRequireNewSystem:                "Require new system",
	tcbDescriptionApplicationAssigned:             "Application assigned",
}

func (d tcbDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", tcbDescriptionToString[d], d, tcbDescriptionDescription[d])
}
