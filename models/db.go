package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Datastore interface {
	HotelFindFirstForGateById(gate int, id uint64) (*Hotel, error)
	HotelsFind() ([]*Hotel, error)
	HotelSave(hotel *Hotel) error
	NewRecord(value interface{}) bool
}

type DB struct {
	*gorm.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	//db.LogMode(true)
	//db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
	// db.DB().SetConnMaxLifetime(time.Hour)
	//Migrate the schema
	//db.AutoMigrate(&hotelModel{})
	return &DB{db}, nil
}