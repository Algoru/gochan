package gochan

import "strconv"

type ChanBool bool

func (cb *ChanBool) UnmarshalJSON(data []byte) error {
	str := string(data)
	value, err := strconv.ParseBool(str)
	if err != nil {
		return err
	}

	if value {
		*cb = true
	} else {
		*cb = false
	}

	return nil
}
