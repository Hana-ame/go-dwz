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
	Init()
	db.AutoMigrate(&Tags{})

	db.Create(&Tags{"tag1", "id1"})
	db.Create(&Tags{"tag1", "id2"})
	db.Create(&Tags{"tag2", "id1"})
	db.Create(&Tags{"tag2", "id2"})
}
