package middleware

import (
	"FeedbackAPI/auth"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
)

// https://withcodeexample.com/secure-authentication-authorization-golang-fiber-guide/
// https://withcodeexample.com/routing-middleware-fiber-scalable-web-apps/

func AuthMiddleware(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "authorization header missing",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid authorization header format, expected 'Bearer <token>'",
		})
	}

	tokenString := parts[1]
	claims, err := auth.ValidateJWT(tokenString)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idParam := ctx.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID in request",
		})
	}

	if claims.UserID != uint(id) {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "you are not authorized to access this resource",
		})
	}

	ctx.Locals("claims", claims)
	return ctx.Next()
}
