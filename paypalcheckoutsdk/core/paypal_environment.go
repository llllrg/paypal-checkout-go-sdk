package core

import "encoding/base64"

const (
	LIVE_API_URL    = "https://api.paypal.com"
	LIVE_WEB_URL    = "https://www.paypal.com"
	SANDBOX_API_URL = "https://api.sandbox.paypal.com"
	SANDBOX_WEB_URL = "https://www.sandbox.paypal.com"
)

type ENV uint

const (
	SANDBOX ENV = iota + 1
	LIVE
)

type PayPalEnvironment struct {
	clientId     string
	clientSecret string
	webUrl       string
	apiUrl       string
	environment  ENV
}

func NewPayPalEnvironment(clientId, clientSecret string, environment ENV) *PayPalEnvironment {
	env := &PayPalEnvironment{
		clientId:     clientId,
		clientSecret: clientSecret,
		webUrl:       LIVE_WEB_URL,
		apiUrl:       LIVE_API_URL,
		environment:  LIVE,
	}
	if env.environment == SANDBOX {
		env.environment = SANDBOX
		env.webUrl = SANDBOX_WEB_URL
		env.apiUrl = SANDBOX_API_URL

	}
	return env

}

func (p *PayPalEnvironment) AuthorizationString() string {
	return base64.StdEncoding.EncodeToString([]byte(p.clientId + ":" + p.clientSecret))
}
