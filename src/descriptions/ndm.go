package descriptions

import "fmt"

type ndmDescription int32

const (
	ndmDescriptionInterruptByRequest ndmDescription = iota + 1
	ndmDescriptionProcessingPriorityRequest
	ndmDescriptionInErrorState
	ndmDescriptionDisconnected
	ndmDescriptionCancelledByOtherRequest
	ndmDescriptionCancelledByHardwareEvents
	ndmDescriptionCancelledByDisconnect
	ndmDescriptionCancelledByUserRequest
	ndmDescriptionOperationDenied
	ndmDescriptionLockedByOtherProcess
	ndmDescriptionNotLocked
	ndmDescriptionAlreadySuspended
	ndmDescriptionNotSuspended
	ndmDescriptionInvalidOperation
	ndmDescriptionNotExclusive
	ndmDescriptionExclusiveByOtherProcess
	ndmDescriptionExclusiveByOwnProcess
	ndmDescriptionBackgroundDisconnected
	ndmDescriptionForegroundConnectionExists
)

var ndmDescriptionToString = map[ndmDescription]string{
	ndmDescriptionInterruptByRequest:         "InterruptByRequest",
	ndmDescriptionProcessingPriorityRequest:  "ProcessingPriorityRequest",
	ndmDescriptionInErrorState:               "InErrorState",
	ndmDescriptionDisconnected:               "Disconnected",
	ndmDescriptionCancelledByOtherRequest:    "CancelledByOtherRequest",
	ndmDescriptionCancelledByHardwareEvents:  "CancelledByHardwareEvents",
	ndmDescriptionCancelledByDisconnect:      "CancelledByDisconnect",
	ndmDescriptionCancelledByUserRequest:     "CancelledByUserRequest",
	ndmDescriptionOperationDenied:            "OperationDenied",
	ndmDescriptionLockedByOtherProcess:       "LockedByOtherProcess",
	ndmDescriptionNotLocked:                  "NotLocked",
	ndmDescriptionAlreadySuspended:           "AlreadySuspended",
	ndmDescriptionNotSuspended:               "NotSuspended",
	ndmDescriptionInvalidOperation:           "InvalidOperation",
	ndmDescriptionNotExclusive:               "NotExclusive",
	ndmDescriptionExclusiveByOtherProcess:    "ExclusiveByOtherProcess",
	ndmDescriptionExclusiveByOwnProcess:      "ExclusiveByOwnProcess",
	ndmDescriptionBackgroundDisconnected:     "BackgroundDisconnected",
	ndmDescriptionForegroundConnectionExists: "ForegroundConnectionExists",
}

var ndmDescriptionDescription = map[ndmDescription]string{
	ndmDescriptionInterruptByRequest:         "Interrupted by request",
	ndmDescriptionProcessingPriorityRequest:  "Already processing priority request",
	ndmDescriptionInErrorState:               "In error state",
	ndmDescriptionDisconnected:               "The network daemon manager is disconnected",
	ndmDescriptionCancelledByOtherRequest:    "Cancelled by other request",
	ndmDescriptionCancelledByHardwareEvents:  "Cancelled by hardware events",
	ndmDescriptionCancelledByDisconnect:      "Cancelled by disconnect",
	ndmDescriptionCancelledByUserRequest:     "Cancelled by user request",
	ndmDescriptionOperationDenied:            "Operation denied",
	ndmDescriptionLockedByOtherProcess:       "Locked by other process",
	ndmDescriptionNotLocked:                  "Not locked",
	ndmDescriptionAlreadySuspended:           "Already suspended",
	ndmDescriptionNotSuspended:               "Not suspended",
	ndmDescriptionInvalidOperation:           "Invalid operation",
	ndmDescriptionNotExclusive:               "Not exclusive",
	ndmDescriptionExclusiveByOtherProcess:    "Exclusive by other process",
	ndmDescriptionExclusiveByOwnProcess:      "Exclusive by own process",
	ndmDescriptionBackgroundDisconnected:     "Background disconnected",
	ndmDescriptionForegroundConnectionExists: "Foreground connection exists",
}

func (d ndmDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", ndmDescriptionToString[d], d, ndmDescriptionDescription[d])
}
