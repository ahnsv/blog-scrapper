package models

import t "time"

type Post struct {
	Title   string
	Content string
	Date    t.Time
	Tags    []string
}
