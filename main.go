package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

func main() {
	b := Budget{
		CampaignID: "Kittens",
		Balance:    100,
		Expires:    time.Now().UTC().Add(time.Hour * 24 * 7),
	}

	fmt.Println(b.TimeLeft())
}
