package person

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//find all people
func Get_all(c *fiber.Ctx) error {
	p := Person{ID: uuid.New(), Name: "pippo", Age: 22, CreatedAt: time.Now().Local().UTC()}
	pp := Person{ID: uuid.New(), Name: "pluto", Age: 34, CreatedAt: time.Now().Local().UTC()}
	ppp := Person{ID: uuid.New(), Name: "topolino", Age: 56, CreatedAt: time.Now().Local().UTC()}
	list := make([]Person, 3)
	list = append(list, p)
	list = append(list, pp)
	list = append(list, ppp)
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
