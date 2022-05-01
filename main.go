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

func (b Budget) TimeLeft() time.Duration { // b is receiver
	return b.Expires.Sub(time.Now().UTC())
}

func (b *Budget) Update(amount float64) { // b is pointer receiver
	b.Balance += amount
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("CampaignID is required")
	}
	if balance <= 0 {
		return nil, fmt.Errorf("Balance must be positive")
	}
	if expires.Before(time.Now().UTC()) {
		return nil, fmt.Errorf("Expires must be in the future")
	}
	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}

	return &b, nil
}

func main() {
	expires := time.Now().UTC().Add(7 * 24 * time.Hour)
	b1, err := NewBudget("Puppies", 500, expires)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", b1)
	}

	b2, err := NewBudget("Donkeys", 0, expires)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", b2)
	}

	b := Budget{
		CampaignID: "Kittens",
		Balance:    100,
		Expires:    time.Now().UTC().Add(time.Hour * 24 * 7),
	}

	b.Update(50)
	fmt.Println(b.TimeLeft())
	fmt.Println(b.Balance)
}
