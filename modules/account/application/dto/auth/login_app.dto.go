package auth

import domSchemaAuth "github.com/d3ta-go/ddd-mod-account/modules/account/domain/schema/auth"

// LoginAppReqDTO type
type LoginAppReqDTO struct {
	domSchemaAuth.LoginAppRequest
}

// LoginAppResDTO type
type LoginAppResDTO struct {
	domSchemaAuth.LoginAppResponse
}
