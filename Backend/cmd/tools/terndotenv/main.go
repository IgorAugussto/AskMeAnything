package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migration/",
		"--config",
		"./internal/store/pgstore/migration/tern.conf",
	)

	//Comando para ver erro caso o sistema não mostre onde está errado
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar tern migrate: %v\n", err)
	}
}
