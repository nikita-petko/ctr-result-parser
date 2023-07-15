package descriptions

import "fmt"

type nsDescription int32

const (
	nsDescriptionRebootNotRequired nsDescription = iota + 1
	nsDescriptionShutdownProcessing
	nsDescriptionDemoLaunchLimited
	nsDescriptionContentNotFound
)

var nsDescriptionToString = map[nsDescription]string{
	nsDescriptionRebootNotRequired:  "RebootNotRequired",
	nsDescriptionShutdownProcessing: "ShutdownProcessing",
	nsDescriptionDemoLaunchLimited:  "DemoLaunchLimited",
	nsDescriptionContentNotFound:    "ContentNotFound",
}

var nsDescriptionDescription = map[nsDescription]string{
	nsDescriptionRebootNotRequired:  "Reboot is not required.",
	nsDescriptionShutdownProcessing: "Shutting down.",
	nsDescriptionDemoLaunchLimited:  "Demo launch is limited.",
	nsDescriptionContentNotFound:    "Content not found.",
}

func (d nsDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", nsDescriptionToString[d], d, nsDescriptionDescription[d])
}
