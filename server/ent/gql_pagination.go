// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
	"github.com/johannmunoz/gql_postgres_go/ent/review"
	"github.com/johannmunoz/gql_postgres_go/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    uuid.UUID `msgpack:"i"`
	Value Value     `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// ManufacturerEdge is the edge representation of Manufacturer.
type ManufacturerEdge struct {
	Node   *Manufacturer `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// ManufacturerConnection is the connection containing edges to Manufacturer.
type ManufacturerConnection struct {
	Edges      []*ManufacturerEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

// ManufacturerPaginateOption enables pagination customization.
type ManufacturerPaginateOption func(*manufacturerPager) error

// WithManufacturerOrder configures pagination ordering.
func WithManufacturerOrder(order *ManufacturerOrder) ManufacturerPaginateOption {
	if order == nil {
		order = DefaultManufacturerOrder
	}
	o := *order
	return func(pager *manufacturerPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultManufacturerOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithManufacturerFilter configures pagination filter.
func WithManufacturerFilter(filter func(*ManufacturerQuery) (*ManufacturerQuery, error)) ManufacturerPaginateOption {
	return func(pager *manufacturerPager) error {
		if filter == nil {
			return errors.New("ManufacturerQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type manufacturerPager struct {
	order  *ManufacturerOrder
	filter func(*ManufacturerQuery) (*ManufacturerQuery, error)
}

func newManufacturerPager(opts []ManufacturerPaginateOption) (*manufacturerPager, error) {
	pager := &manufacturerPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultManufacturerOrder
	}
	return pager, nil
}

func (p *manufacturerPager) applyFilter(query *ManufacturerQuery) (*ManufacturerQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *manufacturerPager) toCursor(m *Manufacturer) Cursor {
	return p.order.Field.toCursor(m)
}

func (p *manufacturerPager) applyCursors(query *ManufacturerQuery, after, before *Cursor) *ManufacturerQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultManufacturerOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *manufacturerPager) applyOrder(query *ManufacturerQuery, reverse bool) *ManufacturerQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultManufacturerOrder.Field {
		query = query.Order(direction.orderFunc(DefaultManufacturerOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Manufacturer.
func (m *ManufacturerQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...ManufacturerPaginateOption,
) (*ManufacturerConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newManufacturerPager(opts)
	if err != nil {
		return nil, err
	}

	if m, err = pager.applyFilter(m); err != nil {
		return nil, err
	}

	conn := &ManufacturerConnection{Edges: []*ManufacturerEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := m.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := m.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	m = pager.applyCursors(m, after, before)
	m = pager.applyOrder(m, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		m = m.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		m = m.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := m.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Manufacturer
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Manufacturer {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Manufacturer {
			return nodes[i]
		}
	}

	conn.Edges = make([]*ManufacturerEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &ManufacturerEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// ManufacturerOrderFieldName orders Manufacturer by name.
	ManufacturerOrderFieldName = &ManufacturerOrderField{
		field: manufacturer.FieldName,
		toCursor: func(m *Manufacturer) Cursor {
			return Cursor{
				ID:    m.ID,
				Value: m.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f ManufacturerOrderField) String() string {
	var str string
	switch f.field {
	case manufacturer.FieldName:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f ManufacturerOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *ManufacturerOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("ManufacturerOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *ManufacturerOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid ManufacturerOrderField", str)
	}
	return nil
}

// ManufacturerOrderField defines the ordering field of Manufacturer.
type ManufacturerOrderField struct {
	field    string
	toCursor func(*Manufacturer) Cursor
}

// ManufacturerOrder defines the ordering of Manufacturer.
type ManufacturerOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *ManufacturerOrderField `json:"field"`
}

// DefaultManufacturerOrder is the default ordering of Manufacturer.
var DefaultManufacturerOrder = &ManufacturerOrder{
	Direction: OrderDirectionAsc,
	Field: &ManufacturerOrderField{
		field: manufacturer.FieldID,
		toCursor: func(m *Manufacturer) Cursor {
			return Cursor{ID: m.ID}
		},
	},
}

// ToEdge converts Manufacturer into ManufacturerEdge.
func (m *Manufacturer) ToEdge(order *ManufacturerOrder) *ManufacturerEdge {
	if order == nil {
		order = DefaultManufacturerOrder
	}
	return &ManufacturerEdge{
		Node:   m,
		Cursor: order.Field.toCursor(m),
	}
}

// ProductEdge is the edge representation of Product.
type ProductEdge struct {
	Node   *Product `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// ProductConnection is the connection containing edges to Product.
type ProductConnection struct {
	Edges      []*ProductEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

// ProductPaginateOption enables pagination customization.
type ProductPaginateOption func(*productPager) error

// WithProductOrder configures pagination ordering.
func WithProductOrder(order *ProductOrder) ProductPaginateOption {
	if order == nil {
		order = DefaultProductOrder
	}
	o := *order
	return func(pager *productPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultProductOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithProductFilter configures pagination filter.
func WithProductFilter(filter func(*ProductQuery) (*ProductQuery, error)) ProductPaginateOption {
	return func(pager *productPager) error {
		if filter == nil {
			return errors.New("ProductQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type productPager struct {
	order  *ProductOrder
	filter func(*ProductQuery) (*ProductQuery, error)
}

func newProductPager(opts []ProductPaginateOption) (*productPager, error) {
	pager := &productPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultProductOrder
	}
	return pager, nil
}

func (p *productPager) applyFilter(query *ProductQuery) (*ProductQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *productPager) toCursor(pr *Product) Cursor {
	return p.order.Field.toCursor(pr)
}

func (p *productPager) applyCursors(query *ProductQuery, after, before *Cursor) *ProductQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultProductOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *productPager) applyOrder(query *ProductQuery, reverse bool) *ProductQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultProductOrder.Field {
		query = query.Order(direction.orderFunc(DefaultProductOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Product.
func (pr *ProductQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...ProductPaginateOption,
) (*ProductConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newProductPager(opts)
	if err != nil {
		return nil, err
	}

	if pr, err = pager.applyFilter(pr); err != nil {
		return nil, err
	}

	conn := &ProductConnection{Edges: []*ProductEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := pr.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := pr.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	pr = pager.applyCursors(pr, after, before)
	pr = pager.applyOrder(pr, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		pr = pr.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		pr = pr.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := pr.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Product
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Product {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Product {
			return nodes[i]
		}
	}

	conn.Edges = make([]*ProductEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &ProductEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// ProductOrderFieldName orders Product by name.
	ProductOrderFieldName = &ProductOrderField{
		field: product.FieldName,
		toCursor: func(pr *Product) Cursor {
			return Cursor{
				ID:    pr.ID,
				Value: pr.Name,
			}
		},
	}
	// ProductOrderFieldUpc orders Product by upc.
	ProductOrderFieldUpc = &ProductOrderField{
		field: product.FieldUpc,
		toCursor: func(pr *Product) Cursor {
			return Cursor{
				ID:    pr.ID,
				Value: pr.Upc,
			}
		},
	}
	// ProductOrderFieldPrice orders Product by price.
	ProductOrderFieldPrice = &ProductOrderField{
		field: product.FieldPrice,
		toCursor: func(pr *Product) Cursor {
			return Cursor{
				ID:    pr.ID,
				Value: pr.Price,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f ProductOrderField) String() string {
	var str string
	switch f.field {
	case product.FieldName:
		str = "NAME"
	case product.FieldUpc:
		str = "UPC"
	case product.FieldPrice:
		str = "PRICE"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f ProductOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *ProductOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("ProductOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *ProductOrderFieldName
	case "UPC":
		*f = *ProductOrderFieldUpc
	case "PRICE":
		*f = *ProductOrderFieldPrice
	default:
		return fmt.Errorf("%s is not a valid ProductOrderField", str)
	}
	return nil
}

// ProductOrderField defines the ordering field of Product.
type ProductOrderField struct {
	field    string
	toCursor func(*Product) Cursor
}

// ProductOrder defines the ordering of Product.
type ProductOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *ProductOrderField `json:"field"`
}

// DefaultProductOrder is the default ordering of Product.
var DefaultProductOrder = &ProductOrder{
	Direction: OrderDirectionAsc,
	Field: &ProductOrderField{
		field: product.FieldID,
		toCursor: func(pr *Product) Cursor {
			return Cursor{ID: pr.ID}
		},
	},
}

// ToEdge converts Product into ProductEdge.
func (pr *Product) ToEdge(order *ProductOrder) *ProductEdge {
	if order == nil {
		order = DefaultProductOrder
	}
	return &ProductEdge{
		Node:   pr,
		Cursor: order.Field.toCursor(pr),
	}
}

// ReviewEdge is the edge representation of Review.
type ReviewEdge struct {
	Node   *Review `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// ReviewConnection is the connection containing edges to Review.
type ReviewConnection struct {
	Edges      []*ReviewEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

// ReviewPaginateOption enables pagination customization.
type ReviewPaginateOption func(*reviewPager) error

// WithReviewOrder configures pagination ordering.
func WithReviewOrder(order *ReviewOrder) ReviewPaginateOption {
	if order == nil {
		order = DefaultReviewOrder
	}
	o := *order
	return func(pager *reviewPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultReviewOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithReviewFilter configures pagination filter.
func WithReviewFilter(filter func(*ReviewQuery) (*ReviewQuery, error)) ReviewPaginateOption {
	return func(pager *reviewPager) error {
		if filter == nil {
			return errors.New("ReviewQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type reviewPager struct {
	order  *ReviewOrder
	filter func(*ReviewQuery) (*ReviewQuery, error)
}

func newReviewPager(opts []ReviewPaginateOption) (*reviewPager, error) {
	pager := &reviewPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultReviewOrder
	}
	return pager, nil
}

func (p *reviewPager) applyFilter(query *ReviewQuery) (*ReviewQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *reviewPager) toCursor(r *Review) Cursor {
	return p.order.Field.toCursor(r)
}

func (p *reviewPager) applyCursors(query *ReviewQuery, after, before *Cursor) *ReviewQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultReviewOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *reviewPager) applyOrder(query *ReviewQuery, reverse bool) *ReviewQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultReviewOrder.Field {
		query = query.Order(direction.orderFunc(DefaultReviewOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Review.
func (r *ReviewQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...ReviewPaginateOption,
) (*ReviewConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newReviewPager(opts)
	if err != nil {
		return nil, err
	}

	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}

	conn := &ReviewConnection{Edges: []*ReviewEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := r.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := r.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	r = pager.applyCursors(r, after, before)
	r = pager.applyOrder(r, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		r = r.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		r = r.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := r.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Review
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Review {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Review {
			return nodes[i]
		}
	}

	conn.Edges = make([]*ReviewEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &ReviewEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// ReviewOrderFieldBody orders Review by body.
	ReviewOrderFieldBody = &ReviewOrderField{
		field: review.FieldBody,
		toCursor: func(r *Review) Cursor {
			return Cursor{
				ID:    r.ID,
				Value: r.Body,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f ReviewOrderField) String() string {
	var str string
	switch f.field {
	case review.FieldBody:
		str = "body"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f ReviewOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *ReviewOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("ReviewOrderField %T must be a string", v)
	}
	switch str {
	case "body":
		*f = *ReviewOrderFieldBody
	default:
		return fmt.Errorf("%s is not a valid ReviewOrderField", str)
	}
	return nil
}

// ReviewOrderField defines the ordering field of Review.
type ReviewOrderField struct {
	field    string
	toCursor func(*Review) Cursor
}

// ReviewOrder defines the ordering of Review.
type ReviewOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *ReviewOrderField `json:"field"`
}

// DefaultReviewOrder is the default ordering of Review.
var DefaultReviewOrder = &ReviewOrder{
	Direction: OrderDirectionAsc,
	Field: &ReviewOrderField{
		field: review.FieldID,
		toCursor: func(r *Review) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

// ToEdge converts Review into ReviewEdge.
func (r *Review) ToEdge(order *ReviewOrder) *ReviewEdge {
	if order == nil {
		order = DefaultReviewOrder
	}
	return &ReviewEdge{
		Node:   r,
		Cursor: order.Field.toCursor(r),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}

	conn := &UserConnection{Edges: []*UserEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := u.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := u.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		u = u.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}

	conn.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

var (
	// UserOrderFieldEmail orders User by email.
	UserOrderFieldEmail = &UserOrderField{
		field: user.FieldEmail,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Email,
			}
		},
	}
	// UserOrderFieldUsername orders User by username.
	UserOrderFieldUsername = &UserOrderField{
		field: user.FieldUsername,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Username,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.field {
	case user.FieldEmail:
		str = "email"
	case user.FieldUsername:
		str = "username"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "email":
		*f = *UserOrderFieldEmail
	case "username":
		*f = *UserOrderFieldUsername
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
