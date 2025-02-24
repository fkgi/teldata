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
	TranslationType uint8           `json:"tt"`
	NatureOfAddress NatureOfAddress `json:"na"`
	NumberingPlan   NumberingPlan   `json:"np"`
	Digits          TBCD            `json:"digits"`
}

func (gt GlobalTitle) String() string {
	return fmt.Sprintf("%s (TT=%d NA=%s, NP=%s)",
		gt.Digits, gt.TranslationType, gt.NatureOfAddress, gt.NumberingPlan)
}

/*
	func (gt GlobalTitle) Bytes() []byte {
		return append(
			[]byte{0x80 | byte(gt.NatureOfAddress<<4) | byte(gt.NumberingPlan)},
			gt.Digits.Bytes()...)
	}
*/
func (gt GlobalTitle) IsEmpty() bool {
	return gt.Digits.Length() == 0
}

/*
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
*/

/*
NumberingPlan parameter for global title.
*/
type NumberingPlan byte

func (np NumberingPlan) String() string {
	switch np {
	case UnknownNP:
		return "unknown"
	case ISDNTelephony:
		return "telephony"
	case Generic:
		return "generic"
	case Data:
		return "data"
	case Telex:
		return "telex"
	case MaritimeMobile:
		return "maritime"
	case LandMobile:
		return "land"
	case ISDNMobile:
		return "mobile"
	case National:
		return "national"
	case Private:
		return "private"
	case ERMES:
		return "ermes"
	case Internet:
		return "internet"
	case NetworkSpecificPlan:
		return "nwspecific"
	case WAP:
		return "wap"
	}
	return fmt.Sprintf("undefined(%d)", np)
}

// MarshalJSON provide custom marshaller
func (np NumberingPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(np.String())
}

// UnmarshalJSON provide custom marshaller
func (np *NumberingPlan) UnmarshalJSON(b []byte) (e error) {
	var s string
	if e = json.Unmarshal(b, &s); e == nil {
		switch s {
		case "unknown":
			*np = UnknownNP
		case "telephony":
			*np = ISDNTelephony
		case "generic":
			*np = Generic
		case "data":
			*np = Data
		case "telex":
			*np = Telex
		case "maritime":
			*np = MaritimeMobile
		case "land":
			*np = LandMobile
		case "mobile":
			*np = ISDNMobile
		case "national":
			*np = National
		case "private":
			*np = Private
		case "ermes":
			*np = ERMES
		case "internet":
			*np = Internet
		case "nwspecific":
			*np = NetworkSpecificPlan
		case "wap":
			*np = WAP
		default:
			e = InvalidDataError{Name: "NP", Bytes: b}
		}
	}
	return
}

const (
	UnknownNP           NumberingPlan = 0x00 // unknown
	ISDNTelephony       NumberingPlan = 0x01 // ISDN/telephony (ITU-T E.163 and E.164)
	Generic             NumberingPlan = 0x02 // generic
	Data                NumberingPlan = 0x03 // data (ITU-T X.121)
	Telex               NumberingPlan = 0x04 // telex (ITU-T F.69)
	MaritimeMobile      NumberingPlan = 0x05 // maritime mobile (ITU-T E.210, E.211)
	LandMobile          NumberingPlan = 0x06 // land mobile (ITU-T E.212)
	ISDNMobile          NumberingPlan = 0x07 // ISDN/mobile (ITU-T E.214)
	National            NumberingPlan = 0x08 // national (MAP spec.)
	Private             NumberingPlan = 0x09 // private (MAP spec.)
	ERMES               NumberingPlan = 0x0a // european radio messaging system (ETSI DE/PS 3 01-3 spec.)
	Internet            NumberingPlan = 0x0d // Internet IP
	NetworkSpecificPlan NumberingPlan = 0x0e // private network or network-specific
	WAP                 NumberingPlan = 0x12 // WAP Client Id
)

/*
NatureOfAddress parameter for global title.
*/
type NatureOfAddress byte

func (na NatureOfAddress) String() string {
	switch na {
	case UnknownNA:
		return "unknown"
	case International:
		return "international"
	case NationalSignificant:
		return "national"
	case NetworkSpecific:
		return "nwspecific"
	case Subscriber:
		return "subscriber"
	case Alphanumeric:
		return "alphanumeric"
	case Abbreviated:
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
			*na = UnknownNA
		case "international":
			*na = International
		case "national":
			*na = NationalSignificant
		case "nwspecific":
			*na = NetworkSpecific
		case "subscriber":
			*na = Subscriber
		case "alphanumeric":
			*na = Alphanumeric
		case "abbreviated":
			*na = Abbreviated
		default:
			e = InvalidDataError{Name: "NoA", Bytes: b}
		}
	}
	return
}

const (
	UnknownNA           NatureOfAddress = 0x00 // unknown
	International       NatureOfAddress = 0x01 // international number
	NationalSignificant NatureOfAddress = 0x02 // national significant number
	NetworkSpecific     NatureOfAddress = 0x03 // reserved for national use (network specific number in MAP spec.)
	Subscriber          NatureOfAddress = 0x04 // subscriber number
	Alphanumeric        NatureOfAddress = 0x05 // alphanumeric address
	Abbreviated         NatureOfAddress = 0x06 // abbreviated (speed dial) number
)
