package controller

import (
	"net/http"

	"github.com/b6226442/systemrepair/entity"
	"github.com/gin-gonic/gin"
)

func CreateRepairInformation(c *gin.Context) {

	var repairinformation entity.RepairInformation

	var equipment entity.Equipment
	var problem entity.Problem
	var urgency entity.Urgency
	var checkin entity.CheckIn

	// ผลลัพธ์ที่ได้จากขะ้น ตอนที่ 10 จะถูก bind เข้าตัวแปร repairinformation
	if err := c.ShouldBindJSON(&repairinformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา CheckIn ด้วย ID
	if tx := entity.DB().Where("id = ?", repairinformation.CheckInID).First(&checkin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkin not found"})
		return
	}

	// 12: ค้นหา Equipment ด้วย ID
	if tx := entity.DB().Where("id = ?", repairinformation.EquipmentID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	// 13: ค้นหา Problem ด้วย ID
	if tx := entity.DB().Where("id = ?", repairinformation.ProblemID).First(&problem); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playlist not found"})
		return
	}

	// 14: ค้นหา Urgency ด้วย ID
	if tx := entity.DB().Where("id = ?", repairinformation.UrgencyID).First(&urgency); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "urgency not found"})
		return
	}

	// 15: สร้าง RepaitInformtion
	ri := entity.RepairInformation{
		CheckIn:   checkin,                    // โยงความสัมพันธ์กับ Entity checkin
		Equipment: equipment,                  // โยงความสัมพันธ์กับ Entity Equipment
		Problem:   problem,                    // โยงความสัมพันธ์กับ Entity Problem
		Urgency:   urgency,                    // โยงความสัมพันธ์กับ Entity Urgency
		Datetime:  repairinformation.Datetime, // ตั้งค่าฟิลด์ Datetime
	}

	// 16: บันทึก
	if err := entity.DB().Create(&ri).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ri})
}

// GET /repair_information/:id
func GetRepairInformation(c *gin.Context) {
	var repairinformation entity.RepairInformation
	id := c.Param("id")
	if err := entity.DB().Preload("CheckIn").Preload("Equipment").Preload("Problem").Preload("Urgency").Raw("SELECT * FROM repair_informations WHERE id = ?", id).Find(&repairinformation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": repairinformation})
}

// List /repair_informations
func ListRepairInformations(c *gin.Context) {
	var repairinformations []entity.RepairInformation
	if err := entity.DB().Preload("CheckIn").Preload("CheckIn.Customer").Preload("CheckIn.Room").Preload("Equipment").Preload("Problem").Preload("Urgency").Raw("SELECT * FROM repair_informations").Find(&repairinformations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": repairinformations})
}

// DELETE /repair_informations/:id
func DeleteRepairInformation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM repair_informations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repairinformation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /repair_informations
func UpdateRepairInformation(c *gin.Context) {
	var repairinformation entity.RepairInformation
	if err := c.ShouldBindJSON(&repairinformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", repairinformation.ID).First(&repairinformation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repairinformation not found"})
		return
	}

	if err := entity.DB().Save(&repairinformation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": repairinformation})
}
