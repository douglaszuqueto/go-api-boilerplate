package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/douglaszuqueto/go-api-boilerplate/http"
	"github.com/douglaszuqueto/go-api-boilerplate/pkg/storage"
	"github.com/douglaszuqueto/go-api-boilerplate/pkg/util"
	"github.com/douglaszuqueto/go-api-boilerplate/pkg/util/graceful"
)

var db storage.UserStorage

func main() {
	fmt.Println("API boilerplate")

	grace := graceful.New()

	storage.Open()

	insertData()

	// Start the api
	go http.Run()

	grace.Wait()
}

func insertData() {
	for i := 1; i <= 10; i++ {
		idString := strconv.Itoa(i)

		log.Println("Inserindo user:", idString)

		password, _ := util.GeneratePassword("password_" + idString)

		user := storage.User{
			Username:  "username_" + idString,
			Password:  password,
			State:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now().Add(time.Hour),
		}

		_, err := storage.GetDB().CreateUser(context.Background(), user)
		if err != nil {
			log.Println("CreateUser err", err)
		}
	}
}
