package main

import (
	"context"
	"fmt"

	"github.com/henomis/lingoose/assistant"
	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)

func main(){
	
	myThread := thread.New()

	myThread.AddMessages(
		thread.NewSystemMessage().AddContent(
			thread.NewTextContent("You are a powerful personal assistant"),
		),
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("Hello! How are you today?"),
		),
	)
	
	myAssistant := assistant.New(ollama.New().WithModel("llama2").WithTemperature(0)).WithThread(myThread)

	err := myAssistant.Run(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println(myAssistant.Thread())
	
}