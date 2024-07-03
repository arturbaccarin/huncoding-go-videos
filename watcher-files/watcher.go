package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	yaml "gopkg.in/yaml.v2"
)

var (
	Configurations User
)

type User struct {
	Username string `json:"username" yaml:"username"`
	Lastname string `json:"lastname" yaml:"lastname"`
	Age      int8   `json:"age" yaml:"age"`
}

func ReadAndMarshalFile() {
	fmt.Println("Reading configuration")

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(data, &Configurations); err != nil {
		panic(err)
	}

	fmt.Println(Configurations)
}

func main() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("File changed, changing configuration")
					ReadAndMarshalFile()
					fmt.Printf("New Configuration: %#v\n", Configurations)
				}

				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("File created, executing proccess")

					files, err := os.ReadDir("./files")
					if err != nil {
						log.Fatal(err)
					}

					for _, file := range files {
						fmt.Println(file.Name(), file.IsDir())
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Add folder
	err = watcher.Add("./files")
	if err != nil {
		log.Fatal(err)
	}

	ReadAndMarshalFile()

	// Block main goroutine forever.
	<-make(chan struct{})
}
