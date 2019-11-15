package main

import (
	"fmt"
	"strconv"
	"text/tabwriter"
	"time"

	"golang.org/x/text/language"
)

type game struct {
	CompatabilityURLs []*address   `json:"compatability_urls"`
	Complete          bool         `json:"complete"`
	Description       string       `json:"description"`
	GamePath          string       `json:"game_path"`
	GameURL           *address     `json:"game_url"`
	ID                string       `json:"id"`
	Image             *image       `json:"image"`
	IsReleased        bool         `json:"is_released"`
	Language          language.Tag `json:"language"`
	LastModified      int64        `json:"last_modified"`
	Path              string       `json:"path"`
	Name              string       `json:"name"`
	ReleaseDateMask   string       `json:"release_date_mask"`
	Timestamp         time.Time    `json:"timestamp"`
	Title             string       `json:"title"`
	Type              string       `json:"type"`
	Unix              int64        `json:"unix"`
	URI               string       `json:"uri"`
	URL               *address     `json:"url"`
}

func getGameCompatability(rawurl string) {}

func marshalGame(g *game) (*[]byte, error) {
	return marshal(g)
}

func newGame(c *compatabilityGame, i *compatabilityItem) (*game, error) {
	var (
		ok bool
	)
	ok = (c != nil) || (i != nil)
	if !ok {
		return nil, fmt.Errorf("*c, *l and *i are nil")
	}
	const (
		template string = "%s%s"
	)
	var (
		complete        bool
		description     string
		g               *game
		gamePath        string
		gameURL         *address
		ID              string
		image           *image
		isReleased      bool
		language        = language.AmericanEnglish
		lastModified    int64
		path            string
		name            string
		releaseDateMask string
		timestamp       time.Time
		title           string
		typeOf          string
		unix            int64
		URI             string
		URL             *address
	)
	complete = (c != nil) && (i != nil)
	if c != nil {
		gamePath = c.Path
		gameURL, _ = newAddress(fmt.Sprintf(template, nintendoURL, c.URL))
		ID = c.ID
		image, _ = newImage(fmt.Sprintf(template, nintendoURL, c.Image))
		isReleased, _ = strconv.ParseBool(c.IsReleased)
		name = c.Name
		path = c.Path
		releaseDateMask = c.ReleaseDateMask
		timestamp, _ = time.Parse("2006-01-02", releaseDateMask)
		timestamp = timestamp.UTC()
		typeOf = c.Type
		unix = timestamp.Unix()
	}
	if i != nil {
		description = i.Description
		lastModified = i.LastModified
		path = i.Path
		title = i.Title
		URL, _ = newAddress(fmt.Sprintf(template, nintendoURL, i.URL))
	}
	URI = normalizeURI(name)
	g = &game{
		Complete:        complete,
		Description:     description,
		GamePath:        gamePath,
		GameURL:         gameURL,
		ID:              ID,
		Image:           image,
		IsReleased:      isReleased,
		Language:        language,
		LastModified:    lastModified,
		Path:            path,
		Name:            name,
		ReleaseDateMask: releaseDateMask,
		Timestamp:       timestamp,
		Title:           title,
		Type:            typeOf,
		Unix:            unix,
		URI:             URI,
		URL:             URL}
	return g, nil
}

func readGame(fullpath string) (*game, error) {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = readFile(fullpath)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return unmarshalGame(b)
}

func tableGame(w *tabwriter.Writer, g *game) error {
	return printlnTable(w, *g)
}

func unmarshalGame(b *[]byte) (*game, error) {
	var (
		err error
		g   game
		ok  bool
	)
	err = unmarshal(b, &g)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &g, err
}

func writeGame(path, folder string, g *game) error {
	var (
		b   *[]byte
		err error
		ok  bool
	)
	b, err = marshalGame(g)
	ok = (err == nil)
	if !ok {
		return err
	}
	return writeJSON(path, folder, g.URI, b)
}
