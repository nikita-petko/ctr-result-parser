package descriptions

import "fmt"

type qtmDescription int32

const (
	qtmDescriptionFatalError qtmDescription = iota + 1
	qtmDescriptionInvalidArgument
	qtmDescriptionIsSleeping
	qtmDescriptionInvalidMemoryAllocation
	qtmDescriptionIrLedInvalidState
	qtmDescriptionCameraInvalidState
	qtmDescriptionInvalidLuminance
	qtmDescriptionCameraExclusive
	qtmDescriptionNotAvailable
)

var qtmDescriptionToString = map[qtmDescription]string{
	qtmDescriptionFatalError:              "FatalError",
	qtmDescriptionInvalidArgument:         "InvalidArgument",
	qtmDescriptionIsSleeping:              "IsSleeping",
	qtmDescriptionInvalidMemoryAllocation: "InvalidMemoryAllocation",
	qtmDescriptionIrLedInvalidState:       "IrLedInvalidState",
	qtmDescriptionCameraInvalidState:      "CameraInvalidState",
	qtmDescriptionInvalidLuminance:        "InvalidLuminance",
	qtmDescriptionCameraExclusive:         "CameraExclusive",
	qtmDescriptionNotAvailable:            "NotAvailable",
}

var qtmDescriptionDescription = map[qtmDescription]string{
	qtmDescriptionFatalError:              "A fatal error has occurred.",
	qtmDescriptionInvalidArgument:         "An invalid argument was passed.",
	qtmDescriptionIsSleeping:              "The device is sleeping.",
	qtmDescriptionInvalidMemoryAllocation: "Memory allocation failed.",
	qtmDescriptionIrLedInvalidState:       "The IR LED is in an invalid state.",
	qtmDescriptionCameraInvalidState:      "The camera is in an invalid state.",
	qtmDescriptionInvalidLuminance:        "The luminance is invalid.",
	qtmDescriptionCameraExclusive:         "The camera is in use by another process.",
	qtmDescriptionNotAvailable:            "The requested function is not available.",
}

func (d qtmDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", qtmDescriptionToString[d], d, qtmDescriptionDescription[d])
}
