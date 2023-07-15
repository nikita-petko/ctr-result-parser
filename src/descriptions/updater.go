package descriptions

import "fmt"

type updaterDescription int32

const (
	updaterDescriptionNotExpectedVersion updaterDescription = iota + 1
	updaterDescriptionHashMismatch
	updaterDescriptionNotExpectedMcuFirmVersion
	updaterDescriptionFailedMcuFirmUpdate
	updaterDescriptionFailedAllocateBuffer
	updaterDescriptionInvalidHashLength
	updaterDescriptionFailedReadContent
	updaterDescriptionNotSupported
	updaterDescriptionFailedUpdateMcu
)

var updaterDescriptionToString = map[updaterDescription]string{
	updaterDescriptionNotExpectedVersion:        "NotExpectedVersion",
	updaterDescriptionHashMismatch:              "HashMismatch",
	updaterDescriptionNotExpectedMcuFirmVersion: "NotExpectedMcuFirmVersion",
	updaterDescriptionFailedMcuFirmUpdate:       "FailedMcuFirmUpdate",
	updaterDescriptionFailedAllocateBuffer:      "FailedAllocateBuffer",
	updaterDescriptionInvalidHashLength:         "InvalidHashLength",
	updaterDescriptionFailedReadContent:         "FailedReadContent",
	updaterDescriptionNotSupported:              "NotSupported",
	updaterDescriptionFailedUpdateMcu:           "FailedUpdateMcu",
}

var updaterDescriptionDescription = map[updaterDescription]string{
	updaterDescriptionNotExpectedVersion:        "The version of the firmware is not as expected.",
	updaterDescriptionHashMismatch:              "The hash of the firmware is not as expected.",
	updaterDescriptionNotExpectedMcuFirmVersion: "The version of the MCU firmware is not as expected.",
	updaterDescriptionFailedMcuFirmUpdate:       "Failed to update the MCU firmware.",
	updaterDescriptionFailedAllocateBuffer:      "Failed to allocate buffer.",
	updaterDescriptionInvalidHashLength:         "The length of the hash is invalid.",
	updaterDescriptionFailedReadContent:         "Failed to read content.",
	updaterDescriptionNotSupported:              "The operation is not supported.",
	updaterDescriptionFailedUpdateMcu:           "Failed to update the MCU.",
}

func (d updaterDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", updaterDescriptionToString[d], d, updaterDescriptionDescription[d])
}
