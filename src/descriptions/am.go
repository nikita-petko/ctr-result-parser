package descriptions

import "fmt"

type amDescription int32

const (
	amDescriptionInvalidCiaFormat amDescription = iota + 1
	amDescriptionInvalidImportPipeState
	amDescriptionInvalidCiaReaderState
	amDescriptionInvalidImportState
	amDescriptionInvalidUpdateState
	amDescriptionTooLargeFileSize
	amDescriptionInvalidContent
	amDescriptionUnsupportedPipeOperation
	amDescriptionImportPipealreadyOpen
	amDescriptionInvalidPipe
	amDescriptionInvalidFirmFormat
	amDescriptionFailedToVerifyFirmware
	amDescriptionFirmwareHashUnmatched
	amDescriptionInvalidFirmwareProgramId
	amDescriptionImportNotBegun
	amDescriptionUnexpectedEsError
	amDescriptionNoSuchDbEntry
	amDescriptionImportIncompleted
	amDescriptionCiaMetaNotFound
	amDescriptionTooManySessions
	amDescriptionInvalidTmdInCia
	amDescriptionInvalidSystemMenuDataSize
	amDescriptionFailedToAddTicketDb
	amDescriptionFailedToGetTicketDb
	amDescriptionInvalidTicketVersion
	amDescriptionNoEsCertifcatesFound
	amDescriptionTooManyCertifcates
	amDescriptionNotFoundJumpId
	amDescriptionImportContextNotFound
	amDescriptionMediaNotSupported
	amDescriptionInvalidArgument
	amDescriptionInvalidInputStream
	amDescriptionInvalidDbFormat
	amDescriptionInternalDataCorrupted
	amDescriptionFilesystemError
	amDescriptionEsError
	amDescriptionVerificationFailed
	amDescriptionNotResumable
	amDescriptionInvalidVersion
	amDescriptionNotSupported
	amDescriptionUnknownEsError
	amDescriptionInvalidSaveDataSize
	amDescriptionDatabaseNotFound
	amDescriptionProgramNotDeletable
	amDescriptionInvalidDatabaseState
	amDescriptionUniqueIdMatch
	amDescriptionUniqueIdMismatch
	amDescriptionHashCheckFailed
	amDescriptionDecryptionFailed
	amDescriptionBkpWriteFailed
	amDescriptionBkpReadFailed
	amDescriptionInvalidBkpFormat
	amDescriptionNoEnoughSpace
	amDescriptionTicketDbOverflow
	amDescriptionInvalidTicketTitleId
	amDescriptionMainContentNotFound
	amDescriptionNoRight
	amDescriptionInvalidTicketSize
	amDescriptionContentNotFound
	amDescriptionInvalidProgramId
	amDescriptionUnsupportedNumContents
	amDescriptionUndeletableContents
	amDescriptionBufferNotEnough
	amDescriptionTicketNotFound
	amDescriptionInvalidTitleCombination
	amDescriptionContextVerificatonFailed
	amDescriptionUnsupportedContextVersion

	amDescriptionEsOk = iota + 100 - 67
	amDescriptionEsFail
	amDescriptionEsInvalid
	amDescriptionEsStorage
	amDescriptionEsStorageSize
	amDescriptionEsCrypto
	amDescriptionEsVerification
	amDescriptionEsDeviceIdMismatch
	amDescriptionEsIssuerNotFound
	amDescriptionEsIncorrectSigType
	amDescriptionEsIncorrectPubkeyType
	amDescriptionEsIncorrectTicketVersion
	amDescriptionEsIncorrectTmdVersion
	amDescriptionEsNoRight
	amDescriptionEsAlignment
)

var amDescriptionToString = map[amDescription]string{
	amDescriptionInvalidCiaFormat:          "InvalidCiaFormat",
	amDescriptionInvalidImportPipeState:    "InvalidImportPipeState",
	amDescriptionInvalidCiaReaderState:     "InvalidCiaReaderState",
	amDescriptionInvalidImportState:        "InvalidImportState",
	amDescriptionInvalidUpdateState:        "InvalidUpdateState",
	amDescriptionTooLargeFileSize:          "TooLargeFileSize",
	amDescriptionInvalidContent:            "InvalidContent",
	amDescriptionUnsupportedPipeOperation:  "UnsupportedPipeOperation",
	amDescriptionImportPipealreadyOpen:     "ImportPipealreadyOpen",
	amDescriptionInvalidPipe:               "InvalidPipe",
	amDescriptionInvalidFirmFormat:         "InvalidFirmFormat",
	amDescriptionFailedToVerifyFirmware:    "FailedToVerifyFirmware",
	amDescriptionFirmwareHashUnmatched:     "FirmwareHashUnmatched",
	amDescriptionInvalidFirmwareProgramId:  "InvalidFirmwareProgramId",
	amDescriptionImportNotBegun:            "ImportNotBegun",
	amDescriptionUnexpectedEsError:         "UnexpectedEsError",
	amDescriptionNoSuchDbEntry:             "NoSuchDbEntry",
	amDescriptionImportIncompleted:         "ImportIncompleted",
	amDescriptionCiaMetaNotFound:           "CiaMetaNotFound",
	amDescriptionTooManySessions:           "TooManySessions",
	amDescriptionInvalidTmdInCia:           "InvalidTmdInCia",
	amDescriptionInvalidSystemMenuDataSize: "InvalidSystemMenuDataSize",
	amDescriptionFailedToAddTicketDb:       "FailedToAddTicketDb",
	amDescriptionFailedToGetTicketDb:       "FailedToGetTicketDb",
	amDescriptionInvalidTicketVersion:      "InvalidTicketVersion",
	amDescriptionNoEsCertifcatesFound:      "NoEsCertifcatesFound",
	amDescriptionTooManyCertifcates:        "TooManyCertifcates",
	amDescriptionNotFoundJumpId:            "NotFoundJumpId",
	amDescriptionImportContextNotFound:     "ImportContextNotFound",
	amDescriptionMediaNotSupported:         "MediaNotSupported",
	amDescriptionInvalidArgument:           "InvalidArgument",
	amDescriptionInvalidInputStream:        "InvalidInputStream",
	amDescriptionInvalidDbFormat:           "InvalidDbFormat",
	amDescriptionInternalDataCorrupted:     "InternalDataCorrupted",
	amDescriptionFilesystemError:           "FilesystemError",
	amDescriptionEsError:                   "EsError",
	amDescriptionVerificationFailed:        "VerificationFailed",
	amDescriptionNotResumable:              "NotResumable",
	amDescriptionInvalidVersion:            "InvalidVersion",
	amDescriptionNotSupported:              "NotSupported",
	amDescriptionUnknownEsError:            "UnknownEsError",
	amDescriptionInvalidSaveDataSize:       "InvalidSaveDataSize",
	amDescriptionDatabaseNotFound:          "DatabaseNotFound",
	amDescriptionProgramNotDeletable:       "ProgramNotDeletable",
	amDescriptionInvalidDatabaseState:      "InvalidDatabaseState",
	amDescriptionUniqueIdMatch:             "UniqueIdMatch",
	amDescriptionUniqueIdMismatch:          "UniqueIdMismatch",
	amDescriptionHashCheckFailed:           "HashCheckFailed",
	amDescriptionDecryptionFailed:          "DecryptionFailed",
	amDescriptionBkpWriteFailed:            "BkpWriteFailed",
	amDescriptionBkpReadFailed:             "BkpReadFailed",
	amDescriptionInvalidBkpFormat:          "InvalidBkpFormat",
	amDescriptionNoEnoughSpace:             "NoEnoughSpace",
	amDescriptionTicketDbOverflow:          "TicketDbOverflow",
	amDescriptionInvalidTicketTitleId:      "InvalidTicketTitleId",
	amDescriptionMainContentNotFound:       "MainContentNotFound",
	amDescriptionNoRight:                   "NoRight",
	amDescriptionInvalidTicketSize:         "InvalidTicketSize",
	amDescriptionContentNotFound:           "ContentNotFound",
	amDescriptionInvalidProgramId:          "InvalidProgramId",
	amDescriptionUnsupportedNumContents:    "UnsupportedNumContents",
	amDescriptionUndeletableContents:       "UndeletableContents",
	amDescriptionBufferNotEnough:           "BufferNotEnough",
	amDescriptionTicketNotFound:            "TicketNotFound",
	amDescriptionInvalidTitleCombination:   "InvalidTitleCombination",
	amDescriptionContextVerificatonFailed:  "ContextVerificatonFailed",
	amDescriptionUnsupportedContextVersion: "UnsupportedContextVersion",
	amDescriptionEsOk:                      "EsOk",
	amDescriptionEsFail:                    "EsFail",
	amDescriptionEsInvalid:                 "EsInvalid",
	amDescriptionEsStorage:                 "EsStorage",
	amDescriptionEsStorageSize:             "EsStorageSize",
	amDescriptionEsCrypto:                  "EsCrypto",
	amDescriptionEsVerification:            "EsVerification",
	amDescriptionEsDeviceIdMismatch:        "EsDeviceIdMismatch",
	amDescriptionEsIssuerNotFound:          "EsIssuerNotFound",
	amDescriptionEsIncorrectSigType:        "EsIncorrectSigType",
	amDescriptionEsIncorrectPubkeyType:     "EsIncorrectPubkeyType",
	amDescriptionEsIncorrectTicketVersion:  "EsIncorrectTicketVersion",
	amDescriptionEsIncorrectTmdVersion:     "EsIncorrectTmdVersion",
	amDescriptionEsNoRight:                 "EsNoRight",
	amDescriptionEsAlignment:               "EsAlignment",
}

var amDescriptionDescription = map[amDescription]string{
	amDescriptionInvalidCiaFormat:          "Invalid CIA format",
	amDescriptionInvalidImportPipeState:    "Invalid import pipe state",
	amDescriptionInvalidCiaReaderState:     "Invalid CIA reader state",
	amDescriptionInvalidImportState:        "Invalid import state",
	amDescriptionInvalidUpdateState:        "Invalid update state",
	amDescriptionTooLargeFileSize:          "Too large file size",
	amDescriptionInvalidContent:            "Invalid content",
	amDescriptionUnsupportedPipeOperation:  "Unsupported pipe operation",
	amDescriptionImportPipealreadyOpen:     "Import pipe already open",
	amDescriptionInvalidPipe:               "Invalid pipe",
	amDescriptionInvalidFirmFormat:         "Invalid firm format",
	amDescriptionFailedToVerifyFirmware:    "Failed to verify firmware",
	amDescriptionFirmwareHashUnmatched:     "Firmware hash unmatched",
	amDescriptionInvalidFirmwareProgramId:  "Invalid firmware program ID",
	amDescriptionImportNotBegun:            "Import not begun",
	amDescriptionUnexpectedEsError:         "Unexpected ES error",
	amDescriptionNoSuchDbEntry:             "No such DB entry",
	amDescriptionImportIncompleted:         "Import incompleted",
	amDescriptionCiaMetaNotFound:           "CIA meta not found",
	amDescriptionTooManySessions:           "Too many sessions",
	amDescriptionInvalidTmdInCia:           "Invalid TMD in CIA",
	amDescriptionInvalidSystemMenuDataSize: "Invalid system menu data size",
	amDescriptionFailedToAddTicketDb:       "Failed to add ticket DB",
	amDescriptionFailedToGetTicketDb:       "Failed to get ticket DB",
	amDescriptionInvalidTicketVersion:      "Invalid ticket version",
	amDescriptionNoEsCertifcatesFound:      "No ES certifcates found",
	amDescriptionTooManyCertifcates:        "Too many certifcates",
	amDescriptionNotFoundJumpId:            "Not found jump ID",
	amDescriptionImportContextNotFound:     "Import context not found",
	amDescriptionMediaNotSupported:         "Media not supported",
	amDescriptionInvalidArgument:           "Invalid argument",
	amDescriptionInvalidInputStream:        "Invalid input stream",
	amDescriptionInvalidDbFormat:           "Invalid DB format",
	amDescriptionInternalDataCorrupted:     "Internal data corrupted",
	amDescriptionFilesystemError:           "Filesystem error",
	amDescriptionEsError:                   "ES error",
	amDescriptionVerificationFailed:        "Verification failed",
	amDescriptionNotResumable:              "Not resumable",
	amDescriptionInvalidVersion:            "Invalid version",
	amDescriptionNotSupported:              "Not supported",
	amDescriptionUnknownEsError:            "Unknown ES error",
	amDescriptionInvalidSaveDataSize:       "Invalid save data size",
	amDescriptionDatabaseNotFound:          "Database not found",
	amDescriptionProgramNotDeletable:       "Program not deletable",
	amDescriptionInvalidDatabaseState:      "Invalid database state",
	amDescriptionUniqueIdMatch:             "Unique ID match",
	amDescriptionUniqueIdMismatch:          "Unique ID mismatch",
	amDescriptionHashCheckFailed:           "Hash check failed",
	amDescriptionDecryptionFailed:          "Decryption failed",
	amDescriptionBkpWriteFailed:            "BKP write failed",
	amDescriptionBkpReadFailed:             "BKP read failed",
	amDescriptionInvalidBkpFormat:          "Invalid BKP format",
	amDescriptionNoEnoughSpace:             "No enough space",
	amDescriptionTicketDbOverflow:          "Ticket DB overflow",
	amDescriptionInvalidTicketTitleId:      "Invalid ticket title ID",
	amDescriptionMainContentNotFound:       "Main content not found",
	amDescriptionNoRight:                   "No right",
	amDescriptionInvalidTicketSize:         "Invalid ticket size",
	amDescriptionContentNotFound:           "Content not found",
	amDescriptionInvalidProgramId:          "Invalid program ID",
	amDescriptionUnsupportedNumContents:    "Unsupported num contents",
	amDescriptionUndeletableContents:       "Undeletable contents",
	amDescriptionBufferNotEnough:           "Buffer not enough",
	amDescriptionTicketNotFound:            "Ticket not found",
	amDescriptionInvalidTitleCombination:   "Invalid title combination",
	amDescriptionContextVerificatonFailed:  "Context verificaton failed",
	amDescriptionUnsupportedContextVersion: "Unsupported context version",
	amDescriptionEsOk:                      "ES OK",
	amDescriptionEsFail:                    "ES fail",
	amDescriptionEsInvalid:                 "ES invalid",
	amDescriptionEsStorage:                 "ES storage",
	amDescriptionEsStorageSize:             "ES storage size",
	amDescriptionEsCrypto:                  "ES crypto",
	amDescriptionEsVerification:            "ES verification",
	amDescriptionEsDeviceIdMismatch:        "ES device ID mismatch",
	amDescriptionEsIssuerNotFound:          "ES issuer not found",
	amDescriptionEsIncorrectSigType:        "ES incorrect sig type",
	amDescriptionEsIncorrectPubkeyType:     "ES incorrect pubkey type",
	amDescriptionEsIncorrectTicketVersion:  "ES incorrect ticket version",
	amDescriptionEsIncorrectTmdVersion:     "ES incorrect TMD version",
	amDescriptionEsNoRight:                 "ES no right",
	amDescriptionEsAlignment:               "ES alignment",
}

func (d amDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", amDescriptionToString[d], d, amDescriptionDescription[d])
}
