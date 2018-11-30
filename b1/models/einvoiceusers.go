package models

import (
	"encoding/xml"
	"time"
)

// EInvoiceUsers ...
// xml files geting values struct
type EInvoiceUsers struct {
	XMLName           xml.Name  `json:"-"`
	Identifier        string    `json:"identifier"`
	Alias             string    `json:"alias"`
	Title             string    `json:"title"`
	Type              string    `json:"type"`
	FirstCreationTime time.Time `json:"firstCreationTime"`
}

// Users ...
// Users xml struct
type Users struct {
	XMLName xml.Name        `xml:"UserList"`
	Users   []EInvoiceUsers `xml:"User"`
}
