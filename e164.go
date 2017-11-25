package teldata

import "encoding/json"

// E164 value
type E164 TBCD

// ParseE164 makes E164 data from string
func ParseE164(s string) (m E164, e error) {
	t, e := ParseTBCD(s)
	if e == nil && t.Length() > 15 {
		e = InvalidDataError{Name: "E164", Bytes: []byte(s)}
	}
	m = E164(t)
	return
}

func (m E164) String() string {
	return TBCD(m).String()
}

// Bytes returns TBCD format byte data
func (m E164) Bytes() []byte {
	r := make([]byte, len(m))
	copy(r, m)
	return r
}

// ToE164 makes E164 from this TBCD value
func ToE164(b []byte) (m E164, e error) {
	m = make([]byte, len(b))
	copy(m, b)
	if m.Length() > 15 {
		e = InvalidDataError{Name: "E164", Bytes: b}
	}
	return
}

// Length of this E164
func (m E164) Length() int {
	return TBCD(m).Length()
}

// MarshalJSON provide custom marshaller
func (m E164) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// UnmarshalJSON provide custom marshaller
func (m *E164) UnmarshalJSON(b []byte) (e error) {
	var s string
	if e = json.Unmarshal(b, &s); e != nil {
		return e
	}
	if len(s) != 0 {
		*m, e = ParseE164(s)
	}
	return e
}
