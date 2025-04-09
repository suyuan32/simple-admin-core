// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionary"
)

// Dictionary Table | 字典信息表
type Dictionary struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// The title shown in the ui | 展示名称 （建议配合i18n）
	Title string `json:"title,omitempty"`
	// The name of dictionary for search | 字典搜索名称
	Name string `json:"name,omitempty"`
	// The description of dictionary | 字典的描述
	Desc string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DictionaryQuery when eager-loading is set.
	Edges        DictionaryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DictionaryEdges holds the relations/edges for other nodes in the graph.
type DictionaryEdges struct {
	// DictionaryDetails holds the value of the dictionary_details edge.
	DictionaryDetails []*DictionaryDetail `json:"dictionary_details,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DictionaryDetailsOrErr returns the DictionaryDetails value or an error if the edge
// was not loaded in eager-loading.
func (e DictionaryEdges) DictionaryDetailsOrErr() ([]*DictionaryDetail, error) {
	if e.loadedTypes[0] {
		return e.DictionaryDetails, nil
	}
	return nil, &NotLoadedError{edge: "dictionary_details"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dictionary) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dictionary.FieldID, dictionary.FieldStatus:
			values[i] = new(sql.NullInt64)
		case dictionary.FieldTitle, dictionary.FieldName, dictionary.FieldDesc:
			values[i] = new(sql.NullString)
		case dictionary.FieldCreatedAt, dictionary.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dictionary fields.
func (d *Dictionary) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dictionary.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = uint64(value.Int64)
		case dictionary.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case dictionary.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case dictionary.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = uint8(value.Int64)
			}
		case dictionary.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				d.Title = value.String
			}
		case dictionary.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case dictionary.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				d.Desc = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Dictionary.
// This includes values selected through modifiers, order, etc.
func (d *Dictionary) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryDictionaryDetails queries the "dictionary_details" edge of the Dictionary entity.
func (d *Dictionary) QueryDictionaryDetails() *DictionaryDetailQuery {
	return NewDictionaryClient(d.config).QueryDictionaryDetails(d)
}

// Update returns a builder for updating this Dictionary.
// Note that you need to call Dictionary.Unwrap() before calling this method if this Dictionary
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dictionary) Update() *DictionaryUpdateOne {
	return NewDictionaryClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dictionary entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dictionary) Unwrap() *Dictionary {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dictionary is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dictionary) String() string {
	var builder strings.Builder
	builder.WriteString("Dictionary(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(d.Title)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(d.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// Dictionaries is a parsable slice of Dictionary.
type Dictionaries []*Dictionary
