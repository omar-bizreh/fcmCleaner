package models

// PushResponse Push notification check response
type PushResponse struct {
	MulticastID int32    `json:"multicast_id"`
	Success     int32    `json:"success"`
	Failed      int32    `json:"failure"`
	Results     []result `json:"results"`
}

type result struct {
	MessageID string `json:"message_id"`
	Error     string `json:"error"`
}
