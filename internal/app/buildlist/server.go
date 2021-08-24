package buildlist

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/Nonne46/Builds-List/internal/app/model"
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
		"toHTML":        toHTML,
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
			c.Abort()
			return
		}

		comments := s.store.Comment().FindByBuildId(build.Id)

		c.HTML(http.StatusOK, "build.tmpl.html", gin.H{
			"title":    build.Name,
			"build":    build,
			"comments": comments,
		})

	})

	s.router.GET("/registration", func(c *gin.Context) {

		c.HTML(http.StatusOK, "register.tmpl.html", gin.H{})
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

	s.router.POST("/addComment", func(c *gin.Context) {

		pageId, err := strconv.Atoi(c.PostForm("pageId"))
		if err != nil {
			c.Abort()
			return
		}
		username := c.PostForm("username")
		commentText := c.PostForm("commentText")

		comment := &model.Comment{
			IdPage:   pageId,
			Username: username,
			Comment:  commentText,
			Time:     time.Now(),
		}

		s.store.Comment().AddComment(comment)

		c.Redirect(http.StatusFound, c.PostForm("pageId"))
	})

	s.router.POST("/addUser", func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")

		u := &model.User{
			Name:     username,
			Email:    email,
			Password: password,
		}

		s.store.User().CreateUser(u)

		c.Redirect(http.StatusFound, "/")
	})

	s.router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.store.Build().GetBuilds())
	})

	s.router.GET("/json/:id", func(c *gin.Context) {
		buildId, _ := strconv.Atoi(c.Param("id"))

		build, err := s.store.Build().FindById(buildId)

		if err != nil {
			c.JSON(404, gin.H{"code": "BUILD_NOT_FOUND", "message": err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, build)

	})

}

func (s *server) loadStatic() {
	s.router.LoadHTMLGlob("./static/templates/*")
	s.router.StaticFile("/favicon.ico", "./static/img/favicon.ico")
	s.router.Static("/assets", "./static")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func toHTML(data string) template.HTML {
	return template.HTML(data)
}
