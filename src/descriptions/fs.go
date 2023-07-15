package descriptions

import "fmt"

type fsDescription int32

const (
	fsDescriptionFindFinished         fsDescription = 10
	fsDescriptionDbmFindFinished      fsDescription = 20
	fsDescriptionDbmFindKeyFinished   fsDescription = 21
	fsDescriptionDbmIterationFinished fsDescription = 22

	fsDescriptionNotFound                fsDescription = 100
	fsDescriptionArchiveNotFound         fsDescription = 101
	fsDescriptionCardPartitionNotFound   fsDescription = 102
	fsDescriptionDbmNotFound             fsDescription = 110
	fsDescriptionDbmKeyNotFound          fsDescription = 111
	fsDescriptionDbmFileNotFound         fsDescription = 112
	fsDescriptionDbmDirectoryNotFound    fsDescription = 113
	fsDescriptionDbmIdNotFound           fsDescription = 114
	fsDescriptionFatNotFound             fsDescription = 120
	fsDescriptionMediaNotFound           fsDescription = 130
	fsDescriptionFatNoStorage            fsDescription = 135
	fsDescriptionCardRomNotFound         fsDescription = 140
	fsDescriptionCardRomNoDevice         fsDescription = 141
	fsDescriptionCardRomUnkown           fsDescription = 142
	fsDescriptionCardBackupNotFound      fsDescription = 150
	fsDescriptionCardBackupType1NoDevice fsDescription = 151
	fsDescriptionCardBackupType2NoDevice fsDescription = 152
	fsDescriptionCardBackupType1Unkown   fsDescription = 153
	fsDescriptionCardBackupType2Unkown   fsDescription = 154
	fsDescriptionSdmcNotFound            fsDescription = 170
	fsDescriptionSdmcNoDevice            fsDescription = 171
	fsDescriptionSdmcUnkown              fsDescription = 172

	fsDescriptionAlreadyExists    fsDescription = 180
	fsDescriptionDbmAlreadyExists fsDescription = 185
	fsDescriptionFatAlreadyExists fsDescription = 190

	fsDescriptionNotEnoughSpace    fsDescription = 200
	fsDescriptionDbmNotEnoughSpace fsDescription = 205
	fsDescriptionFatNotEnoughSpace fsDescription = 210

	fsDescriptionArchiveInvalidated fsDescription = 220

	fsDescriptionOperationDenied          fsDescription = 230
	fsDescriptionResourceLocked           fsDescription = 231
	fsDescriptionCacheError               fsDescription = 235
	fsDescriptionDbmOperationDenied       fsDescription = 240
	fsDescriptionFatOperationDenied       fsDescription = 250
	fsDescriptionWriteProtected           fsDescription = 260
	fsDescriptionSdmcWriteProtected       fsDescription = 265
	fsDescriptionCardWriteProtected       fsDescription = 270
	fsDescriptionMediaAccessError         fsDescription = 280
	fsDescriptionCardRomAccessError       fsDescription = 290
	fsDescriptionCardRomCommError         fsDescription = 291
	fsDescriptionCardBackupCommError      fsDescription = 300
	fsDescriptionCardBackupType1CommError fsDescription = 301
	fsDescriptionCardBackupType2CommError fsDescription = 302
	fsDescriptionNandAccessError          fsDescription = 320
	fsDescriptionNandCommError            fsDescription = 321
	fsDescriptionNandEccFailed            fsDescription = 322
	fsDescriptionSdmcAccessError          fsDescription = 330
	fsDescriptionSdmcCommError            fsDescription = 331
	fsDescriptionSdmcEccFailed            fsDescription = 332
	fsDescriptionBackupNotRequired        fsDescription = 335

	fsDescriptionNotFormatted            fsDescription = 340
	fsDescriptionFormatIsFactoryDefaults fsDescription = 341
	fsDescriptionDbmNotFormatted         fsDescription = 350
	fsDescriptionDbmVersionMismatch      fsDescription = 351

	fsDescriptionBadFormat                fsDescription = 360
	fsDescriptionFatBadFormat             fsDescription = 370
	fsDescriptionFatBrokenEntry           fsDescription = 371
	fsDescriptionCardBackupBadFormat      fsDescription = 380
	fsDescriptionCardBackupType1BadFormat fsDescription = 381
	fsDescriptionCardBackupType2BadFormat fsDescription = 382

	fsDescriptionVerificationFailed    fsDescription = 390
	fsDescriptionHashMismatch          fsDescription = 391
	fsDescriptionVerifySignatureFailed fsDescription = 392
	fsDescriptionInvalidFileFormat     fsDescription = 393
	fsDescriptionVerifyHeaderMacFailed fsDescription = 394
	fsDescriptionVerifyDataHashFailed  fsDescription = 395

	fsDescriptionProgramDataNotFound    fsDescription = 560
	fsDescriptionProgramNotFound        fsDescription = 561
	fsDescriptionSystemMenuDataNotFound fsDescription = 562
	fsDescriptionBannerDataNotFound     fsDescription = 563
	fsDescriptionLogoDataNotFound       fsDescription = 564
	fsDescriptionProgramInfoNotFound    fsDescription = 565
	fsDescriptionContentNotFound        fsDescription = 566
	fsDescriptionNoSuchExeFsSection     fsDescription = 567

	fsDescriptionRequestRetry      fsDescription = 580
	fsDescriptionRequestSameRetry  fsDescription = 581
	fsDescriptionRequestSmallRetry fsDescription = 582

	fsDescriptionInvalidProgramFormat fsDescription = 590
	fsDescriptionInvalidCxiFormat     fsDescription = 591
	fsDescriptionInvalidCciFormat     fsDescription = 592

	fsDescriptionOutOfResource         fsDescription = 600
	fsDescriptionOutOfMemory           fsDescription = 601
	fsDescriptionHandleTableFull       fsDescription = 602
	fsDescriptionDbmOutOfResource      fsDescription = 610
	fsDescriptionDbmTableFull          fsDescription = 611
	fsDescriptionDbmFileEntryFull      fsDescription = 612
	fsDescriptionDbmDirectoryEntryFull fsDescription = 613
	fsDescriptionDbmOutOfMemory        fsDescription = 614
	fsDescriptionFatOutOfResource      fsDescription = 620
	fsDescriptionFatOutOfMemory        fsDescription = 621

	fsDescriptionAccessDenied     fsDescription = 630
	fsDescriptionNotDevelopmentId fsDescription = 631
	fsDescriptionNoContentRights  fsDescription = 632
	fsDescriptionDbmAccessDenied  fsDescription = 640
	fsDescriptionFatAccessDenied  fsDescription = 650

	fsDescriptionInvalidArgument         fsDescription = 700
	fsDescriptionInvalidPositionBase     fsDescription = 701
	fsDescriptionInvalidPathFormat       fsDescription = 702
	fsDescriptionPathTooLong             fsDescription = 703
	fsDescriptionOutOfBounds             fsDescription = 704
	fsDescriptionInvalidPosition         fsDescription = 705
	fsDescriptionDbmInvalidArgument      fsDescription = 710
	fsDescriptionDbmFileNameTooLong      fsDescription = 711
	fsDescriptionDbmDirectoryNameTooLong fsDescription = 712
	fsDescriptionDbmInvalidPathFormat    fsDescription = 713
	fsDescriptionDbmOutOfBounds          fsDescription = 714
	fsDescriptionFatInvalidArgument      fsDescription = 720
	fsDescriptionFatInvalidSize          fsDescription = 721

	fsDescriptionNotInitialized          fsDescription = 730
	fsDescriptionLibraryNotInitialized   fsDescription = 731
	fsDescriptionMovableDataIsNotEnabled fsDescription = 732
	fsDescriptionDbmNotInitialized       fsDescription = 735
	fsDescriptionCardRomNotInitialized   fsDescription = 740

	fsDescriptionAlreadyInitialized fsDescription = 750

	fsDescriptionUnsupportedOperation fsDescription = 760
	fsDescriptionUnsupportedAlignment fsDescription = 761
	fsDescriptionUnsupportedMedia     fsDescription = 762
	fsDescriptionDbmInvalidOperation  fsDescription = 770

	fsDescriptionNotImplemented    fsDescription = 900
	fsDescriptionInvalidKeyType    fsDescription = 901
	fsDescriptionPartitionNotFound fsDescription = 902

	fsDescriptionFatError      fsDescription = 910
	fsDescriptionFatCorruption fsDescription = 911

	fsDescriptionNandError fsDescription = 920
	fsDescriptionNandFatal fsDescription = 921

	fsDescriptionUnknownError fsDescription = 930
)

var fsDescriptionToString = map[fsDescription]string{
	fsDescriptionFindFinished:         "FindFinished",
	fsDescriptionDbmFindFinished:      "DbmFindFinished",
	fsDescriptionDbmFindKeyFinished:   "DbmFindKeyFinished",
	fsDescriptionDbmIterationFinished: "DbmIterationFinished",

	fsDescriptionNotFound:                "NotFound",
	fsDescriptionArchiveNotFound:         "ArchiveNotFound",
	fsDescriptionCardPartitionNotFound:   "CardPartitionNotFound",
	fsDescriptionDbmNotFound:             "DbmNotFound",
	fsDescriptionDbmKeyNotFound:          "DbmKeyNotFound",
	fsDescriptionDbmFileNotFound:         "DbmFileNotFound",
	fsDescriptionDbmDirectoryNotFound:    "DbmDirectoryNotFound",
	fsDescriptionDbmIdNotFound:           "DbmIdNotFound",
	fsDescriptionFatNotFound:             "FatNotFound",
	fsDescriptionMediaNotFound:           "MediaNotFound",
	fsDescriptionFatNoStorage:            "FatNoStorage",
	fsDescriptionCardRomNotFound:         "CardRomNotFound",
	fsDescriptionCardRomNoDevice:         "CardRomNoDevice",
	fsDescriptionCardRomUnkown:           "CardRomUnkown",
	fsDescriptionCardBackupNotFound:      "CardBackupNotFound",
	fsDescriptionCardBackupType1NoDevice: "CardBackupType1NoDevice",
	fsDescriptionCardBackupType2NoDevice: "CardBackupType2NoDevice",
	fsDescriptionCardBackupType1Unkown:   "CardBackupType1Unkown",
	fsDescriptionCardBackupType2Unkown:   "CardBackupType2Unkown",
	fsDescriptionSdmcNotFound:            "SdmcNotFound",
	fsDescriptionSdmcNoDevice:            "SdmcNoDevice",
	fsDescriptionSdmcUnkown:              "SdmcUnkown",

	fsDescriptionAlreadyExists:    "AlreadyExists",
	fsDescriptionDbmAlreadyExists: "DbmAlreadyExists",
	fsDescriptionFatAlreadyExists: "FatAlreadyExists",

	fsDescriptionNotEnoughSpace:    "NotEnoughSpace",
	fsDescriptionDbmNotEnoughSpace: "DbmNotEnoughSpace",
	fsDescriptionFatNotEnoughSpace: "FatNotEnoughSpace",

	fsDescriptionArchiveInvalidated: "ArchiveInvalidated",

	fsDescriptionOperationDenied:          "OperationDenied",
	fsDescriptionResourceLocked:           "ResourceLocked",
	fsDescriptionCacheError:               "CacheError",
	fsDescriptionDbmOperationDenied:       "DbmOperationDenied",
	fsDescriptionFatOperationDenied:       "FatOperationDenied",
	fsDescriptionWriteProtected:           "WriteProtected",
	fsDescriptionSdmcWriteProtected:       "SdmcWriteProtected",
	fsDescriptionCardWriteProtected:       "CardWriteProtected",
	fsDescriptionMediaAccessError:         "MediaAccessError",
	fsDescriptionCardRomAccessError:       "CardRomAccessError",
	fsDescriptionCardRomCommError:         "CardRomCommError",
	fsDescriptionCardBackupCommError:      "CardBackupCommError",
	fsDescriptionCardBackupType1CommError: "CardBackupType1CommError",
	fsDescriptionCardBackupType2CommError: "CardBackupType2CommError",
	fsDescriptionNandAccessError:          "NandAccessError",
	fsDescriptionNandCommError:            "NandCommError",
	fsDescriptionNandEccFailed:            "NandEccFailed",
	fsDescriptionSdmcAccessError:          "SdmcAccessError",
	fsDescriptionSdmcCommError:            "SdmcCommError",
	fsDescriptionSdmcEccFailed:            "SdmcEccFailed",
	fsDescriptionBackupNotRequired:        "BackupNotRequired",

	fsDescriptionNotFormatted:            "NotFormatted",
	fsDescriptionFormatIsFactoryDefaults: "FormatIsFactoryDefaults",
	fsDescriptionDbmNotFormatted:         "DbmNotFormatted",
	fsDescriptionDbmVersionMismatch:      "DbmVersionMismatch",

	fsDescriptionBadFormat:                "BadFormat",
	fsDescriptionFatBadFormat:             "FatBadFormat",
	fsDescriptionFatBrokenEntry:           "FatBrokenEntry",
	fsDescriptionCardBackupBadFormat:      "CardBackupBadFormat",
	fsDescriptionCardBackupType1BadFormat: "CardBackupType1BadFormat",
	fsDescriptionCardBackupType2BadFormat: "CardBackupType2BadFormat",

	fsDescriptionVerificationFailed:    "VerificationFailed",
	fsDescriptionHashMismatch:          "HashMismatch",
	fsDescriptionVerifySignatureFailed: "VerifySignatureFailed",
	fsDescriptionInvalidFileFormat:     "InvalidFileFormat",
	fsDescriptionVerifyHeaderMacFailed: "VerifyHeaderMacFailed",
	fsDescriptionVerifyDataHashFailed:  "VerifyDataHashFailed",

	fsDescriptionProgramDataNotFound:    "ProgramDataNotFound",
	fsDescriptionProgramNotFound:        "ProgramNotFound",
	fsDescriptionSystemMenuDataNotFound: "SystemMenuDataNotFound",
	fsDescriptionBannerDataNotFound:     "BannerDataNotFound",
	fsDescriptionLogoDataNotFound:       "LogoDataNotFound",
	fsDescriptionProgramInfoNotFound:    "ProgramInfoNotFound",
	fsDescriptionContentNotFound:        "ContentNotFound",
	fsDescriptionNoSuchExeFsSection:     "NoSuchExeFsSection",

	fsDescriptionRequestRetry:      "RequestRetry",
	fsDescriptionRequestSameRetry:  "RequestSameRetry",
	fsDescriptionRequestSmallRetry: "RequestSmallRetry",

	fsDescriptionInvalidProgramFormat: "InvalidProgramFormat",
	fsDescriptionInvalidCxiFormat:     "InvalidCxiFormat",
	fsDescriptionInvalidCciFormat:     "InvalidCciFormat",

	fsDescriptionOutOfResource:         "OutOfResource",
	fsDescriptionOutOfMemory:           "OutOfMemory",
	fsDescriptionHandleTableFull:       "HandleTableFull",
	fsDescriptionDbmOutOfResource:      "DbmOutOfResource",
	fsDescriptionDbmTableFull:          "DbmTableFull",
	fsDescriptionDbmFileEntryFull:      "DbmFileEntryFull",
	fsDescriptionDbmDirectoryEntryFull: "DbmDirectoryEntryFull",
	fsDescriptionDbmOutOfMemory:        "DbmOutOfMemory",
	fsDescriptionFatOutOfResource:      "FatOutOfResource",
	fsDescriptionFatOutOfMemory:        "FatOutOfMemory",

	fsDescriptionAccessDenied:     "AccessDenied",
	fsDescriptionNotDevelopmentId: "NotDevelopmentId",
	fsDescriptionNoContentRights:  "NoContentRights",
	fsDescriptionDbmAccessDenied:  "DbmAccessDenied",
	fsDescriptionFatAccessDenied:  "FatAccessDenied",

	fsDescriptionInvalidArgument:         "InvalidArgument",
	fsDescriptionInvalidPositionBase:     "InvalidPositionBase",
	fsDescriptionInvalidPathFormat:       "InvalidPathFormat",
	fsDescriptionPathTooLong:             "PathTooLong",
	fsDescriptionOutOfBounds:             "OutOfBounds",
	fsDescriptionInvalidPosition:         "InvalidPosition",
	fsDescriptionDbmInvalidArgument:      "DbmInvalidArgument",
	fsDescriptionDbmFileNameTooLong:      "DbmFileNameTooLong",
	fsDescriptionDbmDirectoryNameTooLong: "DbmDirectoryNameTooLong",
	fsDescriptionDbmInvalidPathFormat:    "DbmInvalidPathFormat",
	fsDescriptionDbmOutOfBounds:          "DbmOutOfBounds",
	fsDescriptionFatInvalidArgument:      "FatInvalidArgument",
	fsDescriptionFatInvalidSize:          "FatInvalidSize",

	fsDescriptionNotInitialized:          "NotInitialized",
	fsDescriptionLibraryNotInitialized:   "LibraryNotInitialized",
	fsDescriptionMovableDataIsNotEnabled: "MovableDataIsNotEnabled",
	fsDescriptionDbmNotInitialized:       "DbmNotInitialized",
	fsDescriptionCardRomNotInitialized:   "CardRomNotInitialized",

	fsDescriptionAlreadyInitialized: "AlreadyInitialized",

	fsDescriptionUnsupportedOperation: "UnsupportedOperation",
	fsDescriptionUnsupportedAlignment: "UnsupportedAlignment",
	fsDescriptionUnsupportedMedia:     "UnsupportedMedia",
	fsDescriptionDbmInvalidOperation:  "DbmInvalidOperation",

	fsDescriptionNotImplemented:    "NotImplemented",
	fsDescriptionInvalidKeyType:    "InvalidKeyType",
	fsDescriptionPartitionNotFound: "PartitionNotFound",

	fsDescriptionFatError:      "FatError",
	fsDescriptionFatCorruption: "FatCorruption",

	fsDescriptionNandError: "NandError",
	fsDescriptionNandFatal: "NandFatal",

	fsDescriptionUnknownError: "UnknownError",
}

var fsDescriptionDescription = map[fsDescription]string{
	fsDescriptionFindFinished:         "Search finished",
	fsDescriptionDbmFindFinished:      "DBM search completed",
	fsDescriptionDbmFindKeyFinished:   "DBM Key lookup completed.",
	fsDescriptionDbmIterationFinished: "DBM iteration is complete.",

	fsDescriptionNotFound:                "File, archive, etc. not found.",
	fsDescriptionArchiveNotFound:         "Archive not found",
	fsDescriptionCardPartitionNotFound:   "Card partition not found",
	fsDescriptionDbmNotFound:             "DBM can't find files, archives, etc.",
	fsDescriptionDbmKeyNotFound:          "DBM key not found",
	fsDescriptionDbmFileNotFound:         "DBM file not found",
	fsDescriptionDbmDirectoryNotFound:    "DBM directory not found",
	fsDescriptionDbmIdNotFound:           "DBM ID not found",
	fsDescriptionFatNotFound:             "Specified FAT volume or path does not exist.",
	fsDescriptionMediaNotFound:           "Media not found or not recognized.",
	fsDescriptionFatNoStorage:            "No accessible FAT device exists.",
	fsDescriptionCardRomNotFound:         "The card media cannot be found/recognized.",
	fsDescriptionCardRomNoDevice:         "The card is not inserted or is missing.",
	fsDescriptionCardRomUnkown:           "Bad card (wrong ID).",
	fsDescriptionCardBackupNotFound:      "The backup card media cannot be found/recognized.",
	fsDescriptionCardBackupType1NoDevice: "Card Type1 backup device is not inserted or removed.",
	fsDescriptionCardBackupType2NoDevice: "Card Type2 backup device is not inserted or removed.",
	fsDescriptionCardBackupType1Unkown:   "Bad card Type1 backup device (ID bad).",
	fsDescriptionCardBackupType2Unkown:   "Bad card Type2 backup device (ID bad).",
	fsDescriptionSdmcNotFound:            "The SD card cannot be found/recognized.",
	fsDescriptionSdmcNoDevice:            "The SD card is not inserted or is missing.",
	fsDescriptionSdmcUnkown:              "A medium other than an SD card is inserted (MMC, etc.).",

	fsDescriptionAlreadyExists:    "Files, archives, etc. already exists.",
	fsDescriptionDbmAlreadyExists: "The specified DBM file already exists.",
	fsDescriptionFatAlreadyExists: "The specified FAT path already exists.",

	fsDescriptionNotEnoughSpace:    "No space left",
	fsDescriptionDbmNotEnoughSpace: "DBM storage was full.",
	fsDescriptionFatNotEnoughSpace: "Insufficient space on FAT device.",

	fsDescriptionArchiveInvalidated: "Archiving disabled.",

	fsDescriptionOperationDenied:          "Operation denied",
	fsDescriptionResourceLocked:           "Resource locked",
	fsDescriptionCacheError:               "An error occurred during a cache operation.",
	fsDescriptionDbmOperationDenied:       "DBM operation denied",
	fsDescriptionFatOperationDenied:       "FAT operation denied",
	fsDescriptionWriteProtected:           "Writing is prohibited.",
	fsDescriptionSdmcWriteProtected:       "The SD card is locked (write-protected).",
	fsDescriptionCardWriteProtected:       "The card is locked (write-protected).",
	fsDescriptionMediaAccessError:         "An error occurred while accessing the media.",
	fsDescriptionCardRomAccessError:       "Error accessing card rom media.",
	fsDescriptionCardRomCommError:         "Card communication error (such as CRC error).",
	fsDescriptionCardBackupCommError:      "Error accessing card backup media.",
	fsDescriptionCardBackupType1CommError: "Card Type1 backup device communication error (such as CRC error).",
	fsDescriptionCardBackupType2CommError: "Card Type2 backup device communication error (such as CRC error).",
	fsDescriptionNandAccessError:          "Error accessing NAND media.",
	fsDescriptionNandCommError:            "NAND communication error (CRC error, etc.)",
	fsDescriptionNandEccFailed:            "Data lost due to NAND ECC uncorrectable.",
	fsDescriptionSdmcAccessError:          "Error accessing SD card media.",
	fsDescriptionSdmcCommError:            "SD card communication error (such as CRC error).",
	fsDescriptionSdmcEccFailed:            "Data lost due to SD card ECC uncorrectable.",
	fsDescriptionBackupNotRequired:        "Save data does not exist.",

	fsDescriptionNotFormatted:            "Not formatted",
	fsDescriptionFormatIsFactoryDefaults: "Factory default.",
	fsDescriptionDbmNotFormatted:         "DBM not formatted",
	fsDescriptionDbmVersionMismatch:      "Different DBM data versions.",

	fsDescriptionBadFormat:                "Invalid format",
	fsDescriptionFatBadFormat:             "Bad FAT format",
	fsDescriptionFatBrokenEntry:           "The specified FAT path target is broken.",
	fsDescriptionCardBackupBadFormat:      "Bad card backup format",
	fsDescriptionCardBackupType1BadFormat: "Invalid card type1 backup device format",
	fsDescriptionCardBackupType2BadFormat: "Invalid card type2 backup device format",

	fsDescriptionVerificationFailed:    "Validation failed or tampering detected",
	fsDescriptionHashMismatch:          "Hash mismatch",
	fsDescriptionVerifySignatureFailed: "Signature verification failed",
	fsDescriptionInvalidFileFormat:     "Invalid file format",
	fsDescriptionVerifyHeaderMacFailed: "Verification of the header MAC failed",
	fsDescriptionVerifyDataHashFailed:  "Hash verification of real data failed",

	fsDescriptionProgramDataNotFound:    "Program, data in program not found.",
	fsDescriptionProgramNotFound:        "Program not found",
	fsDescriptionSystemMenuDataNotFound: "Does not contain icon data.",
	fsDescriptionBannerDataNotFound:     "Does not contain banner data.",
	fsDescriptionLogoDataNotFound:       "Does not contain logo data.",
	fsDescriptionProgramInfoNotFound:    "ProgramInfo not found",
	fsDescriptionContentNotFound:        "Content not found",
	fsDescriptionNoSuchExeFsSection:     "Specified section not found.",

	fsDescriptionRequestRetry:      "Request retry",
	fsDescriptionRequestSameRetry:  "Retry request from SDMC device.",
	fsDescriptionRequestSmallRetry: "Retry request from card device.",

	fsDescriptionInvalidProgramFormat: "Invalid program format",
	fsDescriptionInvalidCxiFormat:     "Malformed CXI file",
	fsDescriptionInvalidCciFormat:     "Malformed CCI file",

	fsDescriptionOutOfResource:         "Insufficient resources. This error should not occur in the product.",
	fsDescriptionOutOfMemory:           "Failed to allocate memory.",
	fsDescriptionHandleTableFull:       "There is no space in the handle table.",
	fsDescriptionDbmOutOfResource:      "Insufficient DBM resources.",
	fsDescriptionDbmTableFull:          "The DBM table is full.",
	fsDescriptionDbmFileEntryFull:      "The DBM file entry is full.",
	fsDescriptionDbmDirectoryEntryFull: "The DBM directory entry is full.",
	fsDescriptionDbmOutOfMemory:        "DBM failed to allocate memory.",
	fsDescriptionFatOutOfResource:      "Insufficient FAT resources.",
	fsDescriptionFatOutOfMemory:        "FAT failed to allocate memory.",

	fsDescriptionAccessDenied:     "You don't have access. This error should not occur in the product.",
	fsDescriptionNotDevelopmentId: "Invalid program ID.",
	fsDescriptionNoContentRights:  "You don't have permission to start the program.",
	fsDescriptionDbmAccessDenied:  "DBM access denied",
	fsDescriptionFatAccessDenied:  "FAT access denied",

	fsDescriptionInvalidArgument:         "Invalid argument. This error should not occur in the product.",
	fsDescriptionInvalidPositionBase:     "Bad seek position.",
	fsDescriptionInvalidPathFormat:       "Bad path.",
	fsDescriptionPathTooLong:             "Path too long",
	fsDescriptionOutOfBounds:             "Operation out of bounds",
	fsDescriptionInvalidPosition:         "Bad position in archive.",
	fsDescriptionDbmInvalidArgument:      "Bad DBM argument.",
	fsDescriptionDbmFileNameTooLong:      "DBM file name too long",
	fsDescriptionDbmDirectoryNameTooLong: "DBM directory name too long",
	fsDescriptionDbmInvalidPathFormat:    "Bad DBM path",
	fsDescriptionDbmOutOfBounds:          "DBM index out of range.",
	fsDescriptionFatInvalidArgument:      "Bad FAT argument.",
	fsDescriptionFatInvalidSize:          "Bad FAT size",

	fsDescriptionNotInitialized:          "A process that requires initialization was called before it was initialized. This error should not occur in the product.",
	fsDescriptionLibraryNotInitialized:   "File system library not initialized",
	fsDescriptionMovableDataIsNotEnabled: "Movable data is not initialized.",
	fsDescriptionDbmNotInitialized:       "DBM not initialized",
	fsDescriptionCardRomNotInitialized:   "Card not initialized",

	fsDescriptionAlreadyInitialized: "Initialization processing was executed in an already initialized state. This error should not occur in the product.",

	fsDescriptionUnsupportedOperation: "Unsupported function or operation not allowed. This error should not occur in the product.",
	fsDescriptionUnsupportedAlignment: "Unsupported alignment size",
	fsDescriptionUnsupportedMedia:     "This feature is not supported for the specified media.",
	fsDescriptionDbmInvalidOperation:  "Illegal DBM operation",

	fsDescriptionNotImplemented:    "Feature not implemented.",
	fsDescriptionInvalidKeyType:    "An incorrect encryption key is being used.",
	fsDescriptionPartitionNotFound: "A partition that should exist is not found.",

	fsDescriptionFatError:      "FAT error",
	fsDescriptionFatCorruption: "FAT is in an invalid state.",

	fsDescriptionNandError: "NAND error",
	fsDescriptionNandFatal: "NAND may be faulty.",

	fsDescriptionUnknownError: "Unknown error",
}

func (d fsDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", fsDescriptionToString[d], d, fsDescriptionDescription[d])
}
