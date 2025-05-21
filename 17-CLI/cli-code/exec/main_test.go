package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExec(t *testing.T) {

	cmd := exec.Command("./envtest")                        // тут нет парсинга аргументов
	cmd.Env = []string{"USER=gopher", "CITY=Mountain View"} // не передаём свои переменные окружения
	// result, err := cmd.CombinedOutput()
	// require.NoError(t, err)
	// fmt.Println(string(result))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start() // Run?
	require.NoError(t, err)
	err = cmd.Wait()
	require.NoError(t, err)
}

// try add os.Exit(1) into env app
