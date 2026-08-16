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
	pressedMask uint32
	steps       uint32
	p2steps     uint32
	p2choose    int
	isWithBot   bool
	toothInput  int
	maxsteps    uint32
	p2maxsteps  uint32
)

func PrintTooth() {
	fmt.Print("\nЧелюсть крокодила/Crocodile's jaw:\n")
	color1 := "\033[90m" // да, я сделал цветы/yes, i made those fancy colors
	color2 := "\033[37m"
	reset := "\033[0m"
	for i := 1; i <= toothInput; i++ {
		if ((pressedMask >> uint(i)) & 1) == 1 {
			fmt.Print(color1 + "[X] " + reset)
		} else {
			fmt.Printf(color2+"[%d] "+reset, i)
		}
	}
	fmt.Println("\n-------------------------------------------------")
}

func init() {
	rand.Seed(time.Now().UnixNano()) // если у вас go 1.20(или младше, к примеру 1.26), можете удалять эту функцию/if you have go 1.20(or younger, for example 1.26), you can delete this function
}

func GenCrocoTooth(maxInput int) int {
	return rand.Intn(maxInput) + 1
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
		if p2choose < 1 || p2choose > toothInput {
			fmt.Println("Зубов всего ", toothInput, "!/Only ", toothInput, " teeth available!")
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
		tooth = GenCrocoTooth(toothInput)
		if p2steps > p2maxsteps && p2steps != 0 {
			p2maxsteps = p2steps
			time.Sleep(1 * time.Second)
			fmt.Println("УРА/YAY!!! Игрок 2, ваш новый рекорд/Player 2, your new record is: ", maxsteps)
		}
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

	go func() {
		<-sigs
		fmt.Println("\nПока, игрок! Увидимся...")
		fmt.Println("Goodbye, Player! See ya' later...")
		if maxsteps != 0 {
			fmt.Println("Ваш максимальный рекорд на сегодня/Your max record for today: ", maxsteps) // да, это подсчет рекорда за сегодняшнюю сессию/yes, this thing counts record for today' session
		} else {
			fmt.Println("Спешим вас огорчить, но вы не набрали никакого рекорда за сегодня, у вас 0/We are sorry, but you didn't made any record today, you have 0!")
		}
		if isWithBot == false && p2maxsteps != 0{
			fmt.Println("Максимальный счет второго игрока на сегодня/2th Player max record for today: ", p2maxsteps) // то-же самое для игрока 2/same for player 2
		} else if isWithBot == false {
			fmt.Println("Спешим вас огорчить, Игрок 2, но вы не набрали никакого рекорда за сегодня, у вас 0/We are sorry, Player 2, but you didn't made any record today, you have 0!")
		}
		os.Exit(0)
	}()

	for {
		fmt.Print("Вы хотите поиграть с ИИ или с человеком(ИИ/Чел)/You want to play with an AI or with a human?(AI/Human): ")
		_, err := fmt.Scan(&playChoise)
		if err != nil {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Неверный ввод/Incorrect input!")
			continue
		}
		playChs = strings.ToLower(playChoise) // да, я под-оптимизировал эту часть
		if strings.HasPrefix(playChs, "и") || strings.HasPrefix(playChs, "a") {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Окей! Теперь вы с ботом. Удачи!")
			isWithBot = true
			break
		} else if strings.HasPrefix(playChs, "ч") || strings.HasPrefix(playChs, "h") {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Окей! Теперь вы с человеком. Удачи!")
			isWithBot = false
			break
		} else {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Неверный ввод, либо-же неизвестный язык/Incorrect input, or unknown language!")
		}
	}

	for {
		fmt.Print("Сколько зубов/How much teeth? ") // ура, новый выбор!
		_, err := fmt.Scan(&toothInput)
		if err != nil {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Неверный ввод/Non-correct input!")
			continue
		} else if toothInput < 12 {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Слишком мало зубов/Too little teeth!")
			continue
		} else if toothInput > 16 {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Слишком много зубов/Too many teeth!")
			continue
		} else {
			break
		}
	}

	tooth = GenCrocoTooth(toothInput)
	var choose int
	var aisteps int

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
		if choose < 1 || choose > toothInput {
			fmt.Println("Зубов всего ", toothInput, "!/Only ", toothInput, " teeth available!")
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
					botstep = GenCrocoTooth(toothInput)
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
					tooth = GenCrocoTooth(toothInput)
					if steps > maxsteps && steps != 0 {
						maxsteps = steps
						time.Sleep(1 * time.Second)
						fmt.Println("УРА/YAY!!! Ваш новый рекорд/Your new record: ", maxsteps)
					}
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
			tooth = GenCrocoTooth(toothInput)
			if steps > maxsteps && steps != 0 {
				maxsteps = steps
				time.Sleep(1 * time.Second)
				fmt.Println("УРА/YAY!!! Ваш новый рекорд/Your new record: ", maxsteps)
			}
			steps = 0
			p2steps = 0
			aisteps = 0
			pressedMask = 0
			time.Sleep(3 * time.Second)
			fmt.Println("Новая партия/New match!")
		}
	}
}
