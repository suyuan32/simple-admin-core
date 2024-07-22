// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/configuration"
)

// Configuration is the model entity for the Configuration schema.
type Configuration struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Sort Number | 排序编号
	Sort uint32 `json:"sort,omitempty"`
	// State true: normal false: ban | 状态 true 正常 false 禁用
	State bool `json:"state,omitempty"`
	// Configurarion name | 配置名称
	Name string `json:"name,omitempty"`
	// Configuration key | 配置的键名
	Key string `json:"key,omitempty"`
	// Configuraion value | 配置的值
	Value string `json:"value,omitempty"`
	// Configuration category | 配置的分类
	Category string `json:"category,omitempty"`
	// Remark | 备注
	Remark       string `json:"remark,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Configuration) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case configuration.FieldState:
			values[i] = new(sql.NullBool)
		case configuration.FieldID, configuration.FieldSort:
			values[i] = new(sql.NullInt64)
		case configuration.FieldName, configuration.FieldKey, configuration.FieldValue, configuration.FieldCategory, configuration.FieldRemark:
			values[i] = new(sql.NullString)
		case configuration.FieldCreatedAt, configuration.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Configuration fields.
func (c *Configuration) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case configuration.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case configuration.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case configuration.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case configuration.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				c.Sort = uint32(value.Int64)
			}
		case configuration.FieldState:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				c.State = value.Bool
			}
		case configuration.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case configuration.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				c.Key = value.String
			}
		case configuration.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				c.Value = value.String
			}
		case configuration.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				c.Category = value.String
			}
		case configuration.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				c.Remark = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Configuration.
// This includes values selected through modifiers, order, etc.
func (c *Configuration) GetValue(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Configuration.
// Note that you need to call Configuration.Unwrap() before calling this method if this Configuration
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Configuration) Update() *ConfigurationUpdateOne {
	return NewConfigurationClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Configuration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Configuration) Unwrap() *Configuration {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Configuration is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Configuration) String() string {
	var builder strings.Builder
	builder.WriteString("Configuration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", c.Sort))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", c.State))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(c.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(c.Value)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(c.Category)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(c.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Configurations is a parsable slice of Configuration.
type Configurations []*Configuration
