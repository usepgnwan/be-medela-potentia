package controllers

import (
	"be-medela-potentia/app/helpers"
	"be-medela-potentia/app/models"
	connection "be-medela-potentia/conection"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Tags Users
// @Summary Mengambil semua data
// @Produce json
// @Param Page query int false "Posisi page di halaman berapa (Wajib jika Limit di isi)"
// @Param Limit query int false "Data yang di tampikan per halamanan (Wajib jika Page di isi)"
// @Param KeySearch query string false "Key untuk search data spesifik (Direkomendasikan jika search untuk tabel Page & Limit isi juga)"
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Response 200 {object} helpers.Response "Successfully fetched Parameter System"
// @Response 400 {object} helpers.Response "Bad request"
// @Router /api/users [get]
func GetUser(c *fiber.Ctx) error {
	var pag helpers.Pagination

	if err := c.QueryParser(&pag); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helpers.Response{
			Success: false,
			Error:   "Failed parsing query",
		})
	}

	var (
		data       []models.User
		totalRows  int64
		key_search string = pag.KeySearch
	)

	query := connection.DB.Model(&models.User{})
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
			Error:   "Failed to fetch data",
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

// @Tags Users
// @Summary registrasi user
// @Accept json
// @Produce json
// @Param user body models.User true "Buat Akun Baru"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Router /api/users [post]
func PostUser(c *fiber.Ctx) error {
	var data models.User
	var res helpers.Response
	res.Error = "Internal server error"
	res.Success = false

	if err := c.BodyParser(&data); err != nil {
		res.Error = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	valErr, err := helpers.ValidateData(&data)
	if err != nil {
		res.Error = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	if len(valErr) > 0 {
		res.Error = "Validation failed"
		res.Data = valErr
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	password := data.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helpers.Response{
			Error:   "Error while hashing password",
			Success: false,
			Data:    nil,
		})
	}
	data.Password = string(hash)

	var roles models.UserRole
	if err := connection.DB.Model(&models.UserRole{}).Where("id = ?", data.RoleId).First(&roles).Error; err != nil {
		res.Error = "Account Role tidak tersedia"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if errCreate := connection.DB.Create(&data).Error; errCreate != nil {
		res.Error = errCreate.Error()
		code := http.StatusInternalServerError
		return c.Status(code).JSON(res)
	}

	if err := connection.DB.Model(&models.User{}).Preload("UserRole").Where("id = ?", data.ID).First(&data).Error; err != nil {
		fmt.Println("gagal get baru")
	}
	res.Success = true
	res.Error = ""
	res.Data = data
	return c.Status(http.StatusOK).JSON(res)
}

type ResponseJWT struct {
	Message string      `json:"message" form:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Token   *string     `json:"token"`
}

// @Tags Users
// @Summary login ke user
// @Accept json
// @Produce json
// @Accept multipart/form-data
// @Produce json
// @Param data body models.DataLogin true "data login"
// @Success 200 {object} ResponseJWT
// @Failure 400 {object} ResponseJWT
// @Router /api/users/login  [post]
// @Param x-api-key header string true "API secret key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
func UserLogin(c *fiber.Ctx) error {
	var data models.DataLogin
	res := new(ResponseJWT)
	res.Data = nil
	res.Status = false
	res.Token = nil

	if err := c.BodyParser(&data); err != nil {
		res.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	validationErr, err := helpers.ValidateData(&data)

	if err != nil {
		res.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	if len(validationErr) > 0 {
		res.Message = "Validation failed"
		res.Data = validationErr
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	var usercontact = make(map[string]string)

	usercontact["username"] = data.UserContact
	var user models.User
	if err := connection.DB.Model(&models.User{}).Where(usercontact).Preload("UserRole").First(&user).Where(usercontact).Error; err != nil {
		res.Message = err.Error()
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		res.Message = "validation Failed"
		res.Data = map[string]string{"username": "username/password yang dimasukan salah"}
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	claims := &models.JwtUser{
		user.ID,
		user.Name,
		user.Username,
		user.RoleId,
		*user.UserRole,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72 * 30)), // expired on 30 days
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("AUTHSECRETKEY")))
	if err != nil {
		res.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Message = "Login success"
	res.Status = true
	res.Data = user
	res.Token = &t
	return c.Status(http.StatusOK).JSON(res)
}

// @Tags Users
// @Summary Claim token
// @Description Untuk memastikan user yang login sesuai
// @Produce json
// @Param x-api-key header string true "Unique API Key" default(UspGnwnelpsrSVTsQYu8LVRyGcl5m7kmi)
// @Response 200 {object} helpers.Response "Successfully fetched Parameter System"
// @Response 400 {object} helpers.Response "Bad request"
// @Param Authorization header string true "Authorization JWT login" default(Bearer eyJhbGciOiJ...)
// @Router /api/users/check-jwt [get]
func ClaimJwt(c *fiber.Ctx) error {
	claims := c.Locals("user").(*models.JwtUser)
	return c.Status(http.StatusOK).JSON(helpers.Response{
		Success: true,
		Data:    claims,
	})
}
