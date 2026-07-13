package helpers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// Redirect redirects to the given url
func Redirect(ctx *gin.Context, url string) {
	ctx.Header("HX-Redirect", url)
	ctx.Status(http.StatusOK)
}

// Render renders the given component
func Render(ctx *gin.Context, component templ.Component) {
	ctx.Status(http.StatusOK)
	if err := component.Render(ctx, ctx.Writer); err != nil {
		ctx.Abort()
	}
}

// Pagination is used to paginate results.
type Pagination struct {
	// Limit is the maximum number of results to return on this page.
	Limit int
	// Offset is the number of results to skip from the beginning of the results.
	// Typically: (page number - 1) * limit.
	Offset int
}

func NewPagination(limit, offset int) *Pagination {
	return &Pagination{Limit: limit, Offset: offset}
}

// Validate returns an error if the pagination is invalid
func (p *Pagination) Validate() error {
	if p.Limit < 1 {
		return fmt.Errorf("pagination limit must be at least 1")
	}
	if p.Offset < 0 {
		return fmt.Errorf("pagination offset cannot be negative")
	}
	return nil
}
