package handler

import (
	"net/http"

	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"github.com/adityarizkyramadhan/tes_intern_delos/middleware"
	"github.com/adityarizkyramadhan/tes_intern_delos/utils/response"
	"github.com/gin-gonic/gin"
)

type FarmHandler struct {
	FarmUseCase domain.FarmUseCase
	FarmHandler domain.FarmHandler
}

func NewHandlerFarm(useCase domain.FarmUseCase) *FarmHandler {
	return &FarmHandler{
		FarmUseCase: useCase,
	}
}

func (h FarmHandler) Register(c *gin.Context) {
	var input domain.FarmInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Bind JSON fail", err))
		return
	}
	id, err := h.FarmUseCase.Create(&domain.FarmModel{
		Username:        input.Username,
		Contact:         input.Contact,
		Location:        input.Location,
		Leader:          input.Leader,
		Supervisor:      input.Supervisor,
		Name:            input.Name,
		NumberEmployees: input.NumberEmployees,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Create data fail", err))
		return
	}
	token, err := middleware.GenerateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Create token fail", err))
		return
	}
	c.JSON(http.StatusCreated, response.ResponseWhenSuccess("Create data success", gin.H{
		"token": token,
	}))
}

func (h FarmHandler) Update(c *gin.Context) {
	var input domain.FarmInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Bind JSON fail", err))
		return
	}
	idUser := c.MustGet("login").(uint)
	if err := h.FarmUseCase.Update(&domain.FarmModel{
		Username:        input.Username,
		Contact:         input.Contact,
		Location:        input.Location,
		Leader:          input.Leader,
		Supervisor:      input.Supervisor,
		Name:            input.Name,
		NumberEmployees: input.NumberEmployees,
	}, idUser); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Update data fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Update data success", nil))
}

func (h FarmHandler) Delete(c *gin.Context) {
	idUser := c.MustGet("login").(uint)
	if err := h.FarmUseCase.Delete(idUser); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Delete data fail", nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Delete data success", nil))
}

func (h FarmHandler) Read(c *gin.Context) {
	idUser := c.MustGet("login").(uint)
	data, err := h.FarmUseCase.Read(idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Read data fail", nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess("Delete data success", data))
}
