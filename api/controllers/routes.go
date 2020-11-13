package controllers

import "github.com/MAlifiahM/Vutura_Test_Code/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateUser))).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods("DELETE")

	//Transactions routes
	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateTransaction))).Methods("POST")
	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetTransactions))).Methods("GET")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetTransaction))).Methods("GET")
	s.Router.HandleFunc("/transcations/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateTransaction))).Methods("PUT")
	s.Router.HandleFunc("/transactions/{id}", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareAuthentication(s.DeleteTransaction))).Methods("DELETE")
}