package model

import "github.com/eldius/curseforge-client-go/client/types"

type ModCategory struct {
	Name    string
	ID      int64
	ClassID int
}

type ModAuthor struct {
	Name string
	ID   int64
	URL  string
}

type Mod struct {
	Name        string        `db:"name" json:"name"`
	ID          int64         `db:"id" json:"id"`
	ClassID     int64         `db:"class_id" json:"class_id"`
	GameID      int64         `db:"game_id" json:"game_id"`
	Versions    []string      `db:"versions" json:"versions"`
	Category    []ModCategory `db:"category" json:"category"`
	URL         string        `db:"url" json:"url"`
	SourceURL   string        `db:"source_url" json:"source_url"`
	WikiURL     string        `db:"wiki_url" json:"wiki_url"`
	Description string        `db:"description" json:"description"`
	Authors     []ModAuthor   `db:"authors" json:"authors"`
}

func NewMod(md types.ModData) Mod {
	m := new(Mod)
	for _, lf := range md.LatestFiles {
		m.Versions = append(m.Versions, lf.GameVersions...)
	}
	for _, c := range md.Categories {
		m.Category = append(m.Category, ModCategory{
			Name:    c.Name,
			ID:      c.ID,
			ClassID: c.ClassID,
		})
	}
	m.Name = md.Name
	m.URL = md.Links.WebsiteURL
	m.SourceURL = md.Links.SourceURL
	m.WikiURL = md.Links.WikiURL
	m.GameID = md.GameID

	m.ID = md.ID
	m.Description = md.Summary

	for _, a := range md.Authors {
		m.Authors = append(m.Authors, ModAuthor{
			Name: a.Name,
			ID:   a.ID,
			URL:  a.URL,
		})
	}

	m.ClassID = md.ClassID

	return *m
}
