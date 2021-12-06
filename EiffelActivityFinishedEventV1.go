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

// NewActivityFinishedV1 creates a new struct pointer that represents
// major version 1 of EiffelActivityFinishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 1.x.x
// currently known by this SDK.
func NewActivityFinishedV1() (*ActivityFinishedV1, error) {
	var event ActivityFinishedV1
	event.Meta.Type = "EiffelActivityFinishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][1].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityFinishedV1) MarshalJSON() ([]byte, error) {
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
		links = make([]ActFV1Link, 0)
	}
	s := struct {
		Data  *ActFV1Data  `json:"data"`
		Links []ActFV1Link `json:"links"`
		Meta  *ActFV1Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *ActivityFinishedV1) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type ActivityFinishedV1 struct {
	// Mandatory fields
	Data  ActFV1Data   `json:"data"`
	Links []ActFV1Link `json:"links"`
	Meta  ActFV1Meta   `json:"meta"`

	// Optional fields

}

type ActFV1Data struct {
	// Mandatory fields
	Outcome ActFV1DataOutcome `json:"outcome"`

	// Optional fields
	CustomData     []ActFV1DataCustomDatum   `json:"customData,omitempty"`
	PersistentLogs []ActFV1DataPersistentLog `json:"persistentLogs,omitempty"`
}

type ActFV1DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type ActFV1DataOutcome struct {
	// Mandatory fields
	Conclusion ActFV1DataOutcomeConclusion `json:"conclusion"`

	// Optional fields
	Description string `json:"description,omitempty"`
}

type ActFV1DataOutcomeConclusion string

const (
	ActFV1DataOutcomeConclusion_Successful   ActFV1DataOutcomeConclusion = "SUCCESSFUL"
	ActFV1DataOutcomeConclusion_Unsuccessful ActFV1DataOutcomeConclusion = "UNSUCCESSFUL"
	ActFV1DataOutcomeConclusion_Failed       ActFV1DataOutcomeConclusion = "FAILED"
	ActFV1DataOutcomeConclusion_Aborted      ActFV1DataOutcomeConclusion = "ABORTED"
	ActFV1DataOutcomeConclusion_TimedOut     ActFV1DataOutcomeConclusion = "TIMED_OUT"
	ActFV1DataOutcomeConclusion_Inconclusive ActFV1DataOutcomeConclusion = "INCONCLUSIVE"
)

type ActFV1DataPersistentLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields

}

type ActFV1Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type ActFV1Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActFV1MetaSecurity `json:"security,omitempty"`
	Source   ActFV1MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActFV1MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM ActFV1MetaSecuritySDM `json:"sdm,omitempty"`
}

type ActFV1MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type ActFV1MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                     `json:"domainId,omitempty"`
	Host       string                     `json:"host,omitempty"`
	Name       string                     `json:"name,omitempty"`
	Serializer ActFV1MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                     `json:"uri,omitempty"`
}

type ActFV1MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
