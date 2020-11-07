package auth

// ActivateRegistrationRequest represent ActivateRegistration Request
type ActivateRegistrationRequest struct {
	ActivationCode string `json:"activationCode"`
}
