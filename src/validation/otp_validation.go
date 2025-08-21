package validation

import "time"

type QueryOtp struct {
	Page   int    `validate:"omitempty,number,max=50"`
	Limit  int    `validate:"omitempty,number,max=50"`
	Search string `validate:"omitempty,max=50"`
}
type UpdateOtp struct {
	OtpCode   string     `json:"otp_code" validate:"omitempty,len=6,numeric"`
	Purpose   string     `json:"purpose" validate:"omitempty,oneof=login register verify"`
	IsUsed    *bool      `json:"is_used"`
	ExpiresAt *time.Time `json:"expires_at"`
}
