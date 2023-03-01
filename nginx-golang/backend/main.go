package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

Hello world

_####,____,####'
__"###,__,####'
___"####.####'_____,,,,,,,,,,,
____"######"____,########,,___####__,####
______####_____#####""####,__####____####
______####____,####____####_,####____####
______####____,###,____####__,####___####
______####_____'####,,####'__'####,,####
______####_______"######"______"######"





`,
	)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler)

	fmt.Println("Go backend started!")
	log.Fatal(http.ListenAndServe(":80", r))
}
