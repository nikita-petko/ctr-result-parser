package descriptions

import "fmt"

type socketDescription int32

const (
	socketDescriptionFailedToInitializeInterface socketDescription = iota + 1
	socketDescriptionFailedToInitializeSocketCore
	socketDescriptionTooManyRequests
	socketDescriptionPermissionDenied
	socketDescriptionUnkownConfiguration
	socketDescriptionUnkownClient
	socketDescriptionBadDescriptor
	socketDescriptionRequestSessionFull
	socketDescriptionNetworkResetted
	socketDescriptionTooManyProcesses
	socketDescriptionAlreadyRegistered
	socketDescriptionTooManySockets
	socketDescriptionMismatchVersion
	socketDescriptionAddressCollision
	socketDescriptionTimeout
	socketDescriptionOutOfSystemResource
	socketDescriptionInvalidCoreState
	socketDescriptionAborted
)

var socketDescriptionToString = map[socketDescription]string{
	socketDescriptionFailedToInitializeInterface:  "FailedToInitializeInterface",
	socketDescriptionFailedToInitializeSocketCore: "FailedToInitializeSocketCore",
	socketDescriptionTooManyRequests:              "TooManyRequests",
	socketDescriptionPermissionDenied:             "PermissionDenied",
	socketDescriptionUnkownConfiguration:          "UnkownConfiguration",
	socketDescriptionUnkownClient:                 "UnkownClient",
	socketDescriptionBadDescriptor:                "BadDescriptor",
	socketDescriptionRequestSessionFull:           "RequestSessionFull",
	socketDescriptionNetworkResetted:              "NetworkResetted",
	socketDescriptionTooManyProcesses:             "TooManyProcesses",
	socketDescriptionAlreadyRegistered:            "AlreadyRegistered",
	socketDescriptionTooManySockets:               "TooManySockets",
	socketDescriptionMismatchVersion:              "MismatchVersion",
	socketDescriptionAddressCollision:             "AddressCollision",
	socketDescriptionTimeout:                      "Timeout",
	socketDescriptionOutOfSystemResource:          "OutOfSystemResource",
	socketDescriptionInvalidCoreState:             "InvalidCoreState",
	socketDescriptionAborted:                      "Aborted",
}

var socketDescriptionDescription = map[socketDescription]string{
	socketDescriptionFailedToInitializeInterface:  "Failed to initialize socket interface",
	socketDescriptionFailedToInitializeSocketCore: "Failed to initialize socket core",
	socketDescriptionTooManyRequests:              "Too many requests",
	socketDescriptionPermissionDenied:             "Permission denied",
	socketDescriptionUnkownConfiguration:          "Unknown configuration",
	socketDescriptionUnkownClient:                 "Unknown client",
	socketDescriptionBadDescriptor:                "Bad descriptor",
	socketDescriptionRequestSessionFull:           "Request session full",
	socketDescriptionNetworkResetted:              "Network resetted",
	socketDescriptionTooManyProcesses:             "Too many processes",
	socketDescriptionAlreadyRegistered:            "Already registered",
	socketDescriptionTooManySockets:               "Too many sockets",
	socketDescriptionMismatchVersion:              "Mismatch version",
	socketDescriptionAddressCollision:             "Address collision",
	socketDescriptionTimeout:                      "Timeout",
	socketDescriptionOutOfSystemResource:          "Out of system resource",
	socketDescriptionInvalidCoreState:             "Invalid core state",
	socketDescriptionAborted:                      "Aborted",
}

func (d socketDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", socketDescriptionToString[d], d, socketDescriptionDescription[d])
}
