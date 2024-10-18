package tools

import (
	"os"
	"os/exec"
)

func Restart() error {
	// Получаем путь к текущему исполняемому файлу
	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	// Получаем аргументы текущего процесса
	args := os.Args

	// Создаем команду для перезапуска программы
	cmd := exec.Command(execPath, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем новую копию программы
	if err := cmd.Start(); err != nil {
		return err
	}

	// Завершаем текущий процесс
	os.Exit(0)

	return nil // эта строка никогда не будет выполнена, так как процесс завершится
}
