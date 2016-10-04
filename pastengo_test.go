package main

import (
	"testing"
	"os"
)

const TESTDB_FILE = "testdb.sqlite"
const TESTDB_TYPE = "sqlite3"
const TESTCONF_TEMPLATE = "config.example.json"
const TESTCONF_FILE = "config.json"

func TestLoadConfiguration(t *testing.T) {
	_, err := os.Stat(TESTCONF_FILE)
	if err != nil && os.IsNotExist(err) {
		os.Link(TESTCONF_TEMPLATE, TESTCONF_FILE)
	}

	LoadConfiguration()

	var emptyConfiguration Configuration
	if configuration == emptyConfiguration {
		t.Error("expected configuration to be valid after loading")
	}
}

func TestGetDB(t *testing.T) {
	os.Remove(TESTDB_FILE)
	dbType = TESTDB_TYPE
	dbString = TESTDB_FILE

	db := GetDB()
	if db == nil {
		t.Error("expected valid database")
	}
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Error("expected pingable database")
	}
}

func TestCheckDB(t *testing.T) {
	db := GetDB()
	defer db.Close()
	_, err := db.Query("SELECT 1 FROM pastebin LIMIT 1")
	if err == nil {
		t.Error("expected SELECT on non existent table to fail")
	} else if err.Error() != "no such table: pastebin" {
		t.Error("did not expect database error : ", err)
	}

	CheckDB()

	res, err := db.Query("SELECT 1 FROM pastebin LIMIT 1")
	if err != nil && err.Error() != "no such table: pastebin" {
		t.Error("expected database to be populated after CheckDB")
	} else if err != nil {
		t.Error("did not expect database error : ", err)
	} else if res.Err() != nil {
		t.Error("did not expect database row error : ", res.Err())
	}
}

func TestValidPasteId(t *testing.T) {
	paste := "abcde"
	if !ValidPasteId(paste) {
		t.Errorf("expected %s to be a valid paste ID", paste)
	}
	paste = "favicon.ico"
	if ValidPasteId(paste) {
		t.Errorf("expected %s to be an invalid paste ID", paste)
	}
	paste = "1234567890123456789012345678901"
	if ValidPasteId(paste) {
		t.Errorf("expected %s to be a invalid paste ID", paste)
	}
}

func TestGenerateName(t *testing.T) {
	name, err := GenerateName()
	if err != nil {
		t.Error("did not expect error : ", err)
	}
	if !ValidPasteId(name) {
		t.Errorf("expected GenerateName to output a valid paste ID : %s", name)
	}
	name2, err := GenerateName()
	if name == name2 {
		t.Errorf("expected %s and %s to be 2 different GenerateName results", name, name2)
	}

	db := GetDB()
	defer db.Close()
	db.Exec("INSERT INTO pastebin (id) VALUES (?)", name)
	name3, err := GenerateName()
	if name == name3 {
		t.Errorf("expected %s and %s to be 2 different GenerateName results", name, name3)
	}
}

func TestSha1(t *testing.T) {
	input := ""
	expected := "2jmj7l5rSw0yVb_vlWAYkK_YBwk="
	output := Sha1(input)
	if output != expected {
		t.Errorf("expected sha1('') == %s got %s instead", expected, output)
	}

	input = "abcde123456_$^ù*ø"
	expected = "FEpsLSpz-VoymSaM2efFIX4vV44="
	output = Sha1(input)
	if output != expected {
		t.Errorf("expected sha1(%s) == %s go %s instead", input, expected, output)
	}
}