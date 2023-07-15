package descriptions

import "fmt"

type srvDescription int32

const (
	srvDescriptionFailedSynchronization srvDescription = iota + 1
	srvDescriptionNoSuchHandle
	srvDescriptionAlreadyExists
	srvDescriptionNotExists
	srvDescriptionTooLongServiceName
	srvDescriptionNotPermitted
	srvDescriptionInvalidName
	srvDescriptionBufferOverflow
)

var srvDescriptionToString = map[srvDescription]string{
	srvDescriptionFailedSynchronization: "FailedSynchronization",
	srvDescriptionNoSuchHandle:          "NoSuchHandle",
	srvDescriptionAlreadyExists:         "AlreadyExists",
	srvDescriptionNotExists:             "NotExists",
	srvDescriptionTooLongServiceName:    "TooLongServiceName",
	srvDescriptionNotPermitted:          "NotPermitted",
	srvDescriptionInvalidName:           "InvalidName",
	srvDescriptionBufferOverflow:        "BufferOverflow",
}

var srvDescriptionDescription = map[srvDescription]string{
	srvDescriptionFailedSynchronization: "Failed to synchronize with the server",
	srvDescriptionNoSuchHandle:          "No such handle exists",
	srvDescriptionAlreadyExists:         "The handle already exists",
	srvDescriptionNotExists:             "The handle does not exist",
	srvDescriptionTooLongServiceName:    "The service name is too long",
	srvDescriptionNotPermitted:          "The operation is not permitted",
	srvDescriptionInvalidName:           "The service name is invalid",
	srvDescriptionBufferOverflow:        "The buffer is too big",
}

func (d srvDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", srvDescriptionToString[d], d, srvDescriptionDescription[d])
}
