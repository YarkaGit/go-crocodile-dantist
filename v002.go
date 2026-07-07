package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenCrocoTooth() int {
	return rand.Intn(13) + 1
}

func main() {
	var steps int
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Добро пожаловать в игру 'Крокодил Дантист'! Эта версия написана на Go.")
	fmt.Println("Welcome to the 'Crocodile Dentist' game! This version was made in Go.")
	tooth := GenCrocoTooth()
	var choose int
	var pressedMask uint16
	go func() {
		<-sigs
		fmt.Println("\nПока, игрок! Увидимся...")
		fmt.Println("Goodbye, Player! See ya' later...")
		os.Exit(0)
	}()
	for {
		fmt.Print("Выберите зуб/Choose the tooth: ")
		_, err := fmt.Scan(&choose)
		if err != nil {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Некорректный ввод/Non-correct input!")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if choose < 1 || choose > 13 {
			fmt.Println("Зубов всего 13!/Only 13 teeth available!")
			continue
		}
		if (pressedMask >> uint(choose)) & 1 == 1 {
			fmt.Println("Этот зуб уже нажат! Выберите другой./This tooth is already pressed! Choose another one.")
			continue
		}

		pressedMask |= (1 << uint(choose))
		if choose != tooth {
			fmt.Println("Вам повезло!/You got lucky!")
			steps++
			
			// --- ХОД ИИ ---
			fmt.Println("Ход ИИ.../AI is thinking...")
			time.Sleep(1 * time.Second)
			var botstep int
			for {
				botstep = GenCrocoTooth()
				if (pressedMask >> uint(botstep)) & 1 == 0 {
					break
				}
			}
			pressedMask |= (1 << uint(botstep))
			fmt.Println("Ход ИИ: ", botstep)
			
			if botstep == tooth {
				fmt.Println("ИИ проиграл! Вы победили!")
				fmt.Println("AI lost! You won!")
				fmt.Println("Робот держался... ", steps, " ходов!")
				fmt.Println("AI was holding for... ", steps, " steps!")
				tooth = GenCrocoTooth()
				steps = 0
				pressedMask = 0
				time.Sleep(3 * time.Second)
				fmt.Println("Новая партия/New match!")
				continue
			} else {
				fmt.Println("ИИ повезло!/AI got lucky!")
				steps++
			}
		} else {
			fmt.Println("Вы проиграли/Game Over!")
			fmt.Println("Вы продержались... ", steps, " ходов!")
			fmt.Println("You were holding for... ", steps, " steps!")
			tooth = GenCrocoTooth()
			steps = 0
			pressedMask = 0
			time.Sleep(3 * time.Second)
			fmt.Println("Новая партия/New match!")
		}
	}
}
