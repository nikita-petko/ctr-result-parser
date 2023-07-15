package descriptions

import "fmt"

type rdtDescription int32

const (
	rdtDescriptionResetReceived rdtDescription = iota + 1
	rdtDescriptionUntimelyCall
	rdtDescriptionInvalidValue
)

var rdtDescriptionToString = map[rdtDescription]string{
	rdtDescriptionResetReceived: "ResetReceived",
	rdtDescriptionUntimelyCall:  "UntimelyCall",
	rdtDescriptionInvalidValue:  "InvalidValue",
}

var rdtDescriptionDescription = map[rdtDescription]string{
	rdtDescriptionResetReceived: "Thesese results transition to a CLOSED state because a reset signal was received from the partner.",
	rdtDescriptionUntimelyCall:  "Do not call the function in this state.",
	rdtDescriptionInvalidValue:  "Invalid parameter value.",
}

func (d rdtDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", rdtDescriptionToString[d], d, rdtDescriptionDescription[d])
}
