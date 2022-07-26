package transaction

import "startup/users"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User users.User
}
