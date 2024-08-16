package teldata

import (
	"fmt"
)

/*
GlobalTitle is an address used in the SCCP protocol for routing
signaling messages on telecommunications networks.
*/
type GlobalTitle struct {
	NatureOfAddress
	NumberingPlan
	Digits TBCD
}

func (gt *GlobalTitle) String() string {
	return fmt.Sprintf("%s(NA=%s, NP=%s)",
		gt.Digits, gt.NatureOfAddress, gt.NumberingPlan)
}

/*
NumberingPlan parameter for global title.
*/
type NumberingPlan byte

func (np NumberingPlan) String() string {
	switch np {
	case NP_Unknown:
		return "unknown"
	case NP_ISDNTelephony:
		return "ISDN/Telephony"
	case NP_Generic:
		return "generic"
	case NP_Data:
		return "data"
	case NP_Telex:
		return "telex"
	case NP_MaritimeMobile:
		return "maritime mobile"
	case NP_LandMobile:
		return "land mobile"
	case NP_ISDNMobile:
		return "ISDN/mobile"
	case NP_National:
		return "MAP national"
	case NP_Private:
		return "MAP private"
	case NP_ERMES:
		return "european radio messaging system"
	case NP_Internal:
		return "Internet IP"
	case NP_NetworkSpecific:
		return "private network or network-specific"
	case NP_WAP:
		return "WAP Client Id"
	}
	return fmt.Sprintf("undefined(%d)", np)
}

const (
	NP_Unknown         NumberingPlan = 0x00 // unknown
	NP_ISDNTelephony   NumberingPlan = 0x01 // ISDN/telephony (ITU-T E.163 and E.164)
	NP_Generic         NumberingPlan = 0x02 // generic
	NP_Data            NumberingPlan = 0x03 // data (ITU-T X.121)
	NP_Telex           NumberingPlan = 0x04 // telex (ITU-T F.69)
	NP_MaritimeMobile  NumberingPlan = 0x05 // maritime mobile (ITU-T E.210, E.211)
	NP_LandMobile      NumberingPlan = 0x06 // land mobile (ITU-T E.212)
	NP_ISDNMobile      NumberingPlan = 0x07 // ISDN/mobile (ITU-T E.214)
	NP_National        NumberingPlan = 0x08 // national (MAP spec.)
	NP_Private         NumberingPlan = 0x09 // private (MAP spec.)
	NP_ERMES           NumberingPlan = 0x0a // european radio messaging system (ETSI DE/PS 3 01-3 spec.)
	NP_Internal        NumberingPlan = 0x0d // Internet IP
	NP_NetworkSpecific NumberingPlan = 0x0e // private network or network-specific
	NP_WAP             NumberingPlan = 0x12 // WAP Client Id
)

/*
NatureOfAddress parameter for global title.
*/
type NatureOfAddress byte

func (na NatureOfAddress) String() string {
	switch na {
	case NA_Unknown:
		return "unknown"
	case NA_International:
		return "international number"
	case NA_National:
		return "national significant number"
	case NA_NetworkSpecific:
		return "reserved for national use"
	case NA_Subscriber:
		return "subscriber number"
	case NA_Alphanumeric:
		return "apphanumeric address"
	case NA_Abbreviated:
		return "abbreviated (spped dial) number"
	}
	return fmt.Sprintf("undefined(%d)", na)
}

const (
	NA_Unknown         NatureOfAddress = 0x00 // unknown
	NA_International   NatureOfAddress = 0x01 // international number
	NA_National        NatureOfAddress = 0x02 // national significant number
	NA_NetworkSpecific NatureOfAddress = 0x03 // reserved for national use (network specific number in MAP spec.)
	NA_Subscriber      NatureOfAddress = 0x04 // subscriber number
	NA_Alphanumeric    NatureOfAddress = 0x05 // apphanumeric address
	NA_Abbreviated     NatureOfAddress = 0x06 // abbreviated (spped dial) number
)
