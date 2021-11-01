package controller

import (
	"net/http"

	"github.com/b6226442/systemrepair/entity"
	"github.com/gin-gonic/gin"
)

// POST /urgencies
func CreateUrgeny(c *gin.Context) {
	var urgency entity.Urgency
	if err := c.ShouldBindJSON(&urgency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&urgency).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": urgency})
}

// GET /urgency/:id
func GetUrgency(c *gin.Context) {
	var urgency entity.Urgency
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM urgencies WHERE id = ?", id).Scan(&urgency).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": urgency})
}

// GET /urgencies
func ListUrgenies(c *gin.Context) {
	var urgenies []entity.Urgency
	if err := entity.DB().Raw("SELECT * FROM urgencies").Scan(&urgenies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": urgenies})
}

// DELETE /urgenies/:id
func DeleteUrgency(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM urgenies WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "urgency not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /urgenies
func UpdateUrgenies(c *gin.Context) {
	var urgency entity.Urgency
	if err := c.ShouldBindJSON(&urgency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", urgency.ID).First(&urgency); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem not found"})
		return
	}

	if err := entity.DB().Save(&urgency).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": urgency})
}