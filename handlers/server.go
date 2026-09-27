package handlers

import "github.com/gin-gonic/gin"

type HTTPServer struct {
	handlers *HTTPHandlers
}

func NewHTTPServer(handlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{handlers: handlers}
}

func (s *HTTPServer) Start(port string) {
	r := gin.Default()

	r.GET("/equipments", s.handlers.HandlerGetEquipmentsInfo)
	r.GET("/miners", s.handlers.HandlerGetMinersInfo)

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
