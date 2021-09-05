package seeds

import (
	"github.com/bxcodec/faker/v3"
	"github.com/ppal31/mygo/internal/types"
	"math/rand"
	"time"
)

func (s Seed) SeedDb() {
	rand.Seed(time.Now().UnixNano())
	authorIds := seedAuthors(s)
	seedResources(s, authorIds)

}

func seedAuthors(s Seed) []int64 {
	authorIds := make([]int64, 100)
	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO authors(name) VALUES (?)`)
		// execute query
		result, err := stmt.Exec(faker.Name())
		if err != nil {
			panic(err)
		}
		aid, _ := result.LastInsertId()
		authorIds[i] = aid
	}
	return authorIds
}

func seedResources(s Seed, authorIds []int64) {
	rtypes := []types.ResourceType{types.BOOK, types.BLOG, types.VIDEO}
	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO resources(name, rtype, display_name, url, author_id) VALUES (?, ?, ?, ?, ?)`)
		// execute query
		_, err := stmt.Exec(faker.Name(), rtypes[rand.Intn(3)], faker.Name(), faker.URL(), authorIds[rand.Intn(100)])
		if err != nil {
			panic(err)
		}
	}
}
