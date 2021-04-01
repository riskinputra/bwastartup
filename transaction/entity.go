package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID        int
	Amount    int
	Status    string
	Code      string
	User      user.User
	Campaign  campaign.Campaign
	CreatedAt time.Time
	UpdatedAt time.Time
}
