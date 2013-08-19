package main

import (
	"testing"
	"time"
)

func TestDecodeName(t *testing.T) {
	bu := &archive{}
	bu.decodeName("derp_2013-08-19T03:18:25Z")

	if bu.name != "derp" {
		t.Error("name supposed to be derp")
	}

	tme, _ := time.Parse(time.RFC3339, "2013-08-19T03:18:25Z")

	if !bu.createdAt.Equal(tme) {
		t.Error("times are supposed to be the same")
	}
}

func TestArchiveName(t *testing.T) {
	bu := &archive{}
	bu.decodeName("derp_2013-08-19T03:18:25Z")

	if name := bu.archiveName(); name != "derp_2013-08-19T03:18:25Z" {
		t.Error("name supposed to match")
	}
}
