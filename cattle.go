package main

import (
	"log"
	"os"

	dockerapi "github.com/fsouza/go-dockerclient"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	dockerHost := os.Getenv("DOCKER_HOST")
	if dockerHost == "" {
		os.Setenv("DOCKER_HOST", "unix:///tmp/docker.sock")
	}
	docker, err := dockerapi.NewClientFromEnv()
	assert(err)

	events := make(chan *dockerapi.APIEvents)
	assert(docker.AddEventListener(events))
	log.Println("Listening for Docker events ...")

	// Process Docker events
	for msg := range events {
		switch msg.Status {
		case "start":
			cont, err := docker.InspectContainer(msg.Actor.ID)
			assert(err)
			log.Println(cont.Config.Labels)
			for contPort, value := range cont.NetworkSettings.Ports {
				log.Println(contPort)
				log.Println(value)
			}
		case "die":
			// todo
			cont, err := docker.InspectContainer(msg.Actor.ID)
			assert(err)
			log.Println(cont.Config.Labels)
			for contPort, value := range cont.NetworkSettings.Ports {
				log.Println(contPort)
				log.Println(value)
			}
		}
	}
}
