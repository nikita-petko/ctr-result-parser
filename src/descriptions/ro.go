package descriptions

import "fmt"

type roDescription int32

const (
	roDescriptionAlreadyLoaded roDescription = iota + 1
	roDescriptionAtexitNotFound
	roDescriptionBrokenObject
	roDescriptionControlObjectNotFound
	roDescriptionEitNodeNotFound
	roDescriptionInvalidList
	roDescriptionInvalidOffsetRange
	roDescriptionInvalidRegion
	roDescriptionInvalidSign
	roDescriptionInvalidSign2
	roDescriptionInvalidString
	roDescriptionInvalidTarget
	roDescriptionNotLoaded
	roDescriptionNotRegistered
	roDescriptionOutOfSpace
	roDescriptionRangeErrorAtExport
	roDescriptionRangeErrorAtHeader
	roDescriptionRangeErrorAtImport
	roDescriptionRangeErrorAtIndexExport
	roDescriptionRangeErrorAtIndexImport
	roDescriptionRangeErrorAtInternalRelocation
	roDescriptionRangeErrorAtOffsetExport
	roDescriptionRangeErrorAtOffsetImport
	roDescriptionRangeErrorAtReference
	roDescriptionRangeErrorAtSection
	roDescriptionRangeErrorAtSymbolExport
	roDescriptionRangeErrorAtSymbolImport
	roDescriptionRegistrationNotFound
	roDescriptionRwNotSupported
	roDescriptionStaticModule
	roDescriptionTooSmallSize
	roDescriptionTooSmallTarget
	roDescriptionUnknownObjectControl
	roDescriptionUnknownRelocationType
	roDescriptionVeneerRequired
	roDescriptionVerificationFailed
)

var roDescriptionToString = map[roDescription]string{
	roDescriptionAlreadyLoaded:                  "AlreadyLoaded",
	roDescriptionAtexitNotFound:                 "AtexitNotFound",
	roDescriptionBrokenObject:                   "BrokenObject",
	roDescriptionControlObjectNotFound:          "ControlObjectNotFound",
	roDescriptionEitNodeNotFound:                "EitNodeNotFound",
	roDescriptionInvalidList:                    "InvalidList",
	roDescriptionInvalidOffsetRange:             "InvalidOffsetRange",
	roDescriptionInvalidRegion:                  "InvalidRegion",
	roDescriptionInvalidSign:                    "InvalidSign",
	roDescriptionInvalidSign2:                   "InvalidSign2",
	roDescriptionInvalidString:                  "InvalidString",
	roDescriptionInvalidTarget:                  "InvalidTarget",
	roDescriptionNotLoaded:                      "NotLoaded",
	roDescriptionNotRegistered:                  "NotRegistered",
	roDescriptionOutOfSpace:                     "OutOfSpace",
	roDescriptionRangeErrorAtExport:             "RangeErrorAtExport",
	roDescriptionRangeErrorAtHeader:             "RangeErrorAtHeader",
	roDescriptionRangeErrorAtImport:             "RangeErrorAtImport",
	roDescriptionRangeErrorAtIndexExport:        "RangeErrorAtIndexExport",
	roDescriptionRangeErrorAtIndexImport:        "RangeErrorAtIndexImport",
	roDescriptionRangeErrorAtInternalRelocation: "RangeErrorAtInternalRelocation",
	roDescriptionRangeErrorAtOffsetExport:       "RangeErrorAtOffsetExport",
	roDescriptionRangeErrorAtOffsetImport:       "RangeErrorAtOffsetImport",
	roDescriptionRangeErrorAtReference:          "RangeErrorAtReference",
	roDescriptionRangeErrorAtSection:            "RangeErrorAtSection",
	roDescriptionRangeErrorAtSymbolExport:       "RangeErrorAtSymbolExport",
	roDescriptionRangeErrorAtSymbolImport:       "RangeErrorAtSymbolImport",
	roDescriptionRegistrationNotFound:           "RegistrationNotFound",
	roDescriptionRwNotSupported:                 "RwNotSupported",
	roDescriptionStaticModule:                   "StaticModule",
	roDescriptionTooSmallSize:                   "TooSmallSize",
	roDescriptionTooSmallTarget:                 "TooSmallTarget",
	roDescriptionUnknownObjectControl:           "UnknownObjectControl",
	roDescriptionUnknownRelocationType:          "UnknownRelocationType",
	roDescriptionVeneerRequired:                 "VeneerRequired",
	roDescriptionVerificationFailed:             "VerificationFailed",
}

var roDescriptionDescription = map[roDescription]string{
	roDescriptionAlreadyLoaded:                  "The module is already loaded.",
	roDescriptionAtexitNotFound:                 "The atexit function is not found.",
	roDescriptionBrokenObject:                   "The module is broken.",
	roDescriptionControlObjectNotFound:          "The control object is not found.",
	roDescriptionEitNodeNotFound:                "The EIT node is not found.",
	roDescriptionInvalidList:                    "The list is invalid.",
	roDescriptionInvalidOffsetRange:             "The offset range is invalid.",
	roDescriptionInvalidRegion:                  "The region is invalid.",
	roDescriptionInvalidSign:                    "The sign is invalid.",
	roDescriptionInvalidSign2:                   "The sign2 is invalid.",
	roDescriptionInvalidString:                  "The string is invalid.",
	roDescriptionInvalidTarget:                  "The target is invalid.",
	roDescriptionNotLoaded:                      "The module is not loaded.",
	roDescriptionNotRegistered:                  "The module is not registered.",
	roDescriptionOutOfSpace:                     "There is no space.",
	roDescriptionRangeErrorAtExport:             "The range error occurs at export.",
	roDescriptionRangeErrorAtHeader:             "The range error occurs at header.",
	roDescriptionRangeErrorAtImport:             "The range error occurs at import.",
	roDescriptionRangeErrorAtIndexExport:        "The range error occurs at index export.",
	roDescriptionRangeErrorAtIndexImport:        "The range error occurs at index import.",
	roDescriptionRangeErrorAtInternalRelocation: "The range error occurs at internal relocation.",
	roDescriptionRangeErrorAtOffsetExport:       "The range error occurs at offset export.",
	roDescriptionRangeErrorAtOffsetImport:       "The range error occurs at offset import.",
	roDescriptionRangeErrorAtReference:          "The range error occurs at reference.",
	roDescriptionRangeErrorAtSection:            "The range error occurs at section.",
	roDescriptionRangeErrorAtSymbolExport:       "The range error occurs at symbol export.",
	roDescriptionRangeErrorAtSymbolImport:       "The range error occurs at symbol import.",
	roDescriptionRegistrationNotFound:           "The registration is not found.",
	roDescriptionRwNotSupported:                 "The RW is not supported.",
	roDescriptionStaticModule:                   "The module is static.",
	roDescriptionTooSmallSize:                   "The size is too small.",
	roDescriptionTooSmallTarget:                 "The target is too small.",
	roDescriptionUnknownObjectControl:           "The object control is unknown.",
	roDescriptionUnknownRelocationType:          "The relocation type is unknown.",
	roDescriptionVeneerRequired:                 "The veneer is required.",
	roDescriptionVerificationFailed:             "The verification failed.",
}

func (d roDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", roDescriptionToString[d], d, roDescriptionDescription[d])
}
