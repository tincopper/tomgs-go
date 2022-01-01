package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"os"
)

// import "github.com/go-pg/pg/v10/orm"

func main() {

	//open database connection first
	db := pg.Connect(&pg.Options{
		User:     "user01",
		Password: "123456",
		Database: "testdb",
		Addr:     "localhost:5432",   //database address
	})
	defer db.Close()

	//open output file
	out, err := os.Create("./bar.json")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	//execute opy
	//copy := `COPY (SELECT row_to_json(foo) FROM (SELECT * FROM userlist) foo ) TO STDOUT;`
	copy := `COPY userlist ("id", "name") TO STDOUT`

	//_, err = db.CopyTo(out, copy)

	//db.AddQueryHook(&DemoQueryHook{userId: 1})
	conn := db.Conn()
	defer conn.Close()

	_, err = conn.Exec("SET rls.userid = 1")
	if err != nil {
		panic(err)
	}
	_, err = conn.CopyTo(out, copy)
	if err != nil {
		panic(err)
	}

}

type DemoQueryHook struct {
	userId int
}

func (d DemoQueryHook) BeforeQuery(ctx context.Context, event *pg.QueryEvent) (context.Context, error) {
	_, err := event.DB.Exec(fmt.Sprintf("SET rls.userid = %d", d.userId))
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func (d DemoQueryHook) AfterQuery(ctx context.Context, event *pg.QueryEvent) error {
	return nil
}
