package teldata

import (
	"bytes"
	"encoding/json"
	"strings"
)

/*
TBCD is Telephony Binary Coded Decimal value that is defined in TS29.002.

TBCD-STRING ::= OCTET STRING

This type (Telephony Binary Coded Decimal String) is used to represent
several digits from 0 through 9, *, #, a, b, c, two digits per octet,
each digit encoded 0000 to 1001 (0 to 9),
1010 (*), 1011 (#), 1100 (a), 1101 (b) or 1110 (c); 1111 used as filler
when there is an odd number of digits.
bits 8765 of octet n encoding digit 2n bits 4321 of octet n encoding
digit 2(n-1) +1
*/
type TBCD []byte

// ParseTBCD create TBCD value from string
func ParseTBCD(s string) (TBCD, error) {
	if strings.ContainsRune(s, '\x00') {
		return nil, &InvalidDataError{
			Name:  "TBCD",
			Bytes: []byte(s)}
	} else if len(s)%2 != 0 {
		s = s + "\x00"
	}

	r := make([]byte, len(s)/2)
	for i, c := range s {
		var v byte
		switch c {
		case '0':
			v = 0x00
		case '1':
			v = 0x01
		case '2':
			v = 0x02
		case '3':
			v = 0x03
		case '4':
			v = 0x04
		case '5':
			v = 0x05
		case '6':
			v = 0x06
		case '7':
			v = 0x07
		case '8':
			v = 0x08
		case '9':
			v = 0x09
		case '*':
			v = 0x0a
		case '#':
			v = 0x0b
		case 'a', 'A':
			v = 0x0c
		case 'b', 'B':
			v = 0x0d
		case 'c', 'C':
			v = 0x0e
		case '\x00':
			v = 0x0f
		default:
			return r, &InvalidDataError{
				Name:  "TBCD",
				Bytes: []byte(string(c))}
		}
		if i%2 == 1 {
			v = v << 4
		}
		r[i/2] = r[i/2] | v
	}
	return r, nil
}

// NewTBCD create TBCD value from byte slice
func NewTBCD(b []byte) (TBCD, error) {
	for i := range b[:len(b)-1] {
		if b[i]&0xf0 == 0xf0 || b[i]&0x0f == 0x0f {
			return nil, &InvalidDataError{
				Name:  "TBCD",
				Bytes: b}
		}
	}
	if b[len(b)-1]&0x0f == 0x0f {
		return nil, &InvalidDataError{
			Name:  "TBCD",
			Bytes: b}
	}
	t := make([]byte, len(b))
	copy(t, b)
	return t, nil
}

// Length return length of the TBCD digit
func (t TBCD) Length() int {
	ret := len(t) * 2
	if ret != 0 && t[len(t)-1]&0xf0 == 0xf0 {
		ret--
	}
	return ret
}

// String return string value of the TBCD digit
func (t TBCD) String() string {
	if len(t) == 0 {
		return "N/A"
	}
	var b bytes.Buffer
	so := [2]byte{}
	for _, c := range t {
		so[0] = c & 0x0f
		so[1] = (c & 0xf0) >> 4
		for _, s := range so {
			switch s {
			case 0x00:
				b.WriteRune('0')
			case 0x01:
				b.WriteRune('1')
			case 0x02:
				b.WriteRune('2')
			case 0x03:
				b.WriteRune('3')
			case 0x04:
				b.WriteRune('4')
			case 0x05:
				b.WriteRune('5')
			case 0x06:
				b.WriteRune('6')
			case 0x07:
				b.WriteRune('7')
			case 0x08:
				b.WriteRune('8')
			case 0x09:
				b.WriteRune('9')
			case 0x0a:
				b.WriteRune('*')
			case 0x0b:
				b.WriteRune('#')
			case 0x0c:
				b.WriteRune('a')
			case 0x0d:
				b.WriteRune('b')
			case 0x0e:
				b.WriteRune('c')
			case 0x0f:
			}
		}
	}
	return b.String()
}

// Bytes return byte data
func (t TBCD) Bytes() []byte {
	ret := make([]byte, len(t))
	copy(ret, t)
	return ret
}

// MarshalJSON provide custom marshaller
func (t TBCD) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON provide custom marshaller
func (t *TBCD) UnmarshalJSON(b []byte) (e error) {
	var s string
	if e = json.Unmarshal(b, &s); e != nil {
		return e
	}
	if len(s) != 0 {
		*t, e = ParseTBCD(s)
	}
	return e
}
