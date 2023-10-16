package main

import "github.com/meteran/gnext"

func main() {
	r := gnext.Router()
	r.GET("/all-tags-and-categories", get_all_tags_and_categories)
}
