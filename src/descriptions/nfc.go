package descriptions

import "fmt"

type nfcDescription int32

const (
	nfcDescriptionInvalidOperation nfcDescription = iota + 1
	nfcDescriptionInvalidArgument
	nfcDescriptionTimeoutError
	nfcDescriptionConnectionError
	nfcDescriptionUnknownError
	nfcDescriptionNotSupported
	nfcDescriptionBufferSmall
	nfcDescriptionUnknownFormat
	nfcDescriptionNoResources
	nfcDescriptionIsBusy
	nfcDescriptionNfcFunctionError
	nfcDescriptionTagNotFound
	nfcDescriptionNeedRetry
	nfcDescriptionNeedFormat
	nfcDescriptionInternalError
	nfcDescriptionHardwareError
)

var nfcDescriptionToString = map[nfcDescription]string{
	nfcDescriptionInvalidOperation: "InvalidOperation",
	nfcDescriptionInvalidArgument:  "InvalidArgument",
	nfcDescriptionTimeoutError:     "TimeoutError",
	nfcDescriptionConnectionError:  "ConnectionError",
	nfcDescriptionUnknownError:     "UnknownError",
	nfcDescriptionNotSupported:     "NotSupported",
	nfcDescriptionBufferSmall:      "BufferSmall",
	nfcDescriptionUnknownFormat:    "UnknownFormat",
	nfcDescriptionNoResources:      "NoResources",
	nfcDescriptionIsBusy:           "IsBusy",
	nfcDescriptionNfcFunctionError: "NfcFunctionError",
	nfcDescriptionTagNotFound:      "TagNotFound",
	nfcDescriptionNeedRetry:        "NeedRetry",
	nfcDescriptionNeedFormat:       "NeedFormat",
	nfcDescriptionInternalError:    "InternalError",
	nfcDescriptionHardwareError:    "HardwareError",
}

var nfcDescriptionDescription = map[nfcDescription]string{
	nfcDescriptionInvalidOperation: "Invalid operation was requested.",
	nfcDescriptionInvalidArgument:  "Invalid argument was passed.",
	nfcDescriptionTimeoutError:     "Timeout error occurred.",
	nfcDescriptionConnectionError:  "Connection error occurred.",
	nfcDescriptionUnknownError:     "Unknown error occurred.",
	nfcDescriptionNotSupported:     "Not supported.",
	nfcDescriptionBufferSmall:      "Buffer is too small.",
	nfcDescriptionUnknownFormat:    "Unknown format.",
	nfcDescriptionNoResources:      "Not enough resources.",
	nfcDescriptionIsBusy:           "Device is busy.",
	nfcDescriptionNfcFunctionError: "NFC function error occurred.",
	nfcDescriptionTagNotFound:      "Tag was not found.",
	nfcDescriptionNeedRetry:        "Need retry.",
	nfcDescriptionNeedFormat:       "Need format.",
	nfcDescriptionInternalError:    "Internal error occurred.",
	nfcDescriptionHardwareError:    "Hardware error occurred.",
}

func (d nfcDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", nfcDescriptionToString[d], d, nfcDescriptionDescription[d])
}
