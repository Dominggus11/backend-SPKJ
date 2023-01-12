package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "host=trivy-v3-db.internal user=postgres password=71f24c20c6116c28d6d8d790767ac6d56fb7df92b4389c8a dbname=trivy-v3 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=192.168.1.116 user=roy password=programming dbname=spkj port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=localhost user=roy password=programming dbname=spkj port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := "host=dpg-cf08td9a6gdm8jvnnar0-a user=spkj_user password=A0OXnvYLO2rCRrQ2n4iB9AwSS4QS9PMJ dbname=spkj port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "postgres://spkj_user:A0OXnvYLO2rCRrQ2n4iB9AwSS4QS9PMJ@dpg-cf08td9a6gdm8jvnnar0-a.singapore-postgres.render.com/spkj"
	// dsn := "postgres://sfevjnypxdnvkx:f6c66c1d9766ecc14c0613e0861e5e6500f21cd53b52b9131da4a2fd0f32fcb9@ec2-52-207-90-231.compute-1.amazonaws.com:5432/d7buu4l2qd5voc"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connect")
	}

	database.AutoMigrate(&Students{}, &Criterias{}, &Users{})

	DB = database
}
