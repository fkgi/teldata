package teldata

/*
GlobalTitle is an address used in the SCCP protocol for routing
signaling messages on telecommunications networks.
*/
type GlobalTitle struct {
	NatureOfAddress
	NumberingPlan
	Digits TBCD
}

/*
NumberingPlan parameter for global title.
*/
type NumberingPlan byte

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

const (
	NA_Unknown         NatureOfAddress = 0x00 // unknown
	NA_International   NatureOfAddress = 0x01 // international number
	NA_National        NatureOfAddress = 0x02 // national significant number
	NA_NetworkSpecific NatureOfAddress = 0x03 // reserved for national use (network specific number in MAP spec.)
	NA_Subscriber      NatureOfAddress = 0x04 // subscriber number
	NA_Alphanumeric    NatureOfAddress = 0x05 // apphanumeric address
	NA_Abbreviated     NatureOfAddress = 0x06 // abbreviated (spped dial) number
)
