package v1

import (
	"net/http"

	"bitbucket.org/udevs/api_gateway/genproto/position_service"
	"github.com/gin-gonic/gin"
)

// Position godoc
// @ID create-position
// @Router /v1/position [POST]
// @Summary create position
// @Description create position
// @Tags position
// @Accept json
// @Produce json
// @Param position body position_service.CreatePositionRequest true "position"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionId} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreatePosition(c *gin.Context) {
	var position position_service.CreatePositionRequest

	if err := c.BindJSON(&position); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.PositionService().CreatePosition(c.Request.Context(), &position)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// Get-Position godoc
// @ID get-position
// @Router /v1/position [GET]
// @Summary get position
// @Description get position
// @Tags position
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllPosition(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.PositionService().GetAllPosition(
		c.Request.Context(),
		&position_service.GetAllPositionRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all positions", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Position godoc
// @ID update-position
// @Router /v1/position [PUT]
// @Summary update position
// @Description update position
// @Tags position
// @Accept json
// @Produce json
// @Param position body position_service.UpdatePositionRequest true "position"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionId} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdatePosition(c *gin.Context) {
	var position_request position_service.UpdatePositionRequest

	if err := c.BindJSON(&position_request); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.PositionService().UpdatePosition(c.Request.Context(), &position_request)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Position godoc
// @ID delete-position
// @Router /v1/position [DELETE]
// @Summary delete position
// @Description delete position
// @Tags position
// @Accept json
// @Produce json
// @Param position body position_service.PositionId true "position"
// @Success 200 {object} models.ResponseModel{data=position_service.PositionId} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeletePosition(c *gin.Context) {
	var position_id position_service.PositionId

	if err := c.BindJSON(&position_id); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.PositionService().DeletePosition(c.Request.Context(), &position_id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while deleting position", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Position godoc
// @ID getbyid-position
// @Router /v1/position/{id} [GET]
// @Summary getById position
// @Description GetById position
// @Tags position
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.ResponseModel{data=position_service.Position} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetByIdPosition(c *gin.Context) {

	resp, err := h.services.PositionService().GetByIdPosition(
		c.Request.Context(),
		&position_service.PositionId{
			Id: c.Param("id"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all position by given id", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}
