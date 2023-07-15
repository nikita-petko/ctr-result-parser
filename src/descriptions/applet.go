package descriptions

import "fmt"

type appletDescription int32

const (
	appletDescriptionAppletNoAreaToRegister appletDescription = iota + 1
	appletDescriptionAppletParameterBufferNotEmtpy
	appletDescriptionAppletCallerNotFound
	appletDescriptionAppletNotAllowed
	appletDescriptionAppletDifferentMode
	appletDescriptionAppletDifferentVersion
	appletDescriptionAppletShutdownProcessing
	appletDescriptionAppletTransitionBusy
	appletDescriptionAppletVersionMustLaunchDirectly
	appletDescriptionInvalidQuery
	appletDescriptionNoApplications
	appletDescriptionTooOldSystemUpdater
)

var appletDescriptionToString = map[appletDescription]string{
	appletDescriptionAppletNoAreaToRegister:          "AppletNoAreaToRegister",
	appletDescriptionAppletParameterBufferNotEmtpy:   "AppletParameterBufferNotEmtpy",
	appletDescriptionAppletCallerNotFound:            "AppletCallerNotFound",
	appletDescriptionAppletNotAllowed:                "AppletNotAllowed",
	appletDescriptionAppletDifferentMode:             "AppletDifferentMode",
	appletDescriptionAppletDifferentVersion:          "AppletDifferentVersion",
	appletDescriptionAppletShutdownProcessing:        "AppletShutdownProcessing",
	appletDescriptionAppletTransitionBusy:            "AppletTransitionBusy",
	appletDescriptionAppletVersionMustLaunchDirectly: "AppletVersionMustLaunchDirectly",
	appletDescriptionInvalidQuery:                    "InvalidQuery",
	appletDescriptionNoApplications:                  "NoApplications",
	appletDescriptionTooOldSystemUpdater:             "TooOldSystemUpdater",
}

var appletDescriptionDescription = map[appletDescription]string{
	appletDescriptionAppletNoAreaToRegister:          "There is no space in the table used for registration.",
	appletDescriptionAppletParameterBufferNotEmtpy:   "The parameter region is not empty.",
	appletDescriptionAppletCallerNotFound:            "The process that called the applet does not exist.",
	appletDescriptionAppletNotAllowed:                "Access is denied.",
	appletDescriptionAppletDifferentMode:             "Different mode.",
	appletDescriptionAppletDifferentVersion:          "Different version.",
	appletDescriptionAppletShutdownProcessing:        "Shutting down.",
	appletDescriptionAppletTransitionBusy:            "Something else is in the middle of a transition.",
	appletDescriptionAppletVersionMustLaunchDirectly: "A different version that must be started directly without using menus.",
	appletDescriptionInvalidQuery:                    "Invalid query.",
	appletDescriptionNoApplications:                  "No applications.",
	appletDescriptionTooOldSystemUpdater:             "The system update is too old.",
}

func (d appletDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", appletDescriptionToString[d], d, appletDescriptionDescription[d])
}
