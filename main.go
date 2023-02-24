package main

import (
	"fmt"
	"guardrail/gitscan/config"
	"guardrail/gitscan/internal/database"
	"guardrail/gitscan/internal/git"
)

func main() {
	source := "https://github.com/Yoshinaz/test_secret"
	r := git.New("test", source)
	cfg, err := config.LoadConfig()
	db, err := database.New(cfg.DB)
	if err != nil {
		panic(err)
	}
	workingCh := make(chan bool, 16)
	report, err := r.Scan(db, workingCh)

	fmt.Println(report)
}
