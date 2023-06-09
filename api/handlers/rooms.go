package handlers

import (
	"Impact/models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Get By ID Room godoc
// @ID get_by_id_room
// @Router /api/rooms/{id} [GET]
// @Summary Get By ID Room
// @Description Get By ID Room
// @Tags Rooms
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Room "Success Request"
// @Response 404 {object} models.DefaultError "Not found"
func (h *Handler) GetByIDRoom(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{Message: "topilmadi"})
		return
	}
	resp, err := h.storages.Room().GetRoom(context.Background(), idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{Message: "topilmadi"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Get List Rooms godoc
// @ID get_list_rooms
// @Router /api/rooms [GET]
// @Summary Get List Rooms
// @Description Get List Rooms
// @Tags Rooms
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param type query string false "type"
// @Param page query string false "page"
// @Param page_size query string false "page_size"
// @Success 200 {object} models.RoomsResponse "Success Request"
func (h *Handler) GetRoomsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	search := c.Query("search")
	roomType := c.Query("type")
	resp, err := h.storages.Room().GetRooms(context.Background(), models.RoomsRequest{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
		Type:     roomType,
	})
	if err != nil {
		c.JSON(http.StatusOK, models.RoomsResponse{
			Page:     page,
			PageSize: pageSize,
			Count:    0,
			Results:  []models.Room{},
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
