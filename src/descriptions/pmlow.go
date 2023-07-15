package descriptions

import "fmt"

type pmlowDescription int32

const (
	pmlowDescriptionInvalidRomFormat pmlowDescription = iota + 1
	pmlowDescriptionFailedToReadProgramInfo
	pmlowDescriptionProgramNotLoaded
	pmlowDescriptionFailedToVerifyAcidSignature
	pmlowDescriptionInvalidAcid
)

var pmlowDescriptionToString = map[pmlowDescription]string{
	pmlowDescriptionInvalidRomFormat:            "InvalidRomFormat",
	pmlowDescriptionFailedToReadProgramInfo:     "FailedToReadProgramInfo",
	pmlowDescriptionProgramNotLoaded:            "ProgramNotLoaded",
	pmlowDescriptionFailedToVerifyAcidSignature: "FailedToVerifyAcidSignature",
	pmlowDescriptionInvalidAcid:                 "InvalidAcid",
}

var pmlowDescriptionDescription = map[pmlowDescription]string{
	pmlowDescriptionInvalidRomFormat:            "Invalid ROM format",
	pmlowDescriptionFailedToReadProgramInfo:     "Failed to read program info",
	pmlowDescriptionProgramNotLoaded:            "Program not loaded",
	pmlowDescriptionFailedToVerifyAcidSignature: "Failed to verify access control info descriptor signature",
	pmlowDescriptionInvalidAcid:                 "Invalid access control info descriptor",
}

func (d pmlowDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", pmlowDescriptionToString[d], d, pmlowDescriptionDescription[d])
}
