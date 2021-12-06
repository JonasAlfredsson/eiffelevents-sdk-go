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

// NewIssueDefinedV2 creates a new struct pointer that represents
// major version 2 of EiffelIssueDefinedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 2.x.x
// currently known by this SDK.
func NewIssueDefinedV2() (*IssueDefinedV2, error) {
	var event IssueDefinedV2
	event.Meta.Type = "EiffelIssueDefinedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][2].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *IssueDefinedV2) MarshalJSON() ([]byte, error) {
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
		links = make([]IDV2Link, 0)
	}
	s := struct {
		Data  *IDV2Data  `json:"data"`
		Links []IDV2Link `json:"links"`
		Meta  *IDV2Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *IssueDefinedV2) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type IssueDefinedV2 struct {
	// Mandatory fields

	// Optional fields
	Data  IDV2Data   `json:"data,omitempty"`
	Links []IDV2Link `json:"links,omitempty"`
	Meta  IDV2Meta   `json:"meta,omitempty"`
}

type IDV2Data struct {
	// Mandatory fields
	ID      string       `json:"id"`
	Tracker string       `json:"tracker"`
	Type    IDV2DataType `json:"type"`
	URI     string       `json:"uri"`

	// Optional fields
	CustomData []IDV2DataCustomDatum `json:"customData,omitempty"`
	Title      string                `json:"title,omitempty"`
}

type IDV2DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type IDV2DataType string

const (
	IDV2DataType_Bug         IDV2DataType = "BUG"
	IDV2DataType_Improvement IDV2DataType = "IMPROVEMENT"
	IDV2DataType_Feature     IDV2DataType = "FEATURE"
	IDV2DataType_WorkItem    IDV2DataType = "WORK_ITEM"
	IDV2DataType_Requirement IDV2DataType = "REQUIREMENT"
	IDV2DataType_Other       IDV2DataType = "OTHER"
)

type IDV2Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type IDV2Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security IDV2MetaSecurity `json:"security,omitempty"`
	Source   IDV2MetaSource   `json:"source,omitempty"`
	Tags     []string         `json:"tags,omitempty"`
}

type IDV2MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM IDV2MetaSecuritySDM `json:"sdm,omitempty"`
}

type IDV2MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type IDV2MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
