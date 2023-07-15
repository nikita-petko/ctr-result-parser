package descriptions

import "fmt"

type dbgDescription int32

const (
	dbgDescriptionDebugOutputIsDisabled dbgDescription = iota + 1
	dbgDescriptionDebuggerNotPresent
	dbgDescriptionInaccessiblePage
)

var dbgDescriptionToString = map[dbgDescription]string{
	dbgDescriptionDebugOutputIsDisabled: "DebugOutputIsDisabled",
	dbgDescriptionDebuggerNotPresent:    "DebuggerNotPresent",
	dbgDescriptionInaccessiblePage:      "InaccessiblePage",
}

var dbgDescriptionDescription = map[dbgDescription]string{
	dbgDescriptionDebugOutputIsDisabled: "Debug output is disabled",
	dbgDescriptionDebuggerNotPresent:    "Debugger is not connected",
	dbgDescriptionInaccessiblePage:      "An inaccessible page is included",
}

func (d dbgDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", dbgDescriptionToString[d], d, dbgDescriptionDescription[d])
}
