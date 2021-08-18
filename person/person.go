package person

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

//person struct
type Person struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name      string    `db:"name" json:"name"`
	Age       int       `db:"age" json:"age"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Value make the Person struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (p Person) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan make the Person struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (p *Person) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(j, &p)
}
