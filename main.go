package main

import (
	"chris-sul/shorturl/internal/link"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "dbname=url_dev sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// /links
	link.RegisterHandlers(r, db, ctx)
	link.RegisterCatchAll(r, db, ctx)

	r.Run()
}
