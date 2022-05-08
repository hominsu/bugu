// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bugu/app/bugu/service/internal/data/ent/artifact"
	"bugu/app/bugu/service/internal/data/ent/file"
	"bugu/app/bugu/service/internal/data/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ArtifactCreate is the builder for creating a Artifact entity.
type ArtifactCreate struct {
	config
	mutation *ArtifactMutation
	hooks    []Hook
}

// SetFileID sets the "file_id" field.
func (ac *ArtifactCreate) SetFileID(u uuid.UUID) *ArtifactCreate {
	ac.mutation.SetFileID(u)
	return ac
}

// SetArtifactHash sets the "artifact_hash" field.
func (ac *ArtifactCreate) SetArtifactHash(u uuid.UUID) *ArtifactCreate {
	ac.mutation.SetArtifactHash(u)
	return ac
}

// SetArtifactSize sets the "artifact_size" field.
func (ac *ArtifactCreate) SetArtifactSize(i int64) *ArtifactCreate {
	ac.mutation.SetArtifactSize(i)
	return ac
}

// SetArtifactAddr sets the "artifact_addr" field.
func (ac *ArtifactCreate) SetArtifactAddr(s string) *ArtifactCreate {
	ac.mutation.SetArtifactAddr(s)
	return ac
}

// SetMethod sets the "method" field.
func (ac *ArtifactCreate) SetMethod(a artifact.Method) *ArtifactCreate {
	ac.mutation.SetMethod(a)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArtifactCreate) SetCreatedAt(t time.Time) *ArtifactCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableCreatedAt(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ArtifactCreate) SetUpdatedAt(t time.Time) *ArtifactCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ArtifactCreate) SetNillableUpdatedAt(t *time.Time) *ArtifactCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ArtifactCreate) SetID(u uuid.UUID) *ArtifactCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetAffiliatedFileID sets the "affiliated_file" edge to the File entity by ID.
func (ac *ArtifactCreate) SetAffiliatedFileID(id uuid.UUID) *ArtifactCreate {
	ac.mutation.SetAffiliatedFileID(id)
	return ac
}

// SetAffiliatedFile sets the "affiliated_file" edge to the File entity.
func (ac *ArtifactCreate) SetAffiliatedFile(f *File) *ArtifactCreate {
	return ac.SetAffiliatedFileID(f.ID)
}

// AddAffiliatedUserIDs adds the "affiliated_user" edge to the User entity by IDs.
func (ac *ArtifactCreate) AddAffiliatedUserIDs(ids ...uuid.UUID) *ArtifactCreate {
	ac.mutation.AddAffiliatedUserIDs(ids...)
	return ac
}

// AddAffiliatedUser adds the "affiliated_user" edges to the User entity.
func (ac *ArtifactCreate) AddAffiliatedUser(u ...*User) *ArtifactCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ac.AddAffiliatedUserIDs(ids...)
}

// Mutation returns the ArtifactMutation object of the builder.
func (ac *ArtifactCreate) Mutation() *ArtifactMutation {
	return ac.mutation
}

// Save creates the Artifact in the database.
func (ac *ArtifactCreate) Save(ctx context.Context) (*Artifact, error) {
	var (
		err  error
		node *Artifact
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArtifactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtifactCreate) SaveX(ctx context.Context) *Artifact {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtifactCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtifactCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArtifactCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := artifact.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := artifact.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtifactCreate) check() error {
	if _, ok := ac.mutation.FileID(); !ok {
		return &ValidationError{Name: "file_id", err: errors.New(`ent: missing required field "Artifact.file_id"`)}
	}
	if _, ok := ac.mutation.ArtifactHash(); !ok {
		return &ValidationError{Name: "artifact_hash", err: errors.New(`ent: missing required field "Artifact.artifact_hash"`)}
	}
	if _, ok := ac.mutation.ArtifactSize(); !ok {
		return &ValidationError{Name: "artifact_size", err: errors.New(`ent: missing required field "Artifact.artifact_size"`)}
	}
	if _, ok := ac.mutation.ArtifactAddr(); !ok {
		return &ValidationError{Name: "artifact_addr", err: errors.New(`ent: missing required field "Artifact.artifact_addr"`)}
	}
	if _, ok := ac.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "Artifact.method"`)}
	}
	if v, ok := ac.mutation.Method(); ok {
		if err := artifact.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "Artifact.method": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Artifact.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Artifact.updated_at"`)}
	}
	if _, ok := ac.mutation.AffiliatedFileID(); !ok {
		return &ValidationError{Name: "affiliated_file", err: errors.New(`ent: missing required edge "Artifact.affiliated_file"`)}
	}
	return nil
}

func (ac *ArtifactCreate) sqlSave(ctx context.Context) (*Artifact, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ac *ArtifactCreate) createSpec() (*Artifact, *sqlgraph.CreateSpec) {
	var (
		_node = &Artifact{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: artifact.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: artifact.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.ArtifactHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: artifact.FieldArtifactHash,
		})
		_node.ArtifactHash = value
	}
	if value, ok := ac.mutation.ArtifactSize(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: artifact.FieldArtifactSize,
		})
		_node.ArtifactSize = value
	}
	if value, ok := ac.mutation.ArtifactAddr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: artifact.FieldArtifactAddr,
		})
		_node.ArtifactAddr = value
	}
	if value, ok := ac.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: artifact.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: artifact.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: artifact.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.AffiliatedFileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   artifact.AffiliatedFileTable,
			Columns: []string{artifact.AffiliatedFileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FileID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AffiliatedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artifact.AffiliatedUserTable,
			Columns: artifact.AffiliatedUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArtifactCreateBulk is the builder for creating many Artifact entities in bulk.
type ArtifactCreateBulk struct {
	config
	builders []*ArtifactCreate
}

// Save creates the Artifact entities in the database.
func (acb *ArtifactCreateBulk) Save(ctx context.Context) ([]*Artifact, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artifact, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtifactMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtifactCreateBulk) SaveX(ctx context.Context) []*Artifact {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtifactCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtifactCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
