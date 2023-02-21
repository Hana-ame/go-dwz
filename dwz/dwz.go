package dwz

func T() {
	Init()
	db.AutoMigrate(&Tags{})

	db.Create(&Tags{"tag1", "id1"})
	db.Create(&Tags{"tag1", "id2"})
	db.Create(&Tags{"tag2", "id1"})
	db.Create(&Tags{"tag2", "id2"})
}
