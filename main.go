package main

import (
	"github.com/Hana-ame/go-dwz/dwz"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// note tested
	app.Get("/go", func(c *fiber.Ctx) error {
		id := c.Query("id", "/error.html")
		if id == "/error.html" {
			return c.Redirect(id)
		}

		url, err := dwz.GetUrl(id)
		if err != nil {
			return err
		}

		return c.Redirect(url)
	})

	// not tested
	app.Post("/link", func(c *fiber.Ctx) error {
		var o dwz.Link
		err := c.BodyParser(&o)
		if err != nil {
			return err
		}

		err = dwz.AddUrl(o.Url, o.Description)
		if err != nil {
			return err
		}

		return c.SendString("OK")
	})

	// not tested
	app.Delete("/link", func(c *fiber.Ctx) error {
		// var o dwz.Link
		// err := c.BodyParser(&o)
		// if err != nil {
		// 	return err
		// }
		id := c.Query("id", "")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Should have id")
		}

		err := dwz.DelUrl(id)
		if err != nil {
			return err
		}

		return c.SendString("OK")
		// return err
	})

	// not tested
	app.Get("/link", func(c *fiber.Ctx) error {
		id := c.Query("id", "")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Should have id")
		}

		o, err := dwz.GetLink(id)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(o)
	})

	// not tested
	app.Get("/tag/:tag", func(c *fiber.Ctx) error {
		tag := c.Params("tag", "")
		if tag == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Sould have tag")
		}

		// query all
		ol, err := dwz.ReadLinksByTag(tag)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(ol)
	})

	// not tested
	app.Post("/tag/:tag", func(c *fiber.Ctx) error {
		tag := c.Params("tag", "")
		id := c.Query("id", "")
		if tag == "" || id == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Sould have both id and tag")
		}

		// o := &dwz.Tag{tag, id, 0}
		err := dwz.AddTag(tag, id)

		return err
	})

	app.Listen(":3000")
}

/*
	what to do next
	测试一下这些接口是不是能用，
	然后用postman看接口是不是能用。

	然后可以开始写前端了。
*/
