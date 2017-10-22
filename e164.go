package teldata

// E164 value
type E164 struct {
	tbcd TBCD
}

// ParseE164 makes E164 data from string
func ParseE164(s string) (m E164, e error) {
	m.tbcd, e = ParseTBCD(s)
	if e == nil && m.tbcd.Length() > 15 {
		e = InvalidDataError{Name: "E164", Bytes: []byte(s)}
	}
	return
}

func (m E164) String() string {
	return m.tbcd.String()
}

// Bytes returns TBCD format byte data
func (m E164) Bytes() []byte {
	r := make([]byte, len(m.tbcd))
	copy(r, m.tbcd)
	return r
}

// ToE164 makes E164 from this TBCD value
func ToE164(b []byte) (m E164, e error) {
	m.tbcd = make([]byte, len(b))
	copy(m.tbcd, b)
	if m.tbcd.Length() > 15 {
		e = InvalidDataError{Name: "E164", Bytes: b}
	}
	return
}

// Length of this E164
func (m E164) Length() int {
	return m.tbcd.Length()
}
