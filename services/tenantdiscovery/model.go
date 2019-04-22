package tenantdiscovery

// Tenant represents a campaign openid configuration
type Tenant struct {
	ID        string   `json:"_id"`
	Hostnames []string `json:"hostnames"`
	Auth      Auth     `json:"auth"`
}

// Auth represents a openid method from campaign
type Auth struct {
	Method                   string  `json:"method"`
	OAuthAuthorizationServer *string `json:"oAuthAuthorizationServer"`
	OpenIDConfiguration      *string `json:"openIdConfiguration"`
	UserPoolID               *string `json:"userPoolId"`
}

// AuthMethod indicates the auth method options
type AuthMethod string

const (
	//Oidc ...
	Oidc AuthMethod = "oidc"
	//Classic ...
	Classic AuthMethod = "classic"
	// Cognito ...
	Cognito AuthMethod = "cognito"
)
