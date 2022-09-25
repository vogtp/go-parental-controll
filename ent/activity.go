// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vogtp/go-parental-control/ent/activity"
)

// Activity is the model entity for the Activity schema.
type Activity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration     int64 `json:"duration,omitempty"`
	day_activity *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Activity) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case activity.FieldID, activity.FieldDuration:
			values[i] = new(sql.NullInt64)
		case activity.FieldUsername:
			values[i] = new(sql.NullString)
		case activity.ForeignKeys[0]: // day_activity
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Activity", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Activity fields.
func (a *Activity) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case activity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case activity.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				a.Username = value.String
			}
		case activity.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				a.Duration = value.Int64
			}
		case activity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field day_activity", value)
			} else if value.Valid {
				a.day_activity = new(int)
				*a.day_activity = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Activity.
// Note that you need to call Activity.Unwrap() before calling this method if this Activity
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Activity) Update() *ActivityUpdateOne {
	return (&ActivityClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Activity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Activity) Unwrap() *Activity {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Activity is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Activity) String() string {
	var builder strings.Builder
	builder.WriteString("Activity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("username=")
	builder.WriteString(a.Username)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", a.Duration))
	builder.WriteByte(')')
	return builder.String()
}

// Activities is a parsable slice of Activity.
type Activities []*Activity

func (a Activities) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
