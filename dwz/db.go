package dwz

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// this time, make more commmits as possible.
var db *gorm.DB

func Init() (err error) {
	db, err = gorm.Open(sqlite.Open("C:\\workplace\\go-dwz\\dwz.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

// to store the links Object.
type Link struct {
	Id          string `gorm:"primaryKey"` // the short one for link
	Description string `gorm:""`           // intro
	Url         string `gorm:""`           // true URL of the link
	ClickCnt    int    `gorm:""`           // times that the link was clicked
}

func (l *Link) Create(db *gorm.DB) (err error) {
	tx := db.Create(l)
	return tx.Error
}

func (l *Link) Read(db *gorm.DB, id string) (err error) {
	tx := db.First(l, id)
	return tx.Error
}

// to store the tag relationships.
type Tags struct {
	Tag string `gorm:"primaryKey"`
	Id  string `gorm:"primaryKey"`
}

// PRAGMA table_info('tags');
// find that Tag is not primaryKey (SQLite)

// SELECT * FROM sqlite_master WHERE type = 'index';
// find indexes (SQLite)
