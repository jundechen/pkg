// Copyright 2015, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package userjwt

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/brainattica/golang-jwt-authentication-api-sample/core/redis"
	"github.com/dgrijalva/jwt-go"
)

// JWTVerify is a middleware for echo to verify a JWT.
//func JWTVerify(dbrSess dbr.SessionRunner) func(http.Handler) http.Handler {
//
//	/*
//		@todo
//		1. load backend users from DB
//		2. use them to check the valid token, etc
//		3. create a polling service to update the cached backend user instead
//		   of querying for each request the database.
//		4. more stuff
//	*/
//
//	return func(c *echo.Context) error {
//		// Skip WebSocket
//		if (c.Request().Header.Get(echo.Upgrade)) == echo.WebSocket {
//			return nil
//		}
//		token, err := jwt.ParseFromRequest(c.Request(), func(token *jwt.Token) (interface{}, error) {
//			return []byte(`publicKey @todo`), nil
//		})
//		he := echo.NewHTTPError(http.StatusUnauthorized)
//
//		if err != nil {
//			log.Error("backend.JWTVerify.ParseFromRequest", "err", err, "req", c.Request())
//			he.SetCode(http.StatusBadRequest)
//			return he
//		}
//
//		if token.Valid {
//			return nil
//		}
//		// log.Info() ?
//
//		return he
//	}
//}

type blacklist struct{}

func (b blacklist) Set(_ string, _ time.Duration) error { return nil }
func (b blacklist) Has(_ string) bool                   { return false }

// AuthManager main object for handling JWT authentication, generation, blacklists and log outs.
type AuthManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	lastError  error
	// Expire defines the duration when the token is about to expire
	Expire time.Duration
	// SigningMethod sets the signing method pointer
	SigningMethod *jwt.SigningMethodRSA
	// EnableJTI activates the (JWT ID) Claim, a unique identifier. UUID.
	EnableJTI bool
	// Blacklist some kind of backend storage to handle blocked tokens.
	// Default black hole storage.
	Blacklist interface {
		Set(key string, expires time.Duration) error
		Has(key string) bool
	}
}

// NewAuthManager create a new manager. If private key option will not be
// passed then a key pair will be generated if both keys, or one of the two, are/is nil.
// Default expire is one hour. Default signing method is RS512.
func NewAuthManager(opts ...OptionFunc) (*AuthManager, error) {
	a := new(AuthManager)
	for _, opt := range opts {
		opt(a)
	}
	if a.lastError != nil {
		return nil, a.lastError
	}
	if a.privateKey == nil || a.publicKey == nil {
		generatePrivateKey(a)
	}
	if a.lastError != nil {
		return nil, a.lastError
	}
	a.Expire = time.Hour
	a.SigningMethod = jwt.SigningMethodRS512
	a.Blacklist = blacklist{}
	return a
}

// GenerateToken creates a new JSON web token. The claims argument will be
// assigned after the registered claim names exp and iat have been set.
// If EnableJTI is false the returned argument jti is empty.
// For details of the registered claim names please see
// http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html#rfc.section.4.1
func (a *AuthManager) GenerateToken(claims map[string]interface{}) (token, jti string, err error) {
	now := time.Now()
	t := jwt.New(a.SigningMethod)
	t.Claims["exp"] = now.Add(a.Expire).Unix()
	t.Claims["iat"] = now.Unix()
	for k, v := range claims {
		t.Claims[k] = v
	}
	if a.EnableJTI {
		jti = uuid.New()
		t.Claims["jti"] = jti
	}
	token, err = t.SignedString(a.privateKey)
	return
}

func tokenExpiresIn(timestamp interface{}) time.Duration {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}
	return 0
}

func (a *AuthManager) Logout(tokenString string, token *jwt.Token) error {

	return a.Blacklist.Set(tokenString, tokenExpiresIn(token.Claims["exp"]))
}

func (backend *AuthManager) IsInBlacklist(token string) bool {
	redisConn := redis.Connect()
	redisToken, _ := redisConn.GetValue(token)

	if redisToken == nil {
		return false
	}

	return true
}

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := InitJWTAuthenticationBackend()

	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
