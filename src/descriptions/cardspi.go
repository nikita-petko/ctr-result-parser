package descriptions

import "fmt"

type cardspiDescription int32

const (
	cardspiDescriptionNotPermitted cardspiDescription = iota + 1
	cardspiDescriptionCardIreqTimeOut
	cardspiDescriptionCardIreqNotDetected
	cardspiDescriptionInvalidArgument
	cardspiDescriptionDeviceNotFound
	cardspiDescriptionTimeOut
)

var cardspiDescriptionToString = map[cardspiDescription]string{
	cardspiDescriptionNotPermitted:        "NotPermitted",
	cardspiDescriptionCardIreqTimeOut:     "CardIreqTimeOut",
	cardspiDescriptionCardIreqNotDetected: "CardIreqNotDetected",
	cardspiDescriptionInvalidArgument:     "InvalidArgument",
	cardspiDescriptionDeviceNotFound:      "DeviceNotFound",
	cardspiDescriptionTimeOut:             "TimeOut",
}

var cardspiDescriptionDescription = map[cardspiDescription]string{
	cardspiDescriptionNotPermitted:        "The operation is not permitted.",
	cardspiDescriptionCardIreqTimeOut:     "The card interrupt request timed out.",
	cardspiDescriptionCardIreqNotDetected: "The card interrupt request was not detected.",
	cardspiDescriptionInvalidArgument:     "An invalid argument was passed.",
	cardspiDescriptionDeviceNotFound:      "The device was not found.",
	cardspiDescriptionTimeOut:             "The operation timed out.",
}

func (d cardspiDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", cardspiDescriptionToString[d], d, cardspiDescriptionDescription[d])
}
