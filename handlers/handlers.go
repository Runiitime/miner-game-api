package handlers

import (
	"miner-game-api/logic/facility"
	"miner-game-api/logic/info"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandlers struct {
	facil *facility.Facility
	info  *info.Info
}

func NewHTTPHandlers(facil *facility.Facility, info *info.Info) *HTTPHandlers {
	return &HTTPHandlers{
		facil: facil,
		info:  info,
	}
}

/* MINERS */

func HandlerGetSalaryByLevel(c *gin.Context) {}

func HandlerHireMiner(c *gin.Context) {}

func HandlerGetWorkingMiners(c *gin.Context) {}

/*
	HandlerGetMinersInfo

Получение информации о шахтерах для найма
*/
func (h *HTTPHandlers) HandlerGetMinersInfo(c *gin.Context) {
	c.JSON(http.StatusOK, h.info.MinersInfo)
}

/* EQUIPMENTS */

/*
	HandlerGetEquipmentsInfo

Получения информации об оборудовании, которое можно приобрести
*/
func (h *HTTPHandlers) HandlerGetEquipmentsInfo(c *gin.Context) {
	c.JSON(http.StatusOK, h.info.EquipmentsInfo)
}

func HandlerByiEquipment(c *gin.Context) {}

func HandlerGetEquipmentsStatus(c *gin.Context) {}

/* FACILITY */

func HandlerGetFacilityInfo(c *gin.Context) {}
func HandlerStop(c *gin.Context)            {}
