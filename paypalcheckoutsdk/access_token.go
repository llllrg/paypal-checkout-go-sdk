package paypalcheckoutsdk

import "time"

type AccessToken struct {
	Token      string
	TokenType  string
	ExpiresIn  int64
	CreateDate time.Time
}

func NewAccessToken() *AccessToken {
	return &AccessToken{CreateDate: time.Now()}
}

func (a *AccessToken) IsExpired() bool {
	expireDate := a.CreateDate.Add(time.Duration(a.ExpiresIn) * time.Second)
	return time.Now().After(expireDate)
}
