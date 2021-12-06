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

// NewTestSuiteStartedV1 creates a new struct pointer that represents
// major version 1 of EiffelTestSuiteStartedEvent.
// The returned struct has all required meta members populated.
// The event version is set to the most recent 1.x.x
// currently known by this SDK.
func NewTestSuiteStartedV1() (*TestSuiteStartedV1, error) {
	var event TestSuiteStartedV1
	event.Meta.Type = "EiffelTestSuiteStartedEvent"
	event.Meta.ID = uuid.NewString()
	event.Meta.Version = eventTypeTable[event.Meta.Type][1].latestVersion
	event.Meta.Time = time.Now().UnixMilli()
	return &event, nil
}

// MarshalJSON returns the JSON encoding of the event.
func (e *TestSuiteStartedV1) MarshalJSON() ([]byte, error) {
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
		links = make([]TSSV1Link, 0)
	}
	s := struct {
		Data  *TSSV1Data  `json:"data"`
		Links []TSSV1Link `json:"links"`
		Meta  *TSSV1Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *TestSuiteStartedV1) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type TestSuiteStartedV1 struct {
	// Mandatory fields
	Data  TSSV1Data   `json:"data"`
	Links []TSSV1Link `json:"links"`
	Meta  TSSV1Meta   `json:"meta"`

	// Optional fields

}

type TSSV1Data struct {
	// Mandatory fields
	Name string `json:"name"`

	// Optional fields
	Categories []string               `json:"categories,omitempty"`
	CustomData []TSSV1DataCustomDatum `json:"customData,omitempty"`
	LiveLogs   []TSSV1DataLiveLog     `json:"liveLogs,omitempty"`
	Types      []TSSV1DataType        `json:"types,omitempty"`
}

type TSSV1DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type TSSV1DataLiveLog struct {
	// Mandatory fields
	Name string `json:"name"`
	URI  string `json:"uri"`

	// Optional fields

}

type TSSV1DataType string

const (
	TSSV1DataType_Accessibility    TSSV1DataType = "ACCESSIBILITY"
	TSSV1DataType_BackupRecovery   TSSV1DataType = "BACKUP_RECOVERY"
	TSSV1DataType_Compatibility    TSSV1DataType = "COMPATIBILITY"
	TSSV1DataType_Conversion       TSSV1DataType = "CONVERSION"
	TSSV1DataType_DisasterRecovery TSSV1DataType = "DISASTER_RECOVERY"
	TSSV1DataType_Functional       TSSV1DataType = "FUNCTIONAL"
	TSSV1DataType_Installability   TSSV1DataType = "INSTALLABILITY"
	TSSV1DataType_Interoperability TSSV1DataType = "INTEROPERABILITY"
	TSSV1DataType_Localization     TSSV1DataType = "LOCALIZATION"
	TSSV1DataType_Maintainability  TSSV1DataType = "MAINTAINABILITY"
	TSSV1DataType_Performance      TSSV1DataType = "PERFORMANCE"
	TSSV1DataType_Portability      TSSV1DataType = "PORTABILITY"
	TSSV1DataType_Procedure        TSSV1DataType = "PROCEDURE"
	TSSV1DataType_Reliability      TSSV1DataType = "RELIABILITY"
	TSSV1DataType_Security         TSSV1DataType = "SECURITY"
	TSSV1DataType_Stability        TSSV1DataType = "STABILITY"
	TSSV1DataType_Usability        TSSV1DataType = "USABILITY"
)

type TSSV1Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields

}

type TSSV1Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security TSSV1MetaSecurity `json:"security,omitempty"`
	Source   TSSV1MetaSource   `json:"source,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
}

type TSSV1MetaSecurity struct {
	// Mandatory fields

	// Optional fields
	SDM TSSV1MetaSecuritySDM `json:"sdm,omitempty"`
}

type TSSV1MetaSecuritySDM struct {
	// Mandatory fields
	AuthorIdentity  string `json:"authorIdentity"`
	EncryptedDigest string `json:"encryptedDigest"`

	// Optional fields

}

type TSSV1MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string                    `json:"domainId,omitempty"`
	Host       string                    `json:"host,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Serializer TSSV1MetaSourceSerializer `json:"serializer,omitempty"`
	URI        string                    `json:"uri,omitempty"`
}

type TSSV1MetaSourceSerializer struct {
	// Mandatory fields
	ArtifactID string `json:"artifactId"`
	GroupID    string `json:"groupId"`
	Version    string `json:"version"`

	// Optional fields

}
