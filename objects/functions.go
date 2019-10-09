package objects

import (
	"math"
	"math/rand"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"text/template"
	"time"
)

// Округление до двух знаков после запятой
func round(x float64) float64 {
	var rounder float64
	pow := math.Pow(10, 2.0)
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}
	return rounder / pow
}

// FalseOrTrue - Рандомное да или нет
// percent - 10, 30, 50, 90
func FalseOrTrue(percent float64) bool {
	rand.Seed(time.Now().UnixNano()) // Рандомизируем
	if randomNumber := rand.Float64(); randomNumber <= percent/100.0 {
		return true
	} else {
		return false
	}
}

// Open - открытие страницы в браузере по умолчанию
func Open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func printGamerParams() {
	if Gamer.Fatigue > 100.0 {
		Gamer.Fatigue = 100.0
	}
	if Gamer.Thirst > 100.0 {
		Gamer.Thirst = 100.0
		GameOver = true
	}
	if Gamer.Hunger > 100.0 {
		Gamer.Hunger = 100.0
		GameOver = true
	}
	if Gamer.Heals <= 0 {
		Gamer.Heals = 0
		GameOver = true
	}
	if !NightMode {
		if LengthOfDay == 0 {
			EventNow.TextAll += "<h5>День сейчас закончится</h5>"
			EventNow.TextAll += "Наступает темная и опасная ночь...<br>Постарайся поскорее найти укрытие!"
		} else {
			EventNow.TextAll += "<h5>День.<br>До конца дня осталось - " + strconv.FormatFloat(LengthOfDay, 'f', 0, 64) + "</h5>"
		}
	} else {
		if LengthOfNight == 0 {
			EventNow.TextAll += "<h5>Ночь сейчас закончится</h5>"
			EventNow.TextAll += "Наступает новый день.<br>Не упусти возможность подготовиться к следующей ночи"
		} else {
			EventNow.TextAll += "<h5>Ночь.<br>До конца ночи осталось - " + strconv.FormatFloat(LengthOfNight, 'f', 0, 64) + "</h5>"
		}
	}
	EventNow.TextAll += "<h4><b>Твои характеристики:</b></h4>" +
		"Здоровье: " + strconv.FormatFloat(Gamer.Heals, 'f', 0, 64) +
		"<br>Показатель атаки: " + strconv.FormatFloat(Gamer.IndexAttack+Gamer.Gear.IndexAttack, 'f', 0, 64) +
		"<br>Показатель защиты: " + strconv.FormatFloat(Gamer.IndexDefense+Gamer.Gear.IndexDefense, 'f', 0, 64) +
		"<br>Усталость: " + strconv.FormatFloat(Gamer.Fatigue, 'f', 2, 64) +
		", Голод: " + strconv.FormatFloat(Gamer.Hunger, 'f', 2, 64) +
		", Жажда: " + strconv.FormatFloat(Gamer.Thirst, 'f', 2, 64) +
		"<br><b>Негативные статусы:</b>"
	if Gamer.Debuff.BrokenBone || Gamer.Debuff.Cut || Gamer.Debuff.Disease {
		if Gamer.Debuff.BrokenBone {
			EventNow.TextAll += "<br>Сломанная кость"
		}
		if Gamer.Debuff.Cut {
			EventNow.TextAll += "<br>Порез"
		}
		if Gamer.Debuff.Disease {
			EventNow.TextAll += "<br>Болезнь"
		}
	} else {
		EventNow.TextAll += " Нету"
	}
	EventNow.TextAll += "<br><b>Оружие:</b>"
	if Gamer.Gear.Name != "" {
		EventNow.TextAll += " " + Gamer.Gear.Name +
			" (Прочность - " + strconv.Itoa(Gamer.Gear.Durability) + ")<br><br>"
	} else {
		EventNow.TextAll += " Нету<br><br>"
	}
}

func findAction(place string) {
	OneStart++
	switch place {
	case Forest:
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			Gamer.Backpack.Wood += 2
			Gamer.Backpack.Rod += 1
			EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>2 дерева<br>1 прутик<br><br>"
		} else {
			EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
		}
		EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		printGamerParams()
	case Plain:
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			Gamer.Backpack.Stone += 1
			Gamer.Backpack.Rod += 1
			EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>1 камень<br>1 прутик<br><br>"
		} else {
			EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
		}
		EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		printGamerParams()
	case Desert:
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			Gamer.Backpack.Stone += 1
			EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>1 камень<br><br>"
		} else {
			EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
		}
		EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		printGamerParams()
	case Road:
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			Gamer.Backpack.Food.Vegetables += 1
			Gamer.Backpack.Bottle.Water += 1
			Gamer.Backpack.Cloth += 1
			EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>1 овощ<br>1 бутылку с водой<br>1 тряпку<br><br>"
		} else {
			EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
		}
		EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		printGamerParams()
	case Swamp:
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			Gamer.Backpack.Wood += 1
			Gamer.Backpack.Rod += 1
			EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>1 дерево<br>1 прутик<br><br>"
		} else {
			EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
		}
		EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		printGamerParams()
	case Riverside:
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			Gamer.Backpack.Wood += 2
			Gamer.Backpack.Rod += 1
			Gamer.Backpack.Stone += 2
			EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>2 дерева<br>1 прутик<br>2 камня<br><br>"
		} else {
			EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
		}
		EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		printGamerParams()
	case Bridge:
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			Gamer.Backpack.Stone += 1
			EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>1 камень<br><br>"
		} else {
			EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
		}
		EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		printGamerParams()
	case Building:
		switch EventNow.Building {
		case Hospital:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Cloth += 3
				Gamer.Backpack.Bottle.Alcohol += 2
				Gamer.Backpack.Medicament += 1
				Gamer.Backpack.Splint += 1
				Gamer.Backpack.Bandage += 1
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>3 тряпки<br>2 бутылки с алкоголем<br>Лекарство<br>Шину от перелома<br>и 1 бинт<br><br>"
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		case PoliceStation:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Bottle.Alcohol += 1
				Gamer.Backpack.Bottle.Water += 1
				Gamer.Backpack.Food.Fruits += 1
				takeWeapon := false
				if Gamer.Gear.Name == "" {
					Gamer.Gear = *Weapons[Spear]
					takeWeapon = true
				}
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>1 бутылка с алкоголем<br>1 бутылка с водой<br>1 фрукт<br>"
				if takeWeapon {
					EventNow.TextAll += "<br>И оружие - копье!!!<br><br>"
				}
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		case FoodShop:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Food.Fruits += 3
				Gamer.Backpack.Food.Vegetables += 3
				Gamer.Backpack.Food.RawMeat += 3
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>3 фрукта<br>3 овоща<br>3 сырых куска мяса<br><br>"
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		case HardwareStore:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Cloth += 2
				Gamer.Backpack.Rod += 2
				Gamer.Backpack.Wood += 2
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>2 тряпки<br>2 прутика<br>2 дерева<br><br>"
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		case DepartmentStore:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Bottle.Water += 1
				Gamer.Backpack.Medicament += 1
				Gamer.Backpack.Splint += 1
				Gamer.Backpack.Bandage += 1
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>Бутылку с водой<br>Лекарство<br>Шину от перелома<br>и 1 бинт<br><br>"
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		case House:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Bottle.Alcohol += 1
				Gamer.Backpack.Bottle.Water += 1
				Gamer.Backpack.Food.Fruits += 1
				takeWeapon := false
				if Gamer.Gear.Name == "" {
					Gamer.Gear = *Weapons[BigStick]
					takeWeapon = true
				}
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>1 бутылка с алкоголем<br>1 бутылка с водой<br>1 фрукт<br>"
				if takeWeapon {
					EventNow.TextAll += "<br>И оружие - Дубина!!!<br><br>"
				}
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		case Hangar:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Cloth += 2
				Gamer.Backpack.Rod += 2
				Gamer.Backpack.Wood += 2
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>2 тряпки<br>2 прутика<br>2 дерева<br><br><br><br>"
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		case Camp:
			if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
				Gamer.Backpack.Cloth += 2
				Gamer.Backpack.Rod += 2
				Gamer.Backpack.Wood += 2
				EventNow.TextAll = "<h3>Ты добыл припасы!!!</h3>2 тряпки<br>2 прутика<br>2 дерева<br><br>"
			} else {
				EventNow.TextAll = "<h3>Ты ничего не добыл ...</h3>"
			}
			EventNow.TextAll += "Ты по-прежнему находишься тут: " + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
			printGamerParams()
		}
	}
}

func enemyAttackGamer() {
	if EventNow.EnemyAnimal != nil {
		Battle(EventNow.EnemyAnimal, Gamer)
	} else {
		Battle(EventNow.EnemyHuman, Gamer)
	}
}

func gamerLeaveBattle() {
	Gamer.Fatigue += 15.0
	if EventNow.EnemyAnimal != nil {
		if FalseOrTrue(50) {
			Gamer.Heals -= EventNow.EnemyAnimal.IndexAttack / 3
		} else {
			Gamer.Heals -= EventNow.EnemyAnimal.IndexAttack / 4
		}
		EventNow.TextAll = "После побега от " + EventNow.EnemyAnimal.Name + " ты находишься там же.<br>" + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
	} else {
		if FalseOrTrue(50) {
			Gamer.Heals -= EventNow.EnemyHuman.IndexAttack / 3
		} else {
			Gamer.Heals -= EventNow.EnemyHuman.IndexAttack / 4
		}
		EventNow.TextAll = "После побега от " + EventNow.EnemyHuman.Name + " ты находишься там же.<br>" + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
	}
	printGamerParams()
}

func gamerAttackEnemy() {
	if EventNow.EnemyAnimal != nil {
		Battle(Gamer, EventNow.EnemyAnimal)
	} else {
		Battle(Gamer, EventNow.EnemyHuman)
	}
}

func gamerLeaveFindObject() {
	if EventNow.EnemyAnimal != nil {
		EventNow.TextAll = "Ты незаметно ушел от " + EventNow.EnemyAnimal.Name + " и находишься там же.<br>" + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
	} else {
		EventNow.TextAll = "Ты незаметно ушел от " + EventNow.EnemyHuman.Name + " и находишься там же.<br>" + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
	}
	printGamerParams()
}

func craftTemplateGenerate() {
	// Динамически генерируем шаблон под те предметы, на которые хватает ресурсов у игрока
	TemplateCraft = "<html><head><title>Создание предметов</title></head><body>" +
		"<h4><b>Твои характеристики:</b></h4>" +
		"Здоровье: " + strconv.FormatFloat(Gamer.Heals, 'f', 0, 64) +
		"<br>Показатель атаки: " + strconv.FormatFloat(Gamer.IndexAttack+Gamer.Gear.IndexAttack, 'f', 0, 64) +
		"<br>Показатель защиты: " + strconv.FormatFloat(Gamer.IndexDefense+Gamer.Gear.IndexDefense, 'f', 0, 64) +
		"<br>Усталость: " + strconv.FormatFloat(Gamer.Fatigue, 'f', 2, 64) +
		", Голод: " + strconv.FormatFloat(Gamer.Hunger, 'f', 2, 64) +
		", Жажда: " + strconv.FormatFloat(Gamer.Thirst, 'f', 2, 64) +
		"<br><b>Негативные статусы:</b>"
	if Gamer.Debuff.BrokenBone || Gamer.Debuff.Cut || Gamer.Debuff.Disease {
		if Gamer.Debuff.BrokenBone {
			TemplateCraft += "<br>Сломанная кость"
		}
		if Gamer.Debuff.Cut {
			TemplateCraft += "<br>Порез"
		}
		if Gamer.Debuff.Disease {
			TemplateCraft += "<br>Болезнь"
		}
	} else {
		TemplateCraft += " Нету"
	}
	TemplateCraft += "<br><b>Оружие:</b>"
	if Gamer.Gear.Name != "" {
		TemplateCraft += " " + Gamer.Gear.Name +
			" (Прочность - " + strconv.Itoa(Gamer.Gear.Durability) + ")<br><br>"
	} else {
		TemplateCraft += " Нету<br><br>"
	}
	TemplateCraft += "<form action=\"http://127.0.0.1/game?key=action\" method=\"post\">" +
		"<h3>Ты можешь создать следующие предметы:</h3><select name=\"craftMenu\">"
	if Gamer.Backpack.Wood >= Craft[Spear][Wood] &&
		Gamer.Backpack.Stone >= Craft[Spear][Stone] &&
		Gamer.Backpack.Rod >= Craft[Spear][Rod] {
		TemplateCraft += "<option value=Копье>Копье</option>"
	}
	if Gamer.Backpack.Wood >= Craft[BigStick][Wood] &&
		Gamer.Backpack.Rod >= Craft[BigStick][Rod] {
		TemplateCraft += "<option value=Дубина>Дубина</option>"
	}
	if Gamer.Backpack.Wood >= Craft[Shield][Wood] &&
		Gamer.Backpack.Cloth >= Craft[Shield][Cloth] {
		TemplateCraft += "<option value=Щит>Щит</option>"
	}
	if Gamer.Backpack.Wood >= Craft[Slingshot][Wood] &&
		Gamer.Backpack.Rod >= Craft[Slingshot][Rod] &&
		Gamer.Backpack.Cloth >= Craft[Slingshot][Cloth] &&
		Gamer.Backpack.Stone >= Craft[Slingshot][Stone] {
		TemplateCraft += "<option value=Праща>Праща</option>"
	}
	if Gamer.Backpack.Skin >= Craft[Cloth][Skin] {
		TemplateCraft += "<option value=Тряпка>Тряпка</option>"
	}
	if Gamer.Backpack.Cloth >= Craft[Bandage][Cloth] &&
		Gamer.Backpack.Bottle.Alcohol >= Craft[Bandage][Alcohol] {
		TemplateCraft += "<option value=Бинт>Бинт</option>"
	}
	if Gamer.Backpack.Wood >= Craft[Splint][Wood] &&
		Gamer.Backpack.Cloth >= Craft[Splint][Cloth] {
		TemplateCraft += "<option value=Шина>Шина от перелома</option>"
	}
	if EventNow.Bonfire && Gamer.Backpack.Food.RawMeat >= 1 {
		TemplateCraft += "<option value=Стейк>Жареное мясо</option>"
	}
	if Gamer.Backpack.Bottle.Empty >= 1 && EventNow.Place == Riverside {
		TemplateCraft += "<option value=Вода>Набрать воды в бутылку</option>"
	}
	TemplateCraft += "</select><br><br><input type=\"submit\" value=\"Создать\"></form>" +
		"<h4><a href=\"http://127.0.0.1/game?key=action\">Я передумал что-то создавать</a></h4></body></html>"
}

func crafting(item string) {
	OneStart++
	switch item { // Перебор вариантов крафта
	case Spear:
		Gamer.Backpack.Wood -= Craft[Spear][Wood]
		Gamer.Backpack.Stone -= Craft[Spear][Stone]
		Gamer.Backpack.Rod -= Craft[Spear][Rod]
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		if Gamer.Gear.Name != "" {
			EventNow.TextAll += "<h3>Ты создал копье!</h3>Свое предыдущее оружие - " + Gamer.Gear.Name +
				" пришлось выбросить..."
		} else {
			EventNow.TextAll += "<h3>Ты создал копье!</h3>"
		}
		Gamer.Gear = *Weapons[Spear]
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		printGamerParams()
	case BigStick:
		Gamer.Backpack.Wood -= Craft[BigStick][Wood]
		Gamer.Backpack.Rod -= Craft[BigStick][Rod]
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		if Gamer.Gear.Name != "" {
			EventNow.TextAll += "<h3>Ты создал Дубину!</h3>Свое предыдущее оружие - " + Gamer.Gear.Name +
				" пришлось выбросить..."
		} else {
			EventNow.TextAll += "<h3>Ты создал дубину!</h3>"
		}
		Gamer.Gear = *Weapons[BigStick]
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		printGamerParams()
	case Shield:
		Gamer.Backpack.Wood -= Craft[Shield][Wood]
		Gamer.Backpack.Cloth -= Craft[Shield][Cloth]
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		if Gamer.Gear.Name != "" {
			EventNow.TextAll += "<h3>Ты создал щит!</h3>Свое предыдущее оружие - " + Gamer.Gear.Name +
				" пришлось выбросить..."
		} else {
			EventNow.TextAll += "<h3>Ты создал щит!</h3>"
		}
		Gamer.Gear = *Weapons[Shield]
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		printGamerParams()
	case Slingshot:
		Gamer.Backpack.Wood -= Craft[Slingshot][Wood]
		Gamer.Backpack.Rod -= Craft[Slingshot][Rod]
		Gamer.Backpack.Cloth -= Craft[Slingshot][Cloth]
		Gamer.Backpack.Stone -= Craft[Slingshot][Stone]
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		if Gamer.Gear.Name != "" {
			EventNow.TextAll += "<h3>Ты создал пращу!</h3>Свое предыдущее оружие - " + Gamer.Gear.Name +
				" пришлось выбросить..."
		} else {
			EventNow.TextAll += "<h3>Ты создал пращу!</h3>"
		}
		Gamer.Gear = *Weapons[Slingshot]
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		printGamerParams()
	case Cloth:
		Gamer.Backpack.Skin -= Craft[Cloth][Skin]
		Gamer.Backpack.Cloth += 1
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h3>Ты создал тряпку!</h3>"
		printGamerParams()
	case Bandage:
		Gamer.Backpack.Cloth -= Craft[Bandage][Cloth]
		Gamer.Backpack.Bottle.Alcohol -= Craft[Bandage][Alcohol]
		Gamer.Backpack.Bandage += 1
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h3>Ты создал бинт!</h3>"
		printGamerParams()
	case Splint:
		Gamer.Backpack.Wood -= Craft[Splint][Wood]
		Gamer.Backpack.Cloth -= Craft[Splint][Cloth]
		Gamer.Backpack.Splint += 1
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h3>Ты создал шину от перелома!</h3>"
		printGamerParams()
	case RoastedMeat:
		Gamer.Backpack.Food.RawMeat -= 1
		Gamer.Backpack.Food.RoastedMeat += 1
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h3>Ты пожарил сырое мясо и получил вкуснейший стейк!</h3>"
		printGamerParams()
	case "Вода":
		Gamer.Backpack.Bottle.Empty -= 1
		Gamer.Backpack.Bottle.Water += 1
		Gamer.Fatigue += 2.0 * indexForDifficulty[IndexFatigueAction][Difficulty] // Добавляем усталость
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h3>Ты набрал чистой воды в свободную пустую бутылку!</h3>"
		printGamerParams()
	}
}

func EndGame(result string, w http.ResponseWriter, r *http.Request) {
	TemplateEndGame = "<html><head><title>Конец игры</title></head><body>" +
		"<form action=\"http://127.0.0.1/end\" method=\"post\">"
	if GameOver {
		TemplateEndGame += "<h1>Ты проиграл</h1>"
		switch result {
		case "water":
			TemplateEndGame += "<h3>Умер от жажды... Не забывай вовремя пить воду!</h3>"
		case "food":
			TemplateEndGame += "<h3>Умер от голода... Кушать нужно, дружок!</h3>"
		case "heals":
			TemplateEndGame += "<h3>Умер от полученных ранений... Не лезь в бой если не готов!</h3>"
		}
	}
	TemplateEndGame += "<input type=\"submit\" value=\"Конец...\"></form></body></html>"
	t, _ := template.New("").Parse(TemplateEndGame) // Загружаем шаблон конца игры
	t.Execute(w, "")                                // Отображение шаблона
}

func checkEnd(w http.ResponseWriter, r *http.Request) bool {
	if Gamer.Thirst >= 100.0 {
		GameOver = true
		EndGame("water", w, r)
		return true
	}
	if Gamer.Hunger >= 100.0 {
		GameOver = true
		EndGame("food", w, r)
		return true
	}
	if Gamer.Heals <= 0 {
		GameOver = true
		EndGame("heals", w, r)
		return true
	}
	return false
}

func generateTemplateApply() {
	// Динамически генерируем шаблон под те предметы, которые есть у игрока и которые можно применить
	TemplateApply = "<html><head><title>Применение предметов</title></head><body>" +
		"<h4><b>Твои характеристики:</b></h4>" +
		"Здоровье: " + strconv.FormatFloat(Gamer.Heals, 'f', 0, 64) +
		"<br>Показатель атаки: " + strconv.FormatFloat(Gamer.IndexAttack+Gamer.Gear.IndexAttack, 'f', 0, 64) +
		"<br>Показатель защиты: " + strconv.FormatFloat(Gamer.IndexDefense+Gamer.Gear.IndexDefense, 'f', 0, 64) +
		"<br>Усталость: " + strconv.FormatFloat(Gamer.Fatigue, 'f', 2, 64) +
		", Голод: " + strconv.FormatFloat(Gamer.Hunger, 'f', 2, 64) +
		", Жажда: " + strconv.FormatFloat(Gamer.Thirst, 'f', 2, 64) +
		"<br><b>Негативные статусы:</b>"
	if Gamer.Debuff.BrokenBone || Gamer.Debuff.Cut || Gamer.Debuff.Disease {
		if Gamer.Debuff.BrokenBone {
			TemplateApply += "<br>Сломанная кость"
		}
		if Gamer.Debuff.Cut {
			TemplateApply += "<br>Порез"
		}
		if Gamer.Debuff.Disease {
			TemplateApply += "<br>Болезнь"
		}
	} else {
		TemplateApply += " Нету"
	}
	TemplateApply += "<br><b>Оружие:</b>"
	if Gamer.Gear.Name != "" {
		TemplateApply += " " + Gamer.Gear.Name +
			" (Прочность - " + strconv.Itoa(Gamer.Gear.Durability) + ")<br><br>"
	} else {
		TemplateApply += " Нету<br><br>"
	}
	TemplateApply += "<form action=\"http://127.0.0.1/game?key=action\" method=\"post\">" +
		"<h3>Что хочешь сделать:</h3><select name=\"applyMenu\">"
	if Gamer.Backpack.Bottle.Water >= 1 {
		TemplateApply += "<option value=Вода>Выпить воды</option>"
	}
	if Gamer.Backpack.Bandage >= 1 {
		TemplateApply += "<option value=Бинт>Наложить бинт</option>"
	}
	if Gamer.Backpack.Medicament >= 1 {
		TemplateApply += "<option value=Лекарство>Принять лекарство</option>"
	}
	if Gamer.Backpack.Splint >= 1 {
		TemplateApply += "<option value=Шина>Наложить шину</option>"
	}
	if Gamer.Backpack.Food.Fruits >= 1 {
		TemplateApply += "<option value=Фрукты>Съесть фрукт</option>"
	}
	if Gamer.Backpack.Food.RoastedMeat >= 1 {
		TemplateApply += "<option value=Стейк>Съесть стейк</option>"
	}
	if Gamer.Backpack.Food.Vegetables >= 1 {
		TemplateApply += "<option value=Овощ>Съесть овощ</option>"
	}
	TemplateApply += "</select><br><br><input type=\"submit\" value=\"Применить\"></form>" +
		"<h4><a href=\"http://127.0.0.1/game?key=action\">Я передумал что-то принимать</a></h4></body></html>"
}

func apply(item string) {
	OneStart++
	switch item {
	case "Вода":
		Gamer.Backpack.Bottle.Water -= 1
		Gamer.Backpack.Bottle.Empty += 1
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h2>Ты выпил бутылку с водой, уталив жажду.</h2>Пустую бутылку положил обратно в рюкзак."
		if Gamer.Thirst -= 50.0; Gamer.Thirst < 0 {
			Gamer.Thirst = 0
		}
		if Gamer.Heals += 5; Gamer.Heals > 100.0 {
			Gamer.Heals = 100.0
		}
		if Gamer.Fatigue -= 5; Gamer.Fatigue < 0.0 {
			Gamer.Fatigue = 0.0
		}
		printGamerParams()
	case "Бинт":
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		if Gamer.Debuff.Cut {
			Gamer.Backpack.Bandage -= 1
			EventNow.TextAll += "<h2>Ты перевязал рану и остановил кровотечение.</h2>"
			Gamer.Debuff.Cut = false
			if Gamer.Heals += 10; Gamer.Heals > 100.0 {
				Gamer.Heals = 100.0
			}
			if Gamer.Fatigue -= 5; Gamer.Fatigue < 0.0 {
				Gamer.Fatigue = 0.0
			}
		} else {
			EventNow.TextAll += "<h2>У тебя нету ран, требующих перевязки.</h2>Этот предмет в будущем может спасти тебе жизнь.<br>Думай, прежде чем выбрасывать что-то!"
		}
		printGamerParams()
	case "Лекарство":
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		if Gamer.Debuff.Disease {
			Gamer.Backpack.Medicament -= 1
			EventNow.TextAll += "<h2>Ты принял лекарство и остановил распространение болезни.</h2>Чувствуешь себя значительно лучше."
			Gamer.Debuff.Disease = false
			if Gamer.Heals += 10; Gamer.Heals > 100.0 {
				Gamer.Heals = 100.0
			}
			if Gamer.Fatigue -= 5; Gamer.Fatigue < 0.0 {
				Gamer.Fatigue = 0.0
			}
		} else {
			EventNow.TextAll += "<h2>Тебе не надо принимать лекарство!</h2>Оставь полезный предмет на случай нужды!"
		}
		printGamerParams()
	case "Шина":
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		if Gamer.Debuff.BrokenBone {
			Gamer.Backpack.Splint -= 1
			EventNow.TextAll += "<h2>Ты наложил шину на перелом.</h2>Наконец-то можешь ходить без той сильной боли."
			Gamer.Debuff.BrokenBone = false
			if Gamer.Heals += 10; Gamer.Heals > 100.0 {
				Gamer.Heals = 100.0
			}
			if Gamer.Fatigue -= 5; Gamer.Fatigue < 0.0 {
				Gamer.Fatigue = 0.0
			}
		} else {
			EventNow.TextAll += "<h2>У тебя нету перелома!</h2>Оставь шину в покое, тебе делать чтоли нечего?"
		}
		printGamerParams()
	case "Фрукты":
		Gamer.Backpack.Food.Fruits -= 1
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h2>Ты съел фрукт.</h2>Сочный плод частично уталил жажду и голод"
		if Gamer.Heals += 5; Gamer.Heals > 100.0 {
			Gamer.Heals = 100.0
		}
		if Gamer.Thirst -= 8; Gamer.Thirst < 0.0 {
			Gamer.Thirst = 0.0
		}
		if Gamer.Hunger -= 8; Gamer.Hunger < 0.0 {
			Gamer.Hunger = 0.0
		}
		printGamerParams()
	case "Стейк":
		Gamer.Backpack.Food.RoastedMeat -= 1
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h2>Ты съел стейк.</h2>Очень сытный кусок мяса, ты победил голод!"
		if Gamer.Heals += 5; Gamer.Heals > 100.0 {
			Gamer.Heals = 100.0
		}
		if Gamer.Hunger -= 15; Gamer.Hunger < 0.0 {
			Gamer.Hunger = 0.0
		}
		printGamerParams()
	case "Овощ":
		Gamer.Backpack.Food.Vegetables -= 1
		EventNow.TextAll = "<br>Ты находишься там же - " + EventNow.Place
		if EventNow.Place == Building {
			EventNow.TextAll += " - " + EventNow.Building
		}
		EventNow.TextAll += "<h2>Ты съел овощ.</h2>Полезный овощ немного уталил голод и жажду."
		if Gamer.Heals += 5; Gamer.Heals > 100.0 {
			Gamer.Heals = 100.0
		}
		if Gamer.Thirst -= 5; Gamer.Thirst < 0.0 {
			Gamer.Thirst = 0.0
		}
		if Gamer.Hunger -= 10; Gamer.Hunger < 0.0 {
			Gamer.Hunger = 0.0
		}
		printGamerParams()
	}
}

func generateTemplateBackpack() {
	// Динамически генерируем шаблон под те предметы, на которые хватает ресурсов у игрока
	TemplateBackpack = "<html><head><title>Содержимое рюкзака</title></head><body>" +
		"<h3>У тебя в рюкзаке есть:</h3>"
	if Gamer.Backpack.Wood >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Wood) + "шт. дерева<br>"
	}
	if Gamer.Backpack.Stone >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Stone) + "шт. камня<br>"
	}
	if Gamer.Backpack.Splint >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Splint) + "шт. шин для перелома<br>"
	}
	if Gamer.Backpack.Skin >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Skin) + "шт. шкур диких животных<br>"
	}
	if Gamer.Backpack.Rod >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Rod) + "шт. прутиков<br>"
	}
	if Gamer.Backpack.Medicament >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Medicament) + "шт. лекарств<br>"
	}
	if Gamer.Backpack.Cloth >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Cloth) + "шт. тряпок<br>"
	}
	if Gamer.Backpack.Bandage >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Bandage) + "шт. бинтов<br>"
	}
	if Gamer.Backpack.Food.Vegetables >= 1 || Gamer.Backpack.Food.RoastedMeat >= 1 || Gamer.Backpack.Food.Fruits >= 1 {
		TemplateBackpack += "<br><b>Еда:</b><br>"
	}
	if Gamer.Backpack.Food.Vegetables >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Food.Vegetables) + "шт. овощей<br>"
	}
	if Gamer.Backpack.Food.RoastedMeat >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Food.RoastedMeat) + "шт. стейков<br>"
	}
	if Gamer.Backpack.Food.RawMeat >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Food.RawMeat) + "шт. сырого мяса<br>"
	}
	if Gamer.Backpack.Food.Fruits >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Food.Fruits) + "шт. фруктов<br>"
	}
	if Gamer.Backpack.Bottle.Empty >= 1 || Gamer.Backpack.Bottle.Water >= 1 || Gamer.Backpack.Bottle.Alcohol >= 1 {
		TemplateBackpack += "<br><b>Бутылки:</b><br>"
	}
	if Gamer.Backpack.Bottle.Empty >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Bottle.Empty) + "шт. пустых бутылок<br>"
	}
	if Gamer.Backpack.Bottle.Water >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Bottle.Water) + "шт. бутылок с водой<br>"
	}
	if Gamer.Backpack.Bottle.Alcohol >= 1 {
		TemplateBackpack += strconv.Itoa(Gamer.Backpack.Bottle.Alcohol) + "шт. бутылок с алкоголем<br>"
	}
	TemplateBackpack += "<br><h4><a href=\"http://127.0.0.1/game?key=action\">Назад к выбору действия</a></h4></body></html>"
}
