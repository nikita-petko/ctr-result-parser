package descriptions

import "fmt"

type encDescription int32

const (
	encDescriptionNoBufferLeft encDescription = iota + 1
	encDescriptionInvalidParameter
	encDescriptionInvalidFormat
)

var encDescriptionToString = map[encDescription]string{
	encDescriptionNoBufferLeft:     "NoBufferLeft",
	encDescriptionInvalidParameter: "InvalidParameter",
	encDescriptionInvalidFormat:    "InvalidFormat",
}

var encDescriptionDescription = map[encDescription]string{
	encDescriptionNoBufferLeft:     "No buffer left.",
	encDescriptionInvalidParameter: "Invalid parameter.",
	encDescriptionInvalidFormat:    "Invalid format.",
}

func (d encDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", encDescriptionToString[d], d, encDescriptionDescription[d])
}
