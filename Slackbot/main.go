package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// FMT Package
func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

//OS PACKAGE 
func main(){
	// ADDING TOKEN FROM SLACK
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-7511304285524-7522256802321-a39cRP5rp2aGivo26Ag4dWE4")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A07EYMTK1DZ-7505957729749-921fd30b32edb38feaeb3c923a64ba4da96306332f78bf26e53eab80ad601763")

	//CREATE A BOT
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator", 
		
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
	 	
			//Error handling
			if err !=nil {
				println("error")
			}
			age := 2021-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// LOG PACKAGE
	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}