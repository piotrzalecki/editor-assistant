package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)



func main() {

	app := appConfig

	if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
	}
}


func watchInput() error {
	var lastContent string

	log.Println("Watching for user input")
	for {
		currentContent, err := clipboard.ReadAll()
		if err != nil {
			log.Println("Error reading clipboard: ", err)
			time.Sleep(1000 * time.Microsecond)
			continue
		}

		if currentContent != lastContent && currentContent != "" {
			if currentContent == "/e" {
				log.Println("Exit command submitted. Exiting")
				break
			}
			log.Println("Processing pasted text: ", currentContent)
			lastContent = currentContent
			response, err := processPrompt(currentContent)
			
			if err != nil {
				log.Println("Error prompting model: ", err)
				time.Sleep(1000 * time.Microsecond)
				continue
			}

			response = strings.TrimRight(response, "\n\r")
			log.Println("Writing model response to clipboard: ", response)
			err = clipboard.WriteAll(response)
			if err != nil {
				log.Println("Error writing response to clipboard.")
				time.Sleep(1000 * time.Microsecond)
				continue
			}
			lastContent = response
		}

		time.Sleep(1000 * time.Microsecond)
	}
	return nil
}
