package controller

type SubscribeResponse struct {
	SubscriptionStatus string `json:"subscriptionStatus"`
}

type GetCurrentBlockResponse struct {
	CurrentBlock int `json:"currentBlock"`
}
