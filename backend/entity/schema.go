package entity

import (
	"time"

	"gorm.io/gorm"
)

type CheckIn struct {
	gorm.Model
	// FK จาก Room
	RoomID *uint
	Room   Room `gorm:"references:id"`
	// FK จาก Customer
	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	RepairInformations []RepairInformation `gorm:"foreignKey:CheckInID"`
}

type CheckOut struct {
	gorm.Model

	CheckInID *uint   `gorm:"uniqueIndex"`
	CheckIn   CheckIn `gorm:"references:id"`
}

type Customer struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string

	CheckIns []CheckIn `gorm:"foreignKey:CustomerID"`
}

type Room struct {
	gorm.Model
	Roomnumber string
	Records    []CheckIn `gorm:"foreignKey:RoomID"`
}

type Equipment struct {
	gorm.Model
	Name               string
	RepairInformations []RepairInformation `gorm:"foreignKey:EquipmentID"`
}

type Problem struct {
	gorm.Model
	Value              string
	RepairInformations []RepairInformation `gorm:"foreignKey:ProblemID"`
}

type Urgency struct {
	gorm.Model
	Value              string
	RepairInformations []RepairInformation `gorm:"foreignKey:UrgencyID"`
}

type RepairInformation struct {
	gorm.Model
	Datetime time.Time
	// FK จาก CkeckIn
	CheckInID *uint
	CheckIn   CheckIn `gorm:"references:id"`
	// FK จาก Equipment
	EquipmentID *uint
	Equipment   Equipment
	// FK จาก Problem
	ProblemID *uint
	Problem   Problem
	// FK จาก Urgency
	UrgencyID *uint
	Urgency   Urgency
}
