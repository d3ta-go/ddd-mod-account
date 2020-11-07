package auth

import "encoding/json"

// LoginResponse represent Login Response
type LoginResponse struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}

// ToJSON covert to JSON
func (r *LoginResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
