package teldata

import (
	"encoding/json"
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

func (gt GlobalTitle) MarshalJSON() ([]byte, error) {
	tmp := struct {
		NA  string `json:"na"`
		NP  string `json:"np"`
		Dig string `json:"digits"`
	}{
		gt.NatureOfAddress.String(),
		gt.NumberingPlan.String(),
		gt.Digits.String(),
	}
	return json.Marshal(tmp)
}

func (gt *GlobalTitle) UnmarshalJSON(b []byte) (e error) {
	tmp := struct {
		NA  string `json:"na"`
		NP  string `json:"np"`
		Dig string `json:"digits"`
	}{}
	if e = json.Unmarshal(b, &tmp); e != nil {
		return e
	}
	switch tmp.NA {
	case "unknown":
		gt.NatureOfAddress = NA_Unknown
	case "international":
		gt.NatureOfAddress = NA_International
	case "national":
		gt.NatureOfAddress = NA_National
	case "networkspecific":
		gt.NatureOfAddress = NA_NetworkSpecific
	case "subscriber":
		gt.NatureOfAddress = NA_Subscriber
	case "alphanumeric":
		gt.NatureOfAddress = NA_Alphanumeric
	case "abbreviated":
		gt.NatureOfAddress = NA_Abbreviated
	}
	switch tmp.NP {
	case "unknown":
		gt.NumberingPlan = NP_Unknown
	case "telephony":
		gt.NumberingPlan = NP_ISDNTelephony
	case "generic":
		gt.NumberingPlan = NP_Generic
	case "data":
		gt.NumberingPlan = NP_Data
	case "telex":
		gt.NumberingPlan = NP_Telex
	case "maritime":
		gt.NumberingPlan = NP_MaritimeMobile
	case "land":
		gt.NumberingPlan = NP_LandMobile
	case "mobile":
		gt.NumberingPlan = NP_ISDNMobile
	case "national":
		gt.NumberingPlan = NP_National
	case "private":
		gt.NumberingPlan = NP_Private
	case "ermes":
		gt.NumberingPlan = NP_ERMES
	case "internet":
		gt.NumberingPlan = NP_Internet
	case "nwspecific":
		gt.NumberingPlan = NP_NetworkSpecific
	case "wap":
		gt.NumberingPlan = NP_WAP

	}
	gt.Digits, e = ParseTBCD(tmp.Dig)
	return e
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
		return "telephony"
	case NP_Generic:
		return "generic"
	case NP_Data:
		return "data"
	case NP_Telex:
		return "telex"
	case NP_MaritimeMobile:
		return "maritime"
	case NP_LandMobile:
		return "land"
	case NP_ISDNMobile:
		return "mobile"
	case NP_National:
		return "national"
	case NP_Private:
		return "private"
	case NP_ERMES:
		return "ermes"
	case NP_Internet:
		return "internet"
	case NP_NetworkSpecific:
		return "nwspecific"
	case NP_WAP:
		return "wap"
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
	NP_Internet        NumberingPlan = 0x0d // Internet IP
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
		return "international"
	case NA_National:
		return "national"
	case NA_NetworkSpecific:
		return "networkspecific"
	case NA_Subscriber:
		return "subscriber"
	case NA_Alphanumeric:
		return "alphanumeric"
	case NA_Abbreviated:
		return "abbreviated"
	}
	return fmt.Sprintf("undefined(%d)", na)
}

const (
	NA_Unknown         NatureOfAddress = 0x00 // unknown
	NA_International   NatureOfAddress = 0x01 // international number
	NA_National        NatureOfAddress = 0x02 // national significant number
	NA_NetworkSpecific NatureOfAddress = 0x03 // reserved for national use (network specific number in MAP spec.)
	NA_Subscriber      NatureOfAddress = 0x04 // subscriber number
	NA_Alphanumeric    NatureOfAddress = 0x05 // alphanumeric address
	NA_Abbreviated     NatureOfAddress = 0x06 // abbreviated (speed dial) number
)
