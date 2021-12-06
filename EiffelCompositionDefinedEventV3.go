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
	"time"

	"github.com/clarketm/json"
	"github.com/google/uuid"
)

// NewCompositionDefinedV3 creates a new struct pointer that represents
// major version 3 of EiffelCompositionDefinedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 3.x.x
// currently known by this SDK.
func NewCompositionDefinedV3() (*CompositionDefinedV3, error) {
	var event CompositionDefinedV3
	event.Meta.Type = "EiffelCompositionDefinedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][3].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *CompositionDefinedV3) MarshalJSON() ([]byte, error) {
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
		links = make([]CDV3Link, 0)
	}
	s := struct {
		Data  *CDV3Data  `json:"data"`
		Links []CDV3Link `json:"links"`
		Meta  *CDV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *CompositionDefinedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type CompositionDefinedV3 struct {
	// Mandatory fields
	Data  CDV3Data   `json:"data"`
	Links []CDV3Link `json:"links"`
	Meta  CDV3Meta   `json:"meta"`

	// Optional fields

}

type CDV3Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	CustomData []CDV3DataCustomDatum `json:"customData,omitempty"`
	Version    string                `json:"version,omitempty"`
}

type CDV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type CDV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type CDV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security CDV3MetaSecurity `json:"security,omitempty"`
	Source   CDV3MetaSource   `json:"source,omitempty"`
	Tags     []string         `json:"tags,omitempty"`
}

type CDV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection CDV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []CDV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type CDV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       CDV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                 `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type CDV3MetaSecurityIntegrityProtectionAlg string

const (
	CDV3MetaSecurityIntegrityProtectionAlg_HS256 CDV3MetaSecurityIntegrityProtectionAlg = "HS256"
	CDV3MetaSecurityIntegrityProtectionAlg_HS384 CDV3MetaSecurityIntegrityProtectionAlg = "HS384"
	CDV3MetaSecurityIntegrityProtectionAlg_HS512 CDV3MetaSecurityIntegrityProtectionAlg = "HS512"
	CDV3MetaSecurityIntegrityProtectionAlg_RS256 CDV3MetaSecurityIntegrityProtectionAlg = "RS256"
	CDV3MetaSecurityIntegrityProtectionAlg_RS384 CDV3MetaSecurityIntegrityProtectionAlg = "RS384"
	CDV3MetaSecurityIntegrityProtectionAlg_RS512 CDV3MetaSecurityIntegrityProtectionAlg = "RS512"
	CDV3MetaSecurityIntegrityProtectionAlg_ES256 CDV3MetaSecurityIntegrityProtectionAlg = "ES256"
	CDV3MetaSecurityIntegrityProtectionAlg_ES384 CDV3MetaSecurityIntegrityProtectionAlg = "ES384"
	CDV3MetaSecurityIntegrityProtectionAlg_ES512 CDV3MetaSecurityIntegrityProtectionAlg = "ES512"
	CDV3MetaSecurityIntegrityProtectionAlg_PS256 CDV3MetaSecurityIntegrityProtectionAlg = "PS256"
	CDV3MetaSecurityIntegrityProtectionAlg_PS384 CDV3MetaSecurityIntegrityProtectionAlg = "PS384"
	CDV3MetaSecurityIntegrityProtectionAlg_PS512 CDV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type CDV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type CDV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
