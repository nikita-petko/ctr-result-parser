package descriptions

import "fmt"

type dbmDescription int32

const (
	dbmDescriptionOutOfBounds          dbmDescription = dbmDescription(fsDescriptionDbmOutOfBounds)
	dbmDescriptionOutOfStorage         dbmDescription = dbmDescription(fsDescriptionDbmNotEnoughSpace)
	dbmDescriptionKeyNotFound          dbmDescription = dbmDescription(fsDescriptionDbmKeyNotFound)
	dbmDescriptionTableFull            dbmDescription = dbmDescription(fsDescriptionDbmTableFull)
	dbmDescriptionFindKeyFinished      dbmDescription = dbmDescription(fsDescriptionDbmFindKeyFinished)
	dbmDescriptionFindFinished         dbmDescription = dbmDescription(fsDescriptionDbmFindFinished)
	dbmDescriptionFileNotFound         dbmDescription = dbmDescription(fsDescriptionDbmFileNotFound)
	dbmDescriptionFileNameTooLong      dbmDescription = dbmDescription(fsDescriptionDbmFileNameTooLong)
	dbmDescriptionFileFull             dbmDescription = dbmDescription(fsDescriptionDbmFileEntryFull)
	dbmDescriptionDirectoryNotFound    dbmDescription = dbmDescription(fsDescriptionDbmDirectoryNotFound)
	dbmDescriptionDirectoryNameTooLong dbmDescription = dbmDescription(fsDescriptionDbmDirectoryNameTooLong)
	dbmDescriptionDirectoryFull        dbmDescription = dbmDescription(fsDescriptionDbmDirectoryEntryFull)
	dbmDescriptionInvalidPathFormat    dbmDescription = dbmDescription(fsDescriptionDbmInvalidPathFormat)
	dbmDescriptionAlreadyExists        dbmDescription = dbmDescription(fsDescriptionDbmAlreadyExists)
	dbmDescriptionOperationDenied      dbmDescription = dbmDescription(fsDescriptionDbmOperationDenied)
	dbmDescriptionAccessDenied         dbmDescription = dbmDescription(fsDescriptionDbmAccessDenied)
	dbmDescriptionInvalidOperation     dbmDescription = dbmDescription(fsDescriptionDbmInvalidOperation)
	dbmDescriptionIterationFinished    dbmDescription = dbmDescription(fsDescriptionDbmIterationFinished)
	dbmDescriptionOutOfMemory          dbmDescription = dbmDescription(fsDescriptionDbmOutOfMemory)
	dbmDescriptionNotInitialized       dbmDescription = dbmDescription(fsDescriptionDbmNotInitialized)
	dbmDescriptionIdNotFound           dbmDescription = dbmDescription(fsDescriptionDbmIdNotFound)
	dbmDescriptionNotFormatted         dbmDescription = dbmDescription(fsDescriptionDbmNotFormatted)
	dbmDescriptionVersionMismatch      dbmDescription = dbmDescription(fsDescriptionDbmVersionMismatch)
)

var dbmDescriptionToString = map[dbmDescription]string{
	dbmDescriptionOutOfBounds:          "OutOfBounds",
	dbmDescriptionOutOfStorage:         "OutOfStorage",
	dbmDescriptionKeyNotFound:          "KeyNotFound",
	dbmDescriptionTableFull:            "TableFull",
	dbmDescriptionFindKeyFinished:      "FindKeyFinished",
	dbmDescriptionFindFinished:         "FindFinished",
	dbmDescriptionFileNotFound:         "FileNotFound",
	dbmDescriptionFileNameTooLong:      "FileNameTooLong",
	dbmDescriptionFileFull:             "FileFull",
	dbmDescriptionDirectoryNotFound:    "DirectoryNotFound",
	dbmDescriptionDirectoryNameTooLong: "DirectoryNameTooLong",
	dbmDescriptionDirectoryFull:        "DirectoryFull",
	dbmDescriptionInvalidPathFormat:    "InvalidPathFormat",
	dbmDescriptionAlreadyExists:        "AlreadyExists",
	dbmDescriptionOperationDenied:      "OperationDenied",
	dbmDescriptionAccessDenied:         "AccessDenied",
	dbmDescriptionInvalidOperation:     "InvalidOperation",
	dbmDescriptionIterationFinished:    "IterationFinished",
	dbmDescriptionOutOfMemory:          "OutOfMemory",
	dbmDescriptionNotInitialized:       "NotInitialized",
	dbmDescriptionIdNotFound:           "IdNotFound",
	dbmDescriptionNotFormatted:         "NotFormatted",
	dbmDescriptionVersionMismatch:      "VersionMismatch",
}

var dbmDescriptionDescription = map[dbmDescription]string{
	dbmDescriptionOutOfBounds:          fsDescriptionDescription[fsDescriptionDbmOutOfBounds],
	dbmDescriptionOutOfStorage:         fsDescriptionDescription[fsDescriptionDbmNotEnoughSpace],
	dbmDescriptionKeyNotFound:          fsDescriptionDescription[fsDescriptionDbmKeyNotFound],
	dbmDescriptionTableFull:            fsDescriptionDescription[fsDescriptionDbmTableFull],
	dbmDescriptionFindKeyFinished:      fsDescriptionDescription[fsDescriptionDbmFindKeyFinished],
	dbmDescriptionFindFinished:         fsDescriptionDescription[fsDescriptionDbmFindFinished],
	dbmDescriptionFileNotFound:         fsDescriptionDescription[fsDescriptionDbmFileNotFound],
	dbmDescriptionFileNameTooLong:      fsDescriptionDescription[fsDescriptionDbmFileNameTooLong],
	dbmDescriptionFileFull:             fsDescriptionDescription[fsDescriptionDbmFileEntryFull],
	dbmDescriptionDirectoryNotFound:    fsDescriptionDescription[fsDescriptionDbmDirectoryNotFound],
	dbmDescriptionDirectoryNameTooLong: fsDescriptionDescription[fsDescriptionDbmDirectoryNameTooLong],
	dbmDescriptionDirectoryFull:        fsDescriptionDescription[fsDescriptionDbmDirectoryEntryFull],
	dbmDescriptionInvalidPathFormat:    fsDescriptionDescription[fsDescriptionDbmInvalidPathFormat],
	dbmDescriptionAlreadyExists:        fsDescriptionDescription[fsDescriptionDbmAlreadyExists],
	dbmDescriptionOperationDenied:      fsDescriptionDescription[fsDescriptionDbmOperationDenied],
	dbmDescriptionAccessDenied:         fsDescriptionDescription[fsDescriptionDbmAccessDenied],
	dbmDescriptionInvalidOperation:     fsDescriptionDescription[fsDescriptionDbmInvalidOperation],
	dbmDescriptionIterationFinished:    fsDescriptionDescription[fsDescriptionDbmIterationFinished],
	dbmDescriptionOutOfMemory:          fsDescriptionDescription[fsDescriptionDbmOutOfMemory],
	dbmDescriptionNotInitialized:       fsDescriptionDescription[fsDescriptionDbmNotInitialized],
	dbmDescriptionIdNotFound:           fsDescriptionDescription[fsDescriptionDbmIdNotFound],
	dbmDescriptionNotFormatted:         fsDescriptionDescription[fsDescriptionDbmNotFormatted],
	dbmDescriptionVersionMismatch:      fsDescriptionDescription[fsDescriptionDbmVersionMismatch],
}

func (d dbmDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", dbmDescriptionToString[d], d, dbmDescriptionDescription[d])
}
