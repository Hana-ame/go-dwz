package dwz

import (
	"fmt"
	"testing"
)

func TestGetUrl(t *testing.T) {
	Init()
	r, err := GetUrl("13")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

func TestAddUrl(t *testing.T) {
	err := AddUrl("title", "https://example.com", "description")
	t.Error(err)
}

func TestReadLinksByTag(t *testing.T) {
	fmt.Println("=======================")
	Init()

	tag := "public"

	var rst []*Link
	tx := db.Model(&Link{}).Where("tags.tag = ?", tag).Joins("INNER JOIN `tags` on links.id = tags.id").Find(&rst)
	// db.Model(&Link{}).Joins("JOIN `tags` ON tags.id = links.id").Where("tags.tag=?", tag).Find(&rst)
	// tx := db.Raw("select * from links inner join tags on tags.id=links.id where tags.tag=?;", tag).Find(&rst)

	fmt.Println(rst[0], rst[1], tx)
	// var artists []Artist
	// if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
	// 	Joins("JOIN movies on movies.id=artist_movies.movie_id").
	// 	Joins("JOIN languages on movies.language_id=languages.id").
	// 	Where("languages.name=?", "english").
	// 	Group("artists.id").Preload("Movies").Find(&artists).Error; err != nil {
	// 	log.Fatal(err)
	// }

	// var artists []Artist
	// if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
	// 	Joins("JOIN movies on artist_movies.movie_id=movies.id").Where("movies.title=?", "Nayagan").
	// 	Group("artists.id").Find(&artists).Error; err != nil {
	// 	log.Fatal(err)
	// }

	// for _, ar := range artists {
	// 	fmt.Println(ar.Name)
	// }

	/* output
	   Kamal Hassan
	*/

}
