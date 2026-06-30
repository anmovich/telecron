package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
	"strings"
	"micron/DateWork"
)

func PrintAfter(text string, second time.Duration) error{
	time.Sleep(second)
	cmd := exec.Command("notify-send",
							 "Уведомление от программы",
							 text,
)
	err := cmd.Run()
return err
}

func TimeConverter(tim int, typ int) (time.Duration){
	hui := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
	}	
	return time.Duration(tim) * hui[typ -1]
}

func timerMenu(){
	var typ int
	var tim int
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите какой тип времени использовать: ")
	fmt.Println("1. Секунды")
	fmt.Println("2. Минуты")
	fmt.Println("3. Часы")
	fmt.Scan(&typ)
	if typ > 3 || typ < 1{
		fmt.Println("Введено неверное значение")
		return
	}
	fmt.Print("Введите через сколько сотворить действо(число): ")
	fmt.Scan(&tim)
	fmt.Print("Введите что нужно вывести: ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	go func(tim int, text string, typ int){
		if err := PrintAfter(text, TimeConverter(tim, typ)); err != nil{
			fmt.Println(err)
		}
	}(tim, text, typ)

}

func menu(){
	var z int
	fmt.Println("Введите что вы хотите сделать:")
	fmt.Println("1. Поставить таймер")
	fmt.Println("2. Запланировать на определенное время")
	fmt.Println("0. Выйти из программы")
	fmt.Scan(&z)
	switch z{
	case 1: timerMenu()
case 2: datework.DateMenu() 
case 0: os.Exit(0)
	}
}

func main(){
	for{
		menu()	
	}
}
