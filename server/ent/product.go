// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Upc holds the value of the "upc" field.
	Upc string `json:"upc,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges                 ProductEdges `json:"edges"`
	manufacturer_products *uuid.UUID
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Manufacturer holds the value of the manufacturer edge.
	Manufacturer *Manufacturer `json:"manufacturer,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ManufacturerOrErr returns the Manufacturer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) ManufacturerOrErr() (*Manufacturer, error) {
	if e.loadedTypes[0] {
		if e.Manufacturer == nil {
			// The edge manufacturer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: manufacturer.Label}
		}
		return e.Manufacturer, nil
	}
	return nil, &NotLoadedError{edge: "manufacturer"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[1] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldPrice:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldUpc:
			values[i] = new(sql.NullString)
		case product.FieldID:
			values[i] = new(uuid.UUID)
		case product.ForeignKeys[0]: // manufacturer_products
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Product", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldUpc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field upc", values[i])
			} else if value.Valid {
				pr.Upc = value.String
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = int(value.Int64)
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer_products", values[i])
			} else if value.Valid {
				pr.manufacturer_products = new(uuid.UUID)
				*pr.manufacturer_products = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryManufacturer queries the "manufacturer" edge of the Product entity.
func (pr *Product) QueryManufacturer() *ManufacturerQuery {
	return (&ProductClient{config: pr.config}).QueryManufacturer(pr)
}

// QueryReviews queries the "reviews" edge of the Product entity.
func (pr *Product) QueryReviews() *ReviewQuery {
	return (&ProductClient{config: pr.config}).QueryReviews(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", upc=")
	builder.WriteString(pr.Upc)
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}

func (Product) IsEntity() {}
