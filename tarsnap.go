package main

import (
	"bufio"
	"os/exec"
	"strings"
	"time"
)

type tarsnap struct {
	config *Config
}

func (t *tarsnap) backup(s *Set) error {
	ca := time.Now().UTC().Format(time.RFC3339)

	// archive is named archiveName_currentTime
	name := strings.Join([]string{s.Name, nameTimeSep, ca}, "")

	ops := append([]string{"-cf", name}, s.Dirs...)
	if err := exec.Command(t.config.TarsnapLoc, ops...).Run(); err != nil {
		return err
	}

	return nil
}

func (t *tarsnap) archives() ([]*archive, error) {
	cmd := exec.Command(t.config.TarsnapLoc, "--list-archives")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	archives := []*archive{}

	scn := bufio.NewScanner(stdout)
	for scn.Scan() {
		arc := &archive{}
		if err := arc.decodeName(scn.Text()); err != nil {
			continue
		}

		archives = append(archives, arc)
	}

	return archives, nil
}

func (t *tarsnap) delete(a *archive) error {
	cmd := exec.Command(t.config.TarsnapLoc, "-df", a.archiveName())
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
