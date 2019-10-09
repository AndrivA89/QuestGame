package objects

import (
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("").Parse(TemplateStart) // Загружаем шаблон выбора параметров игры
	t.Execute(w, "")                              // Отображение шаблона
}
func EndGamePage(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func GamePage(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key") // Получение параметра адресной строки
	switch key {              // Перебор параметров
	case "start": // Первый запуск
		GamerName = r.FormValue("gamer")                   // Определение имени игрока
		Difficulty = r.FormValue("difficulty")             // Определение уровня сложности
		Initialization(Difficulty)                         // Генерация игрового мира
		t, _ := template.New("Gamer").Parse(TemplateHello) // Загружаем шаблон приветствия
		t.Execute(w, Gamer)                                // Отображение шаблона
		OneStart++
	case "action": // Выбор действия
		if OneStart > 1 {
			EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		} else {
			printGamerParams()
		}
		attackAction := r.FormValue("battle")
		gamerAttack := r.FormValue("gamerBattle")
		craftMenu := r.FormValue("craftMenu")
		item := r.FormValue("applyMenu")
		crafting(craftMenu)   // создание предметов
		apply(item)           // Применение предметов
		switch attackAction { // Перебор действия при атаке игрока врагом
		case "figth":
			enemyAttackGamer()
		case "leave":
			gamerLeaveBattle()
		}
		switch gamerAttack { // Перебор действия при атаке врага игроком
		case "figth":
			gamerAttackEnemy()
		case "leave":
			gamerLeaveFindObject()
		}
		if !checkEnd(w, r) {
			t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон действия
			t.Execute(w, EventNow)                                       // Отображение шаблона
		}
	case "actionDone": // Выполняем выбранное действие
		action := r.FormValue("action")
		switch action {
		case "selectWay": // Выбрали идти дальше
			Gamer.Fatigue += 8.5 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
			Gamer.Thirst += 5.9 * indexForDifficulty[IndexFatigueAction][Difficulty]  // Добавляем жажду
			Gamer.Hunger += 3.5 * indexForDifficulty[IndexFatigueAction][Difficulty]  // Добавляем голод
			WaySelect := WaySelectData{}
			EventNow.generateWays()        // Создание путей
			WaySelect = EventNow.setWays() // Установка путей
			printGamerParams()
			if Gamer.Fatigue >= 99.9 {
				EventNow.TextAll = "<h3>Ты валишься с ног...</h3>Дальше идти просто не было сил и ты завалился спать прямо тут - " + EventNow.Place
				if EventNow.Place == Building {
					EventNow.TextAll += " - " + EventNow.Building
				}
				switch EventNow.Place {
				case Forest, Plain, Swamp, Road, Riverside, Desert, Bridge:
					if EventNow.Bonfire {
						Gamer.Sleep(Bonfire)
						EventNow.TextAll += "<br>Хорошо хоть костер был разведен!<br>Сон получился более комфортным<br>Ты проголодался и хочешь пить..."
						printGamerParams()
					} else {
						Gamer.Sleep(Outdoor)
						EventNow.TextAll += "<br>Ты полежал на холодной земле, свернушись калачиком.<br>Сном это трудно назвать, но все же ты немного отдохнул.<br>Ты хочешь есть и пить..."
						printGamerParams()
					}
				case Building:
					Gamer.Sleep(Haven)
					EventNow.TextAll += "<br>Благодари Бога (наверное) за то, что ты упал без чувств в безопасном месте!<br>Ты выспался и чувствуешь себя отдохнувшим!<br>Ну разве что только проголодался немного и хочешь пить..."
					printGamerParams()
				}
				if !checkEnd(w, r) {
					t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон выбора действия
					t.Execute(w, EventNow)                                       // Отображение шаблона
				}
			} else {
				printGamerParams()
				if !checkEnd(w, r) {
					t, _ := template.New("WaySelect").Parse(TemplateSelectWay) // Загружаем шаблон выбора пути
					t.Execute(w, WaySelect)                                    // Отображение шаблона
				}
			}
		case "find": // Выбрали поиск
			if !EventNow.IndexFind { // если еще не осматривали местность
				// Добавляем усталость
				Gamer.Fatigue += 2.5 * indexForDifficulty[IndexFatigueAction][Difficulty]
				EventNow.IndexFind = true  // Уже обыскивали
				findAction(EventNow.Place) // Запускаем поиск
				if !checkEnd(w, r) {
					t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон действия
					t.Execute(w, EventNow)                                       // Отображение шаблона
				}
			} else { // если уже осматривали местность
				EventNow.TextAll = "<h1>Ты уже осматривал это место!</h1>"
				EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
				if EventNow.Place == Building {
					EventNow.TextAll += " - " + EventNow.Building
				}
				EventNow.TextAll += "<br>"
				printGamerParams()
				if !checkEnd(w, r) {
					t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон действия
					t.Execute(w, EventNow)                                       // Отображение шаблона
				}
			}
		case "craft": // Выбрали создание предметов
			craftTemplateGenerate()
			if !checkEnd(w, r) {
				t, _ := template.New("EventNow").Parse(TemplateCraft) // Загружаем шаблон крафта
				t.Execute(w, EventNow)                                // Отображение шаблона
			}
		case "apply": // Выбрали применение предметов
			generateTemplateApply()
			if !checkEnd(w, r) {
				t, _ := template.New("EventNow").Parse(TemplateApply) // Загружаем шаблон применения предметов
				t.Execute(w, EventNow)                                // Отображение шаблона
			}
		case "backpack": // Выбрали просмотр рюкзака
			generateTemplateBackpack()
			if !checkEnd(w, r) {
				t, _ := template.New("EventNow").Parse(TemplateBackpack) // Загружаем шаблон отображения содержимого рюкзака
				t.Execute(w, EventNow)                                   // Отображение шаблона
			}
		case "bonfire": // Выбрали разжечь костер
			if EventNow.Place != Building { // Если ты не в здании
				if !EventNow.Bonfire { // Если костер уже не горит
					if Gamer.Backpack.Wood >= Craft[Bonfire][Wood] &&
						Gamer.Backpack.Cloth >= Craft[Bonfire][Cloth] &&
						Gamer.Backpack.Stone >= Craft[Bonfire][Stone] {
						// Добавляем усталость
						Gamer.Fatigue += 2.5 * indexForDifficulty[IndexFatigueAction][Difficulty]
						EventNow.Bonfire = true
						Gamer.Backpack.Wood -= Craft[Bonfire][Wood]
						Gamer.Backpack.Cloth -= Craft[Bonfire][Cloth]
						Gamer.Backpack.Stone -= Craft[Bonfire][Stone]
						EventNow.TextAll = "<h3>Ты развел костер!</h3>Можно пожарить мяса или поспать возле огня<br>"
						EventNow.TextAll += "Ты по-прежнему находишься тут - " + EventNow.Place
						printGamerParams()
						if !checkEnd(w, r) {
							t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон выбора пути
							t.Execute(w, EventNow)                                       // Отображение шаблона
						}
					} else {
						EventNow.TextAll = "<h3>Тебе не хватает ресурсов для этого!</h3>"
						EventNow.TextAll += "Ты по-прежнему находишься тут - " + EventNow.Place
						printGamerParams()
						if !checkEnd(w, r) {
							t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон выбора пути
							t.Execute(w, EventNow)                                       // Отображение шаблона
						}
					}
				} else { // Костер был разведен ранее
					EventNow.TextAll = "<h3>Костер уже был разведен.<br>Тебе что нужно два костра?..</h3>"
					EventNow.TextAll += "Ты по-прежнему находишься тут - " + EventNow.Place
					printGamerParams()
					if !checkEnd(w, r) {
						t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон выбора пути
						t.Execute(w, EventNow)                                       // Отображение шаблона
					}
				}
			} else { // Если ты в здании
				EventNow.TextAll = "<h3>Костер в здании разводить не стоит!<br>Ты можешь сжечь строение, придурок!</h3>"
				EventNow.TextAll += "Ты по-прежнему находишься тут - " + EventNow.Building
				printGamerParams()
				if !checkEnd(w, r) {
					t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон выбора пути
					t.Execute(w, EventNow)                                       // Отображение шаблона
				}
			}
		case "sleep": // Выбрали сон
			EventNow.TextAll = ""
			switch EventNow.Place {
			case Forest, Plain, Swamp, Road, Riverside, Desert, Bridge:
				if EventNow.Bonfire {
					Gamer.Sleep(Bonfire)
					EventNow.TextAll += "Ты немного поспал у огня и чувствуешь себя лучше!<br>Ты проголодался и хочешь пить..."
					EventNow.TextAll += "<br>Ты находишься там же - " + EventNow.Place
					printGamerParams()
				} else {
					Gamer.Sleep(Outdoor)
					EventNow.TextAll += "Ты полежал на холодной земле, свернушись калачиком.<br>Сном это трудно назвать, но все же ты немного отдохнул.<br>Ты хочешь есть и пить..."
					EventNow.TextAll += "<br>Ты находишься там же - " + EventNow.Place
					printGamerParams()
				}
			case Building:
				Gamer.Sleep(Haven)
				EventNow.TextAll += "Ты выспался и чувствуешь себя отдохнувшим!<br>Ну разве что только проголодался немного и хочешь пить..."
				EventNow.TextAll += "<br>Ты находишься там же - " + EventNow.Building
				printGamerParams()
			}
			if !checkEnd(w, r) {
				t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон действия
				t.Execute(w, EventNow)                                       // Отображение шаблона
			}
		}
	case "event": // Новое событие после очередного выбора пути
		EventNow = &Event{}                  // Создаем новый экземпляр события - очистка предыдущего
		EventNow.Gamer = Gamer               // Привязываем к событию игрока
		NowSelectWay = r.FormValue("select") // Сохранение пути в текущем событии
		PrevSelectWay = NowSelectWay         // Сохранение пути в предыдущем событии
		switch NowSelectWay {
		case North:
			EventNow.Place = EventNorth.Place
		case West:
			EventNow.Place = EventWest.Place
		case East:
			EventNow.Place = EventEast.Place
		case South:
			EventNow.Place = EventSouth.Place
		}
		if EventNow.Place == Building { // Если место является зданием
			rand.Seed(time.Now().UnixNano()) // Рандомизируем
			//Выбираем случайное здание
			EventNow.Building = OptionsBuildingEvent[rand.Intn(len(OptionsBuildingEvent))]
			EventNow.TextAll += "<br>Это " + EventNow.Building
		}
		EventNow.setTypeEvent()             // Определяем тип события
		EventNow.NewEvent()                 // Новое событие
		if EventNow.TypeEvent == "Attack" { // Если тип события - Атака
			if !checkEnd(w, r) {
				t, _ := template.New("EventNow").Parse(TemplateAttackOnGamer) // Загружаем шаблон атаки игрока врагом
				t.Execute(w, EventNow)                                        // Отображение шаблона
			}
		} else { // Если тип события не атака
			if EventNow.EnemyAnimal != nil || EventNow.EnemyHuman != nil {
				if !checkEnd(w, r) {
					t, _ := template.New("EventNow").Parse(TemplateGamerAttack) // Загружаем шаблон атаки игроком
					t.Execute(w, EventNow)                                      // Отображение шаблона
				}
			} else {
				if !checkEnd(w, r) {
					t, _ := template.New("EventNow").Parse(TemplateSelectAction) // Загружаем шаблон действия
					t.Execute(w, EventNow)                                       // Отображение шаблона
				}
			}
		}
	}
}
