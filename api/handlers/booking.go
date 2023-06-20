package handlers

import (
	"Impact/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Create Booking godoc
// @ID create_booking
// @Router /api/rooms/{id}/book [POST]
// @Summary Create Booking
// @Description Create Booking
// @Tags Bookings
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param booking body models.BookingRequest true "CreateBookingRequest"
// @Success 201 {object} models.BookingResponse "Success Request"
// @Failure 410 {object} models.DefaultError
func (h *Handler) BookingRoom(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusGone, models.DefaultError{Message: "uzr, siz tanlagan vaqtda xona band"})
		return
	}
	var req models.BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusGone, models.DefaultError{Message: "uzr, siz tanlagan vaqtda xona band"})
		return
	}
	if req.Start != req.End {

		resp, err := h.storages.Booking().BookRoom(c.Request.Context(), idInt, req)
		if err != nil {
			c.JSON(http.StatusGone, models.DefaultError{Message: "uzr, siz tanlagan vaqtda xona band"})
			return
		}

		c.JSON(http.StatusCreated, resp)
	} else {
		c.JSON(http.StatusGone, models.DefaultError{Message: "uzr, siz tanlagan vaqtda xona band"})
		return
	}
}
