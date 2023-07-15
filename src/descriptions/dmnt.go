package descriptions

import "fmt"

type dmntDescription int32

const (
	dmntDescriptionMaxHandle dmntDescription = iota + 1
	dmntDescriptionInvalidHandle
	dmntDescriptionInvalidProcessId
	dmntDescriptionInvalidThreadId
	dmntDescriptionNotAuthorized
	dmntDescriptionBusy
	dmntDescriptionAlreadyDone
	dmntDescriptionProcessTerminated
	dmntDescriptionNoEvent
	dmntDescriptionInnaccessiblePage
	dmntDescriptionInvalidDebugeeRegion
	dmntDescriptionReboot
	dmntDescriptionInvalidArgument
	dmntDescriptionInvalidProgamId
)

var dmntDescriptionToString = map[dmntDescription]string{
	dmntDescriptionMaxHandle:            "MaxHandle",
	dmntDescriptionInvalidHandle:        "InvalidHandle",
	dmntDescriptionInvalidProcessId:     "InvalidProcessId",
	dmntDescriptionInvalidThreadId:      "InvalidThreadId",
	dmntDescriptionNotAuthorized:        "NotAuthorized",
	dmntDescriptionBusy:                 "Busy",
	dmntDescriptionAlreadyDone:          "AlreadyDone",
	dmntDescriptionProcessTerminated:    "ProcessTerminated",
	dmntDescriptionNoEvent:              "NoEvent",
	dmntDescriptionInnaccessiblePage:    "InnaccessiblePage",
	dmntDescriptionInvalidDebugeeRegion: "InvalidDebugeeRegion",
	dmntDescriptionReboot:               "Reboot",
	dmntDescriptionInvalidArgument:      "InvalidArgument",
	dmntDescriptionInvalidProgamId:      "InvalidProgamId",
}

var dmntDescriptionDescription = map[dmntDescription]string{
	dmntDescriptionMaxHandle:            "maximum number of handles reached",
	dmntDescriptionInvalidHandle:        "invalid handle",
	dmntDescriptionInvalidProcessId:     "invalid process ID",
	dmntDescriptionInvalidThreadId:      "invalid thread ID",
	dmntDescriptionNotAuthorized:        "not authorized",
	dmntDescriptionBusy:                 "busy",
	dmntDescriptionAlreadyDone:          "already done",
	dmntDescriptionProcessTerminated:    "process terminated",
	dmntDescriptionNoEvent:              "no event",
	dmntDescriptionInnaccessiblePage:    "inaccessible page",
	dmntDescriptionInvalidDebugeeRegion: "invalid debugee region",
	dmntDescriptionReboot:               "reboot",
	dmntDescriptionInvalidArgument:      "invalid argument",
	dmntDescriptionInvalidProgamId:      "invalid program ID",
}

func (d dmntDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", dmntDescriptionToString[d], d, dmntDescriptionDescription[d])
}
