# 🐊 Crocodile Dentist (Go version)

Простая консольная реализация популярной игры "Крокодил Дантист", написанная на **Go**. Испытай свою удачу и постарайся не нажать на больной зуб!

## ✨ Особенности
*   **Многопоточность**: Используются горутины и каналы для обработки системных сигналов.
*   **Чистый выход**: Корректно завершается при нажатии `CTRL+C` с прощальным сообщением.
*   **Счетчик ходов**: Игра считает, сколько зубов ты успел нажать до проигрыша.
*   **Бесконечный цикл**: После проигрыша игра автоматически начинается заново с новым случайным зубом.

## 🚀 Как запустить

Для запуска вам понадобится установленный [Go](https://go.dev).

1. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/YarkaGit/go-crocodile-dantist
   ```
2. Перейдите в папку с проектом:
   ```bash
   cd go-crocodile-dantist
   ```
3. Запустите игру:
   ```bash
   go run main.go
   ```

## 🎮 Правила игры
Введите номер зуба от **1 до 13**. Если вам повезло — крокодил не захлопнет пасть, и вы сможете продолжить. Если попадете на больной зуб — игра окончена!

---
*Сделано с любовью (и зубами) на Go.*
EN:
# 🐊 Crocodile Dentist (Go version)

A simple console implementation of the popular game "Crocodile Dentist," written in **Go**. Try your luck and try not to press a toothache!

## ✨ Features
* **Multithreading**: Uses goroutines and channels to handle system signals.
* **Clean exit**: Cleanly exits when pressing CTRL+C with a farewell message.
* **Turn counter**: The game counts how many teeth you pressed before losing.
* **Infinite loop**: After losing, the game automatically restarts with a new random tooth.

## 🚀 How to run

You'll need [Go](https://go.dev) installed to run this game.

1. Clone the repository:
```bash
git clone https://github.com/YarkaGit/go-crocodile-dantist
```
2. Change to the project folder:
```bash
cd go-crocodile-dantist
```
3. Run the game:
```bash
go run main.go
```

## 🎮 Game Rules
Enter a tooth number from **1 to 13**. If you're lucky, the crocodile won't close its mouth, and you can continue. If you hit a bad tooth, it's game over!

---
*Made with love (and teeth) in Go.*
