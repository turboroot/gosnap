package main

import (
	"log"
	"strings"
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

		log.Printf("backed up %s (%s)\n", s.Name, strings.Join(s.Dirs, ", "))
	}

	archives, err := tarsnap.archives()
	if err != nil {
		log.Println("error getting archives")
	}

	for _, a := range archives {
		if a.age() > time.Duration(tarsnap.config.MaxAge) {
			log.Printf("deleting %s; it exceeds the max age of %s", a.archiveName(), time.Duration(tarsnap.config.MaxAge))

			if err = tarsnap.delete(a); err != nil {
				log.Printf("error deleting %s\n", a.archiveName())
			}
		}
	}
}
