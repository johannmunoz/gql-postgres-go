// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/predicate"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/product"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent/review"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProduct = "Product"
	TypeReview  = "Review"
)

// ProductMutation represents an operation that mutates the Product nodes in the graph.
type ProductMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	clearedFields  map[string]struct{}
	reviews        map[uuid.UUID]struct{}
	removedreviews map[uuid.UUID]struct{}
	clearedreviews bool
	done           bool
	oldValue       func(context.Context) (*Product, error)
	predicates     []predicate.Product
}

var _ ent.Mutation = (*ProductMutation)(nil)

// productOption allows management of the mutation configuration using functional options.
type productOption func(*ProductMutation)

// newProductMutation creates new mutation for the Product entity.
func newProductMutation(c config, op Op, opts ...productOption) *ProductMutation {
	m := &ProductMutation{
		config:        c,
		op:            op,
		typ:           TypeProduct,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProductID sets the ID field of the mutation.
func withProductID(id uuid.UUID) productOption {
	return func(m *ProductMutation) {
		var (
			err   error
			once  sync.Once
			value *Product
		)
		m.oldValue = func(ctx context.Context) (*Product, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Product.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProduct sets the old Product of the mutation.
func withProduct(node *Product) productOption {
	return func(m *ProductMutation) {
		m.oldValue = func(context.Context) (*Product, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProductMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProductMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Product entities.
func (m *ProductMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ProductMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ProductMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Product.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// AddReviewIDs adds the "reviews" edge to the Review entity by ids.
func (m *ProductMutation) AddReviewIDs(ids ...uuid.UUID) {
	if m.reviews == nil {
		m.reviews = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.reviews[ids[i]] = struct{}{}
	}
}

// ClearReviews clears the "reviews" edge to the Review entity.
func (m *ProductMutation) ClearReviews() {
	m.clearedreviews = true
}

// ReviewsCleared reports if the "reviews" edge to the Review entity was cleared.
func (m *ProductMutation) ReviewsCleared() bool {
	return m.clearedreviews
}

// RemoveReviewIDs removes the "reviews" edge to the Review entity by IDs.
func (m *ProductMutation) RemoveReviewIDs(ids ...uuid.UUID) {
	if m.removedreviews == nil {
		m.removedreviews = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.reviews, ids[i])
		m.removedreviews[ids[i]] = struct{}{}
	}
}

// RemovedReviews returns the removed IDs of the "reviews" edge to the Review entity.
func (m *ProductMutation) RemovedReviewsIDs() (ids []uuid.UUID) {
	for id := range m.removedreviews {
		ids = append(ids, id)
	}
	return
}

// ReviewsIDs returns the "reviews" edge IDs in the mutation.
func (m *ProductMutation) ReviewsIDs() (ids []uuid.UUID) {
	for id := range m.reviews {
		ids = append(ids, id)
	}
	return
}

// ResetReviews resets all changes to the "reviews" edge.
func (m *ProductMutation) ResetReviews() {
	m.reviews = nil
	m.clearedreviews = false
	m.removedreviews = nil
}

// Where appends a list predicates to the ProductMutation builder.
func (m *ProductMutation) Where(ps ...predicate.Product) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ProductMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Product).
func (m *ProductMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProductMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProductMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProductMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Product field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProductMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProductMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Product numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProductMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProductMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProductMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Product nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProductMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProductMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.reviews != nil {
		edges = append(edges, product.EdgeReviews)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProductMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case product.EdgeReviews:
		ids := make([]ent.Value, 0, len(m.reviews))
		for id := range m.reviews {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProductMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedreviews != nil {
		edges = append(edges, product.EdgeReviews)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProductMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case product.EdgeReviews:
		ids := make([]ent.Value, 0, len(m.removedreviews))
		for id := range m.removedreviews {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProductMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedreviews {
		edges = append(edges, product.EdgeReviews)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProductMutation) EdgeCleared(name string) bool {
	switch name {
	case product.EdgeReviews:
		return m.clearedreviews
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProductMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Product unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProductMutation) ResetEdge(name string) error {
	switch name {
	case product.EdgeReviews:
		m.ResetReviews()
		return nil
	}
	return fmt.Errorf("unknown Product edge %s", name)
}

// ReviewMutation represents an operation that mutates the Review nodes in the graph.
type ReviewMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	body           *string
	clearedFields  map[string]struct{}
	product        *uuid.UUID
	clearedproduct bool
	done           bool
	oldValue       func(context.Context) (*Review, error)
	predicates     []predicate.Review
}

var _ ent.Mutation = (*ReviewMutation)(nil)

// reviewOption allows management of the mutation configuration using functional options.
type reviewOption func(*ReviewMutation)

// newReviewMutation creates new mutation for the Review entity.
func newReviewMutation(c config, op Op, opts ...reviewOption) *ReviewMutation {
	m := &ReviewMutation{
		config:        c,
		op:            op,
		typ:           TypeReview,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withReviewID sets the ID field of the mutation.
func withReviewID(id uuid.UUID) reviewOption {
	return func(m *ReviewMutation) {
		var (
			err   error
			once  sync.Once
			value *Review
		)
		m.oldValue = func(ctx context.Context) (*Review, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Review.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withReview sets the old Review of the mutation.
func withReview(node *Review) reviewOption {
	return func(m *ReviewMutation) {
		m.oldValue = func(context.Context) (*Review, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ReviewMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ReviewMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Review entities.
func (m *ReviewMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ReviewMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ReviewMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Review.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetBody sets the "body" field.
func (m *ReviewMutation) SetBody(s string) {
	m.body = &s
}

// Body returns the value of the "body" field in the mutation.
func (m *ReviewMutation) Body() (r string, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old "body" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ResetBody resets all changes to the "body" field.
func (m *ReviewMutation) ResetBody() {
	m.body = nil
}

// SetProductID sets the "product" edge to the Product entity by id.
func (m *ReviewMutation) SetProductID(id uuid.UUID) {
	m.product = &id
}

// ClearProduct clears the "product" edge to the Product entity.
func (m *ReviewMutation) ClearProduct() {
	m.clearedproduct = true
}

// ProductCleared reports if the "product" edge to the Product entity was cleared.
func (m *ReviewMutation) ProductCleared() bool {
	return m.clearedproduct
}

// ProductID returns the "product" edge ID in the mutation.
func (m *ReviewMutation) ProductID() (id uuid.UUID, exists bool) {
	if m.product != nil {
		return *m.product, true
	}
	return
}

// ProductIDs returns the "product" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ProductID instead. It exists only for internal usage by the builders.
func (m *ReviewMutation) ProductIDs() (ids []uuid.UUID) {
	if id := m.product; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetProduct resets all changes to the "product" edge.
func (m *ReviewMutation) ResetProduct() {
	m.product = nil
	m.clearedproduct = false
}

// Where appends a list predicates to the ReviewMutation builder.
func (m *ReviewMutation) Where(ps ...predicate.Review) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ReviewMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Review).
func (m *ReviewMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ReviewMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.body != nil {
		fields = append(fields, review.FieldBody)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ReviewMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case review.FieldBody:
		return m.Body()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ReviewMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case review.FieldBody:
		return m.OldBody(ctx)
	}
	return nil, fmt.Errorf("unknown Review field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ReviewMutation) SetField(name string, value ent.Value) error {
	switch name {
	case review.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	}
	return fmt.Errorf("unknown Review field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ReviewMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ReviewMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ReviewMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Review numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ReviewMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ReviewMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ReviewMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Review nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ReviewMutation) ResetField(name string) error {
	switch name {
	case review.FieldBody:
		m.ResetBody()
		return nil
	}
	return fmt.Errorf("unknown Review field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ReviewMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.product != nil {
		edges = append(edges, review.EdgeProduct)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ReviewMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case review.EdgeProduct:
		if id := m.product; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ReviewMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ReviewMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ReviewMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedproduct {
		edges = append(edges, review.EdgeProduct)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ReviewMutation) EdgeCleared(name string) bool {
	switch name {
	case review.EdgeProduct:
		return m.clearedproduct
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ReviewMutation) ClearEdge(name string) error {
	switch name {
	case review.EdgeProduct:
		m.ClearProduct()
		return nil
	}
	return fmt.Errorf("unknown Review unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ReviewMutation) ResetEdge(name string) error {
	switch name {
	case review.EdgeProduct:
		m.ResetProduct()
		return nil
	}
	return fmt.Errorf("unknown Review edge %s", name)
}
