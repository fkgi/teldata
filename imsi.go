package teldata

import "encoding/json"

// IMSI value
type IMSI string

// ParseIMSI makes IMSI data from string
func ParseIMSI(s string) (IMSI, error) {
	for _, r := range s {
		if r < '0' || r > '9' {
			return IMSI(""), InvalidDataError{Name: "IMSI", Bytes: []byte(s)}
		}
	}
	if len(s) < 5 || len(s) > 15 {
		return IMSI(""), InvalidDataError{Name: "IMSI", Bytes: []byte(s)}
	}
	return IMSI(s), nil
}

func (i IMSI) String() string {
	return string(i)
}

func (i IMSI) IsEmpty() bool {
	return len(i) < 5 || len(i) > 15
}

// Bytes returns TBCD format byte data
func (i IMSI) Bytes() []byte {
	t, e := ParseTBCD(string(i))
	if e != nil {
		return []byte{}
	}
	return []byte(t)
}

// DecodeIMSI makes IMSI from this TBCD value
func DecodeIMSI(b []byte) (m IMSI, e error) {
	s := TBCD(b).String()
	return ParseIMSI(s)
}

// Length of this IMSI
func (i IMSI) Length() int {
	return len(i)
}

// MarshalJSON provide custom marshaller
func (i IMSI) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(i))
}

// UnmarshalJSON provide custom marshaller
func (i *IMSI) UnmarshalJSON(b []byte) (e error) {
	var s string
	if e = json.Unmarshal(b, &s); e != nil {
		return e
	}
	if len(s) != 0 {
		*i, e = ParseIMSI(s)
	}
	return e
}

// MCC value of this IMSI
func (i IMSI) MCC() string {
	if len(i) < 3 {
		return ""
	}
	return string(i[0:3])
}

// MNC2 return 2 charactor MNC value of this IMSI
func (i IMSI) MNC2() string {
	if len(i) < 5 {
		return ""
	}
	return string(i[3:5])
}

// MNC3 return 2 charactor MNC value of this IMSI
func (i IMSI) MNC3() string {
	if len(i) < 6 {
		return ""
	}
	return string(i[3:6])
}
