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

// NewAnnouncementPublishedV1 creates a new struct pointer that represents
// major version 1 of EiffelAnnouncementPublishedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 1.x.x
// currently known by this SDK.
func NewAnnouncementPublishedV1() (*AnnouncementPublishedV1, error) {
	var event AnnouncementPublishedV1
	event.Meta.Type = "EiffelAnnouncementPublishedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][1].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *AnnouncementPublishedV1) MarshalJSON() ([]byte, error) {
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
		links = make([]AnnPV1Link, 0)
	}
	s := struct {
		Data  *AnnPV1Data  `json:"data"`
		Links []AnnPV1Link `json:"links"`
		Meta  *AnnPV1Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *AnnouncementPublishedV1) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type AnnouncementPublishedV1 struct {
	// Mandatory fields
	Data  AnnPV1Data   `json:"data"`
	Links []AnnPV1Link `json:"links"`
	Meta  AnnPV1Meta   `json:"meta"`

	// Optional fields

}

type AnnPV1Data struct {
	// Mandatory fields
	Body     string             `json:"body"`
	Heading  string             `json:"heading"`
	Severity AnnPV1DataSeverity `json:"severity"`

	// Optional fields
	CustomData []AnnPV1DataCustomDatum `json:"customData,omitempty"`
	URI        string                  `json:"uri,omitempty"`
}

type AnnPV1DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type AnnPV1DataSeverity string

const (
	AnnPV1DataSeverity_Minor    AnnPV1DataSeverity = "MINOR"
	AnnPV1DataSeverity_Major    AnnPV1DataSeverity = "MAJOR"
	AnnPV1DataSeverity_Critical AnnPV1DataSeverity = "CRITICAL"
	AnnPV1DataSeverity_Blocker  AnnPV1DataSeverity = "BLOCKER"
	AnnPV1DataSeverity_Closed   AnnPV1DataSeverity = "CLOSED"
	AnnPV1DataSeverity_Canceled AnnPV1DataSeverity = "CANCELED"
)

type AnnPV1Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type AnnPV1Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security AnnPV1MetaSecurity `json:"security,omitempty"`
	Source   AnnPV1MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type AnnPV1MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM AnnPV1MetaSecuritySDM `json:"sdm,omitempty"`
}

type AnnPV1MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type AnnPV1MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                     `json:"domainId,omitempty"`
	Host       string                     `json:"host,omitempty"`
	Name       string                     `json:"name,omitempty"`
	Serializer AnnPV1MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                     `json:"uri,omitempty"`
}

type AnnPV1MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
