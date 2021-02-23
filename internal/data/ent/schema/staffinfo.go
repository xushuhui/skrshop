package schema

import "entgo.io/ent"

// StaffInfo holds the schema definition for the StaffInfo entity.
type StaffInfo struct {
	ent.Schema
}

// Fields of the StaffInfo.
func (StaffInfo) Fields() []ent.Field {
	return nil
}

// Edges of the StaffInfo.
func (StaffInfo) Edges() []ent.Edge {
	return nil
}
