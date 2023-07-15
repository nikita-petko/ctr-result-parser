package descriptions

import "fmt"

type mvdDescription int32

/*

   DESCRIPTION_OK = 0,
   DESCRIPTION_STRM_PROCESSED = 1,
   DESCRIPTION_PIC_RDY = 2,
   DESCRIPTION_PIC_DECODED = 3,
   DESCRIPTION_HDRS_RDY = 4,
   DESCRIPTION_ADVANCED_TOOLS = 5,
   DESCRIPTION_PENDING_FLUSH = 6,
   DESCRIPTION_NONREF_PIC_SKIPPED = 7,

   DESCRIPTION_DRIVER_SUCESS_OFFSET = 50,

   DESCRIPTION_SLICE_RDY = 56,

   DESCRIPTION_LIBS_OFFSET = 100,

   DESCRIPTION_MEM_NOT_INITIALIZED = 101,
   DESCRIPTION_MEM_ALREADY_INITIALIZED = 102,
   DESCRIPTION_INPUT_MALLOC_FAIL = 103,
   DESCRIPTION_DONE_NOTHING = 104,

   DESCRIPTION_DRIVER_ERROR_OFFSET_1 = 200,
   DESCRIPTION_DRIVER_ERROR_OFFSET_1_ORG = 0,

   DESCRIPTION_PARAM_ERROR = 201,
   DESCRIPTION_STRM_ERROR = 202,
   DESCRIPTION_NOT_INITIALIZED = 203,
   DESCRIPTION_MEMFAIL = 204,
   DESCRIPTION_INITFAIL = 205,
   DESCRIPTION_HDRS_NOT_RDY = 206,
   DESCRIPTION_STREAM_NOT_SUPPORTED = 208,

   DESCRIPTION_PP_SET_IN_SIZE_INVALID = 264,
   DESCRIPTION_PP_SET_IN_ADDRESS_INVALID = 265,
   DESCRIPTION_PP_SET_IN_FORMAT_INVALID = 266,
   DESCRIPTION_PP_SET_CROP_INVALID = 267,
   DESCRIPTION_PP_SET_ROTATION_INVALID = 268,
   DESCRIPTION_PP_SET_OUT_SIZE_INVALID = 269,
   DESCRIPTION_PP_SET_OUT_ADDRESS_INVALID = 270,
   DESCRIPTION_PP_SET_OUT_FORMAT_INVALID = 271,
   DESCRIPTION_PP_SET_VIDEO_ADJUST_INVALID = 272,
   DESCRIPTION_PP_SET_RGB_BITMASK_INVALID = 273,
   DESCRIPTION_PP_SET_FRAMEBUFFER_INVALID = 274,
   DESCRIPTION_PP_SET_MASK1_INVALID = 275,
   DESCRIPTION_PP_SET_MASK2_INVALID = 276,
   DESCRIPTION_PP_SET_DEINTERLACE_INVALID = 277,
   DESCRIPTION_PP_SET_IN_STRUCT_INVALID = 278,
   DESCRIPTION_PP_SET_IN_RANGE_MAP_INVALID = 279,
   DESCRIPTION_PP_SET_ABLEND_UNSUPPORTED = 280,
   DESCRIPTION_PP_SET_DEINTERLACING_UNSUPPORTED = 281,
   DESCRIPTION_PP_SET_DITHERING_UNSUPPORTED = 282,
   DESCRIPTION_PP_SET_SCALING_UNSUPPORTED = 283,

   DESCRIPTION_PP_BUSY = 328,

   DESCRIPTION_HW_RESERVED = 454,
   DESCRIPTION_HW_TIMEOUT = 455,
   DESCRIPTION_HW_BUS_ERROR = 456,
   DESCRIPTION_SYSTEM_ERROR = 457,
   DESCRIPTION_DWL_ERROR = 458,

   DESCRIPTION_PP_DEC_COMBINED_MODE_ERROR = 712,
   DESCRIPTION_PP_DEC_RUNTIME_ERROR = 713,

   DESCRIPTION_DRIVER_ERROR_OFFSET_2 = 800,
   DESCRIPTION_DRIVER_ERROR_OFFSET_2_ORG = 900,

   DESCRIPTION_EVALUATION_LIMIT_EXCEEDED = 899,
   DESCRIPTION_FORMAT_NOT_SUPPORTED = 900,

   DESCRIPTION_SDK_RESERVED = 1000

*/

const (
	mvdDescriptionOk mvdDescription = iota
	mvdDescriptionStrmProcessed
	mvdDescriptionPicRdy
	mvdDescriptionPicDecoded
	mvdDescriptionHdrsRdy
	mvdDescriptionAdvancedTools
	mvdDescriptionPendingFlush
	mvdDescriptionNonrefPicSkipped
	mvdDescriptionDriverSucessOffset mvdDescription = iota + 50 - 8
	mvdDescriptionSliceRdy           mvdDescription = 56
	mvdDescriptionLibsOffset         mvdDescription = iota + 100 - 10
	mvdDescriptionMemNotInitialized
	mvdDescriptionMemAlreadyInitialized
	mvdDescriptionInputMallocFail
	mvdDescriptionDoneNothing
	mvdDescriptionDriverErrorOffset1 mvdDescription = iota + 200 - 15
	mvdDescriptionParamError
	mvdDescriptionStrmError
	mvdDescriptionNotInitialized
	mvdDescriptionMemfail
	mvdDescriptionInitfail
	mvdDescriptionHdrsNotRdy
	mvdDescriptionStreamNotSupported
	mvdDescriptionPpSetInSizeInvalid mvdDescription = iota + 264 - 24
	mvdDescriptionPpSetInAddressInvalid
	mvdDescriptionPpSetInFormatInvalid
	mvdDescriptionPpSetCropInvalid
	mvdDescriptionPpSetRotationInvalid
	mvdDescriptionPpSetOutSizeInvalid
	mvdDescriptionPpSetOutAddressInvalid
	mvdDescriptionPpSetOutFormatInvalid
	mvdDescriptionPpSetVideoAdjustInvalid
	mvdDescriptionPpSetRgbBitmaskInvalid
	mvdDescriptionPpSetFramebufferInvalid
	mvdDescriptionPpSetMask1Invalid
	mvdDescriptionPpSetMask2Invalid
	mvdDescriptionPpSetDeinterlaceInvalid
	mvdDescriptionPpSetInStructInvalid
	mvdDescriptionPpSetInRangeMapInvalid
	mvdDescriptionPpSetAblendUnsupported
	mvdDescriptionPpSetDeinterlacingUnsupported
	mvdDescriptionPpSetDitheringUnsupported
	mvdDescriptionPpSetScalingUnsupported
	mvdDescriptionPpBusy     mvdDescription = 328
	mvdDescriptionHwReserved mvdDescription = iota + 454 - 45
	mvdDescriptionHwTimeout
	mvdDescriptionHwBusError
	mvdDescriptionSystemError
	mvdDescriptionDwlError
	mvdDescriptionPpDecCombinedModeError  mvdDescription = 712
	mvdDescriptionPpDecRuntimeError       mvdDescription = 713
	mvdDescriptionDriverErrorOffset2      mvdDescription = 800
	mvdDescriptionEvaluationLimitExceeded mvdDescription = 899
	mvdDescriptionFormatNotSupported      mvdDescription = 900
	mvdDescriptionSdkReserved             mvdDescription = 1000
)

var mvdDescriptionToString = map[mvdDescription]string{
	mvdDescriptionOk:                            "Ok",
	mvdDescriptionStrmProcessed:                 "StrmProcessed",
	mvdDescriptionPicRdy:                        "PicRdy",
	mvdDescriptionPicDecoded:                    "PicDecoded",
	mvdDescriptionHdrsRdy:                       "HdrsRdy",
	mvdDescriptionAdvancedTools:                 "AdvancedTools",
	mvdDescriptionPendingFlush:                  "PendingFlush",
	mvdDescriptionNonrefPicSkipped:              "NonrefPicSkipped",
	mvdDescriptionDriverSucessOffset:            "DriverSucessOffset",
	mvdDescriptionSliceRdy:                      "SliceRdy",
	mvdDescriptionLibsOffset:                    "LibsOffset",
	mvdDescriptionMemNotInitialized:             "MemNotInitialized",
	mvdDescriptionMemAlreadyInitialized:         "MemAlreadyInitialized",
	mvdDescriptionInputMallocFail:               "InputMallocFail",
	mvdDescriptionDoneNothing:                   "DoneNothing",
	mvdDescriptionDriverErrorOffset1:            "DriverErrorOffset1",
	mvdDescriptionParamError:                    "ParamError",
	mvdDescriptionStrmError:                     "StrmError",
	mvdDescriptionNotInitialized:                "NotInitialized",
	mvdDescriptionMemfail:                       "Memfail",
	mvdDescriptionInitfail:                      "Initfail",
	mvdDescriptionHdrsNotRdy:                    "HdrsNotRdy",
	mvdDescriptionStreamNotSupported:            "StreamNotSupported",
	mvdDescriptionPpSetInSizeInvalid:            "PpSetInSizeInvalid",
	mvdDescriptionPpSetInAddressInvalid:         "PpSetInAddressInvalid",
	mvdDescriptionPpSetInFormatInvalid:          "PpSetInFormatInvalid",
	mvdDescriptionPpSetCropInvalid:              "PpSetCropInvalid",
	mvdDescriptionPpSetRotationInvalid:          "PpSetRotationInvalid",
	mvdDescriptionPpSetOutSizeInvalid:           "PpSetOutSizeInvalid",
	mvdDescriptionPpSetOutAddressInvalid:        "PpSetOutAddressInvalid",
	mvdDescriptionPpSetOutFormatInvalid:         "PpSetOutFormatInvalid",
	mvdDescriptionPpSetVideoAdjustInvalid:       "PpSetVideoAdjustInvalid",
	mvdDescriptionPpSetRgbBitmaskInvalid:        "PpSetRgbBitmaskInvalid",
	mvdDescriptionPpSetFramebufferInvalid:       "PpSetFramebufferInvalid",
	mvdDescriptionPpSetMask1Invalid:             "PpSetMask1Invalid",
	mvdDescriptionPpSetMask2Invalid:             "PpSetMask2Invalid",
	mvdDescriptionPpSetDeinterlaceInvalid:       "PpSetDeinterlaceInvalid",
	mvdDescriptionPpSetInStructInvalid:          "PpSetInStructInvalid",
	mvdDescriptionPpSetInRangeMapInvalid:        "PpSetInRangeMapInvalid",
	mvdDescriptionPpSetAblendUnsupported:        "PpSetAblendUnsupported",
	mvdDescriptionPpSetDeinterlacingUnsupported: "PpSetDeinterlacingUnsupported",
	mvdDescriptionPpSetDitheringUnsupported:     "PpSetDitheringUnsupported",
	mvdDescriptionPpSetScalingUnsupported:       "PpSetScalingUnsupported",
	mvdDescriptionPpBusy:                        "PpBusy",
	mvdDescriptionHwReserved:                    "HwReserved",
	mvdDescriptionHwTimeout:                     "HwTimeout",
	mvdDescriptionHwBusError:                    "HwBusError",
	mvdDescriptionSystemError:                   "SystemError",
	mvdDescriptionDwlError:                      "DwlError",
	mvdDescriptionPpDecCombinedModeError:        "PpDecCombinedModeError",
	mvdDescriptionPpDecRuntimeError:             "PpDecRuntimeError",
	mvdDescriptionDriverErrorOffset2:            "DriverErrorOffset2",
	mvdDescriptionEvaluationLimitExceeded:       "EvaluationLimitExceeded",
	mvdDescriptionFormatNotSupported:            "FormatNotSupported",
	mvdDescriptionSdkReserved:                   "SdkReserved",
}

var mvdDescriptionDescription = map[mvdDescription]string{
	mvdDescriptionOk:                            "No error",
	mvdDescriptionStrmProcessed:                 "Stream buffer processed",
	mvdDescriptionPicRdy:                        "Picture ready",
	mvdDescriptionPicDecoded:                    "Picture decoded",
	mvdDescriptionHdrsRdy:                       "Headers decoded",
	mvdDescriptionAdvancedTools:                 "Advanced coding tools detected in stream",
	mvdDescriptionPendingFlush:                  "HW is busy but no more stream",
	mvdDescriptionNonrefPicSkipped:              "Non-reference picture skipped in decoding",
	mvdDescriptionDriverSucessOffset:            "DriverSucessOffset",
	mvdDescriptionSliceRdy:                      "HW finished processing a slice",
	mvdDescriptionLibsOffset:                    "LibsOffset",
	mvdDescriptionMemNotInitialized:             "Memory not initialized",
	mvdDescriptionMemAlreadyInitialized:         "Memory already initialized",
	mvdDescriptionInputMallocFail:               "Input allocation failed",
	mvdDescriptionDoneNothing:                   "Nothing was done",
	mvdDescriptionDriverErrorOffset1:            "DriverErrorOffset1",
	mvdDescriptionParamError:                    "Invalid parameters",
	mvdDescriptionStrmError:                     "Invalid stream structure",
	mvdDescriptionNotInitialized:                "Trying to do HW init without HW being initialized",
	mvdDescriptionMemfail:                       "Memory allocation failed",
	mvdDescriptionInitfail:                      "HW init failed",
	mvdDescriptionHdrsNotRdy:                    "Headers information is not ready",
	mvdDescriptionStreamNotSupported:            "The stream is not supported",
	mvdDescriptionPpSetInSizeInvalid:            "PP: Input size is invalid",
	mvdDescriptionPpSetInAddressInvalid:         "PP: Input address is invalid",
	mvdDescriptionPpSetInFormatInvalid:          "PP: Input format is invalid",
	mvdDescriptionPpSetCropInvalid:              "PP: Cropping size is invalid",
	mvdDescriptionPpSetRotationInvalid:          "PP: Rotation is invalid",
	mvdDescriptionPpSetOutSizeInvalid:           "PP: Output size is invalid",
	mvdDescriptionPpSetOutAddressInvalid:        "PP: Output address is invalid",
	mvdDescriptionPpSetOutFormatInvalid:         "PP: Output format is invalid",
	mvdDescriptionPpSetVideoAdjustInvalid:       "PP: Video adjustment is invalid",
	mvdDescriptionPpSetRgbBitmaskInvalid:        "PP: RGB bit mask is invalid",
	mvdDescriptionPpSetFramebufferInvalid:       "PP: Frame buffer is invalid",
	mvdDescriptionPpSetMask1Invalid:             "PP: Mask 1 is invalid",
	mvdDescriptionPpSetMask2Invalid:             "PP: Mask 2 is invalid",
	mvdDescriptionPpSetDeinterlaceInvalid:       "PP: Deinterlacing setting is invalid",
	mvdDescriptionPpSetInStructInvalid:          "PP: Input structure is invalid",
	mvdDescriptionPpSetInRangeMapInvalid:        "PP: Input range mapping is invalid",
	mvdDescriptionPpSetAblendUnsupported:        "PP: Alpha blending is not supported",
	mvdDescriptionPpSetDeinterlacingUnsupported: "PP: Deinterlacing is not supported",
	mvdDescriptionPpSetDitheringUnsupported:     "PP: Dithering is not supported",
	mvdDescriptionPpSetScalingUnsupported:       "PP: Scaling is not supported",
	mvdDescriptionPpBusy:                        "PP: Trying to set PP configuration while PP is busy",
	mvdDescriptionHwReserved:                    "HW: Reserved value",
	mvdDescriptionHwTimeout:                     "HW: Timeout occurred",
	mvdDescriptionHwBusError:                    "HW: A bus error occurred",
	mvdDescriptionSystemError:                   "System: A system error occurred",
	mvdDescriptionDwlError:                      "DWL: A DWL error occurred",
	mvdDescriptionPpDecCombinedModeError:        "PP: Combined mode not allowed",
	mvdDescriptionPpDecRuntimeError:             "PP: Runtime error occurred",
	mvdDescriptionDriverErrorOffset2:            "DriverErrorOffset2",
	mvdDescriptionEvaluationLimitExceeded:       "Evaluation limit exceeded",
	mvdDescriptionFormatNotSupported:            "Format not supported",
	mvdDescriptionSdkReserved:                   "SdkReserved",
}

func (d mvdDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", mvdDescriptionToString[d], d, mvdDescriptionDescription[d])
}
