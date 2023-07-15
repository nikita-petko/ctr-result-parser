package descriptions

import "fmt"

type ldrDescription int32

const (
	ldrDescriptionFailedHostFileOperation ldrDescription = iota + 1
	ldrDescriptionInvalidRomFormat
	ldrDescriptionNoEntry
	ldrDescriptionOutOfMemory
	ldrDescriptionInvalidProgramLaunchInfo
)

var ldrDescriptionToString = map[ldrDescription]string{
	ldrDescriptionFailedHostFileOperation:  "FailedHostFileOperation",
	ldrDescriptionInvalidRomFormat:         "InvalidRomFormat",
	ldrDescriptionNoEntry:                  "NoEntry",
	ldrDescriptionOutOfMemory:              "OutOfMemory",
	ldrDescriptionInvalidProgramLaunchInfo: "InvalidProgramLaunchInfo",
}

var ldrDescriptionDescription = map[ldrDescription]string{
	ldrDescriptionFailedHostFileOperation:  "Failed to perform host file operation",
	ldrDescriptionInvalidRomFormat:         "Invalid ROM format",
	ldrDescriptionNoEntry:                  "Entry function does not exist",
	ldrDescriptionOutOfMemory:              "Out of memory",
	ldrDescriptionInvalidProgramLaunchInfo: "Invalid program launch information",
}

func (d ldrDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", ldrDescriptionToString[d], d, ldrDescriptionDescription[d])
}
