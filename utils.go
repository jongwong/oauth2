package oauth2

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/nilorg/sdk/random"
)

// RequestClientBasic 获取请求中的客户端信息
func RequestClientBasic(r *http.Request) (basic *ClientBasic, err error) {
	username, password, ok := r.BasicAuth()
	if !ok {
		err = ErrInvalidClient
		return
	}
	basic = &ClientBasic{
		ID:     username,
		Secret: password,
	}
	return
}
func writerJSON(w http.ResponseWriter, statusCode int, value interface{}) (err error) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", contentTypeJSON)
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
	err = json.NewEncoder(w).Encode(value)
	return
}

// WriterJSON 写入Json
func WriterJSON(w http.ResponseWriter, value interface{}) {
	err := writerJSON(w, http.StatusOK, value)
	if err != nil {
		panic(err)
	}
	return
}

// WriterError 写入Error
func WriterError(w http.ResponseWriter, err error) {
	statusCode := http.StatusBadRequest
	if code, ok := ErrStatusCodes[err]; ok {
		statusCode = code
	}
	if werr := writerJSON(w, statusCode, &ErrorResponse{
		Error: err.Error(),
	}); werr != nil {
		panic(werr)
	}
}

// RedirectSuccess 重定向成功
func RedirectSuccess(w http.ResponseWriter, r *http.Request, redirectURI *url.URL, code string) {
	query := redirectURI.Query()
	query.Set(CodeKey, code)
	query.Set(StateKey, r.FormValue(StateKey))
	redirectURI.RawQuery = query.Encode()
	http.Redirect(w, r, redirectURI.String(), http.StatusFound)
}

// RedirectError 重定向错误
func RedirectError(w http.ResponseWriter, r *http.Request, redirectURI *url.URL, err error) {
	query := redirectURI.Query()
	query.Set(ErrorKey, err.Error())
	query.Set(StateKey, r.FormValue(StateKey))
	redirectURI.RawQuery = query.Encode()
	http.Redirect(w, r, redirectURI.String(), http.StatusFound)
}

// RandomState 随机State
func RandomState() string {
	return random.AZaz09(6)
}

// RandomCode 随机Code
func RandomCode() string {
	return random.AZaz09(12)
}
