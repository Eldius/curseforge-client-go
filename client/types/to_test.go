package types

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetLatestFile(t *testing.T) {

	var res SingleModResult
	loadTestData(t, "../samples/minecraft_single_mod_response.json", &res)

	f := res.Data.GetLatestFile()

	// ID from sample file
	latestFileID := 3770428

	if latestFileID != f.ID {
		t.Logf("Wronmg file ID: %d (must be %d)", f.ID, latestFileID)
	}

}

func TestGetLatestFileGameVersions(t *testing.T) {
	var res SingleModResult
	loadTestData(t, "../samples/minecraft_single_mod_response.json", &res)

	f := res.Data.GetLatestFile()

	// ID from sample file
	latestFileID := 3770428

	if latestFileID != f.ID {
		t.Logf("Wronmg file ID: %d (must be %d)", f.ID, latestFileID)
	}

}
func loadTestData(t *testing.T, file string, res interface{}) {
	if f, err := os.Open(file); err != nil {
		t.Logf("Failed to read test data from file: %s", err)
		t.FailNow()
	} else {
		if err := json.NewDecoder(f).Decode(&res); err != nil {
			t.Logf("Failed to parse test data from file: %s", err)
			t.FailNow()
		}
	}
}
