package objects

import (
	"math"
	"math/rand"
	"time"
)

// Initialization - Инициализация объектов в игре
func Initialization(difficulty string) {
	Difficulty = difficulty         // Установка уровня сложности для игры
	initWeapon()                    // Создание объектов оружия для игры
	Gamer.NewHuman(true, GamerName) // Создаем игрока
	initObjectAnimal()              // Создание объектов животных
	GenerateEnemy()                 // Создаем врагов и животных в игре
	initFirstEvent()                // Генерация первого события
	LengthOfDay = indexForDifficulty[GenerateLengthOfDay][Difficulty]
	LengthOfNight = indexForDifficulty[GenerateLengthOfNight][Difficulty]
}

// GenerateEnemy - Генерация всех врагов
func GenerateEnemy() {
	var (
		amountAnimalPredator    int // Переменная для определения кол-ва хищных животных
		amountAnimalNotPredator int // Переменная для определения кол-ва травоядных животных
		randomNumber            int // Переменная для случайного числа
	)
	switch Difficulty { // Уровень сложности
	case Easy:
		AmountHuman = int(math.Round(indexForDifficulty[GenerateBotsHuman][Easy]))
		amountAnimalPredator = int(math.Round(indexForDifficulty[GenerateBotsAnimalPredator][Easy]))
		amountAnimalNotPredator = int(math.Round(indexForDifficulty[GenerateBotsAnimalNotPredator][Easy]))
	case Normal:
		AmountHuman = int(math.Round(indexForDifficulty[GenerateBotsHuman][Normal]))
		amountAnimalPredator = int(math.Round(indexForDifficulty[GenerateBotsAnimalPredator][Normal]))
		amountAnimalNotPredator = int(math.Round(indexForDifficulty[GenerateBotsAnimalNotPredator][Normal]))
	case Hard:
		AmountHuman = int(math.Round(indexForDifficulty[GenerateBotsHuman][Hard]))
		amountAnimalPredator = int(math.Round(indexForDifficulty[GenerateBotsAnimalPredator][Hard]))
		amountAnimalNotPredator = int(math.Round(indexForDifficulty[GenerateBotsAnimalNotPredator][Hard]))
	}
	// Создание людей
	for i := 0; i < AmountHuman; i++ {
		EnemyHuman = append(EnemyHuman, &Human{})
		EnemyHuman[i].index = i                                     // Новый объект структуры
		rand.Seed(time.Now().UnixNano())                            // Рандомизируем
		randomNumber = rand.Intn(AmountHuman)                       // Получаем случайное число
		EnemyHuman[i].NewHuman(false, NameEnemyHuman[randomNumber]) // Создаем бота
		// Удаляем использованное имя из массива имен для исключения повторов
		NameEnemyHuman = append(NameEnemyHuman[:randomNumber], NameEnemyHuman[randomNumber+1:]...)
	}
	// Создание травоядных
	for i := 0; i < amountAnimalNotPredator; i++ {
		EnemyAnimal = append(EnemyAnimal, &Animal{})
		EnemyAnimal[i].index = i
		if FalseOrTrue(50) {
			if i > 0 && EnemyAnimal[i-1].Name != Rabbit {
				setParamsAnimal(Rabbit, i)
			} else {
				setParamsAnimal(Pig, i)
			}
		} else {
			if i > 0 && EnemyAnimal[i-1].Name != Pig {
				setParamsAnimal(Pig, i)
			} else {
				setParamsAnimal(Rabbit, i)
			}
		}
	}
	// Создание хищников
	for i := amountAnimalNotPredator; i < amountAnimalNotPredator+amountAnimalPredator; i++ {
		EnemyAnimal = append(EnemyAnimal, &Animal{})
		EnemyAnimal[i].index = i
		if FalseOrTrue(60) {
			if i > 0 && EnemyAnimal[i-1].Name != Wolf {
				setParamsAnimal(Wolf, i)
			} else {
				setParamsAnimal(Bear, i)
			}
		} else {
			if i > 0 && EnemyAnimal[i-1].Name != Bear {
				setParamsAnimal(Bear, i)
			} else {
				setParamsAnimal(Wolf, i)
			}
		}
	}
	AmountAnimal = amountAnimalNotPredator + amountAnimalPredator
}

// Установка параметров животных в зависимости от их типа
func setParamsAnimal(animal string, index int) {
	EnemyAnimal[index].Name = ObjectAnimal[animal].Name
	EnemyAnimal[index].Predator = ObjectAnimal[animal].Predator
	EnemyAnimal[index].Heals = ObjectAnimal[animal].Heals
	EnemyAnimal[index].IndexAttack = ObjectAnimal[animal].IndexAttack
	EnemyAnimal[index].IndexDefense = ObjectAnimal[animal].IndexDefense
	EnemyAnimal[index].Debuff = ObjectAnimal[animal].Debuff
	EnemyAnimal[index].Reward = ObjectAnimal[animal].Reward
}

// Создание объектов оружия для игры
func initWeapon() {
	// Копье
	Weapons[Spear] = &Gear{}
	Weapons[Spear].Name = Spear
	Weapons[Spear].IndexAttack = 35.0
	Weapons[Spear].IndexDefense = 3.5
	Weapons[Spear].Durability = 5
	// Дубина
	Weapons[BigStick] = &Gear{}
	Weapons[BigStick].Name = BigStick
	Weapons[BigStick].IndexAttack = 15.0
	Weapons[BigStick].IndexDefense = 0.5
	Weapons[BigStick].Durability = 3
	// Щит
	Weapons[Shield] = &Gear{}
	Weapons[Shield].Name = Shield
	Weapons[Shield].IndexAttack = 1.0
	Weapons[Shield].IndexDefense = 15
	Weapons[Shield].Durability = 5
	// Праща
	Weapons[Slingshot] = &Gear{}
	Weapons[Slingshot].Name = Slingshot
	Weapons[Slingshot].IndexAttack = 40.0
	Weapons[Slingshot].IndexDefense = 1.5
	Weapons[Slingshot].Durability = 5
}

// Создание объектов животных
func initObjectAnimal() {
	// Объект - Волк
	ObjectAnimal[Wolf] = &Animal{}
	ObjectAnimal[Wolf].Name = Wolf
	ObjectAnimal[Wolf].Predator = true
	ObjectAnimal[Wolf].Heals = 70.0
	ObjectAnimal[Wolf].IndexAttack = 20.0
	ObjectAnimal[Wolf].IndexDefense = 15.0
	ObjectAnimal[Wolf].Reward.Food.RawMeat = 2
	ObjectAnimal[Wolf].Reward.Skin = 2
	ObjectAnimal[Wolf].Debuff.Cut, ObjectAnimal[Wolf].Debuff.BrokenBone, ObjectAnimal[Wolf].Debuff.Disease = false, false, false
	// Объект - Медведь
	ObjectAnimal[Bear] = &Animal{}
	ObjectAnimal[Bear].Name = Bear
	ObjectAnimal[Bear].Predator = true
	ObjectAnimal[Bear].Heals = 100.0
	ObjectAnimal[Bear].IndexAttack = 30.0
	ObjectAnimal[Bear].IndexDefense = 20.0
	ObjectAnimal[Bear].Reward.Food.RawMeat = 3
	ObjectAnimal[Bear].Reward.Skin = 3
	ObjectAnimal[Bear].Reward.Medicament = 1
	ObjectAnimal[Bear].Reward.Splint = 1
	ObjectAnimal[Bear].Debuff.Cut, ObjectAnimal[Bear].Debuff.BrokenBone, ObjectAnimal[Bear].Debuff.Disease = false, false, false
	// Объект - Кабан
	ObjectAnimal[Pig] = &Animal{}
	ObjectAnimal[Pig].Name = Pig
	ObjectAnimal[Pig].Predator = false
	ObjectAnimal[Pig].Heals = 40.0
	ObjectAnimal[Pig].IndexAttack = 10.0
	ObjectAnimal[Pig].IndexDefense = 5.0
	ObjectAnimal[Pig].Reward.Food.RawMeat = 2
	ObjectAnimal[Pig].Reward.Skin = 1
	ObjectAnimal[Pig].Debuff.Cut, ObjectAnimal[Pig].Debuff.BrokenBone, ObjectAnimal[Pig].Debuff.Disease = false, false, false
	// Объект - Заяц
	ObjectAnimal[Rabbit] = &Animal{}
	ObjectAnimal[Rabbit].Name = Rabbit
	ObjectAnimal[Rabbit].Predator = false
	ObjectAnimal[Rabbit].Heals = 25.0
	ObjectAnimal[Rabbit].IndexAttack = 0.0
	ObjectAnimal[Rabbit].IndexDefense = 2.0
	ObjectAnimal[Rabbit].Reward.Food.RawMeat = 1
	ObjectAnimal[Rabbit].Reward.Skin = 1
	ObjectAnimal[Rabbit].Debuff.Cut, ObjectAnimal[Rabbit].Debuff.BrokenBone, ObjectAnimal[Rabbit].Debuff.Disease = false, false, false
}

// Создание первого события
func initFirstEvent() {
	rand.Seed(time.Now().UnixNano()) // Рандомизируем
	// Генерация случайных мест куда может попасть игрок
	EventNow.Place = OptionPlaceEvent[rand.Intn(len(OptionPlaceEvent))]
	EventNow.TextAll = "Произошло что-то страшное, ты не помнишь что именно и где ты находишься...<br>Ты очнулся и видишь, что перед тобой " + EventNow.Place
	if EventNow.Place == Building { // Если место является зданием
		rand.Seed(time.Now().UnixNano()) // Рандомизируем
		//Выбираем случайное здание
		EventNow.Building = OptionsBuildingEvent[rand.Intn(len(OptionsBuildingEvent))]
		EventNow.TextAll += "<br>Это " + EventNow.Building
	}
	EventNow.Gamer = Gamer
}
