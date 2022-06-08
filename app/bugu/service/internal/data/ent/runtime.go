// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/artifact"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/file"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/schema"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	artifactFields := schema.Artifact{}.Fields()
	_ = artifactFields
	// artifactDescCreatedAt is the schema descriptor for created_at field.
	artifactDescCreatedAt := artifactFields[4].Descriptor()
	// artifact.DefaultCreatedAt holds the default value on creation for the created_at field.
	artifact.DefaultCreatedAt = artifactDescCreatedAt.Default.(func() time.Time)
	// artifactDescUpdatedAt is the schema descriptor for updated_at field.
	artifactDescUpdatedAt := artifactFields[5].Descriptor()
	// artifact.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	artifact.DefaultUpdatedAt = artifactDescUpdatedAt.Default.(func() time.Time)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for created_at field.
	fileDescCreatedAt := fileFields[5].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the created_at field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescUpdatedAt is the schema descriptor for updated_at field.
	fileDescUpdatedAt := fileFields[6].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
}
