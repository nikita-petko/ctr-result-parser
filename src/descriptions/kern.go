package descriptions

import "fmt"

type kernDescription int32

const (
	kernDescriptionNoSuchThread kernDescription = iota + 512
	kernDescriptionSessionMustBeLocked
	kernDescriptionNotBound
	kernDescriptionNonManualEvent
	kernDescriptionNonActive
	kernDescriptionObjectNotFound
	kernDescriptionStopProcessingException
	kernDescriptionInvalidParm
	kernDescriptionProcessTerminated
	kernDescriptionOutOfAddressSpace
	kernDescriptionNoSyncedObject
	kernDescriptionNoDebugEvent
	kernDescriptionOverrideState
	kernDescriptionTerminateRequested
	kernDescriptionInconsistencyMemoryBlockRange
)

var kernDescriptionToString = map[kernDescription]string{
	kernDescriptionNoSuchThread:                  "NoSuchThread",
	kernDescriptionSessionMustBeLocked:           "SessionMustBeLocked",
	kernDescriptionNotBound:                      "NotBound",
	kernDescriptionNonManualEvent:                "NonManualEvent",
	kernDescriptionNonActive:                     "NonActive",
	kernDescriptionObjectNotFound:                "ObjectNotFound",
	kernDescriptionStopProcessingException:       "StopProcessingException",
	kernDescriptionInvalidParm:                   "InvalidParm",
	kernDescriptionProcessTerminated:             "ProcessTerminated",
	kernDescriptionOutOfAddressSpace:             "OutOfAddressSpace",
	kernDescriptionNoSyncedObject:                "NoSyncedObject",
	kernDescriptionNoDebugEvent:                  "NoDebugEvent",
	kernDescriptionOverrideState:                 "OverrideState",
	kernDescriptionTerminateRequested:            "TerminateRequested",
	kernDescriptionInconsistencyMemoryBlockRange: "InconsistencyMemoryBlockRange",
}

var kernDescriptionDescription = map[kernDescription]string{
	kernDescriptionNoSuchThread:                  "No such thread",
	kernDescriptionSessionMustBeLocked:           "Called without locking",
	kernDescriptionNotBound:                      "Specified value out of range",
	kernDescriptionNonManualEvent:                "Improper Clear call",
	kernDescriptionNonActive:                     "Improper Clear call",
	kernDescriptionObjectNotFound:                "Operations on Objects not contained in containers",
	kernDescriptionStopProcessingException:       "Stop processing exception",
	kernDescriptionInvalidParm:                   "Invalid parameter",
	kernDescriptionProcessTerminated:             "The specified process has already terminated",
	kernDescriptionOutOfAddressSpace:             "Insufficient contiguous free space in address space",
	kernDescriptionNoSyncedObject:                "Synchronization object not set",
	kernDescriptionNoDebugEvent:                  "No debug events",
	kernDescriptionOverrideState:                 "Page table state overwritten or taken by others",
	kernDescriptionTerminateRequested:            "Thread termination requested",
	kernDescriptionInconsistencyMemoryBlockRange: "Memory block range is inconsistent",
}

func (d kernDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", kernDescriptionToString[d], d, kernDescriptionDescription[d])
}
