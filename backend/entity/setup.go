package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("systemrepair.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema

	database.AutoMigrate(
		&Customer{},
		&Room{},
		&Equipment{},
		&Problem{},
		&Urgency{},
		&RepairInformation{},
		&CheckIn{},
		&CheckOut{},
	)

	db = database

	//Insert Customer data
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Customer{}).Create(&Customer{
		Name:     "Tanapol Pakkhotanang",
		Email:    "Tanapol@gmail.com",
		Password: string(password),
	})

	db.Model(&Customer{}).Create(&Customer{
		Name:     "Name Name",
		Email:    "name@example.com",
		Password: string(password),
	})

	db.Model(&Customer{}).Create(&Customer{
		Name:     "Test User",
		Email:    "test@yahoo.com",
		Password: string(password),
	})

	var tanapol Customer
	var name Customer
	var test Customer
	db.Raw("SELECT * FROM customers WHERE email = ?", "Tanapol@gmail.com").Scan(&tanapol)
	db.Raw("SELECT * FROM customers WHERE email = ?", "name@example.com").Scan(&name)
	db.Raw("SELECT * FROM customers WHERE email = ?", "test@yahoo.com").Scan(&test)

	// ---  Room data
	Room101 := Room{
		Roomnumber: "101",
	}
	db.Model(&Room{}).Create(&Room101)

	Room102 := Room{
		Roomnumber: "102",
	}
	db.Model(&Room{}).Create(&Room102)

	Room103 := Room{
		Roomnumber: "103",
	}
	db.Model(&Room{}).Create(&Room103)

	Room104 := Room{
		Roomnumber: "104",
	}
	db.Model(&Room{}).Create(&Room104)

	Room105 := Room{
		Roomnumber: "105",
	}
	db.Model(&Room{}).Create(&Room105)

	Room106 := Room{
		Roomnumber: "106",
	}
	db.Model(&Room{}).Create(&Room106)

	Room107 := Room{
		Roomnumber: "107",
	}
	db.Model(&Room{}).Create(&Room107)

	Room108 := Room{
		Roomnumber: "108",
	}
	db.Model(&Room{}).Create(&Room108)

	Room109 := Room{
		Roomnumber: "109",
	}
	db.Model(&Room{}).Create(&Room109)

	// ---  CheckIn data
	CheckIntana1 := CheckIn{
		Customer: tanapol,
		Room:     Room102,
	}
	db.Model(&CheckIn{}).Create(&CheckIntana1)

	CheckIntana2 := CheckIn{
		Customer: tanapol,
		Room:     Room103,
	}
	db.Model(&CheckIn{}).Create(&CheckIntana2)

	CheckIntana3 := CheckIn{
		Customer: tanapol,
		Room:     Room101,
	}
	db.Model(&CheckIn{}).Create(&CheckIntana3)

	CheckIntana4 := CheckIn{
		Customer: tanapol,
		Room:     Room101,
	}
	db.Model(&CheckIn{}).Create(&CheckIntana4)

	CheckInname1 := CheckIn{
		Customer: name,
		Room:     Room105,
	}
	db.Model(&CheckIn{}).Create(&CheckInname1)

	CheckInname2 := CheckIn{
		Customer: name,
		Room:     Room106,
	}
	db.Model(&CheckIn{}).Create(&CheckInname2)

	CheckInname3 := CheckIn{
		Customer: name,
		Room:     Room107,
	}
	db.Model(&CheckIn{}).Create(&CheckInname3)

	// checkout data
	CheckOuttana3 := CheckOut{
		CheckIn: CheckIntana3,
	}
	db.Model(&CheckOut{}).Create(&CheckOuttana3)

	CheckOutemp := CheckOut{}
	db.Model(&CheckOut{}).Create(&CheckOutemp)

	// equipment data

	equipdm := Equipment{
		Name: "กระจกสำหรับแต่งตัว (Dressing mirror)",
	}
	db.Model(&Equipment{}).Create(&equipdm)

	equipchair := Equipment{
		Name: "เก้าอี้ (Chair)",
	}
	db.Model(&Equipment{}).Create(&equipchair)

	equipwaterheater := Equipment{
		Name: "เครื่องทำน้ำอุ่น (Water heater)",
	}
	db.Model(&Equipment{}).Create(&equipwaterheater)

	equiplamp := Equipment{
		Name: "โคมไฟ (Lamp)",
	}
	db.Model(&Equipment{}).Create(&equiplamp)

	equipflushtoilet := Equipment{
		Name: "ชักโครก (flush toilet)",
	}
	db.Model(&Equipment{}).Create(&equipflushtoilet)

	equipbed := Equipment{
		Name: "เตียงนอน (Bed)",
	}
	db.Model(&Equipment{}).Create(&equipbed)

	equipfridge := Equipment{
		Name: "ตู้เย็น (Fridge)",
	}
	db.Model(&Equipment{}).Create(&equipfridge)

	equiptable := Equipment{
		Name: "โต๊ะ (Table)",
	}
	db.Model(&Equipment{}).Create(&equiptable)

	equipwardrobe := Equipment{
		Name: "ตู้เสื้อผ้า (Wardrobe)",
	}
	db.Model(&Equipment{}).Create(&equipwardrobe)

	equiptv := Equipment{
		Name: "ทีวี (TV)",
	}
	db.Model(&Equipment{}).Create(&equiptv)

	equipdoor := Equipment{
		Name: "ประตู (Door)",
	}
	db.Model(&Equipment{}).Create(&equipdoor)

	equipshower := Equipment{
		Name: "ฝักบัว (Shower)",
	}
	db.Model(&Equipment{}).Create(&equipshower)

	equipfan := Equipment{
		Name: "พัดลม (Fan)",
	}
	db.Model(&Equipment{}).Create(&equipfan)

	equipFluorescentlamp := Equipment{
		Name: "หลอดฟลูออเรสเซนต์ (Fluorescent lamp)",
	}
	db.Model(&Equipment{}).Create(&equipFluorescentlamp)

	equipac := Equipment{
		Name: "แอร์ (Air conditioner)",
	}
	db.Model(&Equipment{}).Create(&equipac)

	// problem data
	prodefective := Problem{
		Value: "ชำรุด (Defective)",
	}
	db.Model(&Problem{}).Create(&prodefective)

	pronotwork := Problem{
		Value: "ใช้งานไม่ได้ (Not working)",
	}
	db.Model(&Problem{}).Create(&pronotwork)

	// urgency data
	urgent := Urgency{
		Value: "เร่งด่วน (Urgent)",
	}
	db.Model(&Urgency{}).Create(&urgent)

	urfast := Urgency{
		Value: "เร็ว (Fast)",
	}
	db.Model(&Urgency{}).Create(&urfast)

	urmedium := Urgency{
		Value: "ปานกลาง (Medium)",
	}
	db.Model(&Urgency{}).Create(&urmedium)

	urslow := Urgency{
		Value: "ช้า (Slow)",
	}
	db.Model(&Urgency{}).Create(&urslow)

	urvslow := Urgency{
		Value: "ช้ามากๆ (Vary slow)",
	}
	db.Model(&Urgency{}).Create(&urvslow)

	// 1 repair
	db.Model(&RepairInformation{}).Create(&RepairInformation{
		CheckIn:   CheckIntana1,
		Equipment: equipfridge,
		Problem:   pronotwork,
		Urgency:   urfast,
		Datetime:  time.Now(),
	})

	// 2 repair
	db.Model(&RepairInformation{}).Create(&RepairInformation{
		CheckIn:   CheckIntana2,
		Equipment: equiptable,
		Problem:   prodefective,
		Urgency:   urmedium,
		Datetime:  time.Now(),
	})

	// 3 repair
	db.Model(&RepairInformation{}).Create(&RepairInformation{
		CheckIn:   CheckIntana3,
		Equipment: equipwardrobe,
		Problem:   prodefective,
		Urgency:   urslow,
		Datetime:  time.Now(),
	})

}
