package v1

import (
	"net/http"

	"bitbucket.org/udevs/api_gateway/genproto/position_service"
	"github.com/gin-gonic/gin"
)

// Attribute godoc
// @ID create-attribute
// @Router /v1/attribute [POST]
// @Summary create attribute
// @Description create attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body position_service.CreateAttributeRequest true "attribute"
// @Success 200 {object} models.ResponseModel{data=position_service.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateAttribute(c *gin.Context) {
	var attribute position_service.CreateAttributeRequest

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().CreateAttribute(c.Request.Context(), &attribute)
	if err != nil {
		// h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating attribute", err)
		// return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// Get-Attribute godoc
// @ID get-attribute
// @Router /v1/attribute [GET]
// @Summary get attribute
// @Description get attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=position_service.GetAllAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllAttribute(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.AttributeService().GetAllAttribute(
		c.Request.Context(),
		&position_service.GetAllAttributeRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all attributes", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Attribute godoc
// @ID update-attribute
// @Router /v1/attribute [PUT]
// @Summary update attribute
// @Description update attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param Attribute body position_service.Attribute true "attribute"
// @Success 200 {object} models.ResponseModel{data=position_service.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateAttribute(c *gin.Context) {
	var attribute position_service.Attribute

	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().UpdateAttribute(c.Request.Context(), &attribute)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating attribute", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Attribute godoc
// @ID delete-attribute
// @Router /v1/attribute [DELETE]
// @Summary delete attribute
// @Description delete attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body position_service.AttributeId true "attribute"
// @Success 200 {object} models.ResponseModel{data=position_service.AttributeId} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteAttribute(c *gin.Context) {
	var attribute_id position_service.AttributeId

	if err := c.BindJSON(&attribute_id); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AttributeService().DeleteAttribute(c.Request.Context(), &attribute_id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while deleting attribute", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Attribute godoc
// @ID getbyid-attribute
// @Router /v1/attribute/{id} [GET]
// @Summary getById attribute
// @Description GetById attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.ResponseModel{data=position_service.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetById(c *gin.Context) {

	resp, err := h.services.AttributeService().GetByIdAttribute(
		c.Request.Context(),
		&position_service.AttributeId{
			Id: c.Param("id"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all attributes", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}
