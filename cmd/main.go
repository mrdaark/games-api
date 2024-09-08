package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/mrdaark/games-api/app/gamesapi"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/gamesapi.toml", "config file")
}

func main()  {
	config:= gamesapi.NewConfig()

	_, err:= toml.DecodeFile(configPath, config)

	if (err != nil) {
		log.Fatal(err)
	}

	s:= gamesapi.New(config)

	if err:= gamesapi.Start(s); err != nil {
		log.Fatal(err)
	}


	// http.HandleFunc("/healthcheck", healthcheckHandler)
	// http.ListenAndServe("localhost:3000", nil)
}

// func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
// 	if (r.Method != http.MethodGet) {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}	

// 	fmt.Fprintf(w, "ok");
// }
