package descriptions

import "fmt"

type mpDescription int32

const (
	mpDescriptionFailed mpDescription = iota + 1
	mpDescriptionIllegalState
	mpDescriptionInvalidParam
	mpDescriptionNoChild
	mpDescriptionNoEntry
	mpDescriptionOverMaxEntry
	mpDescriptionNoData
	mpDescriptionNoDataset
	mpDescriptionNotAllowed
	mpDescriptionAlreadyInUse
	mpDescriptionClosed
	mpDescriptionNotEnoughMemory
	mpDescriptionNotInitialized
	mpDescriptionAbort
	mpDescriptionWirelessOff
	mpDescriptionOperating
	mpDescriptionSendQueueFull
)

var mpDescriptionToString = map[mpDescription]string{
	mpDescriptionFailed:          "Failed",
	mpDescriptionIllegalState:    "IllegalState",
	mpDescriptionInvalidParam:    "InvalidParam",
	mpDescriptionNoChild:         "NoChild",
	mpDescriptionNoEntry:         "NoEntry",
	mpDescriptionOverMaxEntry:    "OverMaxEntry",
	mpDescriptionNoData:          "NoData",
	mpDescriptionNoDataset:       "NoDataset",
	mpDescriptionNotAllowed:      "NotAllowed",
	mpDescriptionAlreadyInUse:    "AlreadyInUse",
	mpDescriptionClosed:          "Closed",
	mpDescriptionNotEnoughMemory: "NotEnoughMemory",
	mpDescriptionNotInitialized:  "NotInitialized",
	mpDescriptionAbort:           "Abort",
	mpDescriptionWirelessOff:     "WirelessOff",
	mpDescriptionOperating:       "Operating",
	mpDescriptionSendQueueFull:   "SendQueueFull",
}

var mpDescriptionDescription = map[mpDescription]string{
	mpDescriptionFailed:          "An internal error (WL command error) has occurred.",
	mpDescriptionIllegalState:    "You called a function that cannot be called in the current state of the MP library.",
	mpDescriptionInvalidParam:    "Parameter is invalid.",
	mpDescriptionNoChild:         "The specified child machine does not exist. Only used on the parent machine.",
	mpDescriptionNoEntry:         "Cannot connect because no entries have been accepted. Only used on child devices.",
	mpDescriptionOverMaxEntry:    "Cannot connect because the maximum number of devices is being connected. Only used on child devices.",
	mpDescriptionNoData:          "There was no data to process. It happens in the Receive() function.",
	mpDescriptionNoDataset:       "There was no data sharing data to process. Occurs in the DSTryStep() function.",
	mpDescriptionNotAllowed:      "Not allowed. Returns when Start() is used for a channel that is not allowed to be used.",
	mpDescriptionAlreadyInUse:    "The specified channel is already in use.",
	mpDescriptionClosed:          "Failed because termination processing has already been performed.",
	mpDescriptionNotEnoughMemory: "Could not allocate required memory.",
	mpDescriptionNotInitialized:  "The MP library has not been initialized.",
	mpDescriptionAbort:           "Processing interrupted. Returns when calling AbortReceive() or AbortGetIndication() while Receive() or GetIndication() is running.",
	mpDescriptionWirelessOff:     "Failed because the wireless function is off.",
	mpDescriptionOperating:       "Running. (Driver internal error)",
	mpDescriptionSendQueueFull:   "Send queue is full. (Driver internal error)",
}

func (d mpDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", mpDescriptionToString[d], d, mpDescriptionDescription[d])
}
