package dwz

import (
	"time"

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
	Id          string    `gorm:"not null;primaryKey"` // the short one for link
	Title       string    `gorm:"not null;"`           // title
	Description string    `gorm:"not null;"`           // intro
	Url         string    `gorm:"not null;"`           // true URL of the link
	ClickCnt    int       `gorm:"not null;"`           // times that the link was clicked
	CreatedAt   time.Time `gorm:"not null;"`
	Order       int       `gorm:"not null;"` // how to sort https://stackoverflow.com/questions/514943/order-a-mysql-table-by-two-columns
}

func (o *Link) Create(db *gorm.DB) (err error) {
	tx := db.Create(o)
	return tx.Error
}

func (o *Link) Read(db *gorm.DB) (err error) {
	tx := db.First(o)
	return tx.Error
}

// not tested
func (o *Link) Delete(db *gorm.DB) (err error) {
	tx := db.Delete(o)
	return tx.Error
}

// to store the tag relationships.
type Tag struct {
	Tag string `gorm:"primaryKey;not null"`
	Id  string `gorm:"primaryKey;not null;index"`
	Cnt int    `gorm:"not null;"` // times that the tag was tagged
}

func (o *Tag) Create(db *gorm.DB) (err error) {
	tx := db.Create(o)
	return tx.Error
}

func (o *Tag) Read(db *gorm.DB) (err error) {
	tx := db.First(o)
	return tx.Error
}

// not tested
func (o *Tag) Update(db *gorm.DB) (err error) {
	tx := db.Save(o)
	return tx.Error
}

// gorm create index
// https://gorm.io/docs/indexes.html
// not null
// https://stackoverflow.com/questions/43587610/preventing-null-or-empty-string-values-in-the-db

// PRAGMA table_info('tags');
// find that Tag is not primaryKey (SQLite)

// SELECT * FROM sqlite_master WHERE type = 'index';
// find indexes (SQLite)

// sqlite3.exe
// > .open c:/workplace/go-dwz/dwz.db
// download from https://sqlite.org/download.html
