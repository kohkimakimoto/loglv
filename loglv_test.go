package loglv

import (
	"bytes"
	"log"
	"regexp"
	"testing"
)

func TestNormalLog(t *testing.T) {
	var out bytes.Buffer
	log.SetOutput(&out)

	log.Println("testtest")
	if ok, _ := regexp.MatchString("testtest\n$", out.String()); !ok {
		t.Errorf("Unexpected output: %s\n", out.String())
	}
}

func TestLeveledLog(t *testing.T) {
	var out bytes.Buffer

	Init()
	SetLv(LvDebug)
	SetOutput(&out)

	log.Println("testtest")
	if ok, _ := regexp.MatchString("testtest\n$", out.String()); !ok {
		t.Errorf("Unexpected output: %s\n", out.String())
	}
}

func TestLeveledLogDonotOutput(t *testing.T) {
	var out bytes.Buffer

	Init()
	SetLv(LvQuiet)
	SetOutput(&out)

	log.Println("testtest")
	if out.String() != "" {
		t.Errorf("Unexpected output: %s\n", out.String())
	}
}
