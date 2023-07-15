package descriptions

import "fmt"

type bossDescription int32

const (
	bossDescriptionNone bossDescription = iota
	bossDescriptionInvalidPolicy
	bossDescriptionInvalidAction
	bossDescriptionInvalidOption
	bossDescriptionInvalidAppIdList
	bossDescriptionInvalidTaskIdList
	bossDescriptionInvalidStepIdList
	bossDescriptionInvalidNsDataIdList
	bossDescriptionInvalidTaskStatus
	bossDescriptionInvalidPropertyValue
	bossDescriptionInvalidNewArrivalEvent
	bossDescriptionInvalidNewArrivalFlag
	bossDescriptionInvalidOptOutFlag
	bossDescriptionInvalidTaskError
	bossDescriptionInvalidNsDataValue
	bossDescriptionInvalidNsDataInfo
	bossDescriptionInvalidNsDataReadFlag
	bossDescriptionInvalidNsDataTime
	bossDescriptionInvalidNextExecuteTime
	bossDescriptionHttpRequestHeaderPointerNull
	bossDescriptionInvalidPolicyListAvailability
	bossDescriptionInvalidTestModeAvailability
	bossDescriptionInvalidConfigParentalControl
	bossDescriptionInvalidConfigEulaAgreement
	bossDescriptionInvalidPolicyListEnvId
	bossDescriptionInvalidPolicyListUrl
	bossDescriptionInvalidTaskId
	bossDescriptionInvalidTaskStep
	bossDescriptionInvalidPropertyType
	bossDescriptionInvalidUrl
	bossDescriptionInvalidFilePath
	bossDescriptionInvalidTaskPriority
	bossDescriptionInvalidTaskTargetDuration
	bossDescriptionActionCodeOutOfRange
	bossDescriptionInvalidNsDataSeekPosition
	bossDescriptionInvalidMaxHttpRequestHeader
	bossDescriptionInvalidMaxClientCert
	bossDescriptionInvalidMaxRootCa
	bossDescriptionSchedulingPolicyOutOfRange
	bossDescriptionApInfoTypeOutOfRange
	bossDescriptionTaskPermissionOutOfRange
	bossDescriptionWaitFinishTimeout
	bossDescriptionWaitFinishTaskNotDone
	bossDescriptionIpcNotSessionInitialized
	bossDescriptionIpcPropertySizeError
	bossDescriptionIpcTooManyRequests
	bossDescriptionAlreadyInitialized
	bossDescriptionOutOfMemory
	bossDescriptionInvalidNumberOfNsd
	bossDescriptionNsDataInvalidFormat
	bossDescriptionApliNotExist
	bossDescriptionTaskNotExist
	bossDescriptionTaskStepNotExist
	bossDescriptionApliIdAlreadyExist
	bossDescriptionTaskIdAlreadyExist
	bossDescriptionTaskStepAlreadyExist
	bossDescriptionInvalidSequence
	bossDescriptionDatabaseFull
	bossDescriptionCantUnregisterTask
	bossDescriptionTaskNotRegistered
	bossDescriptionInvalidFilehandle
	bossDescriptionInvalidTaskSchedulingPolicy
	bossDescriptionInvalidHttpRequestHeader
	bossDescriptionInvalidHeadtype
	bossDescriptionStorageAccesspermission
	bossDescriptionStorageInsufficiency
	bossDescriptionInvalidAppidStorageNotfound
	bossDescriptionNsdataNotfound
	bossDescriptionInvalidNsdataGetheadSize
	bossDescriptionNsdataListSizeShortage
	bossDescriptionNsdataListUpdated
	bossDescriptionNotConnectApWithLocation
	bossDescriptionNotConnectNetwork
	bossDescriptionInvalidFriendcode
	bossDescriptionFileAccess
	bossDescriptionTaskAlreadyPaused
	bossDescriptionTaskAlreadyResumed
	bossDescriptionUnexpect
	bossDescriptionInvalidStorageParameter bossDescription = iota + 192 - 78
	bossDescriptionCfginfotypeOutOfRange
	bossDescriptionInvalidMaxHttpQuery
	bossDescriptionInvalidMaxDatastoreDst
	bossDescriptionNsalistInvalidFormat
	bossDescriptionNsalistDownloadTaskError
	bossDescriptionNotEnoughSpaceInExtSaveData
)

var bossDescriptionToString = map[bossDescription]string{
	bossDescriptionNone:                          "None",
	bossDescriptionInvalidPolicy:                 "InvalidPolicy",
	bossDescriptionInvalidAction:                 "InvalidAction",
	bossDescriptionInvalidOption:                 "InvalidOption",
	bossDescriptionInvalidAppIdList:              "InvalidAppIdList",
	bossDescriptionInvalidTaskIdList:             "InvalidTaskIdList",
	bossDescriptionInvalidStepIdList:             "InvalidStepIdList",
	bossDescriptionInvalidNsDataIdList:           "InvalidNsDataIdList",
	bossDescriptionInvalidTaskStatus:             "InvalidTaskStatus",
	bossDescriptionInvalidPropertyValue:          "InvalidPropertyValue",
	bossDescriptionInvalidNewArrivalEvent:        "InvalidNewArrivalEvent",
	bossDescriptionInvalidNewArrivalFlag:         "InvalidNewArrivalFlag",
	bossDescriptionInvalidOptOutFlag:             "InvalidOptOutFlag",
	bossDescriptionInvalidTaskError:              "InvalidTaskError",
	bossDescriptionInvalidNsDataValue:            "InvalidNsDataValue",
	bossDescriptionInvalidNsDataInfo:             "InvalidNsDataInfo",
	bossDescriptionInvalidNsDataReadFlag:         "InvalidNsDataReadFlag",
	bossDescriptionInvalidNsDataTime:             "InvalidNsDataTime",
	bossDescriptionInvalidNextExecuteTime:        "InvalidNextExecuteTime",
	bossDescriptionHttpRequestHeaderPointerNull:  "HttpRequestHeaderPointerNull",
	bossDescriptionInvalidPolicyListAvailability: "InvalidPolicyListAvailability",
	bossDescriptionInvalidTestModeAvailability:   "InvalidTestModeAvailability",
	bossDescriptionInvalidConfigParentalControl:  "InvalidConfigParentalControl",
	bossDescriptionInvalidConfigEulaAgreement:    "InvalidConfigEulaAgreement",
	bossDescriptionInvalidPolicyListEnvId:        "InvalidPolicyListEnvId",
	bossDescriptionInvalidPolicyListUrl:          "InvalidPolicyListUrl",
	bossDescriptionInvalidTaskId:                 "InvalidTaskId",
	bossDescriptionInvalidTaskStep:               "InvalidTaskStep",
	bossDescriptionInvalidPropertyType:           "InvalidPropertyType",
	bossDescriptionInvalidUrl:                    "InvalidUrl",
	bossDescriptionInvalidFilePath:               "InvalidFilePath",
	bossDescriptionInvalidTaskPriority:           "InvalidTaskPriority",
	bossDescriptionInvalidTaskTargetDuration:     "InvalidTaskTargetDuration",
	bossDescriptionActionCodeOutOfRange:          "ActionCodeOutOfRange",
	bossDescriptionInvalidNsDataSeekPosition:     "InvalidNsDataSeekPosition",
	bossDescriptionInvalidMaxHttpRequestHeader:   "InvalidMaxHttpRequestHeader",
	bossDescriptionInvalidMaxClientCert:          "InvalidMaxClientCert",
	bossDescriptionInvalidMaxRootCa:              "InvalidMaxRootCa",
	bossDescriptionSchedulingPolicyOutOfRange:    "SchedulingPolicyOutOfRange",
	bossDescriptionApInfoTypeOutOfRange:          "ApInfoTypeOutOfRange",
	bossDescriptionTaskPermissionOutOfRange:      "TaskPermissionOutOfRange",
	bossDescriptionWaitFinishTimeout:             "WaitFinishTimeout",
	bossDescriptionWaitFinishTaskNotDone:         "WaitFinishTaskNotDone",
	bossDescriptionIpcNotSessionInitialized:      "IpcNotSessionInitialized",
	bossDescriptionIpcPropertySizeError:          "IpcPropertySizeError",
	bossDescriptionIpcTooManyRequests:            "IpcTooManyRequests",
	bossDescriptionAlreadyInitialized:            "AlreadyInitialized",
	bossDescriptionOutOfMemory:                   "OutOfMemory",
	bossDescriptionInvalidNumberOfNsd:            "InvalidNumberOfNsd",
	bossDescriptionNsDataInvalidFormat:           "NsDataInvalidFormat",
	bossDescriptionApliNotExist:                  "ApliNotExist",
	bossDescriptionTaskNotExist:                  "TaskNotExist",
	bossDescriptionTaskStepNotExist:              "TaskStepNotExist",
	bossDescriptionApliIdAlreadyExist:            "ApliIdAlreadyExist",
	bossDescriptionTaskIdAlreadyExist:            "TaskIdAlreadyExist",
	bossDescriptionTaskStepAlreadyExist:          "TaskStepAlreadyExist",
	bossDescriptionInvalidSequence:               "InvalidSequence",
	bossDescriptionDatabaseFull:                  "DatabaseFull",
	bossDescriptionCantUnregisterTask:            "CantUnregisterTask",
	bossDescriptionTaskNotRegistered:             "TaskNotRegistered",
	bossDescriptionInvalidFilehandle:             "InvalidFilehandle",
	bossDescriptionInvalidTaskSchedulingPolicy:   "InvalidTaskSchedulingPolicy",
	bossDescriptionInvalidHttpRequestHeader:      "InvalidHttpRequestHeader",
	bossDescriptionInvalidHeadtype:               "InvalidHeadtype",
	bossDescriptionStorageAccesspermission:       "StorageAccesspermission",
	bossDescriptionStorageInsufficiency:          "StorageInsufficiency",
	bossDescriptionInvalidAppidStorageNotfound:   "InvalidAppidStorageNotfound",
	bossDescriptionNsdataNotfound:                "NsdataNotfound",
	bossDescriptionInvalidNsdataGetheadSize:      "InvalidNsdataGetheadSize",
	bossDescriptionNsdataListSizeShortage:        "NsdataListSizeShortage",
	bossDescriptionNsdataListUpdated:             "NsdataListUpdated",
	bossDescriptionNotConnectApWithLocation:      "NotConnectApWithLocation",
	bossDescriptionNotConnectNetwork:             "NotConnectNetwork",
	bossDescriptionInvalidFriendcode:             "InvalidFriendcode",
	bossDescriptionFileAccess:                    "FileAccess",
	bossDescriptionTaskAlreadyPaused:             "TaskAlreadyPaused",
	bossDescriptionTaskAlreadyResumed:            "TaskAlreadyResumed",
	bossDescriptionUnexpect:                      "Unexpect",
	bossDescriptionInvalidStorageParameter:       "InvalidStorageParameter",
	bossDescriptionCfginfotypeOutOfRange:         "CfginfotypeOutOfRange",
	bossDescriptionInvalidMaxHttpQuery:           "InvalidMaxHttpQuery",
	bossDescriptionInvalidMaxDatastoreDst:        "InvalidMaxDatastoreDst",
	bossDescriptionNsalistInvalidFormat:          "NsalistInvalidFormat",
	bossDescriptionNsalistDownloadTaskError:      "NsalistDownloadTaskError",
	bossDescriptionNotEnoughSpaceInExtSaveData:   "NotEnoughSpaceInExtSaveData",
}

var bossDescriptionDescription = map[bossDescription]string{
	bossDescriptionNone:                          "No error.",
	bossDescriptionInvalidPolicy:                 "The policy information pointer is NULL.",
	bossDescriptionInvalidAction:                 "The task action pointer is NULL.",
	bossDescriptionInvalidOption:                 "The task option pointer is NULL, or the option code is invalid.",
	bossDescriptionInvalidAppIdList:              "The pointer for getting the application list is NULL.",
	bossDescriptionInvalidTaskIdList:             "The pointer for getting the task ID list is NULL.",
	bossDescriptionInvalidStepIdList:             "The pointer for getting the step ID list is NULL.",
	bossDescriptionInvalidNsDataIdList:           "The pointer to NS data list information is NULL.",
	bossDescriptionInvalidTaskStatus:             "The task status pointer is NULL.",
	bossDescriptionInvalidPropertyValue:          "The property value pointer is NULL.",
	bossDescriptionInvalidNewArrivalEvent:        "The pointer to the new arrival event is NULL.",
	bossDescriptionInvalidNewArrivalFlag:         "The pointer to the new arrival flag is NULL.",
	bossDescriptionInvalidOptOutFlag:             "The pointer to the OPTOUT flag is NULL.",
	bossDescriptionInvalidTaskError:              "The pointer to the task error information is NULL.",
	bossDescriptionInvalidNsDataValue:            "The pointer to the region that stores NSDATA is NULL.",
	bossDescriptionInvalidNsDataInfo:             "The pointer to the location that stores additional NSDATA information is NULL.",
	bossDescriptionInvalidNsDataReadFlag:         "The pointer to the storage region for the NSDATA read flag is NULL.",
	bossDescriptionInvalidNsDataTime:             "The pointer to the region used to store the NSDATA update time is NULL.",
	bossDescriptionInvalidNextExecuteTime:        "The pointer to the location used to store the minutes value of the next execution time is  NULL.",
	bossDescriptionHttpRequestHeaderPointerNull:  "The HTTP request header pointer is NULL.",
	bossDescriptionInvalidPolicyListAvailability: "The pointer to information that can be used by the policy list is NULL.",
	bossDescriptionInvalidTestModeAvailability:   "The pointer to information usable in test mode is NULL.",
	bossDescriptionInvalidConfigParentalControl:  "The pointer to information on Parental Controls settings is NULL.",
	bossDescriptionInvalidConfigEulaAgreement:    "The pointer to information about whether the user has agreed to the EULA is NULL.",
	bossDescriptionInvalidPolicyListEnvId:        "The pointer to the policy list's environment ID is NULL.",
	bossDescriptionInvalidPolicyListUrl:          "The pointer to the policy list's URL information is NULL.",
	bossDescriptionInvalidTaskId:                 "The task ID pointer is NULL or a zero-length string.",
	bossDescriptionInvalidTaskStep:               "The current step ID was specified during task registration.",
	bossDescriptionInvalidPropertyType:           "The property type is not supported.",
	bossDescriptionInvalidUrl:                    "The URL string pointer is NULL or a zero-length string.",
	bossDescriptionInvalidFilePath:               "The file path string pointer is NULL or a zero-length string.",
	bossDescriptionInvalidTaskPriority:           "The specified task priority is invalid.",
	bossDescriptionInvalidTaskTargetDuration:     "The task duration is invalid.",
	bossDescriptionActionCodeOutOfRange:          "The task action code is out of range.",
	bossDescriptionInvalidNsDataSeekPosition:     "The NS data seek position exceeds the data size.",
	bossDescriptionInvalidMaxHttpRequestHeader:   "The HTTP request header registration count exceeds the maximum.",
	bossDescriptionInvalidMaxClientCert:          "The number of client certificates set exceeds the maximum.",
	bossDescriptionInvalidMaxRootCa:              "The number of root CAs exceeds the maximum that can be set.",
	bossDescriptionSchedulingPolicyOutOfRange:    "The scheduling policy is invalid.",
	bossDescriptionApInfoTypeOutOfRange:          "The AP information type is invalid.",
	bossDescriptionTaskPermissionOutOfRange:      "The task permission information is invalid.",
	bossDescriptionWaitFinishTimeout:             "The WaitFinish function timed out.",
	bossDescriptionWaitFinishTaskNotDone:         "The target task ended with a result other than TASK_DONE in the WaitFinish function.",
	bossDescriptionIpcNotSessionInitialized:      "The IPC session is not initialized.",
	bossDescriptionIpcPropertySizeError:          "There was a property size error during internal IPC processing.",
	bossDescriptionIpcTooManyRequests:            "There are too many IPC requests. (This is provided for future extensibility and cannot be used.)",
	bossDescriptionAlreadyInitialized:            "IPC is already initialized.",
	bossDescriptionOutOfMemory:                   "Insufficient memory.",
	bossDescriptionInvalidNumberOfNsd:            "The maximum number of NSD files in the NSA was exceeded.",
	bossDescriptionNsDataInvalidFormat:           "An invalid NS data format was detected.",
	bossDescriptionApliNotExist:                  "The specified application ID was not found.",
	bossDescriptionTaskNotExist:                  "The specified task ID was not found.",
	bossDescriptionTaskStepNotExist:              "The specified task step was not found.",
	bossDescriptionApliIdAlreadyExist:            "Another application of the same name is already registered.",
	bossDescriptionTaskIdAlreadyExist:            "Another task of the same name is already registered.",
	bossDescriptionTaskStepAlreadyExist:          "Another task step of the same name is already registered.",
	bossDescriptionInvalidSequence:               "A sequence error (such as an attempt to start a task that is currently running) was detected.",
	bossDescriptionDatabaseFull:                  "Storage and tasks cannot be registered because the maximum number of registered application IDs and tasks has been reached.",
	bossDescriptionCantUnregisterTask:            "The task cannot be deleted because of its state (for example, it may currently be running or have already been scheduled).",
	bossDescriptionTaskNotRegistered:             "A task that is supposed to be registered in the database file was not found.",
	bossDescriptionInvalidFilehandle:             "Invalid file handle.",
	bossDescriptionInvalidTaskSchedulingPolicy:   "An invalid keyword was detected in the scheduling policy list.",
	bossDescriptionInvalidHttpRequestHeader:      "More than the maximum number of option HTTP request headers were specified.",
	bossDescriptionInvalidHeadtype:               "An invalid header type was specified in the GetHeaderInfo function.",
	bossDescriptionStorageAccesspermission:       "You do not have storage access rights.",
	bossDescriptionStorageInsufficiency:          "Insufficient storage space.",
	bossDescriptionInvalidAppidStorageNotfound:   "Storage has not been registered for the corresponding application ID.",
	bossDescriptionNsdataNotfound:                "The specified NS data was not found.",
	bossDescriptionInvalidNsdataGetheadSize:      "The header size does not match the header type specified by the GetHeaderInfo function.",
	bossDescriptionNsdataListSizeShortage:        "The NsDataIdList size is insufficient. (The list is not big enough to store all the NSD serial IDs.)",
	bossDescriptionNsdataListUpdated:             "The target NSD group for BOSS storage was updated since the last time a list was obtained.",
	bossDescriptionNotConnectApWithLocation:      "Not connected to an access point.",
	bossDescriptionNotConnectNetwork:             "Not connected to a network.",
	bossDescriptionInvalidFriendcode:             "Invalid friend code error.",
	bossDescriptionFileAccess:                    "Failed to access a file. An abnormality occurred in the system save data when accessing the internal database and settings. This error cannot normally happen. The problem may disappear if you initialize the system save data, but doing so deletes all task information (including the information registered by other applications).",
	bossDescriptionTaskAlreadyPaused:             "Already paused.",
	bossDescriptionTaskAlreadyResumed:            "Already resumed.",
	bossDescriptionUnexpect:                      "An unexpected error occurred.",
	bossDescriptionInvalidStorageParameter:       "An invalid value was set in the parameter specified for storage.",
	bossDescriptionCfginfotypeOutOfRange:         "The configuration information type is out of range.",
	bossDescriptionInvalidMaxHttpQuery:           "The HTTP query registration count exceeds the maximum.",
	bossDescriptionInvalidMaxDatastoreDst:        "The DataStore data transmission destination registration count exceeds the maximum.",
	bossDescriptionNsalistInvalidFormat:          "An error has been detected in the NSA list file format.",
	bossDescriptionNsalistDownloadTaskError:      "The NSA list download result was an error.",
	bossDescriptionNotEnoughSpaceInExtSaveData:   "Expanded save data has insufficient capacity.",
}

func (d bossDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", bossDescriptionToString[d], d, bossDescriptionDescription[d])
}
