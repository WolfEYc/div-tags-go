package internal

import "github.com/meteran/gnext"

func BuildRouter() *gnext.RootRouter {
	r := gnext.Router()
	r.GET("/all-tags-and-categories", getAllTagsAndCategories)

	return r
}
