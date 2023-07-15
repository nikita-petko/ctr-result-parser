package descriptions

import "fmt"

type dlpDescription int32

const (
	dlpDescriptionInternalState dlpDescription = iota + 1
	dlpDescriptionInternalError
	dlpDescriptionAlreadyOccupiedWirelessDevice
	dlpDescriptionWirelessOff
	dlpDescriptionNotFoundServer
	dlpDescriptionServerIsFull
	dlpDescriptionInvalidAccessToMedia
	dlpDescriptionFailedToAccessMedia
	dlpDescriptionChildTooLarge
	dlpDescriptionIncommutable
	dlpDescriptionInvalidDlpVersion
	dlpDescriptionInvalidRegion
)

var dlpDescriptionToString = map[dlpDescription]string{
	dlpDescriptionInternalState:                 "InternalState",
	dlpDescriptionInternalError:                 "InternalError",
	dlpDescriptionAlreadyOccupiedWirelessDevice: "AlreadyOccupiedWirelessDevice",
	dlpDescriptionWirelessOff:                   "WirelessOff",
	dlpDescriptionNotFoundServer:                "NotFoundServer",
	dlpDescriptionServerIsFull:                  "ServerIsFull",
	dlpDescriptionInvalidAccessToMedia:          "InvalidAccessToMedia",
	dlpDescriptionFailedToAccessMedia:           "FailedToAccessMedia",
	dlpDescriptionChildTooLarge:                 "ChildTooLarge",
	dlpDescriptionIncommutable:                  "Incommutable",
	dlpDescriptionInvalidDlpVersion:             "InvalidDlpVersion",
	dlpDescriptionInvalidRegion:                 "InvalidRegion",
}

var dlpDescriptionDescription = map[dlpDescription]string{
	dlpDescriptionInternalState:                 "The internal state is innappropriate for using the API.",
	dlpDescriptionInternalError:                 "An error occured that cannot be handled from the application.",
	dlpDescriptionAlreadyOccupiedWirelessDevice: "The wireless device is already occupied.",
	dlpDescriptionWirelessOff:                   "State where communication is not possible.",
	dlpDescriptionNotFoundServer:                "Cannot find the server.",
	dlpDescriptionServerIsFull:                  "No more clients can connect to the server.",
	dlpDescriptionInvalidAccessToMedia:          "The media to be accessed is not supported.",
	dlpDescriptionFailedToAccessMedia:           "Access to the media failed.",
	dlpDescriptionChildTooLarge:                 "Too much NAND memory is required to import the child program.",
	dlpDescriptionIncommutable:                  "Cannot communicate with the partner.",
	dlpDescriptionInvalidDlpVersion:             "The DLP version is invalid.",
	dlpDescriptionInvalidRegion:                 "The child program region differs from the region of the host.",
}

func (d dlpDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", dlpDescriptionToString[d], d, dlpDescriptionDescription[d])
}
