package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Users struct {
	ent.Schema
}

func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.Int("telegram_id").
			Positive(),

		field.String("username").
			Optional().
			Unique(),

		field.String(("full_name")),

		field.Bool("is_premium").
			Default(false),

		field.Int("birthday_count").Positive(),
	}
}

func (Users) Edges() []ent.Edge {
	return nil
}
