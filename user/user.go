package user

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//user struct
type User struct {
	gorm.Model
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name      string    `db:"name" json:"name"`
	Surname   string    `db:"surname" json:"surname"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	LastLogin time.Time `db:"last_login" json:"last_login"`
}

// Value make the User struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (u User) Value() (driver.Value, error) {
	return json.Marshal(u)
}

// Scan make the User struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (u *User) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(j, &u)
}
