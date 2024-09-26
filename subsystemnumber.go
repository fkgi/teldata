package teldata

import "encoding/json"

type SubsystemNumber byte

const (
	_ SubsystemNumber = iota
	SsnMgmt
	SsnHLR
	SsnVLR
	SsnMSC
	SsnEIR
)

func (n SubsystemNumber) String() string {
	switch n {
	case SsnMgmt:
		return "mgmt"
	case SsnHLR:
		return "hlr"
	case SsnVLR:
		return "vlr"
	case SsnMSC:
		return "msc"
	case SsnEIR:
		return "eir"
	}
	return "unknown"
}

func (n SubsystemNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *SubsystemNumber) UnmarshalJSON(b []byte) (e error) {
	var s string
	e = json.Unmarshal(b, &s)
	switch s {
	case "mgmn":
		*n = SsnMgmt
	case "hlr":
		*n = SsnHLR
	case "vlr":
		*n = SsnVLR
	case "msc":
		*n = SsnMSC
	case "eir":
		*n = SsnEIR
	default:
		*n = 0
	}
	return
}

func (n SubsystemNumber) Uint() uint8 {
	switch n {
	case SsnMgmt:
		return 1
	case SsnHLR:
		return 6
	case SsnVLR:
		return 7
	case SsnMSC:
		return 8
	case SsnEIR:
		return 9
	}
	return 0
}

func ParseSSN(i uint8) SubsystemNumber {
	switch i {
	case 1:
		return SsnMgmt
	case 6:
		return SsnHLR
	case 7:
		return SsnVLR
	case 8:
		return SsnMSC
	case 9:
		return SsnEIR
	}
	return 0
}
