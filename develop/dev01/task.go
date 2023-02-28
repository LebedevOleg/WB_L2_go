package main

import (
	"time"

	"github.com/beevik/ntp"
)

func GetCurrTime() (string, error) {
	res, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	return time.Now().Add(res.ClockOffset).Format(time.UnixDate), err
}

func GetExactTime() (string, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time.String(), err
}
