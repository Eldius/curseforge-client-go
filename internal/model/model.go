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
	Name        string
	ID          int64
	ClassID     int64
	Versions    []string
	Category    []ModCategory
	URL         string
	SourceURL   string
	WikiURL     string
	Description string
	Authors     []ModAuthor
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
