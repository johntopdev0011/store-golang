package customEcho

import (
	"context"
	"emperror.dev/errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GetRequestID Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// GetCtxWithReqID Get ctx with timeout and request id from echo context
func GetCtxWithReqID(c echo.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// GetRequestCtx Get context  with request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// GetConfigPath Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config"
}

// GetIPAddress Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

var allowedImagesContentTypes = map[string]string{
	"image/bmp":                "bmp",
	"image/gif":                "gif",
	"image/png":                "png",
	"image/jpeg":               "jpeg",
	"image/jpg":                "jpg",
	"image/svg+xml":            "svg",
	"image/webp":               "webp",
	"image/tiff":               "tiff",
	"image/vnd.microsoft.icon": "ico",
}

func CheckImageFileContentType(fileContent []byte) (string, error) {
	contentType := http.DetectContentType(fileContent)

	extension, ok := allowedImagesContentTypes[contentType]
	if !ok {
		return "", errors.New("this content type is not allowed")
	}

	return extension, nil
}
