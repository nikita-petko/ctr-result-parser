package descriptions

import "fmt"

type cecDescription int32

const (
	cecDescriptionUnknown cecDescription = iota + 100
	cecDescriptionBoxSizeFull
	cecDescriptionBoxMessNumFull
	cecDescriptionBoxNumFull
	cecDescriptionBoxAlreadyExists
	cecDescriptionMessTooLarge
	cecDescriptionInvalidData
	cecDescriptionInvalidId
	cecDescriptionNotAgreeEula
	cecDescriptionParentalControlCec
	cecDescriptionFileAccessFailed
	cecDescriptionSessionCanceled
)

var cecDescriptionToString = map[cecDescription]string{
	cecDescriptionUnknown:            "Unknown",
	cecDescriptionBoxSizeFull:        "BoxSizeFull",
	cecDescriptionBoxMessNumFull:     "BoxMessNumFull",
	cecDescriptionBoxNumFull:         "BoxNumFull",
	cecDescriptionBoxAlreadyExists:   "BoxAlreadyExists",
	cecDescriptionMessTooLarge:       "MessTooLarge",
	cecDescriptionInvalidData:        "InvalidData",
	cecDescriptionInvalidId:          "InvalidId",
	cecDescriptionNotAgreeEula:       "NotAgreeEula",
	cecDescriptionParentalControlCec: "ParentalControlCec",
	cecDescriptionFileAccessFailed:   "FileAccessFailed",
	cecDescriptionSessionCanceled:    "SessionCanceled",
}

var cecDescriptionDescription = map[cecDescription]string{
	cecDescriptionUnknown:            "Unknown",
	cecDescriptionBoxSizeFull:        "The box size is full.",
	cecDescriptionBoxMessNumFull:     "The number of messages in the box is full.",
	cecDescriptionBoxNumFull:         "The number of boxes is full.",
	cecDescriptionBoxAlreadyExists:   "The box already exists.",
	cecDescriptionMessTooLarge:       "The message is too large.",
	cecDescriptionInvalidData:        "The data is invalid.",
	cecDescriptionInvalidId:          "The ID is invalid.",
	cecDescriptionNotAgreeEula:       "User has not agreed to the EULA.",
	cecDescriptionParentalControlCec: "User has not agreed due to Parental Controls.",
	cecDescriptionFileAccessFailed:   "Failed.",
	cecDescriptionSessionCanceled:    "Session canceled.",
}

func (d cecDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", cecDescriptionToString[d], d, cecDescriptionDescription[d])
}
