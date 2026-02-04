package controllers

import (
	"be-medela-potentia/app/helpers"
	"be-medela-potentia/app/models"
	connection "be-medela-potentia/conection"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Tags Roles
// @Summary Mengambil semua data
// @Produce json
// @Param Page query int false "Posisi page di halaman berapa (Wajib jika Limit di isi)"
// @Param Limit query int false "Data yang di tampikan per halamanan (Wajib jika Page di isi)"
// @Param KeySearch query string false "Key untuk search data spesifik (Direkomendasikan jika search untuk tabel Page & Limit isi juga)"
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Response 200 {object} helpers.Response "Successfully fetched Parameter System"
// @Response 400 {object} helpers.Response "Bad request"
// @Router /api/roles [get]
func GetRole(c *fiber.Ctx) error {
	var pag helpers.Pagination

	if err := c.QueryParser(&pag); err != nil {

		return c.Status(http.StatusBadRequest).JSON(helpers.Response{
			Error:   "Failed parsing query",
			Success: false,
		})
	}

	var (
		data       []models.UserRole
		totalRows  int64
		key_search string = pag.KeySearch
	)

	query := connection.DB.Model(&models.UserRole{})
	if key_search != "" {
		key_search = "%" + key_search + "%"
		query = query.Where("deskripsi ILIKE ?", key_search)
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
