package handlers

import (
	"miner-game-api/helpers"
	"miner-game-api/logic/facility"
	"miner-game-api/logic/info"
	"net/http"
	"slices"
	"strings"

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

/* INFO */

/*
	HandlerGetMinersInfo

Получение информации о шахтерах для найма
*/
func (h *HTTPHandlers) HandlerGetMinersInfo(c *gin.Context) {
	c.JSON(http.StatusOK, h.info.MinersInfo)
}

func (h *HTTPHandlers) HandlerGetSalaryByLevel(c *gin.Context) {
	level := c.Param("level")
	levels := h.info.MinerLevels()

	if !slices.Contains(levels, level) {
		c.String(http.StatusBadRequest, "Список доступных вариантов: %s", strings.Join(levels, ","))
		return
	}

	miner := h.info.MinersInfo[level]

	c.JSON(http.StatusOK, gin.H{"salary": miner.Salary})
}

func (h *HTTPHandlers) HandlerGetMinersLevels(c *gin.Context) {

}

/*
	HandlerGetEquipmentsInfo

Получения информации об оборудовании, которое можно приобрести
*/
func (h *HTTPHandlers) HandlerGetEquipmentsInfo(c *gin.Context) {
	c.JSON(http.StatusOK, h.info.EquipmentsInfo)
}

func HandlerGetWorkingMiners(c *gin.Context) {}

func (h *HTTPHandlers) HandlerBuyEquipment(c *gin.Context) {
	equipmentId := c.Param("equipmentId")
	titles := h.info.EquipmentTitles()

	if !slices.Contains(titles, equipmentId) {
		c.String(http.StatusBadRequest, "Список доступных вариантов: %s", strings.Join(titles, ","))
		return
	}

	eq := h.info.EquipmentsInfo[equipmentId]

	if eq.Price > h.facil.GetBalance() {
		c.String(http.StatusBadRequest, helpers.ErrNotEnoughBalance.Error())
		return
	}

	if err := h.facil.BuyEquipment(equipmentId); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.Status(http.StatusNoContent)
}

func HandlerGetEquipmentsStatus(c *gin.Context) {}

func HandlerGetFacilityInfo(c *gin.Context) {}
func HandlerStop(c *gin.Context)            {}

func (h *HTTPHandlers) HandlerGetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"balance": h.facil.GetBalance()})
}

func (h *HTTPHandlers) HandlerHireMiner(c *gin.Context) {
	level := c.Param("level")
	levels := h.info.MinerLevels()

	if !slices.Contains(levels, level) {
		c.String(http.StatusBadRequest, "Список доступных вариантов: %s", strings.Join(levels, ","))
		return
	}

	miner := h.info.MinersInfo[level]
	if miner.Salary > h.facil.GetBalance() {
		c.String(http.StatusBadRequest, helpers.ErrNotEnoughBalance.Error())
		return
	}

	go h.facil.HireMiner(level)

	c.Status(http.StatusOK)
}
