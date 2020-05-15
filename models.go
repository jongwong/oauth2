package oauth2

import "encoding/json"

// TokenResponse token response.
type TokenResponse struct {
	AccessToken  string      `json:"access_token"`
	TokenType    string      `json:"token_type,omitempty"`
	ExpiresIn    int64       `json:"expires_in"`
	RefreshToken string      `json:"refresh_token,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Scope        string      `json:"scope,omitempty"`
}

// DeviceAuthorizationResponse Device Authorization Response.
// https://tools.ietf.org/html/rfc8628#section-3.2
type DeviceAuthorizationResponse struct {
	DeviceCode            string `json:"device_code"`
	UserCode              string `json:"user_code"`
	VerificationURI       string `json:"verification_uri"`
	VerificationURIQrcode string `json:"verification_uri_qrcode,omitempty"`
	ExpiresIn             int64  `json:"expires_in"`
	Interval              int    `json:"interval"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// CodeValue code值
type CodeValue struct {
	ClientID    string   `json:"client_id"`
	OpenID      string   `json:"open_id"`
	RedirectURI string   `json:"redirect_uri"`
	Scope       []string `json:"scope"`
}

// MarshalBinary json
func (code *CodeValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(code)
}

// UnmarshalBinary json
func (code *CodeValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, code)
}

// DeviceCodeValue device_code值
type DeviceCodeValue struct {
	ClientID   string   `json:"client_id"`
	OpenID     string   `json:"open_id"`
	DeviceCode string   `json:"device_code"`
	UserCode   string   `json:"user_code"`
	Scope      []string `json:"scope"`
}

// MarshalBinary json
func (code *DeviceCodeValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(code)
}

// UnmarshalBinary json
func (code *DeviceCodeValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, code)
}

// ClientBasic 客户端基础
type ClientBasic struct {
	ID     string `json:"client_id"`
	Secret string `json:"client_secret"`
}
