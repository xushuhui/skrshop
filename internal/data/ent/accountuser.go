// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"skrshop/internal/data/ent/accountuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// AccountUser is the model entity for the AccountUser schema.
type AccountUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 创建时间
	CreateAt time.Time `json:"create_at,omitempty"`
	// 状态 1:enable, 0:disable, -1:deleted
	Status int8 `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case accountuser.FieldID, accountuser.FieldStatus:
			values[i] = &sql.NullInt64{}
		case accountuser.FieldEmail, accountuser.FieldPhone, accountuser.FieldPassword:
			values[i] = &sql.NullString{}
		case accountuser.FieldCreateAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type AccountUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountUser fields.
func (au *AccountUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accountuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			au.ID = int(value.Int64)
		case accountuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				au.Email = value.String
			}
		case accountuser.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				au.Phone = value.String
			}
		case accountuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				au.Password = value.String
			}
		case accountuser.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				au.CreateAt = value.Time
			}
		case accountuser.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				au.Status = int8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AccountUser.
// Note that you need to call AccountUser.Unwrap() before calling this method if this AccountUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (au *AccountUser) Update() *AccountUserUpdateOne {
	return (&AccountUserClient{config: au.config}).UpdateOne(au)
}

// Unwrap unwraps the AccountUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (au *AccountUser) Unwrap() *AccountUser {
	tx, ok := au.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccountUser is not a transactional entity")
	}
	au.config.driver = tx.drv
	return au
}

// String implements the fmt.Stringer.
func (au *AccountUser) String() string {
	var builder strings.Builder
	builder.WriteString("AccountUser(")
	builder.WriteString(fmt.Sprintf("id=%v", au.ID))
	builder.WriteString(", email=")
	builder.WriteString(au.Email)
	builder.WriteString(", phone=")
	builder.WriteString(au.Phone)
	builder.WriteString(", password=")
	builder.WriteString(au.Password)
	builder.WriteString(", create_at=")
	builder.WriteString(au.CreateAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", au.Status))
	builder.WriteByte(')')
	return builder.String()
}

// AccountUsers is a parsable slice of AccountUser.
type AccountUsers []*AccountUser

func (au AccountUsers) config(cfg config) {
	for _i := range au {
		au[_i].config = cfg
	}
}
