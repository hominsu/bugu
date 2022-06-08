// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArtifactsColumns holds the columns for the "artifacts" table.
	ArtifactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "file_id", Type: field.TypeUUID},
		{Name: "method", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "affiliated_file_id", Type: field.TypeUUID, Unique: true},
	}
	// ArtifactsTable holds the schema information for the "artifacts" table.
	ArtifactsTable = &schema.Table{
		Name:       "artifacts",
		Columns:    ArtifactsColumns,
		PrimaryKey: []*schema.Column{ArtifactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "artifacts_files_artifact",
				Columns:    []*schema.Column{ArtifactsColumns[5]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "file_sha_1", Type: field.TypeString, Unique: true},
		{Name: "file_size", Type: field.TypeInt64},
		{Name: "file_addr", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"Adposhel", "Agent", "Allaple", "Amonetize", "Androm", "Autorun", "BrowseFox", "Dinwod", "Elex", "Expiro", "Fasong", "HackKMS", "Hlux", "Injector", "InstallCore", "MultiPlug", "Neoreklami", "Neshta", "Other", "Regrun", "Sality", "Snarasite", "Stantinko", "VBA", "VBKrypt", "Vilsel"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserUserFileColumns holds the columns for the "user_user_file" table.
	UserUserFileColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "file_id", Type: field.TypeUUID},
	}
	// UserUserFileTable holds the schema information for the "user_user_file" table.
	UserUserFileTable = &schema.Table{
		Name:       "user_user_file",
		Columns:    UserUserFileColumns,
		PrimaryKey: []*schema.Column{UserUserFileColumns[0], UserUserFileColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_user_file_user_id",
				Columns:    []*schema.Column{UserUserFileColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_user_file_file_id",
				Columns:    []*schema.Column{UserUserFileColumns[1]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserUserArtifactColumns holds the columns for the "user_user_artifact" table.
	UserUserArtifactColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "artifact_id", Type: field.TypeUUID},
	}
	// UserUserArtifactTable holds the schema information for the "user_user_artifact" table.
	UserUserArtifactTable = &schema.Table{
		Name:       "user_user_artifact",
		Columns:    UserUserArtifactColumns,
		PrimaryKey: []*schema.Column{UserUserArtifactColumns[0], UserUserArtifactColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_user_artifact_user_id",
				Columns:    []*schema.Column{UserUserArtifactColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_user_artifact_artifact_id",
				Columns:    []*schema.Column{UserUserArtifactColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArtifactsTable,
		FilesTable,
		UsersTable,
		UserUserFileTable,
		UserUserArtifactTable,
	}
)

func init() {
	ArtifactsTable.ForeignKeys[0].RefTable = FilesTable
	UserUserFileTable.ForeignKeys[0].RefTable = UsersTable
	UserUserFileTable.ForeignKeys[1].RefTable = FilesTable
	UserUserArtifactTable.ForeignKeys[0].RefTable = UsersTable
	UserUserArtifactTable.ForeignKeys[1].RefTable = ArtifactsTable
}
