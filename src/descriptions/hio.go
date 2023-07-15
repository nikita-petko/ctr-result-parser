package descriptions

import "fmt"

type hioDescription int32

const (
	hioDescriptionInvalidSelection hioDescription = iota
	hioDescriptionTooLarge
	hioDescriptionNotAuthorized
	hioDescriptionAlreadyDone
	hioDescriptionInvalidSize
	hioDescriptionInvalidEnumValue
	hioDescriptionInvalidCombination
	hioDescriptionNoData
	hioDescriptionBusy
	hioDescriptionMisalignedAddress
	hioDescriptionMisalignedSize
	hioDescriptionOutOfMemory
	hioDescriptionNotImplemented
	hioDescriptionInvalidAddress
	hioDescriptionInvalidPointer
	hioDescriptionInvalidHandle
	hioDescriptionNotInitialized
	hioDescriptionAlreadyInitialized
	hioDescriptionNotFound
	hioDescriptionCancelRequested
	hioDescriptionAlreadyExists
	hioDescriptionOutOfRange
	hioDescriptionTimeout
	hioDescriptionInvalidResultValue
)

var hioDescriptionToString = map[hioDescription]string{
	hioDescriptionInvalidSelection:   "InvalidSelection",
	hioDescriptionTooLarge:           "TooLarge",
	hioDescriptionNotAuthorized:      "NotAuthorized",
	hioDescriptionAlreadyDone:        "AlreadyDone",
	hioDescriptionInvalidSize:        "InvalidSize",
	hioDescriptionInvalidEnumValue:   "InvalidEnumValue",
	hioDescriptionInvalidCombination: "InvalidCombination",
	hioDescriptionNoData:             "NoData",
	hioDescriptionBusy:               "Busy",
	hioDescriptionMisalignedAddress:  "MisalignedAddress",
	hioDescriptionMisalignedSize:     "MisalignedSize",
	hioDescriptionOutOfMemory:        "OutOfMemory",
	hioDescriptionNotImplemented:     "NotImplemented",
	hioDescriptionInvalidAddress:     "InvalidAddress",
	hioDescriptionInvalidPointer:     "InvalidPointer",
	hioDescriptionInvalidHandle:      "InvalidHandle",
	hioDescriptionNotInitialized:     "NotInitialized",
	hioDescriptionAlreadyInitialized: "AlreadyInitialized",
	hioDescriptionNotFound:           "NotFound",
	hioDescriptionCancelRequested:    "CancelRequested",
	hioDescriptionAlreadyExists:      "AlreadyExists",
	hioDescriptionOutOfRange:         "OutOfRange",
	hioDescriptionTimeout:            "Timeout",
	hioDescriptionInvalidResultValue: "InvalidResultValue",
}

var hioDescriptionDescription = map[hioDescription]string{
	hioDescriptionInvalidSelection:   "An invalid selection was specified.",
	hioDescriptionTooLarge:           "The size is too large.",
	hioDescriptionNotAuthorized:      "Not connected.",
	hioDescriptionAlreadyDone:        "Already connected.",
	hioDescriptionInvalidSize:        "The size specification is invalid.",
	hioDescriptionInvalidEnumValue:   "Invalid member value.",
	hioDescriptionInvalidCombination: "Invalid combination of parameters.",
	hioDescriptionNoData:             "There is no data.",
	hioDescriptionBusy:               "WAIT state.",
	hioDescriptionMisalignedAddress:  "Invalid address alignment.",
	hioDescriptionMisalignedSize:     "Invalid size alignment.",
	hioDescriptionOutOfMemory:        "Out of memory.",
	hioDescriptionNotImplemented:     "Not implemented.",
	hioDescriptionInvalidAddress:     "Invalid address.",
	hioDescriptionInvalidPointer:     "Invalid pointer.",
	hioDescriptionInvalidHandle:      "Invalid handle.",
	hioDescriptionNotInitialized:     "Not initialized.",
	hioDescriptionAlreadyInitialized: "Already initialized.",
	hioDescriptionNotFound:           "Handle does not exist.",
	hioDescriptionCancelRequested:    "The request was canceled.",
	hioDescriptionAlreadyExists:      "Already exists.",
	hioDescriptionOutOfRange:         "Out of range.",
	hioDescriptionTimeout:            "Timeout.",
	hioDescriptionInvalidResultValue: "Invalid result value.",
}

func (d hioDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", hioDescriptionToString[d], d, hioDescriptionDescription[d])
}
