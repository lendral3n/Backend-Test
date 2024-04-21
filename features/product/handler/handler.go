package handler

import (
	"lendra/features/product"
	"lendra/utils/middlewares"
	"lendra/utils/responses"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
}

func New(service product.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

func (handler *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	input := ProductRequest{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot parse JSON", "data": err})
	}

	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	productCore := RequestToCore(input, uint(userIdLogin))

	err = handler.productService.Create(userIdLogin, productCore)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot create product", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Product created successfully", "data": nil})
}

func (handler *ProductHandler) GetAllProduct(c *fiber.Ctx) error {
	products, err := handler.productService.GetAllProduct()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot fetch products", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Products fetched successfully", "data": products})
}

func (handler *ProductHandler) GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot parse id", "data": err})
	}

	product, err := handler.productService.GetProductById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot fetch product", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Product fetched successfully", "data": product})
}

func (handler *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot parse id", "data": err})
	}

	input := ProductRequest{}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot parse JSON", "data": err})
	}

	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	productCore := RequestToCore(input, uint(userIdLogin))


	err = handler.productService.Update(userIdLogin, id, productCore)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot update product", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Product updated successfully", "data": nil})
}

func (handler *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot parse id", "data": err})
	}

	userIdLogin, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.WebResponse("Invalid access token", nil))
	}

	err = handler.productService.Delete(userIdLogin, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Cannot delete product", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Product deleted successfully", "data": nil})
}
