package descriptions

import "fmt"

type cfgDescription int32

const (
	cfgDescriptionNotVerified cfgDescription = iota + 1
	cfgDescriptionVerificationFailed
	cfgDescriptionInvalidNtrSetting
	cfgDescriptionAlreadyLatestVersion
	cfgDescriptionMountContentFailed
	cfgDescriptionInvalidTarget
	cfgDescriptionObsoleteResult cfgDescription = 1023
)

var cfgDescriptionToString = map[cfgDescription]string{
	cfgDescriptionNotVerified:          "NotVerified",
	cfgDescriptionVerificationFailed:   "VerificationFailed",
	cfgDescriptionInvalidNtrSetting:    "InvalidNtrSetting",
	cfgDescriptionAlreadyLatestVersion: "AlreadyLatestVersion",
	cfgDescriptionMountContentFailed:   "MountContentFailed",
	cfgDescriptionInvalidTarget:        "InvalidTarget",
	cfgDescriptionObsoleteResult:       "ObsoleteResult",
}

var cfgDescriptionDescription = map[cfgDescription]string{
	cfgDescriptionNotVerified:          "Signature not verified.",
	cfgDescriptionVerificationFailed:   "Signature verification failed.",
	cfgDescriptionInvalidNtrSetting:    "A valid NTR setting was not found.",
	cfgDescriptionAlreadyLatestVersion: "Already the latest version.",
	cfgDescriptionMountContentFailed:   "Failed to mount content.",
	cfgDescriptionInvalidTarget:        "The target is invalid.",
	cfgDescriptionObsoleteResult:       "Obsolete; the error must be revised to another as soon as possible.",
}

func (d cfgDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", cfgDescriptionToString[d], d, cfgDescriptionDescription[d])
}
