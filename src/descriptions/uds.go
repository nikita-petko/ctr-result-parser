package descriptions

import "fmt"

type udsDescription int32

const (
	udsDescriptionNetworkIsFull udsDescription = iota + 1
	udsDescriptionWifiOff
	udsDescriptionInvalidParams
	udsDescriptionMiscellaneousSystemError
	udsDescriptionMalformedData
	udsDescriptionInvalidSdkVersion
	udsDescriptionSystemError
	udsDescriptionInvalidData
)

var udsDescriptionToString = map[udsDescription]string{
	udsDescriptionNetworkIsFull:            "NetworkIsFull",
	udsDescriptionWifiOff:                  "WifiOff",
	udsDescriptionInvalidParams:            "InvalidParams",
	udsDescriptionMiscellaneousSystemError: "MiscellaneousSystemError",
	udsDescriptionMalformedData:            "MalformedData",
	udsDescriptionInvalidSdkVersion:        "InvalidSdkVersion",
	udsDescriptionSystemError:              "SystemError",
	udsDescriptionInvalidData:              "InvalidData",
}

var udsDescriptionDescription = map[udsDescription]string{
	udsDescriptionNetworkIsFull:            "Could not connect because the maximum number of stations that can connect to the network has been reached.",
	udsDescriptionWifiOff:                  "Failed because the system is in wireless off mode.",
	udsDescriptionInvalidParams:            "A parameter error other than OutOfRange, TooLarge, or NotAuthorized.",
	udsDescriptionMiscellaneousSystemError: "The function failed because of a system error. It may succeed if called again.",
	udsDescriptionMalformedData:            "Detected the possibility of tampering.",
	udsDescriptionInvalidSdkVersion:        "Non-public error: the UDS SDK version is invalid.",
	udsDescriptionSystemError:              "The function failed because of a system error.",
	udsDescriptionInvalidData:              "Detected invalid data.",
}

func (d udsDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", udsDescriptionToString[d], d, udsDescriptionDescription[d])
}
