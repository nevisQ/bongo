package bongo

import (
	"time"
)

type Validater interface {
	Valid() bool
}

type DocumentBase struct {
	Id       interface{} `bson:"_id,omitempty" json:"_id"`
	Created  *time.Time  `bson:"_created,omitempty" json:"_created,omitempty"`
	Modified *time.Time  `bson:"_modified,omitempty" json:"_modified,omitempty"`

	// We want this to default to false without any work. So this will be the opposite of isNew. We want it to be new unless set to existing
	exists bool
}

// Satisfy the new tracker interface
func (d *DocumentBase) SetIsNew(isNew bool) {
	d.exists = !isNew
}

// Is the document new
func (d *DocumentBase) IsNew() bool {
	return !d.exists
}

// Satisfy the document interface
func (d *DocumentBase) GetId() interface{} {
	return d.Id
}

// Sets the ID for the document
func (d *DocumentBase) SetId(id interface{}) {
	d.Id = id
}

// Set's the created date
func (d *DocumentBase) SetCreated(t time.Time) {
	d.Created = &t
}

// Get the created date
func (d *DocumentBase) GetCreated() time.Time {

	if d.Created == nil {
		return time.Time{}
	}

	return *d.Created
}

// Sets the modified date
func (d *DocumentBase) SetModified(t time.Time) {
	d.Modified = &t
}

// Get's the modified date
func (d *DocumentBase) GetModified() time.Time {

	if d.Modified == nil {
		return time.Time{}
	}

	return *d.Modified
}
