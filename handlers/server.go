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

	r.GET("/info/equipments", s.handlers.HandlerGetEquipmentsInfo)
	r.GET("/info/miners", s.handlers.HandlerGetMinersInfo)
	r.GET("/info/miners/:level", s.handlers.HandlerGetSalaryByLevel)

	r.POST("/facility/hire/:level", s.handlers.HandlerHireMiner)
	r.GET("/facility/balance", s.handlers.HandlerGetBalance)

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
