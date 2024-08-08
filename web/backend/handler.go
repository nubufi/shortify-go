package backend

import (
	"database/sql"
	"fmt"
	"net/http"

	"shortify/lib"
	"shortify/web/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render renders the given component and sets the response status code.
// Parameters:
// - ctx: echo.Context
// - statusCode: int
// - t: templ.Component
// Returns: error
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

type FormData struct {
	Url string `form:"url"`
}

// ShortenHandler handles the POST request to shorten a URL.
// Parameters:
// - c: echo.Context
// - db: *sql.DB
// Returns: error
func ShortenHandler(c echo.Context, db *sql.DB) error {
	var formData FormData
	if err := c.Bind(&formData); err != nil {
		return err
	}

	hash, err := lib.NewURLShortener(db).ShortenURL(formData.Url)
	if err != nil {
		return err
	}

	shortenedURL := fmt.Sprintf("%s/%s", c.Request().Host, hash)

	return Render(c, http.StatusOK, templates.ShortenedURL(shortenedURL))
}

// RedirectHandler handles the GET request to redirect to the original URL.
// Parameters:
// - c: echo.Context
// - db: *sql.db
// Returns: error
func RedirectHandler(c echo.Context, db *sql.DB) error {
	shortCode := c.Param("shortCode")
	originalURL, err := lib.NewURLShortener(db).GetOriginalURL(shortCode)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, originalURL)
}

// HomePageHandler handles the GET request to the home page.
// Parameters:
// - c: echo.Context
// Returns: error
func HomePageHandler(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Home())
}
