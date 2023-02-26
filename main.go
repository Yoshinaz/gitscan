package main

import (
	"github.com/gitscan/http"
)

//func main() {
//	source := "https://github.com/Yoshinaz/test_secret"
//	r := repo.New("test", source)
//	cfg, err := config.LoadConfig()
//	db, err := database.New(cfg.DB)
//	if err != nil {
//		panic(err)
//	}
//	workingCh := make(chan bool, 16)
//	report, err := r.Scan(db, workingCh)
//
//	fmt.Println(report)
//}

func main() {
	http.StartServer()
}
