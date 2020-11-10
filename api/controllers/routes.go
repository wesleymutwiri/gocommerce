package controllers

import "github.com/wesleymutwiri/gocommerce/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.GET("/", middlewares.SetMiddlewareJSON(s.Home))
	s.Router.POST("/login", middlewares.SetMiddlewareJSON(s.Login))
	s.Router.POST("/users", middlewares.SetMiddlewareJSON(s.CreateUser))
	s.Router.GET("/users", middlewares.SetMiddlewareJSON(s.GetUsers))
	s.Router.GET("/users/:id", middlewares.SetMiddlewareJSON(s.GetUser))
	s.Router.PUT("/users/:id", middlewares.SetMiddlewareAuthentication(s.UpdateUser))
	s.Router.DELETE("/users/:id", middlewares.SetMiddlewareAuthentication(s.DeleteUser))

}
