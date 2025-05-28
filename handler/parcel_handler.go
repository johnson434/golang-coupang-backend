package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang-coupang-backend.com/m/dto"
	"golang-coupang-backend.com/m/service"
)

type ParcelHandler struct {
	svc *service.ParcelService
}

func NewParcelHandler(s *service.ParcelService) *ParcelHandler {
	return &ParcelHandler{s}
}

func (h *ParcelHandler) RegisterParcelRoutes(r *gin.Engine) {
	log.Print("RegisterParcelRoutes")
	grp := r.Group("/parcels")
	{
		grp.GET("", h.getAllParcels)     // 모든 Parcel 조회
		grp.GET("/:id", h.getParcelByID) // 특정 Parcel 조회
		grp.POST("", h.createParcel)     // Parcel 생성
		// grp.PUT("/:id", h.updateParcel)    // Parcel 수정
		grp.DELETE("/:id", h.deleteParcel) // Parcel 삭제
	}
}

func (h *ParcelHandler) getParcelByID(c *gin.Context) {
	log.Println("getParcelByID")
	idParam := c.Param("id")
	id, validation_err := strconv.Atoi(idParam)
	if validation_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validation_err.Error()})
		return
	}

	p, err := h.svc.GetParcelByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *ParcelHandler) getAllParcels(c *gin.Context) {
	log.Println("getAllParcels")
	parcels, err := h.svc.GetAllParcels(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parcels)
}

func (h *ParcelHandler) createParcel(c *gin.Context) {
	log.Println("createParcel")
	var p dto.CreateParcelRequest
	if err := c.ShouldBindBodyWithJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 요청 받은 데이터를 모델로 변환
	pModel := p.ToModel()
	log.Printf("model.CreatedAt: %d", pModel.CreatedAt)

	// Parcel 생성
	err := h.svc.CreateParcel(c.Request.Context(), pModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// 성공적으로 생성된 경우
	c.JSON(http.StatusCreated, gin.H{"message": "Parcel created successfully"})
}

func (h *ParcelHandler) deleteParcel(c *gin.Context) {
	log.Println("deleteParcel")
	idParam := c.Param("id")
	id, validation_err := strconv.Atoi(idParam)
	if validation_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validation_err.Error()})
		return
	}

	// Parcel 삭제
	err := h.svc.DeleteParcel(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// 성공적으로 삭제된 경우
	c.JSON(http.StatusOK, gin.H{"message": "Parcel deleted successfully"})
}
