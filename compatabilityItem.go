package main

import (
	"text/tabwriter"
)

type compatabilityItem struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

func marshalCompatabilityItem(c *compatabilityItem) (*[]byte, error) {
	return marshal(c)
}

func tableCompatabilityItem(w *tabwriter.Writer, c *compatabilityItem) error {
	return printlnTable(w, *c)
}

func unmarshalCompatabilityItem(b *[]byte) (*compatabilityItem, error) {
	var (
		c   compatabilityItem
		err error
		ok  bool
	)
	err = unmarshal(b, &c)
	ok = (err == nil)
	if !ok {
		return nil, err
	}
	return &c, err
}
