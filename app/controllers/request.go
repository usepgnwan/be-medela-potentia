package controllers

import (
	"be-medela-potentia/app/helpers"
	"be-medela-potentia/app/models"
	connection "be-medela-potentia/conection"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Tags Request
// @Summary create Request
// @Description bearer di perlukan untuk menyimpan siapa yang melakukan request / approve
// @Accept json
// @Produce json
// @Param user body models.Request true "Buat Request"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Param Authorization header string true "Authorization JWT login" default(Bearer eyJhbGciOiJ...)
// @Router /api/request [post]
func PostRequest(c *fiber.Ctx) error {
	var data models.Request
	claims := c.Locals("user").(*models.JwtUser)
	data.UserID = claims.ID
	data.Status = "PENDING"
	current_step := 1
	data.CurrentStep = current_step

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

	workflow_id := data.WorkflowId
	var request models.Request
	if err := connection.DB.Model(&models.Request{}).Where("workflow_id = ?", workflow_id).Where("current_step = ? ", current_step).First(&request).Order("current_step desc").Order("created_at desc").Error; err == nil {
		return c.Status(http.StatusAccepted).JSON(helpers.Response{
			Success: false,
			Error:   "Request sudah ada dalam proses, status terakhir " + request.Status + " Lakukan Approval!",
		})
	}

	var step models.WorkflowStep
	if err := connection.DB.Model(&models.WorkflowStep{}).Where("level =?", current_step).Where("workflow_id =?", data.WorkflowId).First(&step).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(helpers.Response{
			Success: false,
			Error:   "Step workflow tidak tersedia pastikan step sudah di setting",
		})
	}

	// tolak aku ga punya authorized
	if claims.RoleId != step.RoleId {
		return c.Status(http.StatusUnauthorized).JSON(helpers.Response{
			Success: false,
			Error:   fmt.Sprintf("Anda tidak memiliki hak akses"),
		})
	}

	if data.Amount < step.MinAmount {
		return c.Status(http.StatusBadRequest).JSON(helpers.Response{
			Success: false,
			Error:   fmt.Sprintf("Minimal amount tidak terpenuhi. Minimal: %d", step.MinAmount),
		})
	}

	if err := connection.DB.Model(&models.Request{}).Create(&data).Error; err != nil {
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

// @Tags Request
// @Summary reject Request
// @Accept json
// @Produce json
// @Param id path string true "id request (wajib)"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Param Authorization header string true "Authorization JWT login" default(Bearer eyJhbGciOiJ...)
// @Router /api/request/{id}/reject [post]
func RejectRequest(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JwtUser)
	id := c.Params("id")
	var request models.Request
	if err := connection.DB.Model(&models.Request{}).Where("id = ?", id).Order("created_at desc").First(&request).Error; err != nil {
		return c.Status(http.StatusAccepted).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	var step models.WorkflowStep
	if err := connection.DB.Model(&models.WorkflowStep{}).Where("level =?", request.CurrentStep).Where("workflow_id =?", request.WorkflowId).First(&step).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(helpers.Response{
			Success: false,
			Error:   "Step workflow tidak tersedia pastikan step sudah di setting",
		})
	}

	var checkRequest models.Request
	if err := connection.DB.Model(&models.Request{}).Where("workflow_id = ?", request.WorkflowId).Where("current_step = ?", request.CurrentStep).Order("created_at desc").First(&checkRequest).Error; err != nil {
		return c.Status(http.StatusAccepted).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	if checkRequest.Status == "REJECTED" {
		return c.Status(http.StatusAccepted).JSON(helpers.Response{
			Success: false,
			Error:   fmt.Sprintf("Request sudah di tolak"),
		})
	}

	// tolak kalo ga punya authorized
	if claims.RoleId != step.RoleId {
		return c.Status(http.StatusUnauthorized).JSON(helpers.Response{
			Success: false,
			Error:   fmt.Sprintf("Anda tidak memiliki hak akses"),
		})
	}

	// reject by
	var newRequest models.Request
	newRequest.Amount = request.Amount
	newRequest.UserID = claims.ID
	newRequest.Status = "REJECTED"
	newRequest.WorkflowId = request.WorkflowId
	newRequest.UserID = claims.ID
	newRequest.CurrentStep = request.CurrentStep

	if err := connection.DB.Model(&models.Request{}).Create(&newRequest).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(helpers.Response{
		Success: true,
		Data:    newRequest,
	})
}

// @Tags Request
// @Summary reject Request
// @Accept json
// @Produce json
// @Param id path string true "id request (wajib)"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Param Authorization header string true "Authorization JWT login" default(Bearer eyJhbGciOiJ...)
// @Router /api/request/{id}/approve [post]
func ApproveRequest(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JwtUser)
	id := c.Params("id")
	var request models.Request
	if err := connection.DB.Model(&models.Request{}).Where("id = ?", id).Order("created_at desc").First(&request).Error; err != nil {
		return c.Status(http.StatusAccepted).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}
	var checkRequest models.Request

	// ambil curent request terakhir
	if err := connection.DB.Model(&models.Request{}).Where("workflow_id = ?", request.WorkflowId).Order("created_at desc").First(&checkRequest).Error; err != nil {
		return c.Status(http.StatusAccepted).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	// kalo request terakhir status rejected / approved proses approval selesai
	if checkRequest.Status == "REJECTED" || checkRequest.Status == "APPROVED" {
		return c.Status(http.StatusAccepted).JSON(helpers.Response{
			Success: false,
			Error:   "Request sudah selesai, Status : " + checkRequest.Status,
		})
	}
	level_step := checkRequest.CurrentStep + 1
	var step models.WorkflowStep

	// ambil step selanjutnya untuk proses approval di step selanjutnya
	if err := connection.DB.Model(&models.WorkflowStep{}).Where("level =?", level_step).Where("workflow_id =?", request.WorkflowId).First(&step).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(helpers.Response{
			Success: false,
			Error:   "Step workflow tidak tersedia pastikan step sudah di setting",
		})
	}

	// cek akses yang membolehkan request
	if claims.RoleId != step.RoleId {
		return c.Status(http.StatusUnauthorized).JSON(helpers.Response{
			Success: false,
			Error:   fmt.Sprintf("Anda tidak memiliki hak akses"),
		})
	}

	var newRequest models.Request
	current_step := step.Level
	newRequest.Amount = request.Amount
	newRequest.UserID = claims.ID
	newRequest.Status = "PENDING"
	newRequest.WorkflowId = request.WorkflowId
	newRequest.UserID = claims.ID
	newRequest.CurrentStep = current_step

	// cek step selanjutnya masih ada atau engga kalo masih ada APPROVED
	var nextstep models.WorkflowStep
	if err := connection.DB.Model(&models.WorkflowStep{}).Where("level = ?", (current_step+1)).Where("workflow_id =?", request.WorkflowId).First(&nextstep).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRequest.Status = "APPROVED"
		}
	}

	// update untuk approval seblumnya
	if err := connection.DB.Model(&models.Request{}).
		Where("id = ?", checkRequest.ID).
		Update("status", "APPROVED").Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	// simpan request step seblmnya
	if err := connection.DB.Model(&models.Request{}).Create(&newRequest).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(helpers.Response{
		Success: true,
		Data:    newRequest,
	})
}

// @Tags Request
// @Summary Mengambil detail request
// @Description - id adalah id request
// @Description - memunculkan semua request di workflow
// @Param id path string true "id workflow (wajib)"
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Response 200 {object} helpers.Response "Successfully fetched Parameter System"
// @Response 400 {object} helpers.Response "Bad request"
// @Router /api/request/{id} [get]
func GetDetailRequest(c *fiber.Ctx) error {
	id := c.Params("id")

	var request []models.Request
	if err := connection.DB.Model(&models.Request{}).Preload("Workflow", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "user_id").Preload("RequestBy", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "username")
		})
	}).Preload("ApproveBy", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "username").Preload("UserRole")
	}).Where("id = ?", id).Find(&request).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(helpers.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(helpers.Response{
		Success: true,
		Data:    request,
	})
}
