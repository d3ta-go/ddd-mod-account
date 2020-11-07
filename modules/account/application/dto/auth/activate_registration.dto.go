package auth

import domSchemaAuth "github.com/d3ta-go/ddd-mod-account/modules/account/domain/schema/auth"

// ActivateRegistrationReqDTO type
type ActivateRegistrationReqDTO struct {
	domSchemaAuth.ActivateRegistrationRequest
}

// ActivateRegistrationResDTO type
type ActivateRegistrationResDTO struct {
	domSchemaAuth.ActivateRegistrationResponse
}
