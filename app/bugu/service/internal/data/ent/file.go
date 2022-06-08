// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/artifact"
	"github.com/hominsu/bugu/app/bugu/service/internal/data/ent/file"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FileSha1 holds the value of the "file_sha_1" field.
	FileSha1 string `json:"file_sha_1,omitempty"`
	// FileSize holds the value of the "file_size" field.
	FileSize int64 `json:"file_size,omitempty"`
	// FileAddr holds the value of the "file_addr" field.
	FileAddr string `json:"file_addr,omitempty"`
	// Type holds the value of the "type" field.
	Type file.Type `json:"type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges         FileEdges `json:"edges"`
	file_artifact *uuid.UUID
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Artifact holds the value of the artifact edge.
	Artifact *Artifact `json:"artifact,omitempty"`
	// AffiliatedUser holds the value of the affiliated_user edge.
	AffiliatedUser []*User `json:"affiliated_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ArtifactOrErr returns the Artifact value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) ArtifactOrErr() (*Artifact, error) {
	if e.loadedTypes[0] {
		if e.Artifact == nil {
			// The edge artifact was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: artifact.Label}
		}
		return e.Artifact, nil
	}
	return nil, &NotLoadedError{edge: "artifact"}
}

// AffiliatedUserOrErr returns the AffiliatedUser value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) AffiliatedUserOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.AffiliatedUser, nil
	}
	return nil, &NotLoadedError{edge: "affiliated_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldFileSize:
			values[i] = new(sql.NullInt64)
		case file.FieldFileSha1, file.FieldFileAddr, file.FieldType:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case file.FieldID:
			values[i] = new(uuid.UUID)
		case file.ForeignKeys[0]: // file_artifact
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type File", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case file.FieldFileSha1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_sha_1", values[i])
			} else if value.Valid {
				f.FileSha1 = value.String
			}
		case file.FieldFileSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field file_size", values[i])
			} else if value.Valid {
				f.FileSize = value.Int64
			}
		case file.FieldFileAddr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_addr", values[i])
			} else if value.Valid {
				f.FileAddr = value.String
			}
		case file.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				f.Type = file.Type(value.String)
			}
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field file_artifact", values[i])
			} else if value.Valid {
				f.file_artifact = new(uuid.UUID)
				*f.file_artifact = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryArtifact queries the "artifact" edge of the File entity.
func (f *File) QueryArtifact() *ArtifactQuery {
	return (&FileClient{config: f.config}).QueryArtifact(f)
}

// QueryAffiliatedUser queries the "affiliated_user" edge of the File entity.
func (f *File) QueryAffiliatedUser() *UserQuery {
	return (&FileClient{config: f.config}).QueryAffiliatedUser(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", file_sha_1=")
	builder.WriteString(f.FileSha1)
	builder.WriteString(", file_size=")
	builder.WriteString(fmt.Sprintf("%v", f.FileSize))
	builder.WriteString(", file_addr=")
	builder.WriteString(f.FileAddr)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", f.Type))
	builder.WriteString(", created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
