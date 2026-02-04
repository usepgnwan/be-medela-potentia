package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Generate request ID
		requestID := uuid.New().String()

		// Catat waktu mulai request
		startTime := time.Now()

		// Set request ID ke header
		c.Set("X-Request-ID", requestID)

		// Lanjutkan ke handler berikutnya
		err := c.Next()

		// Catat waktu selesai request
		endTime := time.Now()

		// Hitung processing time
		processingTime := endTime.Sub(startTime)

		// Ambil informasi request
		method := c.Method()
		path := c.Path()
		statusCode := c.Response().StatusCode()

		// Log informasi request
		log.Printf(`{"request_id": "%s", "method": "%s", "path": "%s", "status_code": %d, "processing_time": "%s"}`,
			requestID, method, path, statusCode, processingTime)

		return err
	}
}
