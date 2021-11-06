// Package spec contains struct and interface specification for honcheonui plugins
package spec

import (
	"time"

	"github.com/gofrs/uuid"
)

// HoncheonuiResource is a struct for containing resource data
type HoncheonuiResource struct {
	Provider           string
	Type               string
	OriginalID         string
	UUID               uuid.UUID
	Name               string
	Notes              string
	GroupID            string
	ResourceCreatedAt  time.Time
	ResourceModifiedAt time.Time
	IPAddress          string
	Location           string
	IsConn             bool
	IsOn               bool
	Attributes         map[string]string
	IntegerAttributes  map[string]int
	Tags               []string
	UserIDs            []string
	Raw                interface{}
}

// HoncheonuiStatus is a struct for containing status data
type HoncheonuiStatus struct {
	OriginalID string
	IsConn     bool
	IsOn       bool
}

// HoncheonuiNotification is a struct for containing notification data
type HoncheonuiNotification struct {
	Provider    string
	Type        string
	OriginalID  string
	GroupID     string
	UserID      string
	Title       string
	Content     string
	Category    string
	IssuedBy    string
	IsOpen      bool
	IssuedAt    time.Time
	ModifiedAt  time.Time
	ResourceIDs []string
	UserIDs     []string
}

// Provider is an interface for provider plugins
type Provider interface {
	Init() error
	CheckAccount(string, string) (int, int, error)
	GetResources(string, string) ([]interface{}, error)
	GetStatuses(string, string) ([]interface{}, error)
	GetNotifications(string, string, time.Time) ([]interface{}, error)
}
