package main

import (
	"fmt"
	"strings"
	"time"
)

type archive struct {
	createdAt time.Time
	name      string
}

type ArchiveNameError struct {
	name string
}

func (e *ArchiveNameError) Error() string {
	return fmt.Sprintf("%s is could not be parsed", e.name)
}

func (a *archive) decodeName(n string) error {
	dsh := strings.Index(n, "_")
	if dsh == -1 {
		return &ArchiveNameError{n}
	}

	a.name = n[:dsh]

	var err error
	a.createdAt, err = time.Parse(time.RFC3339, n[dsh+1:])
	if err != nil {
		return &ArchiveNameError{n}
	}

	return nil
}

func (a *archive) archiveName() string {
	return a.name + nameTimeSep + a.createdAt.Format(time.RFC3339)
}

func (a *archive) age() time.Duration {
	return time.Now().Sub(a.createdAt)
}
