package handlers

import (
	"miner-game-api/logic/equipments"
	"miner-game-api/logic/facility"
	"miner-game-api/logic/miner"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

type HTTPHandlers struct {
	facil          *facility.Facility
	equipmentStore *equipments.Equipments
	minersHR       *miner.Miners
}

func NewHTTPHandlers(facil *facility.Facility, eq *equipments.Equipments, mnrs *miner.Miners) *HTTPHandlers {
	return &HTTPHandlers{
		facil:          facil,
		equipmentStore: eq,
		minersHR:       mnrs,
	}
}

/* INFO */

/*
	HandlerGetMinersInfo

Получение информации о шахтерах для найма
*/
func (h *HTTPHandlers) HandlerGetMinersInfo(c *gin.Context) {
	c.JSON(http.StatusOK, h.minersHR.Miners)
}

func (h *HTTPHandlers) HandlerGetMinersLevels(c *gin.Context) {
	c.JSON(http.StatusOK, h.minersHR.Levels)
}

func (h *HTTPHandlers) HandlerGetSalaryByLevel(c *gin.Context) {
	level := c.Param("level")
	levels := h.minersHR.Levels

	if !slices.Contains(levels, level) {
		c.String(http.StatusBadRequest, "Список доступных вариантов: %s", strings.Join(levels, ","))
		return
	}

	c.JSON(http.StatusOK, gin.H{"salary": h.minersHR.SalaryByLevel(level)})
}

/*
	HandlerGetEquipmentsInfo

Получения информации об оборудовании, которое можно приобрести
*/
func (h *HTTPHandlers) HandlerGetEquipmentsInfo(c *gin.Context) {
	sold := c.Query("sold")

	if sold == "" {
		c.JSON(http.StatusOK, h.equipmentStore.Equipments)
		return
	}

	items, err := h.equipmentStore.GetEquipmentsByStatus(sold)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *HTTPHandlers) HandlerGetEquipmentsPositions(c *gin.Context) {
	c.JSON(http.StatusOK, h.equipmentStore.Titles)
}

func (h *HTTPHandlers) HandlerGetWorkingMiners(c *gin.Context) {
	c.JSON(http.StatusOK, h.facil.GetWorkingMiners())
}
func (h *HTTPHandlers) HandlerBuyEquipment(c *gin.Context) {
	id := c.Param("id")

	titles := h.equipmentStore.Titles

	if !slices.Contains(titles, id) {
		c.String(http.StatusBadRequest, "Список доступных вариантов: %s", strings.Join(titles, ","))
		return
	}

	eq, err := h.equipmentStore.BuyCheck(id, h.facil.GetBalance())
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.facil.BuyEquipment(id, eq.Price); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func HandlerGetFacilityInfo(c *gin.Context) {}

func (h *HTTPHandlers) HandlerStop(c *gin.Context) {
	h.facil.Stop()
}

func (h *HTTPHandlers) HandlerGetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"balance": h.facil.GetBalance()})
}

func (h *HTTPHandlers) HandlerHireMiner(c *gin.Context) {
	level := c.Param("level")
	levels := h.minersHR.Levels

	if !slices.Contains(levels, level) {
		c.String(http.StatusBadRequest, "Список доступных вариантов: %s", strings.Join(levels, ","))
		return
	}

	if err := h.minersHR.HireCheck(level, h.facil.GetBalance()); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	go h.facil.HireMiner(level)

	c.Status(http.StatusOK)
}
