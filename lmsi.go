package teldata

import (
	"encoding/hex"
	"encoding/json"
)

// LMSI value
type LMSI [4]byte

func (l LMSI) String() string {
	return hex.EncodeToString(l[:])
}

func (l LMSI) IsEmpty() bool {
	return l[0] == 0 && l[1] == 0 && l[2] == 0 && l[3] == 0
}

// Bytes returns byte data
func (l LMSI) Bytes() []byte {
	return l[:]
}

// DecodeLMSI makes IMSI from this TBCD value
func DecodeLMSI(b []byte) (l LMSI, e error) {
	if len(b) != 4 {
		e = InvalidDataError{Name: "LMSI", Bytes: b}
	} else {
		for i := range l {
			l[i] = b[i]
		}
	}
	return
}

// MarshalJSON provide custom marshaller
func (l LMSI) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

// UnmarshalJSON provide custom marshaller
func (l *LMSI) UnmarshalJSON(b []byte) (e error) {
	var s string
	if e = json.Unmarshal(b, &s); e != nil {
	} else if b, e = hex.DecodeString(s); e == nil {
		*l, e = DecodeLMSI(b)
	}
	return
}
