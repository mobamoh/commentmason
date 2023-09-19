package main

import (
	"fmt"

	"github.com/mobamoh/commentmason/internal/comment"
	"github.com/mobamoh/commentmason/internal/db"
	transportHttp "github.com/mobamoh/commentmason/internal/transport/http"
)

// Run is responsible for the instantiation
// and startup of our application
func Run() error {
	fmt.Println("starting up the app...")
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	if err = db.MigrateDB(); err != nil {
		return err
	}

	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err = httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("### CommentMason API ###")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
