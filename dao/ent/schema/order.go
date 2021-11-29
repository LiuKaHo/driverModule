package schema

import "entgo.io/ent"
import "entgo.io/ent/schema/field"

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
