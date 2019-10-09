package objects

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Sleep - действие сон
func (h *Human) Sleep(place string) {
	var index float64 // Переменная для определения индекса комфорта
	index = indexComfortSleep[Difficulty][place]
	if h.Heals += 5.0 * index; h.Heals > 100.0 {
		h.Heals = 100.0 // Лечим персонажа
	}
	if h.Fatigue -= 20.0 * index; h.Fatigue < 0.0 {
		h.Fatigue = 0.0 // Уменьшаем усталость
	}
	if h.Hunger += 5.0; h.Hunger > 100.0 {
		h.Hunger = 100.0 // Увеличиваем чувство голода
	}
	if h.Thirst += 10.0; h.Thirst > 100.0 {
		h.Thirst = 100.0 // Увеличиваем уровень жажды
	}
	if !NightMode { // если сейчас день, то "проматываем" количество часов
		// равное количеству часов ночи на самом тяжелом уровне сложности, чтобы точно закончить ночь
		if LengthOfDay -= indexForDifficulty[GenerateLengthOfNight][Hard]; LengthOfDay < 0 {
			NightMode = true // устанавливаем день и ставим количество часов дня по уровню сложности
			LengthOfNight = indexForDifficulty[GenerateLengthOfNight][Difficulty]
		}
	} else { // если сейчас ночь, то "проматываем" количество часов
		// равное количеству часов ночи на самом тяжелом уровне сложности, чтобы точно закончить ночь
		if LengthOfNight -= indexForDifficulty[GenerateLengthOfNight][Hard]; LengthOfNight < 0 {
			NightMode = false // устанавливаем ночь и ставим количество часов ночи по уровню сложности
			LengthOfDay = indexForDifficulty[GenerateLengthOfDay][Difficulty]
		}
	}

}

// NewHuman - Создание начальных условий для объекта Human
func (h *Human) NewHuman(gamer bool, name string) {
	var (
		index              float64 // Переменная для определения индекса сложности
		randIndex          = 0     // Переменная для случайного увеличения показателя
		percentItem        float64 // Процент получения доп. предмета
		indexAttackDefense float64 // Переменная для случайного увеличения атаки и защиты
	)
	rand.Seed(time.Now().UnixNano())                    // Рандомизируем
	indexAttackDefense = rand.Float64()                 // Индекс случайного увеличения атаки и защиты
	indexAttackDefense = math.Round(indexAttackDefense) // Округление

	switch gamer {
	case true: // Объект является игроком
		Gamer.index = -1 // Метка для внутренних функций
		index = indexForDifficulty[GenerateGamerParams][Difficulty]
		percentItem = indexForDifficulty[PercentItemGamer][Difficulty]
	case false: // Объект является ботом
		index = indexForDifficulty[GenerateBotsParams][Difficulty]
		percentItem = indexForDifficulty[PercentItemEnemy][Difficulty]
	}
	h.Name, h.Heals = name, 50.0*index
	if h.Heals > 100.0 {
		h.Heals = 100.0
	}
	h.Thirst, h.Hunger, h.Fatigue = 0.0, 0.0, 0.0
	h.Debuff.Cut, h.Debuff.BrokenBone, h.Debuff.Disease = false, false, false
	h.IndexAttack, h.IndexDefense = math.Round(10.0*index)+5*indexAttackDefense, math.Round(2.0*index)+5*indexAttackDefense
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Bottle.Empty = int(math.Round(0.2*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Bottle.Water = int(math.Round(0.3*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Bottle.Alcohol = int(math.Round(0.3*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Wood, h.Backpack.Stone = int(math.Round(1.0*index))+randIndex, int(math.Round(0.7*index))+randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Rod, h.Backpack.Cloth = int(math.Round(0.3*index))+randIndex, int(math.Round(0.3*index))+randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Skin = int(math.Round(0.2*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Medicament = int(math.Round(0.2*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Bandage = int(math.Round(0.2*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Splint = int(math.Round(0.2*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Food.RoastedMeat, h.Backpack.Food.Fruits = int(math.Round(0.3*index))+randIndex, int(math.Round(0.3*index))+randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Food.RawMeat = int(math.Round(0.3*index)) + randIndex
	randIndex = 0
	if FalseOrTrue(percentItem) {
		randIndex = 1
	}
	h.Backpack.Food.Vegetables = int(math.Round(0.3*index)) + randIndex
	if FalseOrTrue(percentItem) {
		h.Gear = *Weapons[BigStick] // Выдаем самое слабое оружие
	}
}

// Генерация нового события
func (e *Event) NewEvent() {
	OneStart++
	switch NightMode {
	case true: // Ночное событие
		if LengthOfNight != 0 {
			LengthOfNight-- // уменьшаем ночное время
		} else {
			NightMode = false
			LengthOfDay = indexForDifficulty[GenerateLengthOfDay][Difficulty]
		}
		switch EventNow.Place {
		case Forest:
			if FalseOrTrue(50) {
				EventNow.TextAll = "Ты забрел в ночной, темный лес!<br>"
				if FalseOrTrue(50) {
					EventNow.TextAll += "Здесь очень страшно и слышится жуткий вой...<br>Разведи костер или уходи отсюда скорее<br>"
				} else {
					EventNow.TextAll += "Здесь явно кто-то или что-то прячется<br>Если ты потеряешь бдительность, то оно покажет себя!<br>"
				}
			}
			switch EventNow.TypeEvent { // Перебор типа события
			case Attack: // Событие - атака
				if !EventNow.setEnemy(75) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find: // Событие - находка
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Bandage += 1
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежал 1 бинт"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
				printGamerParams()
			}
		case Plain:
			EventNow.TextAll = "Ты подошел к равнине!<br>Ночью здесь практически ничего не видно.<br>Если ты разведешь огонь, то будешь как на ладони!"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(70) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find: // Событие - находка
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Splint += 1
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежала шина от перелома"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!<br>"
				printGamerParams()
			}
		case Desert:
			EventNow.TextAll = "Ты оказался в холодной ночной пустыне. <br>Не лучшее место, чтобы провести ночь...<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(60) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем когда-то видимо что-то лежало...<br>Сейчас же там пусто!<br>"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
				printGamerParams()
			}
		case Road:
			EventNow.TextAll = "Ты оказался прямо на автомобильной дороге<br>Ночью любая тень здесь напоминает чудовище или силуэт врага...<br>Следовало бы развести огонь!<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(50) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Cloth += 1
					Gamer.Backpack.Bottle.Water += 1
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало немного припасов:<br>1 тряпка<br>1 бутылка с водой<br>1 сырое мясо"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
				printGamerParams()
			}
		case Swamp:
			EventNow.TextAll = "Ты попал на болотные топи...<br>Тут и днем довольно страшно, а сейчас это самое худшее место на земле!<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(80) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Но к сожалению в нем уже кто-то порылся..."
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
				printGamerParams()
			}
		case Riverside:
			EventNow.TextAll = "Судя по звуку воды - ты подошел к берегу бурной реки<br>Здесь можно набрать воды в пустые бутылки<br>Но лучше развести сначала костер и переждать ночь.<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(65) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Food.Fruits += 2
					EventNow.TextAll += "<h3>Ты нашел яблоневое дерево прямо у воды!</h3>Собрал все яблоки с дерева, к сожалению, их было всего 2"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
				printGamerParams()
			}
		case Bridge:
			EventNow.TextAll = "Ты пришел на старый обветшалый мост!<br>Темно и практически ничего не видно...<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(60) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Skin += 1
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежала шкура дикого зверя"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
				printGamerParams()
			}
		case Building:
			switch EventNow.Building {
			case Hospital:
				EventNow.TextAll = "Ты оказался у госпиталя!<br>Отличное место, чтобы скрыться от ужасов ночи! <br>Скорее забегай в него<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(30) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Medicament += 1
						Gamer.Backpack.Splint += 1
						Gamer.Backpack.Bandage += 1
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало ОЧЕНЬ много припасов:<br>Лекарство<br>Шина от перелома<br>и 1 бинт"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			case PoliceStation:
				EventNow.TextAll = "Ты подошел к полицейскому участку!<br>Отличное место, чтобы скрыться от ужасов ночи! <br>Скорее забегай в него<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(30) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Bottle.Alcohol += 1
						takeWeapon := false
						if Gamer.Gear.Name == "" {
							Gamer.Gear = *Weapons[Spear]
							takeWeapon = true
						}
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало немного припасов:<br>1 бутылка с алкоголем"
						if takeWeapon {
							EventNow.TextAll += "<br>И оружие - копье!!!"
						}
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			case FoodShop:
				EventNow.TextAll = "Ты пришел в продуктовый магазин!<br>Здесь можно переночевать, а также набрать продуктов питания в дорогу<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Food.Fruits += 1
						Gamer.Backpack.Food.Vegetables += 1
						Gamer.Backpack.Food.RawMeat += 1
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало много съедобных припасов:<br>Фрукт<br>Овощ<br>и сырой кусок мяса"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			case HardwareStore:
				EventNow.TextAll = "Ты оказался рядом со строительным магазином!<br>Отличное место, чтобы провести ночь с комфортом.<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Но кто-то явно был тут до тебя..."
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			case DepartmentStore:
				EventNow.TextAll = "Ты возле универсального магазина!<br>Можно переночевать и скрыться от диких животных, бродящих в ночи.<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Его совсем недавно кто-то разорил! Берегись, этот кто-то явно где-то рядом..."
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			case House:
				EventNow.TextAll = "Ты подошел к чьему-то бывшему дому!<br>Скорее забегай внутрь, чтобы переночевать здесь!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(35) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Food.Fruits += 1
						takeWeapon := false
						if Gamer.Gear.Name == "" {
							Gamer.Gear = *Weapons[BigStick]
							takeWeapon = true
						}
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>На столе ты увидел фрукт и забрал его!"
						if takeWeapon {
							EventNow.TextAll += "<br>А в углу стояла Дубина!!!"
						}
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			case Hangar:
				EventNow.TextAll = "Ты находишься возле заброшенного ангара!<br>Страшно, темно, но тут можно спокойно провести ночь!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Его обнесли под чистую еще задолго до твоего появления!"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			case Camp:
				EventNow.TextAll = "Ты стоишь возле бывшего лагеря. Угли в костре еще тлеют, но сейчас тут никого нет.<br>Отличное место для ночлега.<br>"
				EventNow.Bonfire = true
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(50) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Cloth += 2
						Gamer.Backpack.Rod += 1
						Gamer.Backpack.Wood += 2
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Прямо возле костра лежало:<br>2 тряпки<br>1 прутик<br>и 2 дерева"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно попытаться обыскать местность, развести костер и переночевать.<br>Цени такие моменты - это большая редкость, тем более ночью!"
					printGamerParams()
				}
			}
		}
	case false: // Дневное событие
		if LengthOfDay != 0 {
			LengthOfDay-- // уменьшаем дневное время
		} else {
			NightMode = true
			LengthOfNight = indexForDifficulty[GenerateLengthOfNight][Difficulty]
		}
		switch EventNow.Place {
		case Forest:
			if FalseOrTrue(50) {
				EventNow.TextAll = "Ты забрел в дремучий лес!<br>"
				if FalseOrTrue(50) {
					EventNow.TextAll += "Здесь очень страшно и слышится жуткий вой...<br>"
				} else {
					EventNow.TextAll += "От него веет сыростью и холодом...<br>"
				}
			} else {
				EventNow.TextAll = "Ты подошел к красивому и светлому лесу!<br>"
				if FalseOrTrue(50) {
					EventNow.TextAll += "Здесь можно попробовать найти припасы.<br>"
				} else {
					EventNow.TextAll += "В нем даже переночевать не страшно!<br>"
				}
			}
			switch EventNow.TypeEvent { // Перебор типа события
			case Attack: // Событие - атака
				if !EventNow.setEnemy(75) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find: // Событие - находка
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Только вот он полностью вычищен..."
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
				printGamerParams()
			}
		case Plain:
			EventNow.TextAll = "Ты подошел к равнине!<br>Хороший обзор во все стороны, сильный ветер и мелкие кучки различного мусора"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(70) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find: // Событие - находка
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Splint += 1
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Кто-то спрятал в нем шину от перелома"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
				printGamerParams()
			}
		case Road:
			EventNow.TextAll = "Ты оказался прямо на автомобильной дороге<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(50) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Cloth += 1
					Gamer.Backpack.Bottle.Water += 1
					Gamer.Backpack.Food.RawMeat += 1
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало много припасов:<br>1 тряпка<br>1 бутылка с водой<br>1 сырое мясо"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
				printGamerParams()
			}
		case Swamp:
			EventNow.TextAll = "Ты попал на болотные топи...<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(80) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>На что ты рассчитывал? Тут уже давно все выпотрошили!"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
				printGamerParams()
			}
		case Riverside:
			EventNow.TextAll = "Ты подошел к живописному берегу бурной реки.<br>Не упусти шанс и наполни свои пустые бутылки чистой водой!<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(65) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Food.Fruits += 1
					Gamer.Backpack.Food.Vegetables += 1
					EventNow.TextAll += "<h3>Ты увидел плодово-ягодную плантацию у реки!!!</h3>Нашел 1 фрукт и овощ!"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
				printGamerParams()
			}
		case Bridge:
			EventNow.TextAll = "Ты пришел на старый обветшалый мост!<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(60) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					Gamer.Backpack.Bandage += 1
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежал бинт"
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
				printGamerParams()
			}
		case Desert:
			EventNow.TextAll = "Ты оказался в знойной пустыне. <br>Зачем пришел сюда? Что ты тут забыл?<br>"
			switch EventNow.TypeEvent {
			case Attack:
				if !EventNow.setEnemy(60) { // Враги кончились
					EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
					EventNow.TypeEvent = "Find"
				}
			case Find:
				if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
					EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Или это был просто очередной мираж?..."
					printGamerParams()
				} else {
					if !EventNow.setEnemyFind(75) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						printGamerParams()
					}
				}
			case Null: // Ничего не произошло
				EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
				printGamerParams()
			}
		case Building:
			switch EventNow.Building {
			case Hospital:
				EventNow.TextAll = "Ты оказался у госпиталя!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(30) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Cloth += 2
						Gamer.Backpack.Bottle.Alcohol += 3
						Gamer.Backpack.Medicament += 1
						Gamer.Backpack.Bandage += 1
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало ОЧЕНЬ много припасов:<br>2 тряпки<br>3 бутылки с алкоголем<br>2 лекарства<br>1 бинт"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			case PoliceStation:
				EventNow.TextAll = "Ты подошел к полицейскому участку!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(30) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						takeWeapon := false
						if Gamer.Gear.Name == "" {
							Gamer.Gear = *Weapons[Spear]
							takeWeapon = true
						}
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>"
						if takeWeapon {
							EventNow.TextAll += "В нем лежало оружие - копье!!!"
						} else {
							EventNow.TextAll += "Увы, но там оказалось пусто..."
						}
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			case FoodShop:
				EventNow.TextAll = "Ты пришел в продуктовый магазин!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Food.Fruits += 2
						Gamer.Backpack.Food.Vegetables += 1
						Gamer.Backpack.Food.RawMeat += 1
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало ОЧЕНЬ много припасов:<br>2 фрукта<br>2 овощ<br>и 1 сырой кусок мяса"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			case HardwareStore:
				EventNow.TextAll = "Ты оказался рядом со строительным магазином!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>И снова пусто.... В другой раз повезет!"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			case DepartmentStore:
				EventNow.TextAll = "Ты возле универсального магазина!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Bottle.Water += 1
						Gamer.Backpack.Food.RoastedMeat += 1
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежала бутылка с водой и стейк"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			case House:
				EventNow.TextAll = "Ты подошел к чьему-то бывшему дому!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(35) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						takeWeapon := false
						if Gamer.Gear.Name == "" {
							Gamer.Gear = *Weapons[BigStick]
							takeWeapon = true
						}
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>Его полностью обчистили прямо перед твоим приходом..."
						if takeWeapon {
							EventNow.TextAll += "<br>Но дубину почему-то не тронули!!!"
						}
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			case Hangar:
				EventNow.TextAll = "Ты находишься возле заброшенного ангара!<br>"
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(40) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						Gamer.Backpack.Cloth += 1
						Gamer.Backpack.Rod += 2
						Gamer.Backpack.Wood += 2
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>В нем лежало много припасов:<br>1 тряпка<br>2 прутика<br>и 2 дерева"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			case Camp:
				EventNow.TextAll = "Ты стоишь возле бывшего лагеря. Угли в костре еще тлеют!<br>"
				EventNow.Bonfire = true
				switch EventNow.TypeEvent {
				case Attack:
					if !EventNow.setEnemy(50) { // Враги кончились
						EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
						EventNow.TypeEvent = "Find"
					}
				case Find:
					if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
						EventNow.TextAll += "<h3>Ты увидел тайник!!!</h3>К сожалению он был совсем рядом с костром и поэтому все сгорело!"
						printGamerParams()
					} else {
						if !EventNow.setEnemyFind(75) { // Враги кончились
							EventNow.TextAll += "<br><br>Ты заметил какое-то движение недалеко от своей позиции...<br>Похоже, что тебе показалось, можно двигаться дальше.<br>"
							printGamerParams()
						}
					}
				case Null: // Ничего не произошло
					EventNow.TextAll += "<h3>Ни врагов, ни тайников!</h3>Можно спокойно обыскать местность, отдохнуть и расслабиться<br>Цени такие моменты - это большая редкость!"
					printGamerParams()
				}
			}
		}
	}
}

func (e *Event) generateWays() {
	// Генерация случайных мест куда может попасть игрок
	rand.Seed(time.Now().UnixNano()) // Рандомизируем
	EventNorth.Place = OptionPlaceEvent[rand.Intn(len(OptionPlaceEvent))]
	rand.Seed(time.Now().UnixNano()) // Рандомизируем
	EventSouth.Place = OptionPlaceEvent[rand.Intn(len(OptionPlaceEvent))]
	rand.Seed(time.Now().UnixNano()) // Рандомизируем
	EventWest.Place = OptionPlaceEvent[rand.Intn(len(OptionPlaceEvent))]
	rand.Seed(time.Now().UnixNano()) // Рандомизируем
	EventEast.Place = OptionPlaceEvent[rand.Intn(len(OptionPlaceEvent))]
}

func (e *Event) setTypeEvent() {
	if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
		EventNow.TypeEvent = Find
	} else {
		if FalseOrTrue(indexForDifficulty[PercentTypeEvent][Difficulty]) {
			EventNow.TypeEvent = Attack
		} else {
			EventNow.TypeEvent = Null
		}
	}
}

func (e *Event) setWays() WaySelectData {
	data := WaySelectData{}
	switch PrevSelectWay {
	case North:
		data.One = West
		data.Two = East
		data.Three = North
		data.OneText = EventWest.Place
		data.TwoText = EventEast.Place
		data.ThreeText = EventNorth.Place
	case South:
		data.One = South
		data.Two = West
		data.Three = East
		data.OneText = EventSouth.Place
		data.TwoText = EventWest.Place
		data.ThreeText = EventEast.Place
	case West:
		data.One = North
		data.Two = West
		data.Three = South
		data.OneText = EventNorth.Place
		data.TwoText = EventWest.Place
		data.ThreeText = EventSouth.Place
	case East:
		data.One = North
		data.Two = East
		data.Three = South
		data.OneText = EventNorth.Place
		data.TwoText = EventEast.Place
		data.ThreeText = EventSouth.Place
	default:
		data.One = West
		data.Two = East
		data.Three = North
		data.OneText = EventWest.Place
		data.TwoText = EventEast.Place
		data.ThreeText = EventNorth.Place
	}
	return data
}

// Функция установки врага в событие
// percentAnimal - это % того, что врагом будет животное
// возвращает true в случае успеха и false если врагов не осталось
func (e *Event) setEnemy(percentAnimal float64) bool {
	if FalseOrTrue(percentAnimal) {
		if AmountAnimal >= 1 {
			for i := 0; i < len(EnemyAnimal); i++ {
				if EnemyAnimal[i].Predator {
					EventNow.EnemyAnimal = EnemyAnimal[i]
					EventNow.TextAll += "<h3>На тебя напал дикий зверь - " + EventNow.EnemyAnimal.Name + "</h3>" +
						"<br>Здоровье: " + strconv.FormatFloat(EventNow.EnemyAnimal.Heals, 'f', 0, 64) +
						"<br>Показатель атаки: " + strconv.FormatFloat(EventNow.EnemyAnimal.IndexAttack, 'f', 0, 64) +
						"<br>Показатель защиты: " + strconv.FormatFloat(EventNow.EnemyAnimal.IndexDefense, 'f', 0, 64)
					printGamerParams()
					return true
				}
			}
		} else {
			return false
		}
	} else {
		if AmountHuman >= 1 {
			EventNow.EnemyHuman = EnemyHuman[0]
			EventNow.TextAll += "<h3>На тебя напал другой выживший.</h3>Ты узнаешь этого человека, это " +
				EventNow.EnemyHuman.Name +
				"<br>Здоровье: " + strconv.FormatFloat(EventNow.EnemyHuman.Heals, 'f', 0, 64) +
				"<br>Показатель атаки: " + strconv.FormatFloat(EventNow.EnemyHuman.IndexAttack, 'f', 0, 64) +
				"<br>Показатель защиты: " + strconv.FormatFloat(EventNow.EnemyHuman.IndexDefense, 'f', 0, 64)
			printGamerParams()
			return true
		} else {
			return false
		}
	}
	return false
}

// Функция установки найденной цели в событие
// percentAnimal - это % того, что целью будет животное
// возвращает true в случае успеха и false если врагов не осталось
func (e *Event) setEnemyFind(percentAnimal float64) bool {
	if FalseOrTrue(percentAnimal) {
		if AmountAnimal >= 1 {
			for i := 0; i < len(EnemyAnimal); i++ {
				EventNow.EnemyAnimal = EnemyAnimal[i]
				EventNow.TextAll += "<h3>Ты незаметно подкрался к " + EventNow.EnemyAnimal.Name + "</h3>" +
					"Здоровье: " + strconv.FormatFloat(EventNow.EnemyAnimal.Heals, 'f', 0, 64) +
					"<br>Показатель атаки: " + strconv.FormatFloat(EventNow.EnemyAnimal.IndexAttack, 'f', 0, 64) +
					"<br>Показатель защиты: " + strconv.FormatFloat(EventNow.EnemyAnimal.IndexDefense, 'f', 0, 64)
				printGamerParams()
				return true
			}
		} else {
			return false
		}
	} else {
		if AmountHuman >= 1 {
			EventNow.EnemyHuman = EnemyHuman[0]
			EventNow.TextAll += "<h3>Ты незаметно подкрался к другому выжившему.</h3>Ты узнаешь этого человека, это " +
				EventNow.EnemyHuman.Name +
				"<br>Здоровье: " + strconv.FormatFloat(EventNow.EnemyHuman.Heals, 'f', 0, 64) +
				"<br>Показатель атаки: " + strconv.FormatFloat(EventNow.EnemyHuman.IndexAttack, 'f', 0, 64) +
				"<br>Показатель защиты: " + strconv.FormatFloat(EventNow.EnemyHuman.IndexDefense, 'f', 0, 64)
			printGamerParams()
			return true
		} else {
			return false
		}
	}
	return false
}
