package auth

import "encoding/json"

// RegisterResponse represent Register Response
type RegisterResponse struct {
	Email          string `json:"email"`
	ActivationCode string `json:"-"`
}

// ToJSON covert to JSON
func (r *RegisterResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
