package descriptions

import "fmt"

type vctlDescription int32

const (
	vctlDescriptionNotInitialized vctlDescription = iota + 1
	vctlDescriptionAlreadyInitialized
	vctlDescriptionUnsupportedVersion
	vctlDescriptionInvalidArgument
	vctlDescriptionNotEnoughSpace
	vctlDescriptionFatalError vctlDescription = 900
)

var vctlDescriptionToString = map[vctlDescription]string{
	vctlDescriptionNotInitialized:     "NotInitialized",
	vctlDescriptionAlreadyInitialized: "AlreadyInitialized",
	vctlDescriptionUnsupportedVersion: "UnsupportedVersion",
	vctlDescriptionInvalidArgument:    "InvalidArgument",
	vctlDescriptionNotEnoughSpace:     "NotEnoughSpace",
	vctlDescriptionFatalError:         "FatalError",
}

var vctlDescriptionDescription = map[vctlDescription]string{
	vctlDescriptionNotInitialized:     "A process that requires initialization was called before it was initialized.",
	vctlDescriptionAlreadyInitialized: "The initialization process was called while already initialized.",
	vctlDescriptionUnsupportedVersion: "Unsupported version.",
	vctlDescriptionInvalidArgument:    "Invalid argument.",
	vctlDescriptionNotEnoughSpace:     "Insufficient space.",
	vctlDescriptionFatalError:         "A fatal error occurred.",
}

func (d vctlDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", vctlDescriptionToString[d], d, vctlDescriptionDescription[d])
}
