package descriptions

import "fmt"

type nwmDescription int32

const (
	nwmDescriptionSdioInitFailure nwmDescription = iota + 1
	nwmDescriptionModuleInitFailure
	nwmDescriptionModuleEventFailure
	nwmDescriptionOtherFirmwareError
	nwmDescriptionCommandFailure
	nwmDescriptionInvalidState
	nwmDescriptionStateMismatch
	nwmDescriptionDuplicateRxEntry
	nwmDescriptionRxEntryNotFound
	nwmDescriptionBssNotFound
	nwmDescriptionLinkingError
	nwmDescriptionWifiOff
	nwmDescriptionAlreadyInState
	nwmDescriptionOperating
	nwmDescriptionInfraFirmwareError
	nwmDescriptionSapFirmwareError
	nwmDescriptionCecFirmwareError
)

var nwmDescriptionToString = map[nwmDescription]string{
	nwmDescriptionSdioInitFailure:    "SdioInitFailure",
	nwmDescriptionModuleInitFailure:  "ModuleInitFailure",
	nwmDescriptionModuleEventFailure: "ModuleEventFailure",
	nwmDescriptionOtherFirmwareError: "OtherFirmwareError",
	nwmDescriptionCommandFailure:     "CommandFailure",
	nwmDescriptionInvalidState:       "InvalidState",
	nwmDescriptionStateMismatch:      "StateMismatch",
	nwmDescriptionDuplicateRxEntry:   "DuplicateRxEntry",
	nwmDescriptionRxEntryNotFound:    "RxEntryNotFound",
	nwmDescriptionBssNotFound:        "BssNotFound",
	nwmDescriptionLinkingError:       "LinkingError",
	nwmDescriptionWifiOff:            "WifiOff",
	nwmDescriptionAlreadyInState:     "AlreadyInState",
	nwmDescriptionOperating:          "Operating",
	nwmDescriptionInfraFirmwareError: "InfraFirmwareError",
	nwmDescriptionSapFirmwareError:   "SapFirmwareError",
	nwmDescriptionCecFirmwareError:   "CecFirmwareError",
}

var nwmDescriptionDescription = map[nwmDescription]string{
	nwmDescriptionSdioInitFailure:    "Failed to initialize SDIO",
	nwmDescriptionModuleInitFailure:  "Failed to initialize the wireless module.",
	nwmDescriptionModuleEventFailure: "An illegal event has occurred from the wireless module.",
	nwmDescriptionOtherFirmwareError: "Radio module firmware error on non-Infra/SoftAP/CEC",
	nwmDescriptionCommandFailure:     "Command operation to the wireless module failed.",
	nwmDescriptionInvalidState:       "Error due to state inconsistency.",
	nwmDescriptionStateMismatch:      "The requested state and final transition destination state do not match.",
	nwmDescriptionDuplicateRxEntry:   "Error in radio reception entry registration.",
	nwmDescriptionRxEntryNotFound:    "Error deleting radio receive entry.",
	nwmDescriptionBssNotFound:        "Destination BSS was not found.",
	nwmDescriptionLinkingError:       "Error during connection sequence.",
	nwmDescriptionWifiOff:            "An attempt was made to perform a wireless operation with WiFi turned off.",
	nwmDescriptionAlreadyInState:     "Already in the desired state.",
	nwmDescriptionOperating:          "API is running. (for Scan/Measure Channel)",
	nwmDescriptionInfraFirmwareError: "Infra radio module firmware error",
	nwmDescriptionSapFirmwareError:   "SoftAP radio module firmware error",
	nwmDescriptionCecFirmwareError:   "CEC radio module firmware error",
}

func (d nwmDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", nwmDescriptionToString[d], d, nwmDescriptionDescription[d])
}
