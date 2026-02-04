package controllers

import (
	"be-medela-potentia/app/helpers"
	"be-medela-potentia/app/models"
	connection "be-medela-potentia/conection"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Tags Workflow Step
// @Summary Mengambil detail by id workflow step
// @Param id path string true "id (wajib)"
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Response 200 {object} helpers.Response "Successfully fetched Parameter System"
// @Response 400 {object} helpers.Response "Bad request"
// @Router /api/workflows-step/{id} [get]
func GetDetailWorkflowStep(c *fiber.Ctx) error {
	id := c.Params("id")

	var workflow models.WorkflowStep
	if err := connection.DB.Model(&models.WorkflowStep{}).Preload("Workflow").Preload("Actor").First(&workflow, "id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusNotFound).JSON(helpers.Response{
		Success: true,
		Data:    workflow,
	})
}

// @Tags Workflow Step
// @Summary create Workflow
// @Accept json
// @Produce json
// @Param user body models.WorkflowStep true "Buat Workflow"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Param Authorization header string true "Authorization JWT login" default(Bearer eyJhbGciOiJ...)
// @Router /api/workflows-step [post]
func PostWorkflowStep(c *fiber.Ctx) error {
	var data models.WorkflowStep
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	validationErr, err := helpers.ValidateData(&data)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	if len(validationErr) > 0 {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   "Validation failed",
			Data:    validationErr,
		})
	}

	if err := connection.DB.Model(&models.WorkflowStep{}).Create(&data).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(helpers.Response{
		Success: true,
		Data:    data,
	})
}
