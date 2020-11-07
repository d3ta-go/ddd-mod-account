package auth

import domSchemaAuth "github.com/d3ta-go/ddd-mod-account/modules/account/domain/schema/auth"

// RegisterReqDTO type
type RegisterReqDTO struct {
	domSchemaAuth.RegisterRequest
}

// RegisterResDTO type
type RegisterResDTO struct {
	Email string `json:"email"`
}
