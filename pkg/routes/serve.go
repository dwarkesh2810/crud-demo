package routes

import (
	"fmt"
	"log"
	"os"
)

func Serve() {
	r := GetRouter()

	err := r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))) //listen and serve on 0.0.0.0:8080 for windows "loclahost:8000"

	if err != nil {
		log.Fatal("error in routing")
		return
	}
}
