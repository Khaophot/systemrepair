package controller

import (
	"net/http"

	"github.com/b6226442/systemrepair/entity"
	"github.com/gin-gonic/gin"
)

// POST /rooms
func CreateCheckIn(c *gin.Context) {
	
	var checkin entity.CheckIn
	if err := c.ShouldBindJSON(&checkin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

// GET /check_ins/:id
func GetCheckIn(c *gin.Context) {
	var checkin entity.CheckIn
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM check_ins WHERE id = ?", id).Scan(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkin})
}

// List /check_in/reserved/:id
func ListCheckInsReservedByCustomer(c *gin.Context) {
	var ckeckin []entity.CheckIn
	id := c.Param("id")
	if err := entity.DB().Preload("Room").
	Raw("SELECT ci.id, ci.room_id,ci.customer_id, co.check_in_id FROM check_ins ci LEFT join check_outs co On ci.id = co.check_in_id Where ci.customer_id = ? AND co.check_in_id IS NULL" , id).Find(&ckeckin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ckeckin})
}

// GET /check_in
func ListCheckIns(c *gin.Context) {
	var checkins []entity.CheckIn
	if err := entity.DB().Raw("SELECT * FROM check_ins").Scan(&checkins).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkins})
}

// Delete /check_ins/:id
func DeleteCheckin(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM check_ins WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check_in not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /check_ins
func UpdateCheckin(c *gin.Context) {
	var checkin entity.CheckIn
	if err := c.ShouldBindJSON(&checkin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", checkin.ID).First(&checkin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check_in not found"})
		return
	}

	if err := entity.DB().Save(&checkin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checkin})
}