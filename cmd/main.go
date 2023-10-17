package main

import "github.com/wolfeyc/div-tags-go/internal"

func main() {
	router := internal.BuildRouter()
	_ = router.Run()
}
