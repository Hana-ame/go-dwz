package dwz

import "math/rand"

// import "crypto/rand"

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func T() {

	Init()

}

// not tested
func GenID(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		ai := rand.Intn(len(alphabet))
		b[i] = alphabet[ai]
	}
	return string(b)
}

func GetLink(id string) (*Link, error) { // works well
	l := &Link{
		Id: id,
	}
	err := l.Read(db)
	if err != nil {
		return nil, err
	}

	// TODO：update ClickCnt

	return l, nil
}

func GetUrl(id string) (string, error) {
	l, err := GetLink(id)
	if err != nil {
		return "", err
	}
	return l.Url, nil
}

func AddUrl(url, description string) error { // works well
	// var id string
	// var err error
	id := GenID(4)
	for _, err := GetUrl(id); err == nil; {
		id = GenID(4)
		_, err = GetUrl(id)
	}

	// create a link
	link := &Link{
		Id:          id,
		Url:         url,
		Description: description,
		ClickCnt:    0,
	}
	if err := link.Create(db); err != nil {
		return err
	}

	return nil
}

// not tested
func DelUrl(id string) error {
	link := &Link{
		Id:          id,
		Url:         "",
		Description: "",
		ClickCnt:    0,
	}
	return link.Delete(db)
}

// not tested
func AddTag(tag, id string) error {
	t := &Tag{tag, id, 0}

	err := t.Read(db)
	if err == nil { // record found
		t.Cnt += 1
		t.Update(db)
	}

	err = t.Create(db)
	if err != nil {
		return err
	}

	return nil
}
