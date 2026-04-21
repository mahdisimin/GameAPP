package main

import "gameapp/router"

func main() {
	if err := router.ResloveRouter(); err != nil {
		panic(err)
	}
}
