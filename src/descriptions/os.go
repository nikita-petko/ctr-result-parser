package descriptions

import (
	"fmt"

	"github.com/nikita-petko/ctr-result-parser/nn"
)

type osDescription int32

const (
	osDescriptionFailedToAllocateMemory osDescription = iota + 1
	osDescriptionFailedToAllocateSharedMemory
	osDescriptionFailedToAllocateThread
	osDescriptionFailedToAllocateMutex
	osDescriptionFailedToAllocateSemaphore
	osDescriptionFailedToAllocateEvent
	osDescriptionFailedToAllocateTimer
	osDescriptionFailedToAllocatePort
	osDescriptionFailedToAllocateSession
	osDescriptionExceedMemoryLimit
	osDescriptionExceedSharedMemoryLimit
	osDescriptionExceedThreadLimit
	osDescriptionExceedMutexLimit
	osDescriptionExceedSemaphoreLimit
	osDescriptionExceedEventLimit
	osDescriptionExceedTimerLimit
	osDescriptionExceedPortLimit
	osDescriptionExceedSessionLimit
	osDescriptionMaxHandle
	osDescriptionInnaccessiblePage
	osDescriptionAbandoned
	osDescriptionAbandon1
	osDescriptionAbandon2
	osDescriptionInvalidProcessId
	osDescriptionInvalidThreadId
	osDescriptionSessionClosed
	osDescriptionInvalidMessage = iota + 2
	osDescriptionManualResetEventRequired
	osDescriptionTooLongName
	osDescriptionNotOwned
	osDescriptionProcessTerminated
	osDescriptionInvalidTlsIndex
	osDescriptionNoRunnableProcessor
	osDescriptionNoSession
	osDescriptionUsingRegion
	osDescriptionAlreadyReceived
	osDescriptionCancelRequested
	osDescriptionNotReceived
	osDescriptionAbandoned3
	osDescriptionDeliverArgNotReady
	osDescriptionDeliverArgOverSize
	osDescriptionInvalidDeliverArg
	osDescriptionIAmOwner
	osDescriptionExceedsSharedLimit
	osDescriptionUnexpectedPermission
	osDescriptionInvalidTag
	osDescriptionInvalidFormat
	osDescriptionOtherHandle
	osDescriptionFailedToAllocateAddressArbiter
	osDescriptionExceedAddressArbiterLimit
	osDescriptionOverPortCapacity
	osDescriptionNotMapped
	osDescriptionBufferTooFlagmented
	osDescriptionNoAddressSpace
	osDescriptionExceedTlsLimit
	osDescriptionCantStart
	osDescriptionLocked
	osDescriptionNotFinalized
	osDescriptionObsoleteResult = osDescription(nn.MAX_DESCRIPTION)
)

var osDescriptionToString = map[osDescription]string{
	osDescriptionFailedToAllocateMemory:         "FailedToAllocateMemory",
	osDescriptionFailedToAllocateSharedMemory:   "FailedToAllocateSharedMemory",
	osDescriptionFailedToAllocateThread:         "FailedToAllocateThread",
	osDescriptionFailedToAllocateMutex:          "FailedToAllocateMutex",
	osDescriptionFailedToAllocateSemaphore:      "FailedToAllocateSemaphore",
	osDescriptionFailedToAllocateEvent:          "FailedToAllocateEvent",
	osDescriptionFailedToAllocateTimer:          "FailedToAllocateTimer",
	osDescriptionFailedToAllocatePort:           "FailedToAllocatePort",
	osDescriptionFailedToAllocateSession:        "FailedToAllocateSession",
	osDescriptionExceedMemoryLimit:              "ExceedMemoryLimit",
	osDescriptionExceedSharedMemoryLimit:        "ExceedSharedMemoryLimit",
	osDescriptionExceedThreadLimit:              "ExceedThreadLimit",
	osDescriptionExceedMutexLimit:               "ExceedMutexLimit",
	osDescriptionExceedSemaphoreLimit:           "ExceedSemaphoreLimit",
	osDescriptionExceedEventLimit:               "ExceedEventLimit",
	osDescriptionExceedTimerLimit:               "ExceedTimerLimit",
	osDescriptionExceedPortLimit:                "ExceedPortLimit",
	osDescriptionExceedSessionLimit:             "ExceedSessionLimit",
	osDescriptionMaxHandle:                      "MaxHandle",
	osDescriptionInnaccessiblePage:              "InnaccessiblePage",
	osDescriptionAbandoned:                      "Abandoned",
	osDescriptionAbandon1:                       "Abandon1",
	osDescriptionAbandon2:                       "Abandon2",
	osDescriptionInvalidProcessId:               "InvalidProcessId",
	osDescriptionInvalidThreadId:                "InvalidThreadId",
	osDescriptionSessionClosed:                  "SessionClosed",
	osDescriptionInvalidMessage:                 "InvalidMessage",
	osDescriptionManualResetEventRequired:       "ManualResetEventRequired",
	osDescriptionTooLongName:                    "TooLongName",
	osDescriptionNotOwned:                       "NotOwned",
	osDescriptionProcessTerminated:              "ProcessTerminated",
	osDescriptionInvalidTlsIndex:                "InvalidTlsIndex",
	osDescriptionNoRunnableProcessor:            "NoRunnableProcessor",
	osDescriptionNoSession:                      "NoSession",
	osDescriptionUsingRegion:                    "UsingRegion",
	osDescriptionAlreadyReceived:                "AlreadyReceived",
	osDescriptionCancelRequested:                "CancelRequested",
	osDescriptionNotReceived:                    "NotReceived",
	osDescriptionAbandoned3:                     "Abandoned3",
	osDescriptionDeliverArgNotReady:             "DeliverArgNotReady",
	osDescriptionDeliverArgOverSize:             "DeliverArgOverSize",
	osDescriptionInvalidDeliverArg:              "InvalidDeliverArg",
	osDescriptionIAmOwner:                       "IAmOwner",
	osDescriptionExceedsSharedLimit:             "ExceedsSharedLimit",
	osDescriptionUnexpectedPermission:           "UnexpectedPermission",
	osDescriptionInvalidTag:                     "InvalidTag",
	osDescriptionInvalidFormat:                  "InvalidFormat",
	osDescriptionOtherHandle:                    "OtherHandle",
	osDescriptionFailedToAllocateAddressArbiter: "FailedToAllocateAddressArbiter",
	osDescriptionExceedAddressArbiterLimit:      "ExceedAddressArbiterLimit",
	osDescriptionOverPortCapacity:               "OverPortCapacity",
	osDescriptionNotMapped:                      "NotMapped",
	osDescriptionBufferTooFlagmented:            "BufferTooFlagmented",
	osDescriptionNoAddressSpace:                 "NoAddressSpace",
	osDescriptionExceedTlsLimit:                 "ExceedTlsLimit",
	osDescriptionCantStart:                      "CantStart",
	osDescriptionLocked:                         "Locked",
	osDescriptionNotFinalized:                   "NotFinalized",
	osDescriptionObsoleteResult:                 "ObsoleteResult",
}

var osDescriptionDescription = map[osDescription]string{
	osDescriptionFailedToAllocateMemory:         "Reached the physical memory limit",
	osDescriptionFailedToAllocateSharedMemory:   "Reached the shared memory limit",
	osDescriptionFailedToAllocateThread:         "Reached the thread limit",
	osDescriptionFailedToAllocateMutex:          "Reached the mutex limit",
	osDescriptionFailedToAllocateSemaphore:      "Reached the semaphore limit",
	osDescriptionFailedToAllocateEvent:          "Reached the event limit",
	osDescriptionFailedToAllocateTimer:          "Reached the timer limit",
	osDescriptionFailedToAllocatePort:           "Reached the port limit",
	osDescriptionFailedToAllocateSession:        "Reached the session limit",
	osDescriptionExceedMemoryLimit:              "Reached the physical memory allocation limit",
	osDescriptionExceedSharedMemoryLimit:        "Reached the shared memory allocation limit",
	osDescriptionExceedThreadLimit:              "Reached the thread allocation limit",
	osDescriptionExceedMutexLimit:               "Reached the mutex allocation limit",
	osDescriptionExceedSemaphoreLimit:           "Reached the semaphore allocation limit",
	osDescriptionExceedEventLimit:               "Reached the event allocation limit",
	osDescriptionExceedTimerLimit:               "Reached the timer allocation limit",
	osDescriptionExceedPortLimit:                "Reached the port allocation limit",
	osDescriptionExceedSessionLimit:             "Reached the session allocation limit",
	osDescriptionMaxHandle:                      "Reached the maximum number of handles",
	osDescriptionInnaccessiblePage:              "An inaccessible page is included",
	osDescriptionAbandoned:                      "An object was abandoned",
	osDescriptionAbandon1:                       "No longer used",
	osDescriptionAbandon2:                       "No longer used",
	osDescriptionInvalidProcessId:               "Invalid process ID",
	osDescriptionInvalidThreadId:                "Invalid thread ID",
	osDescriptionSessionClosed:                  "The session was closed",
	osDescriptionInvalidMessage:                 "Invalid message",
	osDescriptionManualResetEventRequired:       "A manual reset event is required",
	osDescriptionTooLongName:                    "The name is too long",
	osDescriptionNotOwned:                       "The mutex is not owned",
	osDescriptionProcessTerminated:              "The specified process already terminated",
	osDescriptionInvalidTlsIndex:                "An unallocated TLS index was specified",
	osDescriptionNoRunnableProcessor:            "Affinity mask prohibits running on all processors",
	osDescriptionNoSession:                      "Not a newly connected session",
	osDescriptionUsingRegion:                    "Region in use",
	osDescriptionAlreadyReceived:                "Already received",
	osDescriptionCancelRequested:                "Cancel requested",
	osDescriptionNotReceived:                    "Not received",
	osDescriptionAbandoned3:                     "No longer used",
	osDescriptionDeliverArgNotReady:             "DeliverArg not ready",
	osDescriptionDeliverArgOverSize:             "DeliverArg over size",
	osDescriptionInvalidDeliverArg:              "Invalid DeliverArg",
	osDescriptionIAmOwner:                       "I am owner",
	osDescriptionExceedsSharedLimit:             "Exceeds shared limit",
	osDescriptionUnexpectedPermission:           "Unexpected permission",
	osDescriptionInvalidTag:                     "Invalid IPC message tag",
	osDescriptionInvalidFormat:                  "Invalid format",
	osDescriptionOtherHandle:                    "Closed by other handle",
	osDescriptionFailedToAllocateAddressArbiter: "Failed to allocate address arbiter",
	osDescriptionExceedAddressArbiterLimit:      "Reached the address arbiter limit",
	osDescriptionOverPortCapacity:               "Over IPC port capacity",
	osDescriptionNotMapped:                      "Not mapped",
	osDescriptionBufferTooFlagmented:            "Buffer too fragmented",
	osDescriptionNoAddressSpace:                 "No address space",
	osDescriptionExceedTlsLimit:                 "Reached the TLS limit",
	osDescriptionCantStart:                      "Canceled",
	osDescriptionLocked:                         "Locked",
	osDescriptionNotFinalized:                   "Not finalized",
	osDescriptionObsoleteResult:                 "Obsolete; this error must be changed to another error as soon as possible.",
}

func (d osDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", osDescriptionToString[d], d, osDescriptionDescription[d])
}
