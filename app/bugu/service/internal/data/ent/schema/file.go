package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("file_sha_1").Unique(),
		field.Int64("file_size"),
		field.String("file_addr"),
		field.Enum("type").
			Values("Adposhel", "Agent", "Allaple", "Amonetize", "Androm",
				"Autorun", "BrowseFox", "Dinwod", "Elex", "Expiro",
				"Fasong", "HackKMS", "Hlux", "Injector", "InstallCore",
				"MultiPlug", "Neoreklami", "Neshta", "Other", "Regrun",
				"Sality", "Snarasite", "Stantinko", "VBA", "VBKrypt",
				"Vilsel").
			Default("Other"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artifact", Artifact.Type).
			Unique(),
		edge.From("affiliated_user", User.Type).
			Ref("user_file"),
	}
}
