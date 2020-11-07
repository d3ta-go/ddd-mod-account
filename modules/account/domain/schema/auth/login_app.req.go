package auth

// LoginAppRequest represent LoginApp Request
type LoginAppRequest struct {
	ClientKey string `json:"clientKey"`
	SecretKey string `json:"secretKey"`
}
