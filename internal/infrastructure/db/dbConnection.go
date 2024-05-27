package db

import (
	"fmt"
	"time"

	"github.com/ashkarax/student_data_managing/internal/config"
	"github.com/ashkarax/student_data_managing/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config config.DataBase) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC() // Set NowFunc to UTC time
		},
	})
	if dberr != nil {
		return nil, dberr
	}

	// Table Creation
	if err := DB.AutoMigrate(&domain.Student{}); err != nil {
		fmt.Println("----------***********----------", err)
		return nil, err
	}

	return DB, nil
}
