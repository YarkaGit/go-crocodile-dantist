package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	tooth       int
	pressedMask uint16
	steps       int
	p2steps     int
	p2choose    int
	isWithBot   bool
)

func PrintTooth() {
	fmt.Print("\nЧелюсть крокодила/Crocodile's jaw:\n")
	for i := 1; i <= 13; i++ {
		if ((pressedMask >> uint(i)) & 1) == 1 {
			fmt.Print("[X] ")
		} else {
			fmt.Printf("[%d] ", i)
		}
	}
	fmt.Println("\n-------------------------------------------------")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenCrocoTooth() int {
	return rand.Intn(13) + 1
}

func p2Play() {
	for {
		PrintTooth()
		fmt.Print("[ИГРОК 2]: Выберите зуб/Choose the tooth: ")
		_, err := fmt.Scan(&p2choose)
		if err != nil {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("[ИГРОК 2/PLAYER 2]: Некорректный ввод/Non-correct input!")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if p2choose < 1 || p2choose > 13 {
			fmt.Println("[ИГРОК 2/PLAYER 2]: Зубов всего 13!/Only 13 teeth available!")
			continue
		}
		if (pressedMask>>uint(p2choose))&1 == 1 {
			fmt.Println("[ИГРОК 2/PLAYER 2]: Этот зуб уже нажат! Выберите другой./This tooth is already pressed! Choose another one.")
			continue
		}
		break
	}
	pressedMask |= (1 << uint(p2choose))

	if p2choose != tooth {
		fmt.Println("[ИГРОК 2/PLAYER 2]: Вам повезло!/You got lucky!")
		p2steps++
	} else {
		fmt.Println("Игрок 2, вы проиграли/Player 2, it's Game Over!")
		fmt.Println("Вы продержались... ", p2steps, " ходов!")
		fmt.Println("You were holding for... ", p2steps, " steps!")
		tooth = GenCrocoTooth()
		steps = 0
		p2steps = 0
		pressedMask = 0
		time.Sleep(3 * time.Second)
		fmt.Println("Новая партия/New match!")
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	var playChoise string
	var playChs string

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Добро пожаловать в игру 'Крокодил Дантист'! Эта версия написана на Go.")
	fmt.Println("Welcome to the 'Crocodile Dentist' game! This version was made in Go.")

	for {
		fmt.Print("Вы хотите поиграть с ИИ или с другим человеком(ИИ/Чел)/You want to play with an AI or with a human?(AI/Human): ")
		_, err := fmt.Scan(&playChoise)
		if err != nil {
			fmt.Println("Неверный ввод/Incorrect input!")
			continue
		}
		playChs = strings.ToLower(playChoise)
		if playChs == "ии" || playChs == "ai" {
			fmt.Println("Окей! Теперь вы с ботом. Удачи!")
			isWithBot = true
			break
		} else if playChs == "чел" || playChs == "человек" || playChs == "human" {
			fmt.Println("Окей! Теперь вы с человеком. Удачи!")
			isWithBot = false
			break
		} else {
			fmt.Println("Неверный ввод, либо-же неизвестный язык/Incorrect input, or unknown language!")
		}
	}

	tooth = GenCrocoTooth()
	var choose int
	var aisteps int

	go func() {
		<-sigs
		fmt.Println("\nПока, игрок! Увидимся...")
		fmt.Println("Goodbye, Player! See ya' later...")
		os.Exit(0)
	}()

	for {
		PrintTooth()
		fmt.Print("[ИГРОК 1]: Выберите зуб/Choose the tooth: ")
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
		if (pressedMask>>uint(choose))&1 == 1 {
			fmt.Println("Этот зуб уже нажат! Выберите другой./This tooth is already pressed! Choose another one.")
			continue
		}

		pressedMask |= (1 << uint(choose))

		if choose != tooth {
			fmt.Println("[ИГРОК 1]: Вам повезло!/You got lucky!")
			steps++

			if isWithBot {
				fmt.Println("Ход ИИ.../AI is thinking...")
				time.Sleep(1 * time.Second)
				var botstep int
				for {
					botstep = GenCrocoTooth()
					if (pressedMask>>uint(botstep))&1 == 0 {
						break
					}
					time.Sleep(10 * time.Millisecond) // сделал так что-бы проц не страдал/i made this way to stop the CPU from torturing itself
				}
				pressedMask |= (1 << uint(botstep))
				fmt.Println("Ход ИИ: ", botstep)

				if botstep == tooth {
					fmt.Println("ИИ проиграл! Вы победили!")
					fmt.Println("AI lost! You won!")
					fmt.Println("Робот держался... ", aisteps, " ходов!")
					fmt.Println("AI was holding for... ", aisteps, " steps!")
					tooth = GenCrocoTooth()
					steps = 0
					aisteps = 0
					pressedMask = 0
					time.Sleep(3 * time.Second)
					fmt.Println("Новая партия/New match!")
				} else {
					fmt.Println("ИИ повезло!/AI got lucky!")
					aisteps++
				}
			} else {
				p2Play()
			}
		} else {
			fmt.Println("Вы проиграли/Game Over!")
			fmt.Println("Вы продержались... ", steps, " ходов!")
			fmt.Println("You were holding for... ", steps, " steps!")
			tooth = GenCrocoTooth()
			steps = 0
			p2steps = 0
			aisteps = 0
			pressedMask = 0
			time.Sleep(3 * time.Second)
			fmt.Println("Новая партия/New match!")
		}
	}
}
