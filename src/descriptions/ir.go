package descriptions

import "fmt"

type irDescription int32

const (
	irDescriptionMachineSleep irDescription = iota + 1
	irDescriptionFatalError
	irDescriptionSignatureNotFound
	irDescriptionDifferentSessionID
	irDescriptionInvalidCRC
	irDescriptionFollowingDataNotExist
	irDescriptionFramingError
	irDescriptionOverrunError
	irDescriptionPerformanceError
	irDescriptionModuleOtherError
	irDescriptionAlreadyConnected
	irDescriptionAlreadyTryingToConnect
	irDescriptionNotConnected
	irDescriptionBufferFull
	irDescriptionBufferInsufficient
	irDescriptionPacketFull
	irDescriptionTimeout
	irDescriptionPeripheral
	irDescriptionPeripheralDataNotExist
	irDescriptionCannotConfirmID
	irDescriptionInvalidData
)

var irDescriptionToString = map[irDescription]string{
	irDescriptionMachineSleep:           "MachineSleep",
	irDescriptionFatalError:             "FatalError",
	irDescriptionSignatureNotFound:      "SignatureNotFound",
	irDescriptionDifferentSessionID:     "DifferentSessionID",
	irDescriptionInvalidCRC:             "InvalidCRC",
	irDescriptionFollowingDataNotExist:  "FollowingDataNotExist",
	irDescriptionFramingError:           "FramingError",
	irDescriptionOverrunError:           "OverrunError",
	irDescriptionPerformanceError:       "PerformanceError",
	irDescriptionModuleOtherError:       "ModuleOtherError",
	irDescriptionAlreadyConnected:       "AlreadyConnected",
	irDescriptionAlreadyTryingToConnect: "AlreadyTryingToConnect",
	irDescriptionNotConnected:           "NotConnected",
	irDescriptionBufferFull:             "BufferFull",
	irDescriptionBufferInsufficient:     "BufferInsufficient",
	irDescriptionPacketFull:             "PacketFull",
	irDescriptionTimeout:                "Timeout",
	irDescriptionPeripheral:             "Peripheral",
	irDescriptionPeripheralDataNotExist: "PeripheralDataNotExist",
	irDescriptionCannotConfirmID:        "CannotConfirmID",
	irDescriptionInvalidData:            "InvalidData",
}

var irDescriptionDescription = map[irDescription]string{
	irDescriptionMachineSleep:           "Sleeping because the system is in Sleep Mode.",
	irDescriptionFatalError:             "The IR module may be malfunctioning.",
	irDescriptionSignatureNotFound:      "Cannot find a signature indicating the start of the packet.",
	irDescriptionDifferentSessionID:     "The session ID is different.",
	irDescriptionInvalidCRC:             "Received data has invalid CRC.",
	irDescriptionFollowingDataNotExist:  "Incomplete data.",
	irDescriptionFramingError:           "The received data has a framing error.",
	irDescriptionOverrunError:           "The receiving process could not keep pace with the device's data receive speed.",
	irDescriptionPerformanceError:       "An error related to system performance or CTR performance at a high baud-rate setting occurred.",
	irDescriptionModuleOtherError:       "Unknown error in the IR module.",
	irDescriptionAlreadyConnected:       "Already connected.",
	irDescriptionAlreadyTryingToConnect: "Already trying to connect.",
	irDescriptionNotConnected:           "Not connected.",
	irDescriptionBufferFull:             "The buffer is full.",
	irDescriptionBufferInsufficient:     "The buffer is too small.",
	irDescriptionPacketFull:             "The packet is full.",
	irDescriptionTimeout:                "A timeout has occurred.",
	irDescriptionPeripheral:             "Communication target error.",
	irDescriptionPeripheralDataNotExist: "Data was not found in the communication target.",
	irDescriptionCannotConfirmID:        "Could not confirm the ID for communication.",
	irDescriptionInvalidData:            "Invalid data was received. (This could be a packet replay attack.)",
}

func (d irDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", irDescriptionToString[d], d, irDescriptionDescription[d])
}
