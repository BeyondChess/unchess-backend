package server

import (
	"github.com/markbates/goth/gothic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	// Gin uses : for parameters in routes
	r.GET("/auth/:provider/callback", s.AuthCallbackHandler)
	r.GET("/auth/:provider", s.AuthHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) AuthCallbackHandler(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) AuthHandler(c *gin.Context) {

	gothic.BeginAuthHandler(c.Writer, c.Request)
}
