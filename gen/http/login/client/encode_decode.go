// Code generated by goa v3.10.2, DO NOT EDIT.
//
// login HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/codeready-toolchain/sandbox-auth/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	login "github.com/codeready-toolchain/sandbox-auth/gen/login"
	goahttp "goa.design/goa/v3/http"
)

// BuildLoginRequest instantiates a HTTP request object with method and path
// set to call the "login" service "login" endpoint
func (c *Client) BuildLoginRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginLoginPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("login", "login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginRequest returns an encoder for requests sent to the login login
// server.
func EncodeLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*login.LoginCriteria)
		if !ok {
			return goahttp.ErrInvalidType("login", "login", "*login.LoginCriteria", v)
		}
		if p.Referer != nil {
			head := *p.Referer
			req.Header.Set("Referer", head)
		}
		values := req.URL.Query()
		if p.Redirect != nil {
			values.Add("redirect", *p.Redirect)
		}
		if p.APIClient != nil {
			values.Add("api_client", *p.APIClient)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeLoginResponse returns a decoder for responses returned by the login
// login endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusTemporaryRedirect:
			var (
				body LoginTemporaryRedirectResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("login", "login", err)
			}
			var (
				location *string
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw != "" {
				location = &locationRaw
			}
			res := NewLoginResultTemporaryRedirect(&body, location)
			tmp := "redirect"
			res.Outcome = &tmp
			return res, nil
		case http.StatusUnauthorized:
			var (
				body LoginUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("login", "login", err)
			}
			res := NewLoginResultUnauthorized(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("login", "login", resp.StatusCode, string(body))
		}
	}
}

// BuildCallbackRequest instantiates a HTTP request object with method and path
// set to call the "login" service "callback" endpoint
func (c *Client) BuildCallbackRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CallbackLoginPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("login", "callback", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCallbackRequest returns an encoder for requests sent to the login
// callback server.
func EncodeCallbackRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*login.LoginCallbackCriteria)
		if !ok {
			return goahttp.ErrInvalidType("login", "callback", "*login.LoginCallbackCriteria", v)
		}
		values := req.URL.Query()
		if p.Code != nil {
			values.Add("code", *p.Code)
		}
		if p.State != nil {
			values.Add("state", *p.State)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeCallbackResponse returns a decoder for responses returned by the login
// callback endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCallbackResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusTemporaryRedirect:
			var (
				body CallbackTemporaryRedirectResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("login", "callback", err)
			}
			var (
				location *string
			)
			locationRaw := resp.Header.Get("Location")
			if locationRaw != "" {
				location = &locationRaw
			}
			res := NewCallbackLoginResultTemporaryRedirect(&body, location)
			tmp := "redirect"
			res.Outcome = &tmp
			return res, nil
		case http.StatusUnauthorized:
			var (
				body CallbackUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("login", "callback", err)
			}
			res := NewCallbackLoginResultUnauthorized(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("login", "callback", resp.StatusCode, string(body))
		}
	}
}
