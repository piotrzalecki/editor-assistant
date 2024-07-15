package main

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)







func processPrompt(prompt string) (response string, err error) {

	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		return response, err
	}

	ctx := context.Background()
	response, err = llms.GenerateFromSinglePrompt(ctx, llm, "<<" + prompt + ">>")
	if err != nil {
		return response, err
	}

	return response, nil
}