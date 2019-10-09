package objects

var (
	Gamer         = &Human{} // Создаем пустой объект Human для игрока
	OneStart      = 0        // Переменная для первого запуска
	GamerName     string     // Имя для игрока
	AmountAnimal  int        // Количество животных
	AmountHuman   int        // Количество людей
	LengthOfDay   float64    // Длительность дня
	LengthOfNight float64    // Длительность ночи
	NightMode     bool       // Режим события
	GameOver      bool       // Конец игры
	GameWin       bool       // Победа
)

// Константы и переменные для настройки различных аспектов игры
const (
	Bonfire = "Костер"              // Костер
	Haven   = "Укрытие"             // Укрытие
	Outdoor = "На открытом воздухе" // На открытом воздухе
) // Места отдыха
const (
	Easy   = "Easy"   // Легкий режим
	Normal = "Normal" // Нормальный режим
	Hard   = "Hard"   // Сложный режим
) // Режим игры
const (
	GenerateBotsHuman             = "GenerateBotsHuman"             // Режим генерации людей врагов
	GenerateBotsAnimalPredator    = "GenerateBotsAnimalPredator"    // Режим генерации хищных животных
	GenerateBotsAnimalNotPredator = "GenerateBotsAnimalNotPredator" // Режим генерации травоядных животных
	GenerateBotsParams            = "GenerateBotsParams"            // Режим генерации игрока
	GenerateGamerParams           = "GenerateGamerParams"           // Режим генерации параметров игрока
	PercentItemGamer              = "PercentItemGamer"              // Процент получения доп. параметра игроком
	PercentItemEnemy              = "PercentItemEnemy"              // Процент получения доп. параметра врагом
	PercentTypeEvent              = "PercentTypeEvent"              // Процент на тип события
	GenerateLengthOfDay           = "GenerateLengthOfDay"           // Режим генерации длительности дня
	GenerateLengthOfNight         = "GenerateLengthOfNight"         // Режим генерации длительности ночи
	IndexFatigueAction            = "IndexFatigueAction"            // Индекс усталости при действии
) // Параметры для генерации мира и различных условий игры
const (
	North = "Север"  // Север
	West  = "Запад"  // Запад
	East  = "Восток" // Восток
	South = "Юг"     // Юг
) // Стороны света
const (
	Spear     = "Копье"  // Копье
	Shield    = "Щит"    // Щит
	Slingshot = "Праща"  // Праща
	BigStick  = "Дубина" // Толстая палка
) // Оружие
const (
	Wood        = "Дерево"    // Дерево
	Stone       = "Камень"    // Камень
	Rod         = "Прутик"    // Прутик
	Cloth       = "Ткань"     // Ткань
	Skin        = "Шкура"     // Шкура
	Alcohol     = "Алкоголь"  // Бутылка со спиртом
	Bandage     = "Бинт"      // Бинт
	Splint      = "Шина"      // Шина для перелома
	RawMeat     = "Мясо"      // Сырое мясо
	RoastedMeat = "Стейк"     // Стейк из мяса
	Fruits      = "Фрукты"    // Фрукты
	Vegetables  = "Овощи"     // Овощи
	Medicament  = "Лекарство" // Лекарство
) // Предметы для крафта и находок
const (
	Bear   = "Медведь" // Медведь
	Rabbit = "Кролик"  // Кролик
	Wolf   = "Волк"    // Волк
	Pig    = "Кабан"   // Кабан
) // Возможные животные в игре
const (
	Forest          = "Лес"                       // Лес
	Plain           = "Равнина"                   // Равнина
	Building        = "Строение"                  // Строение
	Swamp           = "Болото"                    // Болото
	Riverside       = "Берег реки"                // Берег реки
	Desert          = "Пустыня"                   // Пустыня
	Bridge          = "Мост"                      // Мост
	Hospital        = "Больница"                  // Больница
	PoliceStation   = "Полицейский участок"       // Полицейский участок
	FoodShop        = "Продовольственный магазин" // Продовольственный магазин
	HardwareStore   = "Строительный магазин"      // Строительный магазин
	DepartmentStore = "Универсальный магазин"     // Универсальный магазин
	House           = "Дом"                       // Дом
	Hangar          = "Ангар"                     // Ангар
	Camp            = "Лагерь"                    // Лагерь
	Road            = "Дорога"                    // Дорога

	Attack = "Attack" // Атака
	Find   = "Find"   // Находка
	Theft  = "Theft"  // Воровство
	Null   = "Null"   // Ничего
) // Места и типы событий

var (
	Difficulty     string                               // Difficulty - Уровень сложности
	PrevSelectWay  string                               // Предыдущий выбор пути Игроком
	NowSelectWay   string                               // Текущий выбор пути
	EnemyHuman     []*Human                             // EnemyHuman - Список ботов людей
	EnemyAnimal    []*Animal                            // EnemyAnimal - Список ботов животных
	Weapons                  = make(map[string]*Gear)   // Weapons - Список оружия в игре
	ObjectAnimal             = make(map[string]*Animal) // Возможные животные в игре
	EventNow                 = &Event{}                 // Переменная для текущего события
	EventNorth               = &Event{}                 // Переменная для события на севере
	EventSouth               = &Event{}                 // Переменная для события на юге
	EventWest                = &Event{}                 // Переменная для события на западе
	EventEast                = &Event{}                 // Переменная для события на востоке
	NameEnemyHuman           = []string{                // NameEnemyHuman - Список имен для ботов людей
		"Jhon", "Сергей", "Василий", "Tom", "Jack", "Donald", "Константин", "Николай", "Артем",
		"Jane", "Екатерина", "Margaret", "Тамара", "Елена", "Hovard", "Jessica", "Максим"}
	// Варианты для генерации места события
	OptionPlaceEvent = []string{Forest, Plain, Building, Swamp, Riverside, Bridge, Road, Desert}
	// Типы событий
	OptionsTypeEvents = []string{Attack, Find, Theft}
	// Варианты строений
	OptionsBuildingEvent = []string{Hospital, PoliceStation, FoodShop,
		HardwareStore, DepartmentStore, House, Hangar, Camp}
	// Рецепты для создания снаряжения
	Craft map[string]map[string]int = map[string]map[string]int{
		Spear:     {Wood: 3, Stone: 1, Rod: 1},
		BigStick:  {Wood: 2, Rod: 1},
		Shield:    {Wood: 3, Cloth: 1},
		Slingshot: {Wood: 1, Rod: 1, Cloth: 1, Stone: 1},
		Cloth:     {Skin: 1},
		Bandage:   {Cloth: 1, Alcohol: 1},
		Splint:    {Wood: 1, Cloth: 1},
		Bonfire:   {Wood: 2, Cloth: 1, Stone: 1},
	}
	// Индексы для уровня сложности
	indexForDifficulty map[string]map[string]float64 = map[string]map[string]float64{
		PercentItemGamer: { // Процент удачи получения доп. предмета Игроком
			Easy:   40.0,
			Normal: 30.0,
			Hard:   20.0,
		},
		PercentItemEnemy: { // Процент удачи получения доп. предмета врагом
			Easy:   20.0,
			Normal: 30.0,
			Hard:   40.0,
		},
		PercentTypeEvent: { // Процент удачи на тип события
			Easy:   55.0,
			Normal: 50.0,
			Hard:   40.0,
		},
		IndexFatigueAction: { // Индекс усталости от действия
			Easy:   0.5,
			Normal: 1.0,
			Hard:   1.5,
		},
		GenerateBotsHuman: { // Количество ботов людей
			Easy:   3.0,
			Normal: 5.0,
			Hard:   7.0,
		},
		GenerateBotsParams: { // Параметры ботов
			Easy:   2.0,
			Normal: 1.8,
			Hard:   1.8,
		},
		GenerateGamerParams: { // Парметры игрока
			Easy:   2.5,
			Normal: 2.0,
			Hard:   1.8,
		},
		GenerateBotsAnimalPredator: { // Количество хищников
			Easy:   2.0,
			Normal: 3.0,
			Hard:   4.0,
		},
		GenerateBotsAnimalNotPredator: { // Количество травоядных
			Easy:   5.0,
			Normal: 3.0,
			Hard:   2.0,
		},
		GenerateLengthOfDay: { // Длительность дня
			Easy:   7.0,
			Normal: 5.0,
			Hard:   3.0,
		},
		GenerateLengthOfNight: { // Длительность ночи
			Easy:   1.0,
			Normal: 2.0,
			Hard:   3.0,
		},
	}
	// Индексы для определения комфорта сна
	indexComfortSleep map[string]map[string]float64 = map[string]map[string]float64{
		Easy: {
			Haven:   2.0,
			Bonfire: 1.5,
			Outdoor: 1.0,
		},
		Normal: {
			Haven:   1.5,
			Bonfire: 1.3,
			Outdoor: 1.0,
		},
		Hard: {
			Haven:   1.2,
			Bonfire: 1.0,
			Outdoor: 0.8,
		},
	}
)

type WaySelectData struct {
	One       string
	OneText   string
	Two       string
	TwoText   string
	Three     string
	ThreeText string
}

// Human - Объект человек
type Human struct {
	index        int      // Индекс для внутренних функций
	Name         string   // Имя персонажа
	Heals        float64  // Показатель здоровья
	Thirst       float64  // Жажда
	Hunger       float64  // Голод
	Fatigue      float64  // Усталость
	IndexAttack  float64  // Показатель атаки
	IndexDefense float64  // Показатель защиты
	Debuff       Debuff   // Негативные статусы
	Backpack     Backpack // Рюкзак
	Gear         Gear     // Оружие
}

// Animal - Объект животное
type Animal struct {
	index        int      // Индекс для внутренних функций
	Name         string   // Название животного
	Predator     bool     // Хищник или травоядное
	Heals        float64  // Показатель здоровья
	IndexAttack  float64  // Показатель атаки
	IndexDefense float64  // Показатель защиты
	Debuff       Debuff   // Негативные статусы
	Reward       Backpack // Награда за убийство животного
}

// Backpack - рюкзак с возможными предметами
type Backpack struct {
	Wood       int    // Дерево
	Stone      int    // Камень
	Rod        int    // Прутик
	Medicament int    // Лекарство
	Cloth      int    // Ткань
	Bandage    int    // Бинт
	Splint     int    // Шина для перелома
	Skin       int    // Шкура
	Food       Food   // Еда
	Bottle     Bottle // Бутылки
}

// Food - структура различной пищи
type Food struct {
	RawMeat     int // Сырое мясо
	RoastedMeat int // Стейк из мяса
	Fruits      int // Фрукты
	Vegetables  int // Овощи
}

// Bottle - структура бутылок с различным наполнением
type Bottle struct {
	Empty   int // Пустая
	Water   int // С водой
	Alcohol int // Со спиртом
}

// Debuff - структура негативных статусов
type Debuff struct {
	Cut        bool // Порез
	BrokenBone bool // Перелом
	Disease    bool // Болезнь
}

// Gear - структура снаряжения
type Gear struct {
	Name         string  // Название
	IndexAttack  float64 // Показатель атаки
	IndexDefense float64 // Показатель брони
	Durability   int     // Прочность
}

// Event - Структура, определяющая событие
type Event struct {
	Place       string  // Место события
	Building    string  // Строение события
	TypeEvent   string  // Тип события
	TextAll     string  // Полный текст события
	EnemyAnimal *Animal // Враг животное
	EnemyHuman  *Human  // Враг человек
	Gamer       *Human  // Игрок
	Bonfire     bool    // Разведен ли костер
	IndexFind   bool    // Осматривал ли ты уже местность
}
