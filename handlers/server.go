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
	r.GET("/equipments/positions", s.handlers.HandlerGetEquipmentsPositions)
	r.GET("/miners", s.handlers.HandlerGetMinersInfo)
	r.GET("/miners/:level", s.handlers.HandlerGetSalaryByLevel)
	r.GET("/miners/levels", s.handlers.HandlerGetMinersLevels)

	r.POST("/facility/miner/:level", s.handlers.HandlerHireMiner)
	r.GET("/facility/miner/working", s.handlers.HandlerGetWorkingMiners)

	r.POST("/facility/equipment/:id", s.handlers.HandlerBuyEquipment)

	r.GET("/facility/balance", s.handlers.HandlerGetBalance)
	r.GET("/facility/stop", s.handlers.HandlerStop)

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
