package nn

import (
	"fmt"
)

// Constants
const (
	// MASK_FAIL_BIT is a mask for the fail bit. (Most significant bit of u32)
	MASK_FAIL_BIT uint32 = 0x80000000

	// SIZE_DESCRIPTION is the size of the description field in the error code.
	SIZE_DESCRIPTION int32 = 10

	// SIZE_MODULE is the size of the module field in the error code.
	SIZE_MODULE int32 = 8

	// SIZE_RESERVE is the size of the reserved space in the error code.
	SIZE_RESERVE int32 = 3

	// SIZE_SUMMARY is the size of the summary field in the error code.
	SIZE_SUMMARY int32 = 6

	// SIZE_LEVEL is the size of the level field in the error code.
	SIZE_LEVEL int32 = 5

	// SHIFTS_DESCRIPTION is the number of shifts to the right to get the description field.
	SHIFTS_DESCRIPTION int32 = 0

	// SHIFTS_MODULE is the number of shifts to the right to get the module field.
	SHIFTS_MODULE int32 = SHIFTS_DESCRIPTION + SIZE_DESCRIPTION

	// SHIFTS_RESERVE is the number of shifts to the right to get the reserved space.
	SHIFTS_RESERVE int32 = SHIFTS_MODULE + SIZE_MODULE

	// SHIFTS_SUMMARY is the number of shifts to the right to get the summary field.
	SHIFTS_SUMMARY int32 = SHIFTS_RESERVE + SIZE_RESERVE

	// SHIFTS_LEVEL is the number of shifts to the right to get the level field.
	SHIFTS_LEVEL int32 = SHIFTS_SUMMARY + SIZE_SUMMARY

	// MASK_NEGATIVE_LEVEL is a mask for checking negative level.
	MASK_NEGATIVE_LEVEL uint32 = 0xFFFFFFE0 // Inline computation of this caused an overflow error which was why it was moved to a constant.

	// MASK_DESCRIPTION is a mask for the description field.
	MASK_DESCRIPTION uint32 = (^uint32(0)) >> (32 - SIZE_DESCRIPTION) << SHIFTS_DESCRIPTION

	// MASK_MODULE is a mask for the module field.
	MASK_MODULE uint32 = (^uint32(0)) >> (32 - SIZE_MODULE) << SHIFTS_MODULE

	// MASK_SUMMARY is a mask for the summary field.
	MASK_SUMMARY uint32 = (^uint32(0)) >> (32 - SIZE_SUMMARY) << SHIFTS_SUMMARY

	// MASK_LEVEL is a mask for the level field.
	MASK_LEVEL uint32 = (^uint32(0)) >> (32 - SIZE_LEVEL) << SHIFTS_LEVEL

	// MAX_DESCRIPTION is the maximum value of the description field.
	MAX_DESCRIPTION uint32 = (^uint32(0)) >> (32 - SIZE_DESCRIPTION)

	// MAX_MODULE is the maximum value of the module field.
	MAX_MODULE uint32 = (^uint32(0)) >> (32 - SIZE_MODULE)

	// MAX_SUMMARY is the maximum value of the summary field.
	MAX_SUMMARY uint32 = (^uint32(0)) >> (32 - SIZE_SUMMARY)

	// MAX_LEVEL is the maximum value of the level field.
	MAX_LEVEL uint32 = (^uint32(0)) >> (32 - SIZE_LEVEL)
)

// Level Enumerated type indicating the severity of an error. (Do not use this type for error handling.)
type Level int32

// Constants for the Level enum.
const (
	// LevelInfo Success. There is additional information available.
	LevelInfo Level = -iota + 1

	// LevelSuccess Success. There is no additional information available.
	LevelSuccess

	// LevelFatal Fatal system level error. Software recovery not possible. Jump to support center guide sequence.
	LevelFatal

	// LevelReset Unexpected error. The best course of action is to transtion to the reset sequence. Reinitializing the module is not likely to lead to recovery,
	// but it may be possible to recover by resetting.
	LevelReset

	// LevelReinit Error that requires the module to be reinitialized.
	LevelReinit

	// LevelUsage Programming error. Indicates a programming error that always occurs with the argument values that were specified, regardless of the particular
	// internal state of the library.
	LevelUsage

	// LevelPermanent Common error. You cannot try again.
	LevelPermanent

	// LevelTemporary Temporary failure. You can try again immediately with the same argument. Succeeds if the number of attempts is relatively low.
	LevelTemporary

	// LevelStatus Expected failure.
	LevelStatus

	// LevelEnd End of the list.
	LevelEnd
)

// Level to string map.
var levelToString = map[Level]string{
	LevelInfo:      "Info",
	LevelSuccess:   "Success",
	LevelFatal:     "Fatal",
	LevelReset:     "Reset",
	LevelReinit:    "Reinit",
	LevelUsage:     "Usage",
	LevelPermanent: "Permanent",
	LevelTemporary: "Temporary",
	LevelStatus:    "Status",
}

// Level description map.
var levelDescription = map[Level]string{
	LevelInfo:      "Success. There is additional information available.",
	LevelSuccess:   "Success. There is no additional information available.",
	LevelFatal:     "Fatal system level error. Software recovery not possible. Jump to support center guide sequence.",
	LevelReset:     "Unexpected error. The best course of action is to transtion to the reset sequence. Reinitializing the module is not likely to lead to recovery, but it may be possible to recover by resetting.",
	LevelReinit:    "Error that requires the module to be reinitialized.",
	LevelUsage:     "Programming error. Indicates a programming error that always occurs with the argument values that were specified, regardless of the particular internal state of the library.",
	LevelPermanent: "Common error. You cannot try again.",
	LevelTemporary: "Temporary failure. You can try again immediately with the same argument. Succeeds if the number of attempts is relatively low.",
	LevelStatus:    "Expected failure.",
}

// ToString returns the string representation of the Level.
func (level Level) ToString() string {
	return fmt.Sprintf("%s (%d) - %s", levelToString[level], level, levelDescription[level])
}

// Summary Enumerated type that indicates an overview of the errors. (Do not use this type for error handling.)
type Summary int32

// Constants for the Summary enum.
const (
	// SummarySuccess Succeeded.
	SummarySuccess Summary = iota

	// SummaryNothingHappened Nothing happens.
	SummaryNothingHappened

	// SummaryWouldBlock The process may be blocked.
	SummaryWouldBlock

	// SummaryOutOfResource The resources needed for the process could not be allocated.
	SummaryOutOfResource

	// SummaryNotFound The object does not exist.
	SummaryNotFound

	// SummaryInvalidState The requested process cannot be executed with the current internal state.
	SummaryInvalidState

	// SummaryNotSupported Not currently supported by the SDK.
	SummaryNotSupported

	// SummaryInvalidArgument The argument is invalid.
	SummaryInvalidArgument

	// SummaryWrongArgument A parameter other than the argument is invalid.
	SummaryWrongArgument

	// SummaryCancelled The process was cancelled.
	SummaryCancelled

	// SummaryStatusChanged The status changed. (Example: Internal status changed while the process was running.)
	SummaryStatusChanged

	// SummaryInternal Error that is used within the library.
	SummaryInternal

	// SummaryInvalidResultValue Not used.
	SummaryInvalidResultValue = Summary(MAX_SUMMARY)
)

// Summary to string map.
var summaryToString = map[Summary]string{
	SummarySuccess:         "Success",
	SummaryNothingHappened: "Nothing Happened",
	SummaryWouldBlock:      "Would Block",
	SummaryOutOfResource:   "Out Of Resource",
	SummaryNotFound:        "Not Found",
	SummaryInvalidState:    "Invalid State",
	SummaryNotSupported:    "Not Supported",
	SummaryInvalidArgument: "Invalid Argument",
	SummaryWrongArgument:   "Wrong Argument",
	SummaryCancelled:       "Cancelled",
	SummaryStatusChanged:   "Status Changed",
	SummaryInternal:        "Internal",
}

// Summary description map.
var summaryDescription = map[Summary]string{
	SummarySuccess:         "Succeeded.",
	SummaryNothingHappened: "Nothing happens.",
	SummaryWouldBlock:      "The process may be blocked.",
	SummaryOutOfResource:   "The resources needed for the process could not be allocated.",
	SummaryNotFound:        "The object does not exist.",
	SummaryInvalidState:    "The requested process cannot be executed with the current internal state.",
	SummaryNotSupported:    "Not currently supported by the SDK.",
	SummaryInvalidArgument: "The argument is invalid.",
	SummaryWrongArgument:   "A parameter other than the argument is invalid.",
	SummaryCancelled:       "The process was cancelled.",
	SummaryStatusChanged:   "The status changed. (Example: Internal status changed while the process was running.)",
	SummaryInternal:        "Error that is used within the library.",
}

// ToString returns the string representation of the Summary.
func (summary Summary) ToString() string {
	return fmt.Sprintf("%s (%d) - %s", summaryToString[summary], summary, summaryDescription[summary])
}

// Module Enumerated type that indicates the modules in which errors occurred. (Do not use this type for error handling.)
//
// The enumerated type definitions may change in the future.
//
// With the exception of ModuleApplication, these values are defined for use inside the SDK. Do not use.
type Module int32

// Constants for the Module enum.
const (
	ModuleCommon Module = iota
	ModuleKernel        // kern
	ModuleUtil
	ModuleFileServer   // fs srv
	ModuleLoaderServer // ldr srv
	ModuleTcb
	ModuleOs
	ModuleDbg
	ModuleDmnt
	ModulePdn
	ModuleGx
	ModuleI2c
	ModuleGpio
	ModuleDd
	ModuleCodec
	ModuleSpi
	ModulePxi
	ModuleFs
	ModuleDi
	ModuleHid
	ModuleCamera // cam
	ModulePi
	ModulePm
	ModulePmLow // pm low
	ModuleFsi
	ModuleSrv
	ModuleNdm
	ModuleNwm
	ModuleSocket // soc
	ModuleLdr    // ldr
	ModuleAcc
	ModuleRomFs
	ModuleAm
	ModuleHio
	ModuleUpdater // upd
	ModuleMic
	ModuleFnd
	ModuleMp
	ModuleMpWL // mp wl
	ModuleAc
	ModuleHttp
	ModuleDsp
	ModuleSnd
	ModuleDlp
	ModuleHioLow // hio low
	ModuleCsnd
	ModuleSsl
	ModuleAmLow // am low
	ModuleNex
	ModuleFriends // frd
	ModuleRdt
	ModuleApplet // app
	ModuleNim
	ModulePtm
	ModuleMidi
	ModuleMc
	ModuleSwc
	ModuleFatFs
	ModuleNgc
	ModuleCard    // card
	ModuleCardNor // card nor
	ModuleSdmc
	ModuleBoss
	ModuleDbm
	ModuleCfg
	ModulePs
	ModuleCec
	ModuleIr
	ModuleUds
	ModulePl
	ModuleCup
	ModuleGyroscope // gyro
	ModuleMcu
	ModuleNs
	ModuleNews
	ModuleRo
	ModuleGd
	ModuleCardSpi // card spi
	ModuleEc
	ModuleWebBrs // web brs
	ModuleTest
	ModuleEnc
	ModulePia
	ModuleAct
	ModuleVctl
	ModuleOlv
	ModuleNeia
	ModuleNpns

	ModuleReserved0
	ModuleReserved1

	ModuleAvd
	ModuleL2b
	ModuleMvd
	ModuleNfc
	ModuleUart
	ModuleQtm
	ModuleNfp
	ModuleNpt

	// ModuleApplication Value defined for use in application.
	ModuleApplication = Module(MAX_MODULE) - 1

	// ModuleInvalidResultValue Not used.
	ModuleInvalidResultValue = Module(MAX_MODULE)
)

// Module to string map.
var moduleToString = map[Module]string{
	ModuleCommon:       "cmn",
	ModuleKernel:       "kern",
	ModuleUtil:         "util",
	ModuleFileServer:   "fs srv",
	ModuleLoaderServer: "ldr srv",
	ModuleTcb:          "tcb",
	ModuleOs:           "os",
	ModuleDbg:          "dbg",
	ModuleDmnt:         "dmnt",
	ModulePdn:          "pdn",
	ModuleGx:           "gx",
	ModuleI2c:          "i2c",
	ModuleGpio:         "gpio",
	ModuleDd:           "dd",
	ModuleCodec:        "codec",
	ModuleSpi:          "spi",
	ModulePxi:          "pxi",
	ModuleFs:           "fs",
	ModuleDi:           "di",
	ModuleHid:          "hid",
	ModuleCamera:       "cam",
	ModulePi:           "pi",
	ModulePm:           "pm",
	ModulePmLow:        "pm low",
	ModuleFsi:          "fsi",
	ModuleSrv:          "srv",
	ModuleNdm:          "ndm",
	ModuleNwm:          "nwm",
	ModuleSocket:       "soc",
	ModuleLdr:          "ldr",
	ModuleAcc:          "acc",
	ModuleRomFs:        "romfs",
	ModuleAm:           "am",
	ModuleHio:          "hio",
	ModuleUpdater:      "upd",
	ModuleMic:          "mic",
	ModuleFnd:          "fnd",
	ModuleMp:           "mp",
	ModuleMpWL:         "mpwl",
	ModuleAc:           "ac",
	ModuleHttp:         "http",
	ModuleDsp:          "dsp",
	ModuleSnd:          "snd",
	ModuleDlp:          "dlp",
	ModuleHioLow:       "hio low",
	ModuleCsnd:         "csnd",
	ModuleSsl:          "ssl",
	ModuleAmLow:        "am low",
	ModuleNex:          "nex",
	ModuleFriends:      "frd",
	ModuleRdt:          "rdt",
	ModuleApplet:       "app",
	ModuleNim:          "nim",
	ModulePtm:          "ptm",
	ModuleMidi:         "midi",
	ModuleMc:           "mc",
	ModuleSwc:          "swc",
	ModuleFatFs:        "fatfs",
	ModuleNgc:          "ngc",
	ModuleCard:         "card",
	ModuleCardNor:      "card nor",
	ModuleSdmc:         "sdmc",
	ModuleBoss:         "boss",
	ModuleDbm:          "dbm",
	ModuleCfg:          "cfg",
	ModulePs:           "ps",
	ModuleCec:          "cec",
	ModuleIr:           "ir",
	ModuleUds:          "uds",
	ModulePl:           "pl",
	ModuleCup:          "cup",
	ModuleGyroscope:    "gyro",
	ModuleMcu:          "mcu",
	ModuleNs:           "ns",
	ModuleNews:         "news",
	ModuleRo:           "ro",
	ModuleGd:           "gd",
	ModuleCardSpi:      "card spi",
	ModuleEc:           "ec",
	ModuleWebBrs:       "web brs",
	ModuleTest:         "test",
	ModuleEnc:          "enc",
	ModulePia:          "pia",
	ModuleAct:          "act",
	ModuleVctl:         "vctl",
	ModuleOlv:          "olv",
	ModuleNeia:         "neia",
	ModuleNpns:         "npns",

	ModuleReserved0: "reserved0",
	ModuleReserved1: "reserved1",

	ModuleAvd:  "avd",
	ModuleL2b:  "l2b",
	ModuleMvd:  "mvd",
	ModuleNfc:  "nfc",
	ModuleUart: "uart",
	ModuleQtm:  "qtm",
	ModuleNfp:  "nfp",
	ModuleNpt:  "npt",

	Module(ModuleApplication): "application",
}

// Module description map. These descriptions are not official. Most are guesses of what the acronyms stand for.
var moduleDescription = map[Module]string{
	ModuleCommon:       "Common",
	ModuleKernel:       "Kernel",
	ModuleUtil:         "Util",
	ModuleFileServer:   "File Server",
	ModuleLoaderServer: "Loader Server",
	ModuleTcb:          "Tcb",
	ModuleOs:           "OS",
	ModuleDbg:          "Debug",
	ModuleDmnt:         "Debug Monitor",
	ModulePdn:          "Power Domain",
	ModuleGx:           "Graphics Driver",
	ModuleI2c:          "Inter-Integrated Circuit Driver",
	ModuleGpio:         "General Purpose Input/Output",
	ModuleDd:           "Daemon Driver",
	ModuleCodec:        "Codec Driver",
	ModuleSpi:          "Serial Peripheral Interface Driver",
	ModulePxi:          "Process eXecution Interface",
	ModuleFs:           "File System",
	ModuleDi:           "Device Interface",
	ModuleHid:          "Human Interface Device",
	ModuleCamera:       "Camera",
	ModulePi:           "Process Interface",
	ModulePm:           "Process Manager",
	ModulePmLow:        "Process Manager Low",
	ModuleFsi:          "File System Interface",
	ModuleSrv:          "Services",
	ModuleNdm:          "Network Daemon Manager",
	ModuleNwm:          "Network Wireless Manager",
	ModuleSocket:       "Socket",
	ModuleLdr:          "Loader",
	ModuleAcc:          "Account Management",
	ModuleRomFs:        "Read-Only File System",
	ModuleAm:           "Application Manager",
	ModuleHio:          "Host IO",
	ModuleUpdater:      "Updater",
	ModuleMic:          "Microphone",
	ModuleFnd:          "Foundation",
	ModuleMp:           "Multiplayer",
	ModuleMpWL:         "Multiplayer Wireless",
	ModuleAc:           "Automatic Connection",
	ModuleHttp:         "Hypertext Transfer Protocol",
	ModuleDsp:          "Digital Signal Processor",
	ModuleSnd:          "Sound",
	ModuleDlp:          "Download Play",
	ModuleHioLow:       "Host IO Low",
	ModuleCsnd:         "Codec Sound",
	ModuleSsl:          "Secure Sockets Layer",
	ModuleAmLow:        "Application Manager Low",
	ModuleNex:          "NEX Game Networking",
	ModuleFriends:      "Friends",
	ModuleRdt:          "Reliable Data Transport",
	ModuleApplet:       "Applet",
	ModuleNim:          "Nintendo Internet Module",
	ModulePtm:          "Power/Time Management",
	ModuleMidi:         "Musical Instrument Digital Interface",
	ModuleMc:           "Media Control",
	ModuleSwc:          "Software Controller",
	ModuleFatFs:        "File Allocation Table File System",
	ModuleNgc:          "Name Generation Control",
	ModuleCard:         "Card",
	ModuleCardNor:      "Card NOR",
	ModuleSdmc:         "Secure Digital Memory Card",
	ModuleBoss:         "Background Operating System Service",
	ModuleDbm:          "Database Manager",
	ModuleCfg:          "Configuration",
	ModulePs:           "Process Signatures",
	ModuleCec:          "Chance Encounter Communication",
	ModuleIr:           "Infrared",
	ModuleUds:          "User Datagram Service",
	ModulePl:           "Play Log",
	ModuleCup:          "Card Update",
	ModuleGyroscope:    "Gyroscope",
	ModuleMcu:          "Microcontroller",
	ModuleNs:           "Network System", // Another IS system module.
	ModuleNews:         "News",
	ModuleRo:           "Read-Only",
	ModuleGd:           "Graphics Data Flow",
	ModuleCardSpi:      "Card Serial Peripheral Interface",
	ModuleEc:           "Nintendo E-Shop",
	ModuleWebBrs:       "Web Browser",
	ModuleTest:         "Test",
	ModuleEnc:          "Encryption",
	ModulePia:          "PIA",
	ModuleAct:          "Account",
	ModuleVctl:         "Version Control",
	ModuleOlv:          "OLV",
	ModuleNeia:         "NEIA",
	ModuleNpns:         "NPNS",

	ModuleReserved0: "Reserved0",
	ModuleReserved1: "Reserved1",

	ModuleAvd:  "AVD",
	ModuleL2b:  "Line To Block",
	ModuleMvd:  "Movie Decoder",
	ModuleNfc:  "Near Field Communication",
	ModuleUart: "Universal Asynchronous Receiver/Transmitter",
	ModuleQtm:  "Quartet",
	ModuleNfp:  "Near Field Proximity",
	ModuleNpt:  "NPT",

	ModuleApplication: "Application",
}

// ToString returns the string representation of the Module.
func (module Module) ToString() string {
	return fmt.Sprintf("%s (%d) - %s", moduleToString[module], module, moduleDescription[module])
}

// Description Enumerated type that indicates the details of errors. (Do not use this type for error handling.)
//
// Negative values are common to all SDK libraries. Positive values are defined independently by each library.
type Description uint32

// Constants for the Description enum.
const (
	// DescriptionSuccess Succeeded.
	DescriptionSuccess Description = iota

	// DescriptionInvalidSelection An invalid value was specified when a specifiable value is discrete
	DescriptionInvalidSelection = Description(MAX_DESCRIPTION) - 23

	// DescriptionTooLarge The value is too large
	DescriptionTooLarge = Description(MAX_DESCRIPTION) - 22

	// DescriptionNotAuthorized Unauthorized operation
	DescriptionNotAuthorized = Description(MAX_DESCRIPTION) - 21

	// DescriptionAlreadyDone The internal status has already been specified
	DescriptionAlreadyDone = Description(MAX_DESCRIPTION) - 20

	// DescriptionInvalidSize Invalid size
	DescriptionInvalidSize = Description(MAX_DESCRIPTION) - 19

	// DescriptionInvalidEnumValue The value is outside the range for enum values
	DescriptionInvalidEnumValue = Description(MAX_DESCRIPTION) - 18

	// DescriptionInvalidCombination Invalid parameter combination
	DescriptionInvalidCombination = Description(MAX_DESCRIPTION) - 17

	// DescriptionNoData No data
	DescriptionNoData = Description(MAX_DESCRIPTION) - 16

	// DescriptionBusy Could not be run because another process was already being performed
	DescriptionBusy = Description(MAX_DESCRIPTION) - 15

	// DescriptionMisalignedAddress Invalid address alignment
	DescriptionMisalignedAddress = Description(MAX_DESCRIPTION) - 14

	// DescriptionMisalignedSize Invalid size alignment
	DescriptionMisalignedSize = Description(MAX_DESCRIPTION) - 13

	// DescriptionOutOfMemory Insufficient memory
	DescriptionOutOfMemory = Description(MAX_DESCRIPTION) - 12

	// DescriptionNotImplemented Not yet implemented
	DescriptionNotImplemented = Description(MAX_DESCRIPTION) - 11

	// DescriptionInvalidAddress Invalid address
	DescriptionInvalidAddress = Description(MAX_DESCRIPTION) - 10

	// DescriptionInvalidPointer Invalid pointer
	DescriptionInvalidPointer = Description(MAX_DESCRIPTION) - 9

	// DescriptionInvalidHandle Invalid handle
	DescriptionInvalidHandle = Description(MAX_DESCRIPTION) - 8

	// DescriptionNotInitialized Not initialized
	DescriptionNotInitialized = Description(MAX_DESCRIPTION) - 7

	// DescriptionAlreadyInitialized Already initialized
	DescriptionAlreadyInitialized = Description(MAX_DESCRIPTION) - 6

	// DescriptionNotFound The object does not exist
	DescriptionNotFound = Description(MAX_DESCRIPTION) - 5

	// DescriptionCancelRequested Request canceled
	DescriptionCancelRequested = Description(MAX_DESCRIPTION) - 4

	// DescriptionAlreadyExists The object already exists
	DescriptionAlreadyExists = Description(MAX_DESCRIPTION) - 3

	// DescriptionOutOfRange The value is outside of the defined range
	DescriptionOutOfRange = Description(MAX_DESCRIPTION) - 2

	// DescriptionTimeout The process timed out
	DescriptionTimeout = Description(MAX_DESCRIPTION) - 1

	// DescriptionInvalidResultValue These values are not used
	DescriptionInvalidResultValue = Description(MAX_DESCRIPTION) - 0
)

// Description to string map.
var descriptionToString = map[Description]string{
	DescriptionSuccess:            "Success",
	DescriptionInvalidSelection:   "Invalid Selection",
	DescriptionTooLarge:           "Too Large",
	DescriptionNotAuthorized:      "Not Authorized",
	DescriptionAlreadyDone:        "Already Done",
	DescriptionInvalidSize:        "Invalid Size",
	DescriptionInvalidEnumValue:   "Invalid Enum Value",
	DescriptionInvalidCombination: "Invalid Combination",
	DescriptionNoData:             "No Data",
	DescriptionBusy:               "Busy",
	DescriptionMisalignedAddress:  "Misaligned Address",
	DescriptionMisalignedSize:     "Misaligned Size",
	DescriptionOutOfMemory:        "Out Of Memory",
	DescriptionNotImplemented:     "Not Implemented",
	DescriptionInvalidAddress:     "Invalid Address",
	DescriptionInvalidPointer:     "Invalid Pointer",
	DescriptionInvalidHandle:      "Invalid Handle",
	DescriptionNotInitialized:     "Not Initialized",
	DescriptionAlreadyInitialized: "Already Initialized",
	DescriptionNotFound:           "Not Found",
	DescriptionCancelRequested:    "Cancel Requested",
	DescriptionAlreadyExists:      "Already Exists",
	DescriptionOutOfRange:         "Out Of Range",
	DescriptionTimeout:            "Timeout",
	DescriptionInvalidResultValue: "Invalid Result Value",
}

// Description description map.
var descriptionDescription = map[Description]string{
	DescriptionSuccess:            "Succeeded.",
	DescriptionInvalidSelection:   "An invalid value was specified when a specifiable value is discrete.",
	DescriptionTooLarge:           "The value is too large.",
	DescriptionNotAuthorized:      "Unauthorized operation.",
	DescriptionAlreadyDone:        "The internal status has already been specified.",
	DescriptionInvalidSize:        "Invalid size.",
	DescriptionInvalidEnumValue:   "The value is outside the range for enum values.",
	DescriptionInvalidCombination: "Invalid parameter combination.",
	DescriptionNoData:             "No data.",
	DescriptionBusy:               "Could not be run because another process was already being performed.",
	DescriptionMisalignedAddress:  "Invalid address alignment.",
	DescriptionMisalignedSize:     "Invalid size alignment.",
	DescriptionOutOfMemory:        "Insufficient memory.",
	DescriptionNotImplemented:     "Not yet implemented.",
	DescriptionInvalidAddress:     "Invalid address.",
	DescriptionInvalidPointer:     "Invalid pointer.",
	DescriptionInvalidHandle:      "Invalid handle.",
	DescriptionNotInitialized:     "Not initialized.",
	DescriptionAlreadyInitialized: "Already initialized.",
	DescriptionNotFound:           "The object does not exist.",
	DescriptionCancelRequested:    "Request canceled.",
	DescriptionAlreadyExists:      "The object already exists.",
	DescriptionOutOfRange:         "The value is outside of the defined range.",
	DescriptionTimeout:            "The process timed out.",
	DescriptionInvalidResultValue: "These values are not used.",
}

// ToString returns the string representation of the Description.
func (description Description) ToString() string {
	return fmt.Sprintf("%s (%d) - %s", descriptionToString[description], description, descriptionDescription[description])
}

// Result Represents the result of an operation.
type Result uint32

// NewResult Creates a new Result from the specified error code.
func NewResult(errorCode uint32) Result {
	return Result(errorCode)
}

// NewResultFromFields Creates a new Result from the specified fields.
func NewResultFromFields(level Level, summary Summary, module Module, description int32) Result {
	return Result(
		((uint32(level) << SHIFTS_LEVEL) & MASK_LEVEL) |
			((uint32(summary) << SHIFTS_SUMMARY) & MASK_SUMMARY) |
			((uint32(module) << SHIFTS_MODULE) & MASK_MODULE) |
			((uint32(description) << SHIFTS_DESCRIPTION) & MASK_DESCRIPTION))
}

func (result Result) getCodeBits(mask uint32, shift int32) uint32 {
	return (uint32(result) & mask) >> shift
}

// IsFailure Returns true if the Result is a failure.
func (result Result) IsFailure() bool {
	return (uint32(result) & MASK_FAIL_BIT) != 0
}

// IsSuccess Returns true if the Result is a success.
func (result Result) IsSuccess() bool {
	return !result.IsFailure()
}

// GetLevel Returns the Level of the Result.
func (result Result) GetLevel() Level {
	if (uint32(result) & MASK_FAIL_BIT) != 0 {
		return Level(result.getCodeBits(MASK_LEVEL, SHIFTS_LEVEL) | MASK_NEGATIVE_LEVEL)
	} else {
		return Level(result.getCodeBits(MASK_LEVEL, SHIFTS_LEVEL))
	}
}

// GetSummary Returns the Summary of the Result.
func (result Result) GetSummary() Summary {
	return Summary(result.getCodeBits(MASK_SUMMARY, SHIFTS_SUMMARY))
}

// GetModule Returns the Module of the Result.
func (result Result) GetModule() Module {
	return Module(result.getCodeBits(MASK_MODULE, SHIFTS_MODULE))
}

// GetDescription Returns the Description of the Result.
func (result Result) GetDescription() int32 {
	return int32(result.getCodeBits(MASK_DESCRIPTION, SHIFTS_DESCRIPTION))
}

// GetValue Returns the value of the Result.
func (result Result) GetValue() uint32 {
	return uint32(result)
}

// Equals Returns true if the Result is equal to the specified Result.
func (result Result) Equals(other Result) bool {
	return result == other
}

// ToString Returns the string representation of the Result.
func (result Result) ToString(withDescription bool) string {
	if result.IsFailure() {
		r := fmt.Sprintf("Failure (0x%X)\nLevel: %s\nSummary: %s\nModule: %s", result.GetValue(), result.GetLevel().ToString(), result.GetSummary().ToString(), result.GetModule().ToString())
		if withDescription {
			r += fmt.Sprintf("\nDescription: %s", Description(result.GetDescription()).ToString())
		}

		return r
	} else {
		r := fmt.Sprintf("Success (0x%X)\nLevel: %s\nSummary: %s\nModule: %s", result.GetValue(), result.GetLevel().ToString(), result.GetSummary().ToString(), result.GetModule().ToString())
		if withDescription {
			r += fmt.Sprintf("\nDescription: %s", Description(result.GetDescription()).ToString())
		}

		return r
	}
}

// ResultSuccess Represents a successful Result.
var ResultSuccess = NewResult(0)

// MakeInfoResult Creates a Result with the specified Summary and Description.
func MakeInfoResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelInfo, summary, module, description)
}

// MakeFatalResult Creates a Result with the specified Summary and Description.
func MakeFatalResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelFatal, summary, module, description)
}

// MakeResetResult Creates a Result with the specified Summary and Description.
func MakeResetResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelReset, summary, module, description)
}

// MakeReinitResult Creates a Result with the specified Summary and Description.
func MakeReinitResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelReinit, summary, module, description)
}

// MakeUsageResult Creates a Result with the specified Summary and Description.
func MakeUsageResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelUsage, summary, module, description)
}

// MakePermanentResult Creates a Result with the specified Summary and Description.
func MakePermanentResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelPermanent, summary, module, description)
}

// MakeTemporaryResult Creates a Result with the specified Summary and Description.
func MakeTemporaryResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelTemporary, summary, module, description)
}

// MakeStatusResult Creates a Result with the specified Summary and Description.
func MakeStatusResult(summary Summary, module Module, description int32) Result {
	return NewResultFromFields(LevelStatus, summary, module, description)
}
