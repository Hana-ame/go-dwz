package dwz

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	Init()
	db.AutoMigrate(&Link{})

	db.Create(&Link{Id: "1", Url: "Url1"})
	db.Create(&Link{Id: "2", Url: "Url3"})
	db.Create(&Link{Id: "3", Url: "Url2"})
	// db.Create(&Link{Id: "D42"})
	// db.Create(&Link{Id: "D42", Description: "d"})
	// db.Create(&Link{Id: "D42"})

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
	db.Create(&Tag{"tag1-", "", 44})
	db.Create(&Tag{"tag2-", "id1-", 0})
	db.Create(&Tag{"tag2-", "id2-", 0})

	// 可以使用多个键进行查找，但是查找到的东西是个谜。
	// 可能是因为这里用的是复合主键的关系。总之试验了以下
	// "", "", 44 可以找到。
	// "", "id", xx 找到的是主键2的对应
	// "tag", "id", xx 会汇报找得到找不到。
	// tg := &Tag{Tag: "", Id: "id2-", Cnt: 44}
	tg := &Tag{Tag: "tag3-", Id: "id3-", Cnt: 0}

	fmt.Println(tg.Read(db))

	fmt.Println(tg)

}
