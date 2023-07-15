package descriptions

import "fmt"

type gdDescription int32

const (
	gdDescriptionSuccess gdDescription = iota
	gdDescriptionInvalidParameter
	gdDescriptionNullParameter
	gdDescriptionOutOfRange
	gdDescriptionOutOfMemory
	gdDescriptionExtMemoryAllocationFailed
	gdDescriptionInvalidMemoryRegion
	gdDescriptionInvalidFunctionCall
	gdDescriptionAlreadyReleased
	gdDescriptionInputlayoutInvalidStreamSlots
	gdDescriptionTexture2dInvalidResolution
	gdDescriptionTexture2dInvalidSubregionResolution
	gdDescriptionTexture2dInvalidFormat
	gdDescriptionTexture2dInvalidMemoryLayout
	gdDescriptionTexture2dInvalidMemoryLocation
	gdDescriptionTexture2dInvalidMiplevelIndex
	gdDescriptionTexture2dInvalidMiplevelIndexForMipmapAutogeneration
	gdDescriptionTexture2dInvalidFormatForCreatingRenderTarget
	gdDescriptionTexture2dInvalidFormatForCreatingDepthStencilTarget
	gdDescriptionDifferentRenderTargetAndDepthStencilTargetResolution
	gdDescriptionTexture2dInvalidTextureUnitId
	gdDescriptionTexture2dInvalidOffset
	gdDescriptionResourceAlreadyMapped
	gdDescriptionResourceWasNotMapped
	gdDescriptionNoTextureBound
	gdDescriptionNoTextureCoordinates
	gdDescriptionInvalidShaderUniformName
	gdDescriptionInvalidShaderUniform
	gdDescriptionInvalidShaderSignature
	gdDescriptionInvalidShaderOperation
	gdDescriptionGeometryShaderIncompatibleWithImmediateDraw
	gdDescriptionSystemNoCmdListBind
	gdDescriptionSystemInvalidCmdListBind
	gdDescriptionSystemReceiveErrorFromGlgetError
	gdDescriptionSystemNoPacketToRecord
	gdDescriptionSystemNoPacketRecorded
	gdDescriptionSystemAPacketIsAlreadyBeingRecorded
	gdDescriptionSystemInvalidPacketId
	gdDescriptionSystemRequestListInsertionIncompatibleWithJump
)

var gdDescriptionToString = map[gdDescription]string{
	gdDescriptionSuccess:                                              "Success",
	gdDescriptionInvalidParameter:                                     "InvalidParameter",
	gdDescriptionNullParameter:                                        "NullParameter",
	gdDescriptionOutOfRange:                                           "OutOfRange",
	gdDescriptionOutOfMemory:                                          "OutOfMemory",
	gdDescriptionExtMemoryAllocationFailed:                            "ExtMemoryAllocationFailed",
	gdDescriptionInvalidMemoryRegion:                                  "InvalidMemoryRegion",
	gdDescriptionInvalidFunctionCall:                                  "InvalidFunctionCall",
	gdDescriptionAlreadyReleased:                                      "AlreadyReleased",
	gdDescriptionInputlayoutInvalidStreamSlots:                        "InputlayoutInvalidStreamSlots",
	gdDescriptionTexture2dInvalidResolution:                           "Texture2dInvalidResolution",
	gdDescriptionTexture2dInvalidSubregionResolution:                  "Texture2dInvalidSubregionResolution",
	gdDescriptionTexture2dInvalidFormat:                               "Texture2dInvalidFormat",
	gdDescriptionTexture2dInvalidMemoryLayout:                         "Texture2dInvalidMemoryLayout",
	gdDescriptionTexture2dInvalidMemoryLocation:                       "Texture2dInvalidMemoryLocation",
	gdDescriptionTexture2dInvalidMiplevelIndex:                        "Texture2dInvalidMiplevelIndex",
	gdDescriptionTexture2dInvalidMiplevelIndexForMipmapAutogeneration: "Texture2dInvalidMiplevelIndexForMipmapAutogeneration",
	gdDescriptionTexture2dInvalidFormatForCreatingRenderTarget:        "Texture2dInvalidFormatForCreatingRenderTarget",
	gdDescriptionTexture2dInvalidFormatForCreatingDepthStencilTarget:  "Texture2dInvalidFormatForCreatingDepthStencilTarget",
	gdDescriptionDifferentRenderTargetAndDepthStencilTargetResolution: "DifferentRenderTargetAndDepthStencilTargetResolution",
	gdDescriptionTexture2dInvalidTextureUnitId:                        "Texture2dInvalidTextureUnitId",
	gdDescriptionTexture2dInvalidOffset:                               "Texture2dInvalidOffset",
	gdDescriptionResourceAlreadyMapped:                                "ResourceAlreadyMapped",
	gdDescriptionResourceWasNotMapped:                                 "ResourceWasNotMapped",
	gdDescriptionNoTextureBound:                                       "NoTextureBound",
	gdDescriptionNoTextureCoordinates:                                 "NoTextureCoordinates",
	gdDescriptionInvalidShaderUniformName:                             "InvalidShaderUniformName",
	gdDescriptionInvalidShaderUniform:                                 "InvalidShaderUniform",
	gdDescriptionInvalidShaderSignature:                               "InvalidShaderSignature",
	gdDescriptionInvalidShaderOperation:                               "InvalidShaderOperation",
	gdDescriptionGeometryShaderIncompatibleWithImmediateDraw:          "GeometryShaderIncompatibleWithImmediateDraw",
	gdDescriptionSystemNoCmdListBind:                                  "SystemNoCmdListBind",
	gdDescriptionSystemInvalidCmdListBind:                             "SystemInvalidCmdListBind",
	gdDescriptionSystemReceiveErrorFromGlgetError:                     "SystemReceiveErrorFromGlgetError",
	gdDescriptionSystemNoPacketToRecord:                               "SystemNoPacketToRecord",
	gdDescriptionSystemNoPacketRecorded:                               "SystemNoPacketRecorded",
	gdDescriptionSystemAPacketIsAlreadyBeingRecorded:                  "SystemAPacketIsAlreadyBeingRecorded",
	gdDescriptionSystemInvalidPacketId:                                "SystemInvalidPacketId",
	gdDescriptionSystemRequestListInsertionIncompatibleWithJump:       "SystemRequestListInsertionIncompatibleWithJump",
}

var gdDescriptionDescription = map[gdDescription]string{
	gdDescriptionSuccess:                                              "Represents success.",
	gdDescriptionInvalidParameter:                                     "A parameter is invalid.",
	gdDescriptionNullParameter:                                        "NULL has been specified for a parameter.",
	gdDescriptionOutOfRange:                                           "A parameter is out of range.",
	gdDescriptionOutOfMemory:                                          "Insufficient FCRAM memory region.",
	gdDescriptionExtMemoryAllocationFailed:                            "Insufficient VRAM memory region.",
	gdDescriptionInvalidMemoryRegion:                                  "The memory region is invalid.",
	gdDescriptionInvalidFunctionCall:                                  "Incorrect function call.",
	gdDescriptionAlreadyReleased:                                      "The resource has already been released.",
	gdDescriptionInputlayoutInvalidStreamSlots:                        "The input layout stream slots are invalid.",
	gdDescriptionTexture2dInvalidResolution:                           "Invalid 2D texture resolution.",
	gdDescriptionTexture2dInvalidSubregionResolution:                  "The region is invalid.",
	gdDescriptionTexture2dInvalidFormat:                               "Invalid 2D texture format.",
	gdDescriptionTexture2dInvalidMemoryLayout:                         "The memory layout is invalid.",
	gdDescriptionTexture2dInvalidMemoryLocation:                       "The memory location is invalid.",
	gdDescriptionTexture2dInvalidMiplevelIndex:                        "Invalid 2D texture mipmap level index.",
	gdDescriptionTexture2dInvalidMiplevelIndexForMipmapAutogeneration: "Invalid auto-generated 2D texture mipmap level index.",
	gdDescriptionTexture2dInvalidFormatForCreatingRenderTarget:        "Invalid 2D texture format for creating the render target.",
	gdDescriptionTexture2dInvalidFormatForCreatingDepthStencilTarget:  "Invalid 2D texture format for creating the depth stencil target.",
	gdDescriptionDifferentRenderTargetAndDepthStencilTargetResolution: "The render target and the depth stencil target are not compatible.",
	gdDescriptionTexture2dInvalidTextureUnitId:                        "The texture unit ID is invalid.",
	gdDescriptionTexture2dInvalidOffset:                               "The texture offset is invalid.",
	gdDescriptionResourceAlreadyMapped:                                "The resource has already been mapped.",
	gdDescriptionResourceWasNotMapped:                                 "The resource is not mapped.",
	gdDescriptionNoTextureBound:                                       "A texture unit was specified, but no texture has been set.",
	gdDescriptionNoTextureCoordinates:                                 "A texture unit was specified, but no texture coordinates have been set.",
	gdDescriptionInvalidShaderUniformName:                             "The shader uniform name is invalid.",
	gdDescriptionInvalidShaderUniform:                                 "The shader uniform is invalid.",
	gdDescriptionInvalidShaderSignature:                               "The shader signature is invalid.",
	gdDescriptionInvalidShaderOperation:                               "The shader is invalid.",
	gdDescriptionGeometryShaderIncompatibleWithImmediateDraw:          "Currently, the geometry shader cannot use immediate draw.",
	gdDescriptionSystemNoCmdListBind:                                  "The command list is not bound.",
	gdDescriptionSystemInvalidCmdListBind:                             "The bound command list is invalid.",
	gdDescriptionSystemReceiveErrorFromGlgetError:                     "A GetError error occurred in the function. (The error might have occurred due to a previous nngx or gd function call.)",
	gdDescriptionSystemNoPacketToRecord:                               "There is no data to save.",
	gdDescriptionSystemNoPacketRecorded:                               "The save process has not started.",
	gdDescriptionSystemAPacketIsAlreadyBeingRecorded:                  "The save process has already started.",
	gdDescriptionSystemInvalidPacketId:                                "The packet ID is invalid.",
	gdDescriptionSystemRequestListInsertionIncompatibleWithJump:       "All requested insertions in the request list are incompatible with jump recording.",
}

func (d gdDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", gdDescriptionToString[d], d, gdDescriptionDescription[d])
}
