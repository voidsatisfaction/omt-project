package botService

import (
	"errors"
	"strconv"
)

func isHour(h string) bool {
	hNum, err := strconv.Atoi(h)
	if err != nil {
		return false
	}
	if hNum >= 0 && hNum <= 24 {
		return true
	}
	return false
}

func isMin(m string) bool {
	mNum, err := strconv.Atoi(m)
	if err != nil {
		return false
	}
	if mNum >= 0 && mNum <= 59 {
		return true
	}
	return false
}

func isColon(cp string) bool {
	return cp[0] == ':'
}

func (a *Action) ValidateTime() error {
	payloads := a.Payloads
	str := payloads[0]
	v := len(str)
	if v > 5 || v < 4 {
		return errors.New("error")
	}
	if v == 4 {
		str = "0" + str
	}

	hourPart := str[0:2]
	colonPart := str[2:3]
	minPart := str[3:5]

	if !isHour(hourPart) || !isMin(minPart) || !isColon(colonPart) {
		return errors.New("error")
	}
	return nil
}
