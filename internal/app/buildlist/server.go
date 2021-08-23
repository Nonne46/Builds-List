package buildlist

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/Nonne46/Builds-List/internal/app/store"
	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
	store  store.Store
}

func newServer(store store.Store, db *sql.DB) *server {
	s := &server{
		router: gin.Default(),
		store:  store,
	}

	s.configureRouter()
	s.loadStatic()

	return s
}

func (s *server) configureRouter() {
	s.router.SetFuncMap(template.FuncMap{
		"formatAsDate":  formatAsDate,
		"countComments": s.store.Comment().CountByBuildId,
	})

	s.router.GET("/", func(c *gin.Context) {
		builds := s.store.Build().GetBuilds()

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title":  "Build list",
			"total":  len(builds),
			"builds": builds,
		})
	})

	s.router.GET("/:id", func(c *gin.Context) {
		buildId, _ := strconv.Atoi(c.Param("id"))

		build, err := s.store.Build().FindById(buildId)

		if err != nil {
			c.JSON(404, gin.H{"code": "BUILD_NOT_FOUND", "message": err.Error()})
		}
		comments := s.store.Comment().FindByBuildId(build.Id)

		c.HTML(http.StatusOK, "build.tmpl.html", gin.H{
			"title":    build.Name,
			"build":    build,
			"comments": comments,
		})

	})

	s.router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.store.Build().GetBuilds())
	})

	s.router.GET("/json/:id", func(c *gin.Context) {
		buildId, _ := strconv.Atoi(c.Param("id"))

		build, err := s.store.Build().FindById(buildId)

		if err != nil {
			c.JSON(404, gin.H{"code": "BUILD_NOT_FOUND", "message": err.Error()})
		} else {
			c.JSON(http.StatusOK, build)
		}
	})

	s.router.POST("/", func(c *gin.Context) {
		query := c.PostForm("Search")
		builds := s.store.Build().FindBySearch(query)

		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title":  "Build list",
			"total":  len(builds),
			"builds": builds,
		})
	})

}

func (s *server) loadStatic() {
	s.router.LoadHTMLGlob("./static/templates/*")
	s.router.StaticFile("/style.css", "./static/css/style.css")
	s.router.StaticFile("/background.jpg", "./static/img/background.jpg")
	s.router.StaticFile("/favicon.ico", "./static/img/favicon.ico")
	s.router.StaticFile("/jqs.js", "./static/js/jquery-3.6.0.slim.min.js")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
