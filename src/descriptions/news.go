package descriptions

import "fmt"

type newsDescription int32

const (
	newsDescriptionNone newsDescription = iota
	newsDescriptionInvalidSubjectSize
	newsDescriptionInvalidMessageSize
	newsDescriptionInvalidPictureSize
	newsDescriptionInvalidPicture
)

var newsDescriptionToString = map[newsDescription]string{
	newsDescriptionNone:               "None",
	newsDescriptionInvalidSubjectSize: "InvalidSubjectSize",
	newsDescriptionInvalidMessageSize: "InvalidMessageSize",
	newsDescriptionInvalidPictureSize: "InvalidPictureSize",
	newsDescriptionInvalidPicture:     "InvalidPicture",
}

var newsDescriptionDescription = map[newsDescription]string{
	newsDescriptionNone:               "No error.",
	newsDescriptionInvalidSubjectSize: "The length of the subject line is too large.",
	newsDescriptionInvalidMessageSize: "The message body is too large.",
	newsDescriptionInvalidPictureSize: "The image is too large.",
	newsDescriptionInvalidPicture:     "Invalid image.",
}

func (d newsDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", newsDescriptionToString[d], d, newsDescriptionDescription[d])
}
