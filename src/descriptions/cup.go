package descriptions

import "fmt"

type cupDescription int32

const (
	cupDescriptionUpdateCancelled cupDescription = iota + 1
	cupDescriptionNotInitialized
	cupDescriptionAlreadyInitialized
	cupDescriptionAlreadyFinalized
	cupDescriptionAlreadyFinished
	cupDescriptionAlreadyCancelled
	cupDescriptionAlreadyUpdating
	cupDescriptionNotStarted
	cupDescriptionNotSupported
	cupDescriptionUpdatePartitionNotFound
	cupDescriptionUpdateNotRequired
	cupDescriptionInvalidUpdatePartitionFormat
	cupDescriptionUpdateInfoNotFound
)

var cupDescriptionToString = map[cupDescription]string{
	cupDescriptionUpdateCancelled:              "UpdateCancelled",
	cupDescriptionNotInitialized:               "NotInitialized",
	cupDescriptionAlreadyInitialized:           "AlreadyInitialized",
	cupDescriptionAlreadyFinalized:             "AlreadyFinalized",
	cupDescriptionAlreadyFinished:              "AlreadyFinished",
	cupDescriptionAlreadyCancelled:             "AlreadyCancelled",
	cupDescriptionAlreadyUpdating:              "AlreadyUpdating",
	cupDescriptionNotStarted:                   "NotStarted",
	cupDescriptionNotSupported:                 "NotSupported",
	cupDescriptionUpdatePartitionNotFound:      "UpdatePartitionNotFound",
	cupDescriptionUpdateNotRequired:            "UpdateNotRequired",
	cupDescriptionInvalidUpdatePartitionFormat: "InvalidUpdatePartitionFormat",
	cupDescriptionUpdateInfoNotFound:           "UpdateInfoNotFound",
}

var cupDescriptionDescription = map[cupDescription]string{
	cupDescriptionUpdateCancelled:              "The update was cancelled.",
	cupDescriptionNotInitialized:               "The update was not initialized.",
	cupDescriptionAlreadyInitialized:           "The update was already initialized.",
	cupDescriptionAlreadyFinalized:             "The update was already finalized.",
	cupDescriptionAlreadyFinished:              "The update was already finished.",
	cupDescriptionAlreadyCancelled:             "The update was already cancelled.",
	cupDescriptionAlreadyUpdating:              "The update was already updating.",
	cupDescriptionNotStarted:                   "The update was not started.",
	cupDescriptionNotSupported:                 "The update is not supported.",
	cupDescriptionUpdatePartitionNotFound:      "The update partition was not found.",
	cupDescriptionUpdateNotRequired:            "The update is not required.",
	cupDescriptionInvalidUpdatePartitionFormat: "The update partition format is invalid.",
	cupDescriptionUpdateInfoNotFound:           "The update info was not found.",
}

func (d cupDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", cupDescriptionToString[d], d, cupDescriptionDescription[d])
}
