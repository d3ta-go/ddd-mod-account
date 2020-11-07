package auth

import domSchemaAuth "github.com/d3ta-go/ddd-mod-account/modules/account/domain/schema/auth"

// LoginReqDTO type
type LoginReqDTO struct {
	domSchemaAuth.LoginRequest
}

// LoginResDTO type
type LoginResDTO struct {
	domSchemaAuth.LoginResponse
}
