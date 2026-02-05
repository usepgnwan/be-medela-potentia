package controllers

import (
	"be-medela-potentia/app/helpers"
	"be-medela-potentia/app/models"
	connection "be-medela-potentia/conection"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Tags Workflow
// @Summary Mengambil semua data
// @Produce json
// @Param Page query int false "Posisi page di halaman berapa (Wajib jika Limit di isi)"
// @Param Limit query int false "Data yang di tampikan per halamanan (Wajib jika Page di isi)"
// @Param KeySearch query string false "Key untuk search data spesifik (Direkomendasikan jika search untuk tabel Page & Limit isi juga)"
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Response 200 {object} helpers.Response "Successfully fetched Parameter System"
// @Response 400 {object} helpers.Response "Bad request"
// @Router /api/workflows [get]
func GetWorkflow(c *fiber.Ctx) error {
	var pag helpers.Pagination

	if err := c.QueryParser(&pag); err != nil {

		return c.Status(http.StatusBadRequest).JSON(helpers.Response{
			Error:   "Failed parsing query",
			Success: false,
		})
	}

	var (
		data       []models.Workflow
		totalRows  int64
		key_search string = pag.KeySearch
	)

	query := connection.DB.Model(&models.Workflow{}).Preload("WorkflowStep", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Actor")
	}).Preload("RequestBy", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "role_id", "username").Preload("UserRole")
	}).Preload("Request")
	if key_search != "" {
		key_search = "%" + key_search + "%"
		query = query.Where("name ILIKE ?", key_search)
	}

	countQuery := query.Session(&gorm.Session{})
	if err := countQuery.Count(&totalRows).Error; err != nil {
		return nil
	}

	if pag.Limit > 0 {
		query = query.Scopes(helpers.Paginate(pag.GetOffset(), pag.Limit))
	}

	query.Find(&data)

	err := query.Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	var result helpers.ResponsePaginate = helpers.ResponsePaginate{
		Response: helpers.Response{
			Success: true,
			Data:    data,
		},
	}

	if pag.Limit == 0 && totalRows == 1 {
		result.Data = data
	}

	if pag.Limit > 0 {
		result.Meta = &helpers.MetaPaginate{
			Total:       totalRows,
			CurrentPage: int64(pag.Page),
			Limit:       int64(pag.Limit),
			Pages:       (totalRows + int64(pag.Limit) - 1) / int64(pag.Limit),
		}
	}

	return c.JSON(result)
}

// @Tags Workflow
// @Summary Mengambil detail by id workflow
// @Param id path string true "id (wajib)"
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Response 200 {object} helpers.Response "Successfully fetched Parameter System"
// @Response 400 {object} helpers.Response "Bad request"
// @Router /api/workflows/{id} [get]
func GetDetailWorkflow(c *fiber.Ctx) error {
	id := c.Params("id")

	var workflow models.Workflow
	if err := connection.DB.Model(&models.Workflow{}).Preload("WorkflowStep", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Actor")
	}).Preload("RequestBy", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "username").Preload("UserRole")
	}).Preload("Request", func(db *gorm.DB) *gorm.DB {
		return db.Preload("ApproveBy", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "username").Preload("UserRole")
		})
	}).First(&workflow, "id = ?", id).Error; err != nil {
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

// @Tags Workflow
// @Summary create Workflow
// @Accept json
// @Produce json
// @Param user body models.Workflow true "Buat Workflow"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Param Authorization header string true "Authorization JWT login" default(Bearer eyJhbGciOiJ...)
// @Router /api/workflows [post]
func PostWorkflow(c *fiber.Ctx) error {
	var data models.Workflow
	claims := c.Locals("user").(*models.JwtUser)
	data.UserID = claims.ID
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
		return c.Status(http.StatusBadRequest).JSON(helpers.Response{
			Success: false,
			Error:   "Validation failed",
			Data:    validationErr,
		})
	}

	if err := connection.DB.Model(&models.Workflow{}).Create(&data).Error; err != nil {
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
