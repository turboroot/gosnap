package main

import (
	"log"
	"time"
)

const (
	nameTimeSep = "_"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("could not load config, %s", err)
	}

	tarsnap := &tarsnap{config: config}

	for _, s := range tarsnap.config.BackupSets {
		if err := tarsnap.backup(s); err != nil {
			log.Printf("error backing up %s\n", s.Name)
			continue
		}

		log.Printf("backed up %s\n", s.Name)
	}

	archives, err := tarsnap.archives()
	if err != nil {
		log.Println("error getting archives")
	}

	for _, a := range archives {
		if a.age() > time.Duration(tarsnap.config.MaxAge) {
			log.Printf("%s is older than %s; deleting", a.archiveName(), tarsnap.config.MaxAge)
			tarsnap.delete(a)
		}
	}
}
