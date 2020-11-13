package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/MAlifiahM/Vutura_Test_Code/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var transactions = []models.Transaction{
	models.Transaction{
		Product: "Sampo",
		Price: 10000,
	},
	models.Transaction{
		Product: "Minuman",
		Price: 10000,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Transaction{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Transaction{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Transaction{}).AddForeignKey("id_user", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		transactions[i].IDUser = users[i].ID

		err = db.Debug().Model(&models.Transaction{}).Create(&transactions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed transactions table: %v", err)
		}
	}
}