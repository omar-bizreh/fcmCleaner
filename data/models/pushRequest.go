package models

// PushRequest Push notification request
type PushRequest struct {
	PushTokens []string `json:"registration_ids"`
	DryRun     bool     `json:"dry_run"`
}
