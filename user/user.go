package user

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"os/user"

	"github.com/OmarAouini/go-fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

//user struct
type User struct {
	ID       int            `gorm:"primary_key"`
	Name     sql.NullString `gorm:"column:name" json:"name"`
	Surname  sql.NullString `gorm:"column:surname" json:"surname"`
	Username sql.NullString `gorm:"column:username" json:"username"`
	Password sql.NullString `gorm:"column:password" json:"password"`
	Email    sql.NullString `gorm:"column:email" json:"email"`
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

//CRUD

//find all users
func Get_all(c *fiber.Ctx) error {
	var users []user.User
	database.Pool.Find(&users)
	if len(users) == 0 {
		return c.JSON([]User{})
	}
	fmt.Println(c.JSON(users))
	return c.JSON(users)
}

//find single user
func Find(c *fiber.Ctx) error {
	var user User
	var username = c.Params("username")
	database.Pool.First(user.Username.String == username)
	return c.JSON(user)
}

//create a new user
func Create(c *fiber.Ctx) error {
	// var user User
	return c.JSON("dcc")
}

//update an existing user
func Update(c *fiber.Ctx) error {
	return c.JSON("vcdx")
}

//delete an existing user
func Delete(c *fiber.Ctx) error {
	return c.JSON("sdvsd")
}
