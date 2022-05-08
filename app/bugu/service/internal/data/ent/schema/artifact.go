package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Artifact holds the schema definition for the Artifact entity.
type Artifact struct {
	ent.Schema
}

// Fields of the Artifact.
func (Artifact) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("file_id", uuid.UUID{}),
		field.UUID("artifact_hash", uuid.UUID{}).Unique(),
		field.Int64("artifact_size"),
		field.String("artifact_addr"),
		field.Enum("method").
			Values("Adposhel", "Agent", "Allaple", "Amonetize", "Androm",
				"Autorun", "BrowseFox", "Dinwod", "Elex", "Expiro",
				"Fasong", "HackKMS", "Hlux", "Injector", "InstallCore",
				"MultiPlug", "Neoreklami", "Neshta", "Other", "Regrun",
				"Sality", "Snarasite", "Stantinko", "VBA", "VBKrypt",
				"Vilsel"),
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

// Edges of the Artifact.
func (Artifact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("affiliated_file", File.Type).
			Ref("artifact").
			Field("file_id").
			Unique().
			Required(),
		edge.From("affiliated_user", User.Type).
			Ref("user_artifact"),
	}
}
