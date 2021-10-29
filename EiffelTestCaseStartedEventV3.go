// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
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

// THIS FILE IS AUTOMATICALLY GENERATED AND MUST NOT BE EDITED BY HAND.

package eiffelevents

import (
	"github.com/clarketm/json"
)

// MarshalJSON returns the JSON encoding of the event.
func (e *TestCaseStartedV3) MarshalJSON() ([]byte, error) {
	// The standard encoding/json package doesn't honor omitempty for
	// non-pointer structs (it doesn't recurse into values, only examines
	// the immediate value). This is a not terribly elegant way of making
	// sure that this struct is marshaled by github.com/clarketm/json
	// without the infinite loop we'd get by just passing the struct to
	// github.com/clarketm/json.Marshal.
	//
	// Make sure the links slice is non-null so that non-initialized slices
	// get serialized as "[]" instead of "null".
	links := e.Links
	if links == nil {
		links = make([]TCSV3Link, 0)
	}
	s := struct {
		Data  *TCSV3Data  `json:"data"`
		Links []TCSV3Link `json:"links"`
		Meta  *TCSV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *TestCaseStartedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type TestCaseStartedV3 struct {
	// Mandatory fields
	Data  TCSV3Data   `json:"data"`
	Links []TCSV3Link `json:"links"`
	Meta  TCSV3Meta   `json:"meta"`

	// Optional fields

}

type TCSV3Data struct {
	// Mandatory fields

	// Optional fields
	CustomData []TCSV3DataCustomDatum `json:"customData,omitempty"`
	Executor   string                 `json:"executor,omitempty"`
	LiveLogs   []TCSV3DataLiveLog     `json:"liveLogs,omitempty"`
}

type TCSV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TCSV3DataLiveLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields
	MediaType string   `json:"mediaType,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type TCSV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type TCSV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TCSV3MetaSecurity `json:"security,omitempty"`
	Source   TCSV3MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TCSV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection TCSV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []TCSV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type TCSV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       TCSV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                  `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type TCSV3MetaSecurityIntegrityProtectionAlg string

const (
	TCSV3MetaSecurityIntegrityProtectionAlg_HS256 TCSV3MetaSecurityIntegrityProtectionAlg = "HS256"
	TCSV3MetaSecurityIntegrityProtectionAlg_HS384 TCSV3MetaSecurityIntegrityProtectionAlg = "HS384"
	TCSV3MetaSecurityIntegrityProtectionAlg_HS512 TCSV3MetaSecurityIntegrityProtectionAlg = "HS512"
	TCSV3MetaSecurityIntegrityProtectionAlg_RS256 TCSV3MetaSecurityIntegrityProtectionAlg = "RS256"
	TCSV3MetaSecurityIntegrityProtectionAlg_RS384 TCSV3MetaSecurityIntegrityProtectionAlg = "RS384"
	TCSV3MetaSecurityIntegrityProtectionAlg_RS512 TCSV3MetaSecurityIntegrityProtectionAlg = "RS512"
	TCSV3MetaSecurityIntegrityProtectionAlg_ES256 TCSV3MetaSecurityIntegrityProtectionAlg = "ES256"
	TCSV3MetaSecurityIntegrityProtectionAlg_ES384 TCSV3MetaSecurityIntegrityProtectionAlg = "ES384"
	TCSV3MetaSecurityIntegrityProtectionAlg_ES512 TCSV3MetaSecurityIntegrityProtectionAlg = "ES512"
	TCSV3MetaSecurityIntegrityProtectionAlg_PS256 TCSV3MetaSecurityIntegrityProtectionAlg = "PS256"
	TCSV3MetaSecurityIntegrityProtectionAlg_PS384 TCSV3MetaSecurityIntegrityProtectionAlg = "PS384"
	TCSV3MetaSecurityIntegrityProtectionAlg_PS512 TCSV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type TCSV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type TCSV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
