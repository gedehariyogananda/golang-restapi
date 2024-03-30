package bookcontroller

import (
	"test/golang/config"
	"test/golang/models"

	"github.com/gofiber/fiber/v2"
)

// config.DB == konfirgurasi set state masuk DATABASE nya
// models.Book == ini hanya nama table di model nya

func Index(c *fiber.Ctx) error {

	// index == var books tampung array
	// lalu simpan ke DATABASE

	var books []models.Book
	config.DB.Find(&books)

	return c.JSON(books)
}

func Show(c *fiber.Ctx) error {

	// ambil id nya dulu dengan params
	// init model book dalam variable books
	// lalu ambil data dari database berdasarkan id dengan config.DB.First(inisialisasi model book, id nya)

	var books models.Book
	id := c.Params("id")
	handler := config.DB.First(&books, id)
	if handler.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "data not found",
		})
	}
	return c.JSON(books)

}

func Store(c *fiber.Ctx) error {

	// init model book dalam variabel books
	// lalu bodtparser == isian json dalam postman masukkan model inisiate nya
	var books models.Book
	if err := c.BodyParser(&books); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// lalu simpan di Database
	if err := config.DB.Create(&books); err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    books,
		"message": "berhasil created data",
	})

}

func Update(c *fiber.Ctx) error {

	// ambil id dengan parrams
	// init models book dalam variabel
	// bodyparser dengan model yang ingin diganti
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// update data berdasarkan id lalu simpan di database
	// dicek dengan rows affected
	if err := config.DB.Where("id = ?", id).Updates(&book); err.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error,
		})
	}

	return c.JSON(fiber.Map{
		"message": "berhasil diupdated",
	})

}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")

	var book models.Book

	if err := config.DB.Where("id = ?", id).Delete(&book); err.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "data tidak ada",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data berhasil diaposition",
	})
}
