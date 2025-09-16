package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList []Bin

func main() {
	createBin()
}

func createBin() (BinList, error) {
	var name string
	var private bool
	var privateInput string
	var binList []Bin

	fmt.Println("Введите название файла")
	fmt.Scanln(&name)
	if name == "" {
		return nil, errors.New("Переданно пустое имя")
	}
	fmt.Println("Вы ходите загрузить закрытый файл? (да/нет)")
	fmt.Scanln(&privateInput)
	privateInput = fmt.Sprintf(strings.ToLower(privateInput))
	if privateInput == "да" {
		private = true
	} else if privateInput == "нет" {
		private = false
	} else {
		fmt.Println("Неверное значение параметра приватности")
		return nil, errors.New("invalid private input")
	}

	newBin := Bin{
		id:        name,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}

	binList = append(binList, newBin)

	fmt.Println(binList)

	return binList, nil
}
