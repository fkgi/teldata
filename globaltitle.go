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
	NatureOfAddress `json:"na"`
	NumberingPlan   `json:"np"`
	Digits          TBCD `json:"digits"`
}

func (gt *GlobalTitle) String() string {
	return fmt.Sprintf("%s(NA=%s, NP=%s)",
		gt.Digits, gt.NatureOfAddress, gt.NumberingPlan)
}

func (gt *GlobalTitle) Bytes() []byte {
	return append(
		[]byte{0x80 | byte(gt.NatureOfAddress<<4) | byte(gt.NumberingPlan)},
		gt.Digits.Bytes()...)
}

func DecodeGlobalTitle(data []byte) (a GlobalTitle, e error) {
	if len(data) == 0 {
		e = InvalidDataError{Name: "global title", Bytes: data}
	} else {
		a.NatureOfAddress = NatureOfAddress(data[0]&0x70) >> 4
		a.NumberingPlan = NumberingPlan(data[0] & 0x0f)
		a.Digits = make(TBCD, len(data)-1)
		copy(a.Digits, data[1:])
	}
	return
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

// MarshalJSON provide custom marshaller
func (np NumberingPlan) MarshalJSON() ([]byte, error) {
	fmt.Println("hoge")
	return json.Marshal(np.String())
}

// UnmarshalJSON provide custom marshaller
func (np *NumberingPlan) UnmarshalJSON(b []byte) (e error) {
	var s string
	if e = json.Unmarshal(b, &s); e == nil {
		switch s {
		case "unknown":
			*np = NP_Unknown
		case "telephony":
			*np = NP_ISDNTelephony
		case "generic":
			*np = NP_Generic
		case "data":
			*np = NP_Data
		case "telex":
			*np = NP_Telex
		case "maritime":
			*np = NP_MaritimeMobile
		case "land":
			*np = NP_LandMobile
		case "mobile":
			*np = NP_ISDNMobile
		case "national":
			*np = NP_National
		case "private":
			*np = NP_Private
		case "ermes":
			*np = NP_ERMES
		case "internet":
			*np = NP_Internet
		case "nwspecific":
			*np = NP_NetworkSpecific
		case "wap":
			*np = NP_WAP
		default:
			e = InvalidDataError{Name: "NP", Bytes: b}
		}
	}
	return
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

// MarshalJSON provide custom marshaller
func (na NatureOfAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(na.String())
}

// UnmarshalJSON provide custom marshaller
func (na *NatureOfAddress) UnmarshalJSON(b []byte) (e error) {
	var s string
	if e = json.Unmarshal(b, &s); e == nil {
		switch s {
		case "unknown":
			*na = NA_Unknown
		case "international":
			*na = NA_International
		case "national":
			*na = NA_National
		case "networkspecific":
			*na = NA_NetworkSpecific
		case "subscriber":
			*na = NA_Subscriber
		case "alphanumeric":
			*na = NA_Alphanumeric
		case "abbreviated":
			*na = NA_Abbreviated
		default:
			e = InvalidDataError{Name: "NoA", Bytes: b}
		}
	}
	return
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
