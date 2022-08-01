package v1

import (
	"net/http"

	"bitbucket.org/udevs/api_gateway/genproto/company_service"
	"github.com/gin-gonic/gin"
)

// Company godoc
// @ID create-company
// @Router /v1/company [POST]
// @Summary create company
// @Description create company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company_service.CreateCompanyRequest true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Create(c *gin.Context) {
	var company company_service.CreateCompanyRequest

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Create(c.Request.Context(), &company)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// Get-Company godoc
// @ID get-company
// @Router /v1/company [GET]
// @Summary get company
// @Description get company
// @Tags company
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=company_service.GetAllCompanyRequest} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAll(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.CompanyService().GetAll(
		c.Request.Context(),
		&company_service.GetAllCompanyRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all professions", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Company godoc
// @ID update-company
// @Router /v1/company [PUT]
// @Summary update company
// @Description update company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company_service.Company true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Update(c *gin.Context) {
	var company company_service.Company

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Update(c.Request.Context(), &company)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Company godoc
// @ID delete-company
// @Router /v1/company [DELETE]
// @Summary delete company
// @Description delete company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company_service.Id true "company"
// @Success 200 {object} models.ResponseModel{data=company_service.Id} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Delete(c *gin.Context) {
	var id company_service.Id

	if err := c.BindJSON(&id); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CompanyService().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while deleting company", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Company godoc
// @ID getById-company
// @Router /v1/company/{id} [GET]
// @Summary getById company
// @Description GetById company
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.ResponseModel{data=company_service.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetByIdCompany(c *gin.Context) {

	resp, err := h.services.CompanyService().GetById(
		c.Request.Context(),
		&company_service.Id{
			Id: c.Param("id"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting company by given id", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}
