package descriptions

import (
	"fmt"

	"github.com/nikita-petko/ctr-result-parser/nn"
)

// GetDescription returns the string representation of the Description enum.
func GetDescription(result nn.Result) string {
	// If description is greater than MASK_DESCRIPTION - 23, it is a SDK description.
	if result.GetDescription() > int32(nn.MASK_DESCRIPTION-23) && result.GetDescription() < int32(nn.MASK_DESCRIPTION) {
		return nn.Description(result.GetDescription()).ToString()
	}

	switch result.GetModule() {
	case nn.ModuleKernel:
		return kernDescription(result.GetDescription()).toString()
	case nn.ModuleTcb:
		return tcbDescription(result.GetDescription()).toString()
	case nn.ModuleFnd:
		return fndDescription(result.GetDescription()).toString()
	case nn.ModuleOs:
		return osDescription(result.GetDescription()).toString()
	case nn.ModuleDbg:
		return dbgDescription(result.GetDescription()).toString()
	case nn.ModuleDmnt:
		return dmntDescription(result.GetDescription()).toString()
	case nn.ModulePdn:
		return pdnDescription(result.GetDescription()).toString()
	case nn.ModuleDd:
		return ddDescription(result.GetDescription()).toString()
	case nn.ModuleFs:
		return fsDescription(result.GetDescription()).toString()
	case nn.ModuleCamera:
		return cameraDescription(result.GetDescription()).toString()
	case nn.ModulePmLow:
		return pmlowDescription(result.GetDescription()).toString()
	case nn.ModuleSrv:
		return srvDescription(result.GetDescription()).toString()
	case nn.ModuleNdm:
		return ndmDescription(result.GetDescription()).toString()
	case nn.ModuleNwm:
		return nwmDescription(result.GetDescription()).toString()
	case nn.ModuleSocket:
		return socketDescription(result.GetDescription()).toString()
	case nn.ModuleLdr:
		return ldrDescription(result.GetDescription()).toString()
	case nn.ModuleAm:
		return amDescription(result.GetDescription()).toString()
	case nn.ModuleHio:
		return hioDescription(result.GetDescription()).toString()
	case nn.ModuleUpdater:
		return updaterDescription(result.GetDescription()).toString()
	case nn.ModuleMic:
		return micDescription(result.GetDescription()).toString()
	case nn.ModuleMp:
		return mpDescription(result.GetDescription()).toString()
	case nn.ModuleAc:
		return acDescription(result.GetDescription()).toString()
	case nn.ModuleHttp:
		return httpDescription(result.GetDescription()).toString()
	case nn.ModuleDsp:
		return dspDescription(result.GetDescription()).toString()
	case nn.ModuleSnd:
		return sndDescription(result.GetDescription()).toString()
	case nn.ModuleDlp:
		return dlpDescription(result.GetDescription()).toString()
	case nn.ModuleCsnd:
		return csndDescription(result.GetDescription()).toString()
	case nn.ModuleSsl:
		return sslDescription(result.GetDescription()).toString()
	case nn.ModuleFriends:
		return friendsDescription(result.GetDescription()).toString()
	case nn.ModuleRdt:
		return rdtDescription(result.GetDescription()).toString()
	case nn.ModuleApplet:
		return appletDescription(result.GetDescription()).toString()
	case nn.ModuleNim:
		return nimDescription(result.GetDescription()).toString()
	case nn.ModulePtm:
		return ptmDescription(result.GetDescription()).toString()
	case nn.ModuleMidi:
		return midiDescription(result.GetDescription()).toString()
	case nn.ModuleBoss:
		return bossDescription(result.GetDescription()).toString()
	case nn.ModuleDbm:
		return dbmDescription(result.GetDescription()).toString()
	case nn.ModuleCfg:
		return cfgDescription(result.GetDescription()).toString()
	case nn.ModuleCec:
		return cecDescription(result.GetDescription()).toString()
	case nn.ModuleIr:
		return irDescription(result.GetDescription()).toString()
	case nn.ModuleUds:
		return udsDescription(result.GetDescription()).toString()
	case nn.ModulePl:
		return plDescription(result.GetDescription()).toString()
	case nn.ModuleCup:
		return cupDescription(result.GetDescription()).toString()
	case nn.ModuleMcu:
		return mcuDescription(result.GetDescription()).toString()
	case nn.ModuleNs:
		return nsDescription(result.GetDescription()).toString()
	case nn.ModuleNews:
		return newsDescription(result.GetDescription()).toString()
	case nn.ModuleRo:
		return roDescription(result.GetDescription()).toString()
	case nn.ModuleGd:
		return gdDescription(result.GetDescription()).toString()
	case nn.ModuleCardSpi:
		return cardspiDescription(result.GetDescription()).toString()
	case nn.ModuleEc:
		return ecDescription(result.GetDescription()).toString()
	case nn.ModuleWebBrs:
		return webbrsDescription(result.GetDescription()).toString()
	case nn.ModuleEnc:
		return encDescription(result.GetDescription()).toString()
	case nn.ModuleVctl:
		return vctlDescription(result.GetDescription()).toString()
	case nn.ModuleL2b:
		return l2bDescription(result.GetDescription()).toString()
	case nn.ModuleMvd:
		return mvdDescription(result.GetDescription()).toString()
	case nn.ModuleNfc:
		return nfcDescription(result.GetDescription()).toString()
	case nn.ModuleQtm:
		return qtmDescription(result.GetDescription()).toString()
	default:
		if result.GetDescription() != int32(nn.DescriptionSuccess) && result.GetDescription() < int32(nn.MASK_DESCRIPTION-23) {
			return fmt.Sprintf("UnknownDescription (%d) - This description is not defined in any module.", result.GetDescription())
		}

		return nn.Description(result.GetDescription()).ToString()
	}
}
