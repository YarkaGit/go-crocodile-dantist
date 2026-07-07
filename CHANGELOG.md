# Чейнджлог / Changelog — go-crocodile-dantist

---

## [0.03] — Мультиплеер, Отрисовка UI и Оптимизация / Multiplayer, UI Rendering, and Optimization

### Added (Добавлено)
*   **Режим мультиплеера:** Добавлена полноценная игра для двух человек на одном ПК (Игрок 1 против Игрока 2) с выбором режима на старте (`ИИ/Чел`). / **Multiplayer Mode:** Added full support for two players on a single PC (Player 1 vs Player 2) with a mode selector (`AI/Human`) at launch.
*   **Консольный интерфейс (TUI):** Создана функция `PrintTooth()`, которая динамически отрисовывает состояние челюсти крокодила перед каждым ходом (выбитые зубы помечаются как `[X]`). / **Console UI (TUI):** Created the `PrintTooth()` function, which dynamically renders the state of the crocodile's jaw before each turn (pressed teeth are marked as `[X]`).

### Changed (Изменено)
*   **Глобальное состояние:** Архитектура переведена на пакетные переменные (глобальный scope) для удобного взаимодействия между основным циклом и функцией `p2Play()`. / **Global State:** Refactored the architecture to use package-level variables (global scope) for seamless interaction between the main game loop and the `p2Play()` function.

### Fixed (Исправлено)
*   **Разделение счетчиков ходов:** Исправлена критическая ошибка версии 0.02, где ходы игрока и ИИ суммировались в одну переменную. Теперь статистика разделена на три независимых счетчика (`steps`, `p2steps`, `aisteps`), что обеспечивает честный подсчет очков. / **Separated Turn Counters:** Fixed a critical bug from v0.02 where player and AI turns were mixed into a single variable. Stats are now split into three independent counters (`steps`, `p2steps`, `aisteps`), ensuring accurate score tracking.
*   **Защита от CPU Throttling:** В цикл генерации хода ИИ добавлена микро-задержка `time.Sleep(10 * time.Millisecond)`, предотвращающая 100% загрузку ядра процессора при поиске свободного зуба на поздних этапах игры. / **CPU Throttling Protection:** Added a micro-delay `time.Sleep(10 * time.Millisecond)` inside the AI turn generator loop, preventing 100% CPU core utilization when searching for an available tooth during the late-game phase.

---

## [0.02] — Добавление ИИ-противника / AI Opponent Introduction

### Added (Добавлено)
*   **Режим против ИИ:** Игра автоматически превратилась в дуэль с компьютером. Бот ходит сразу после успешного хода игрока. / **AI Matchups:** The game automatically turns into a duel against the computer. The bot takes its turn right after a successful player move.
*   **Имитация мышления:** Добавлена искусственная задержка `time.Sleep(1 * time.Second)` перед ходом ИИ для создания эффекта «размышления» робота. / **Thinking Simulation:** Added an artificial delay using `time.Sleep(1 * time.Second)` before the AI's move to create a realistic "thinking" effect.
*   **Интеллектуальный выбор:** ИИ научился проверять битовую маску `pressedMask` через цикл, чтобы никогда не нажимать на уже выбранные зубы. / **Smart Choice:** The AI learned to check the `pressedMask` bitmask via a loop to ensure it never picks an already pressed tooth.
*   **Локализация:** Добавлен перевод для реплик и состояния ИИ (`AI is thinking...`, `AI got lucky!`). / **Localization:** Added translations for AI actions and states (`AI is thinking...`, `AI got lucky!`).

### Known Issues / Bugs (Известные проблемы в этой версии)
*   **Общий счетчик ходов:** Переменная `steps` некорректно смешивает ходы человека и компьютера вместе при игре с ИИ. / **Shared Turn Counter:** The `steps` variable incorrectly mixes player and computer moves together during AI matches.
*   **Высокая нагрузка на процессор (Busy Wait):** Цикл генерации хода ИИ работает без задержек выполнения, нагружая ядро процессора на полную мощность во время подбора оставшихся зубов. / **High CPU Load (Busy Wait):** The AI turn generator loop runs with no execution delays, shifting the CPU core to maximum utilization when searching for available teeth.

---

## [0.01] — Первая стабильная версия / Initial Stable Release

### Added (Добавлено)
*   **Игровой движок:** Реализована базовая механика игры «Крокодил Дантист» в бесконечном цикле. / **Core Game Engine:** Implemented the foundational mechanics of the "Crocodile Dentist" game running inside an infinite loop.
*   **Генерация случайного зуба:** Добавлен генератор случайного «больного» зуба от 1 до 13 с инициализацией через `rand.Seed`. / **Random Tooth Generation:** Added a random generator for the "painful" tooth (ranges from 1 to 13) initialized via `rand.Seed`.
*   **Битовая маска:** Реализована проверка уже нажатых зубов через эффективную маску `uint16` (`pressedMask`). / **Bitmask Logic:** Implemented a system to track pressed teeth efficiently using a `uint16` bitmask (`pressedMask`).
*   **Многопоточность (Сигналы):** Добавлена горутина для перехвата сигналов закрытия (`Ctrl+C`, `SIGTERM`) для корректного выхода из консоли. / **Multithreading & Signals:** Added a goroutine to capture termination signals (`Ctrl+C`, `SIGTERM`) for a graceful console exit.
*   **Счетчик ходов:** Добавлено сохранение и вывод количества успешных шагов до проигрыша. / **Turn Counter:** Tracked and displayed the amount of successful moves made before losing.
*   **Локализация:** Текстовый интерфейс частично переведен на два языка (RU / EN). / **Localization:** The text-based UI has been partially translated into two languages (RU / EN).
*   **Защита от сбоев ввода:** Реализована очистка буфера при вводе некорректных символов (букв вместо чисел), предотвращающая зацикливание консоли. / **Input Crash Protection:** Implemented input buffer flushing when non-numeric characters are entered, preventing infinite console spam.
*   **Валидация диапазона:** Добавлен строгий запрет на ввод чисел вне игрового диапазона от 1 до 13. / **Range Validation:** Restricted inputs to valid numbers strictly between 1 and 13.
