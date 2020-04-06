package developer_content_docs

import (
	"dragonback/lib/models/controller"
	"dragonback/lib/models/repo"
	"dragonback/lib/models/swaggerdocs"
	"github.com/labstack/echo/v4"
	"net/http"
)

var handler = controller.New("/developer/content/docs").

	// returns an index of items
	Handler(
		controller.GET,
		"",
		swaggerdocs.New().
			SetDescription("Fetches a list of markdown documentation files available").
			AddResponse(http.StatusOK, "successful", []string{}, nil),
		func () echo.HandlerFunc {
			files, _ := repo.New(repo.CDragonDocs).Documentation()

			return func(c echo.Context) error {
				var paths []string
				for _, file := range files {
					paths = append(paths, file.URL())
				}
				return c.JSON(http.StatusOK, paths)
			}
	}()).

	// returns a markdown page
	Handler(
		controller.GET,
		"/:filename",
		swaggerdocs.New().
			SetDescription("Fetches a markdown documentation file").
			SetResponseContentType("text/markdown").
			AddParamPath("", "filename", "The name of one of the available files").
			AddResponse(http.StatusOK, "successful", nil, nil),
		func () echo.HandlerFunc {
			files, _ := repo.New(repo.CDragonDocs).Documentation()

			return func(c echo.Context) error {
				for _, file := range files {
					if file.URL() == c.Param("filename") {
						c.Response().Header().Set(echo.HeaderContentType, "text/markdown")
						return c.String(http.StatusOK, file.Content())
					}
				}
				return c.NoContent(404)
			}
	}())
