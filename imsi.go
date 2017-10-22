package teldata

// IMSI value
type IMSI struct {
	string
}

// ParseIMSI makes IMSI data from string
func ParseIMSI(s string) (IMSI, error) {
	for _, r := range s {
		if r < '0' || r > '9' {
			return IMSI{}, InvalidDataError{Name: "IMSI", Bytes: []byte(s)}
		}
	}
	if len(s) < 5 || len(s) > 15 {
		return IMSI{}, InvalidDataError{Name: "IMSI", Bytes: []byte(s)}
	}
	return IMSI{s}, nil
}

func (i IMSI) String() string {
	return i.string
}

// Length of this IMSI
func (i IMSI) Length() int {
	return len(i.string)
}

// MCC value of this IMSI
func (i IMSI) MCC() string {
	return i.string[0:3]
}

// MNC2 return 2 charactor MNC value of this IMSI
func (i IMSI) MNC2() string {
	return i.string[3:5]
}

// MNC3 return 2 charactor MNC value of this IMSI
func (i IMSI) MNC3() string {
	return i.string[3:6]
}
