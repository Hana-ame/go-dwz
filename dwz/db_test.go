package dwz

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	Init()
	db.AutoMigrate(&Link{})

	db.Create(&Link{Id: "D42"})
	db.Create(&Link{Id: "D42"})
	db.Create(&Link{Id: "D42", Description: "d"})
	db.Create(&Link{Id: "D42"})

	var l Link
	db.First(&l)

	fmt.Print(l)
}

func TestTags(t *testing.T) {
	fmt.Println("Testing")
	Init()
	db.AutoMigrate(&Tag{})

	db.Create(&Tag{"tag1-", "id1-", 0})
	db.Create(&Tag{"tag1-", "id2-", 0})
	db.Create(&Tag{"tag2-", "id1-", 0})
	db.Create(&Tag{"tag2-", "id2-", 0})

	tg := &Tag{Tag: "tag1-", Id: "id2-", Cnt: 5}
	tg.Read(db, "")

	fmt.Println(tg)

}
