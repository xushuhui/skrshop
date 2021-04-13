// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"skrshop/internal/data/ent/productattrvalue"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ProductAttrValue is the model entity for the ProductAttrValue schema.
type ProductAttrValue struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 状态 1:enable, 0:disable, -1:deleted
	Status int8 `json:"status,omitempty"`
	// 创建时间
	CreateAt time.Time `json:"create_at,omitempty"`
	// 更新时间
	UpdateAt time.Time `json:"update_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductAttrValue) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case productattrvalue.FieldID, productattrvalue.FieldStatus:
			values[i] = &sql.NullInt64{}
		case productattrvalue.FieldCreateAt, productattrvalue.FieldUpdateAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductAttrValue", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductAttrValue fields.
func (pav *ProductAttrValue) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productattrvalue.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pav.ID = int(value.Int64)
		case productattrvalue.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pav.Status = int8(value.Int64)
			}
		case productattrvalue.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				pav.CreateAt = value.Time
			}
		case productattrvalue.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				pav.UpdateAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ProductAttrValue.
// Note that you need to call ProductAttrValue.Unwrap() before calling this method if this ProductAttrValue
// was returned from a transaction, and the transaction was committed or rolled back.
func (pav *ProductAttrValue) Update() *ProductAttrValueUpdateOne {
	return (&ProductAttrValueClient{config: pav.config}).UpdateOne(pav)
}

// Unwrap unwraps the ProductAttrValue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pav *ProductAttrValue) Unwrap() *ProductAttrValue {
	tx, ok := pav.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductAttrValue is not a transactional entity")
	}
	pav.config.driver = tx.drv
	return pav
}

// String implements the fmt.Stringer.
func (pav *ProductAttrValue) String() string {
	var builder strings.Builder
	builder.WriteString("ProductAttrValue(")
	builder.WriteString(fmt.Sprintf("id=%v", pav.ID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", pav.Status))
	builder.WriteString(", create_at=")
	builder.WriteString(pav.CreateAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(pav.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProductAttrValues is a parsable slice of ProductAttrValue.
type ProductAttrValues []*ProductAttrValue

func (pav ProductAttrValues) config(cfg config) {
	for _i := range pav {
		pav[_i].config = cfg
	}
}
