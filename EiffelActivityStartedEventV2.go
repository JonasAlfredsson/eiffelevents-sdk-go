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

// NewActivityStartedV2 creates a new struct pointer that represents
// major version 2 of EiffelActivityStartedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewActivityStartedV2() (*ActivityStartedV2, error) {
	var event ActivityStartedV2
	event.Meta.Type = "EiffelActivityStartedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *ActivityStartedV2) MarshalJSON() ([]byte, error) {
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
		links = make([]ActSV2Link, 0)
	}
	s := struct {
		Data  *ActSV2Data  `json:"data"`
		Links []ActSV2Link `json:"links"`
		Meta  *ActSV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *ActivityStartedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type ActivityStartedV2 struct {
	// Mandatory fields
	Data  ActSV2Data   `json:"data"`
	Links []ActSV2Link `json:"links"`
	Meta  ActSV2Meta   `json:"meta"`

	// Optional fields

}

type ActSV2Data struct {
	// Mandatory fields

	// Optional fields
	CustomData   []ActSV2DataCustomDatum `json:"customData,omitempty"`
	ExecutionURI string                  `json:"executionUri,omitempty"`
	LiveLogs     []ActSV2DataLiveLog     `json:"liveLogs,omitempty"`
}

type ActSV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type ActSV2DataLiveLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields

}

type ActSV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type ActSV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ActSV2MetaSecurity `json:"security,omitempty"`
	Source   ActSV2MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ActSV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM ActSV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type ActSV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type ActSV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
