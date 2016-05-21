// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
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

package mwjwt

import "github.com/corestoreio/csfw/util/errors"

var errContextJWTNotFound = errors.NewNotFoundf(`[mwjwt] Cannot extract token nor an error from the context`)

const errServiceUnsupportedScope = "[mwjwt] Service does not support this: %s. Only default or website scope are allowed."

const errTokenParseNotValidOrBlackListed = "[mwjwt] Token not valid or black listed"

const errScopedConfigMissingSigningMethod = "[mwjwt] Incomplete configuration for %s. Missing Signing Method and its Key."

const errConfigNotFound = "[mwjwt] Cannot find JWT configuration for %s"

const errUnknownSigningMethod = "[mwjwt] Unknown signing method - Have: %q Want: %q"

const errUnknownSigningMethodOptions = "[mwjwt] Unknown signing method - Have: %q Want: ES, HS or RS"

const errKeyEmpty = "[mwjwt] Provided key argument is empty"

// ErrTokenBlacklisted returned by the middleware if the token can be found
// within the black list.
const errTokenBlacklisted = "[mwjwt] Token has been black listed"

// ErrTokenInvalid returned by the middleware to make understandable that
// a token has been invalidated.
const errTokenInvalid = "[mwjwt] Token has become invalid"

const errStoreNotFound = "[mwjwt] Store not found in token claim"