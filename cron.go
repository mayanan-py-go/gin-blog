package main

import (
	"fmt"
	"gin_log/models"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	log.Println("Staring...")
	c := cron.New()
	e1, err := c.AddFunc("* * * * *", func() {
		log.Println("Run models.CleanAllTag")
		models.CleanAllTag()
	})
	if err != nil {
		c.Stop()
		fmt.Println(e1, err)
	}
	e2, err := c.AddFunc("* * * * *", func() {
		log.Println("Run models.CleanAllArticle")
		models.CleanAllArticle()
	})
	if err != nil {
		c.Stop()
		fmt.Println(e2, err)
	}

	c.Start()

	timer := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second * 10)
		}
	}
}
