package descriptions

import "fmt"

type midiDescription int32

const (
	midiDescriptionAlreadyOpened midiDescription = iota + 1
	midiDescriptionNotOpened
	midiDescriptionBufferOverflow
	midiDescriptionDeviceFifoFull
	midiDescriptionDeviceFrameError
	midiDescriptionDeviceInvalidDataLength
	midiDescriptionUnknownDevice
)

var midiDescriptionToString = map[midiDescription]string{
	midiDescriptionAlreadyOpened:           "AlreadyOpened",
	midiDescriptionNotOpened:               "NotOpened",
	midiDescriptionBufferOverflow:          "BufferOverflow",
	midiDescriptionDeviceFifoFull:          "DeviceFifoFull",
	midiDescriptionDeviceFrameError:        "DeviceFrameError",
	midiDescriptionDeviceInvalidDataLength: "DeviceInvalidDataLength",
	midiDescriptionUnknownDevice:           "UnknownDevice",
}

var midiDescriptionDescription = map[midiDescription]string{
	midiDescriptionAlreadyOpened:           "The device is already opened.",
	midiDescriptionNotOpened:               "The device is not opened.",
	midiDescriptionBufferOverflow:          "The buffer is full.",
	midiDescriptionDeviceFifoFull:          "The device FIFO is full.",
	midiDescriptionDeviceFrameError:        "The device frame error.",
	midiDescriptionDeviceInvalidDataLength: "The device invalid data length.",
	midiDescriptionUnknownDevice:           "The device is unknown.",
}

func (d midiDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", midiDescriptionToString[d], d, midiDescriptionDescription[d])
}
