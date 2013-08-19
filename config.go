package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type maxAge time.Duration

func (m *maxAge) UnmarshalJSON(b []byte) error {
	dur, err := time.ParseDuration(string(b[1 : len(b)-1]))
	if err != nil {
		return err
	}

	*m = maxAge(dur)

	return nil
}

type Config struct {
	TarsnapLoc string `json:"tarsnapLoc"`
	MaxAge     maxAge `json:"maxAge"`
	BackupSets []*Set `json:"backupSets"`
}

type Set struct {
	Name string   `json:"name"`
	Dirs []string `json:"dirs"`
}

func LoadConfig(f string) (*Config, error) {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var config Config

	if err = json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
