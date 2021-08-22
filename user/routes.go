package user

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//find all people
func Get_all(c *fiber.Ctx) error {
	u := User{ID: uuid.New(), Name: "pippo", CreatedAt: time.Now().Local().UTC()}
	uu := User{ID: uuid.New(), Name: "pluto", CreatedAt: time.Now().Local().UTC()}
	uuu := User{ID: uuid.New(), Name: "topolino", CreatedAt: time.Now().Local().UTC()}
	list := make([]User, 3)
	list = append(list, u)
	list = append(list, uu)
	list = append(list, uuu)
	return c.JSON(list)
}

func Find(c *fiber.Ctx) error {
	msg := fmt.Sprintf("hello, %s", c.Params("name"))
	return c.JSON(msg)
}

func Create(c *fiber.Ctx) error {
	return c.JSON("dcc")
}

func Update(c *fiber.Ctx) error {
	return c.JSON("vcdx")
}

func Delete(c *fiber.Ctx) error {
	return c.JSON("sdvsd")
}
