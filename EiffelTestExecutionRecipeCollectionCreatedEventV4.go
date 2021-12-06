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

// NewTestExecutionRecipeCollectionCreatedV4 creates a new struct pointer that represents
// major version 4 of EiffelTestExecutionRecipeCollectionCreatedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 4.x.x
// currently known by this SDK.
func NewTestExecutionRecipeCollectionCreatedV4() (*TestExecutionRecipeCollectionCreatedV4, error) {
	var event TestExecutionRecipeCollectionCreatedV4
	event.Meta.Type = "EiffelTestExecutionRecipeCollectionCreatedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][4].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestExecutionRecipeCollectionCreatedV4) MarshalJSON() ([]byte, error) {
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
		links = make([]TERCCV4Link, 0)
	}
	s := struct {
		Data  *TERCCV4Data  `json:"data"`
		Links []TERCCV4Link `json:"links"`
		Meta  *TERCCV4Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *TestExecutionRecipeCollectionCreatedV4) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type TestExecutionRecipeCollectionCreatedV4 struct {
	// Mandatory fields
	Data  TERCCV4Data   `json:"data"`
	Links []TERCCV4Link `json:"links"`
	Meta  TERCCV4Meta   `json:"meta"`

	// Optional fields

}

type TERCCV4Data struct {
	// Mandatory fields
	SelectionStrategy TERCCV4DataSelectionStrategy `json:"selectionStrategy"`

	// Optional fields
	Batches    []TERCCV4DataBatch       `json:"batches,omitempty"`
	BatchesURI string                   `json:"batchesUri,omitempty"`
	CustomData []TERCCV4DataCustomDatum `json:"customData,omitempty"`
}

type TERCCV4DataBatch struct {
	// Mandatory fields
	Priority int64                    `json:"priority"`
	Recipes  []TERCCV4DataBatchRecipe `json:"recipes"`

	// Optional fields
	Dependencies []TERCCV4DataBatchDependency `json:"dependencies,omitempty"`
	Name         string                       `json:"name,omitempty"`
}

type TERCCV4DataBatchDependency struct {
	// Mandatory fields
	Dependency string `json:"dependency"`
	Dependent  string `json:"dependent"`

	// Optional fields

}

type TERCCV4DataBatchRecipe struct {
	// Mandatory fields
	ID       string                         `json:"id"`
	TestCase TERCCV4DataBatchRecipeTestCase `json:"testCase"`

	// Optional fields
	Constraints []TERCCV4DataBatchRecipeConstraint `json:"constraints,omitempty"`
}

type TERCCV4DataBatchRecipeConstraint struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TERCCV4DataBatchRecipeTestCase struct {
	// Mandatory fields
	ID string `json:"id"`

	// Optional fields
	Tracker string `json:"tracker,omitempty"`
	URI     string `json:"uri,omitempty"`
}

type TERCCV4DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TERCCV4DataSelectionStrategy struct {
	// Mandatory fields
	ID string `json:"id"`

	// Optional fields
	Tracker string `json:"tracker,omitempty"`
	URI     string `json:"uri,omitempty"`
}

type TERCCV4Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type TERCCV4Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TERCCV4MetaSecurity `json:"security,omitempty"`
	Source   TERCCV4MetaSource   `json:"source,omitempty"`
	Tags     []string            `json:"tags,omitempty"`
}

type TERCCV4MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection TERCCV4MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []TERCCV4MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type TERCCV4MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       TERCCV4MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                    `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type TERCCV4MetaSecurityIntegrityProtectionAlg string

const (
	TERCCV4MetaSecurityIntegrityProtectionAlg_HS256 TERCCV4MetaSecurityIntegrityProtectionAlg = "HS256"
	TERCCV4MetaSecurityIntegrityProtectionAlg_HS384 TERCCV4MetaSecurityIntegrityProtectionAlg = "HS384"
	TERCCV4MetaSecurityIntegrityProtectionAlg_HS512 TERCCV4MetaSecurityIntegrityProtectionAlg = "HS512"
	TERCCV4MetaSecurityIntegrityProtectionAlg_RS256 TERCCV4MetaSecurityIntegrityProtectionAlg = "RS256"
	TERCCV4MetaSecurityIntegrityProtectionAlg_RS384 TERCCV4MetaSecurityIntegrityProtectionAlg = "RS384"
	TERCCV4MetaSecurityIntegrityProtectionAlg_RS512 TERCCV4MetaSecurityIntegrityProtectionAlg = "RS512"
	TERCCV4MetaSecurityIntegrityProtectionAlg_ES256 TERCCV4MetaSecurityIntegrityProtectionAlg = "ES256"
	TERCCV4MetaSecurityIntegrityProtectionAlg_ES384 TERCCV4MetaSecurityIntegrityProtectionAlg = "ES384"
	TERCCV4MetaSecurityIntegrityProtectionAlg_ES512 TERCCV4MetaSecurityIntegrityProtectionAlg = "ES512"
	TERCCV4MetaSecurityIntegrityProtectionAlg_PS256 TERCCV4MetaSecurityIntegrityProtectionAlg = "PS256"
	TERCCV4MetaSecurityIntegrityProtectionAlg_PS384 TERCCV4MetaSecurityIntegrityProtectionAlg = "PS384"
	TERCCV4MetaSecurityIntegrityProtectionAlg_PS512 TERCCV4MetaSecurityIntegrityProtectionAlg = "PS512"
)

type TERCCV4MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type TERCCV4MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
