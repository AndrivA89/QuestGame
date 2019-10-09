package objects

import (
	"strconv"
)

// Battle - функция боя между Attacker и Target
func Battle(Attacker interface{}, Target interface{}) {
	damageToTarget := 0.0      // Урон защитнику
	damageToAttacker := 0.0    // Урон атакующему
	indexDebuffAttacker := 1.0 // Ослабляющие факторы атакующего
	indexDebuffTarget := 1.0   // Ослабляющие факторы цели
	switch Attacker.(type) {
	case *Human: // Атакующий - Человек
		switch Target.(type) {
		case *Human: // Человек атакует Человека
			indexDebuffAttacker = indexDebuffObject(Attacker.(*Human))
			indexDebuffTarget = indexDebuffObject(Target.(*Human))
			damageToTarget = Attacker.(*Human).IndexAttack*indexDebuffAttacker + Attacker.(*Human).Gear.IndexAttack*indexDebuffAttacker - Attacker.(*Human).Fatigue/7 - Target.(*Human).IndexDefense*indexDebuffTarget - Target.(*Human).Gear.IndexDefense*indexDebuffTarget
			if damageToTarget < 0.0 {
				damageToTarget = 0.0
			}
			Target.(*Human).Heals -= round(damageToTarget)
			Target.(*Human).Fatigue += 5.0
			// Определение негативного статуса цели
			if damageToTarget > 21 && (Target.(*Human).Debuff.Cut || Target.(*Human).Debuff.BrokenBone) && FalseOrTrue(50) {
				Target.(*Human).Debuff.Disease = true
			}
			if Target.(*Human).Debuff.Cut && Attacker.(*Human).Debuff.Disease && damageToTarget > 15 {
				Target.(*Human).Debuff.Disease = true
			}
			if damageToTarget > 15 && FalseOrTrue(40) {
				Target.(*Human).Debuff.Cut = true
			}
			if damageToTarget > 20 && FalseOrTrue(40) {
				Target.(*Human).Debuff.BrokenBone = true
			}
			damageToAttacker = (Target.(*Human).IndexAttack*indexDebuffTarget+Target.(*Human).Gear.IndexAttack*indexDebuffTarget)/2 - Target.(*Human).Fatigue/7 - Attacker.(*Human).IndexDefense*indexDebuffAttacker - Attacker.(*Human).Gear.IndexDefense*indexDebuffAttacker
			if damageToAttacker < 0.0 {
				damageToAttacker = 0.0
			}
			Attacker.(*Human).Heals -= round(damageToAttacker)
			if FalseOrTrue(40) {
				Attacker.(*Human).Fatigue += 5.0
			} else {
				Attacker.(*Human).Fatigue += 10.0
			}
			// Определение негативного статуса атакующего
			if Attacker.(*Human).Debuff.Cut && Target.(*Human).Debuff.Disease && damageToAttacker > 15 {
				Attacker.(*Human).Debuff.Disease = true
			}
			if damageToAttacker > 15 && FalseOrTrue(40) {
				Attacker.(*Human).Debuff.Cut = true
			}
			if damageToAttacker > 20 && FalseOrTrue(40) {
				Attacker.(*Human).Debuff.BrokenBone = true
			}
			// Уменьшение прочности оружия атакующего
			if Attacker.(*Human).Gear.Durability--; Attacker.(*Human).Gear.Durability <= 0 {
				Attacker.(*Human).Gear = Gear{}
			}
			// Уменьшение прочности оружия цели
			if Target.(*Human).Gear.Durability--; Target.(*Human).Gear.Durability <= 0 {
				Target.(*Human).Gear = Gear{}
			}
			if Target.(*Human).index < 0 { // Если Target это Игрок
				checkAfterBattle(Attacker.(*Human))
			} else {
				checkAfterBattle(Target.(*Human))
			}
		case *Animal: // Человек атакует Животное
			indexDebuffAttacker = indexDebuffObject(Attacker.(*Human))
			indexDebuffTarget = indexDebuffObject(Target.(*Animal))
			damageToTarget = Attacker.(*Human).IndexAttack*indexDebuffAttacker + Attacker.(*Human).Gear.IndexAttack*indexDebuffAttacker - Attacker.(*Human).Fatigue/7 - Target.(*Animal).IndexDefense*indexDebuffTarget
			if damageToTarget < 0.0 {
				damageToTarget = 0.0
			}
			Target.(*Animal).Heals -= round(damageToTarget)
			// Определение негативного статуса цели
			if damageToTarget > 21 && (Target.(*Animal).Debuff.Cut || Target.(*Animal).Debuff.BrokenBone) && FalseOrTrue(50) {
				Target.(*Animal).Debuff.Disease = true
			}
			if Target.(*Animal).Debuff.Cut && Attacker.(*Human).Debuff.Disease && damageToTarget > 15 {
				Target.(*Animal).Debuff.Disease = true
			}
			if damageToTarget > 15 && FalseOrTrue(40) {
				Target.(*Animal).Debuff.Cut = true
			}
			if damageToTarget > 20 && FalseOrTrue(40) {
				Target.(*Animal).Debuff.BrokenBone = true
			}
			damageToAttacker = (Target.(*Animal).IndexAttack*indexDebuffTarget)/2 - Attacker.(*Human).IndexDefense*indexDebuffAttacker - Attacker.(*Human).Gear.IndexDefense*indexDebuffAttacker
			if damageToAttacker < 0.0 {
				damageToAttacker = 0.0
			}
			Attacker.(*Human).Heals -= round(damageToAttacker)
			if FalseOrTrue(40) {
				Attacker.(*Human).Fatigue += 5.0
			} else {
				Attacker.(*Human).Fatigue += 10.0
			}
			// Определение негативного статуса атакующего
			if Attacker.(*Human).Debuff.Cut && Target.(*Animal).Debuff.Disease && damageToAttacker > 15 {
				Attacker.(*Human).Debuff.Disease = true
			}
			if damageToAttacker > 15 && FalseOrTrue(40) {
				Attacker.(*Human).Debuff.Cut = true
			}
			if damageToAttacker > 20 && FalseOrTrue(40) {
				Attacker.(*Human).Debuff.BrokenBone = true
			}
			// Уменьшение прочности оружия атакующего
			if Attacker.(*Human).Gear.Durability--; Attacker.(*Human).Gear.Durability <= 0 {
				Attacker.(*Human).Gear = Gear{}
			}
			checkAfterBattle(Target.(*Animal))
		}
	case *Animal: // Атакующий - Животное
		switch Target.(type) {
		case *Human: // Животное атакует человека
			indexDebuffAttacker = indexDebuffObject(Attacker.(*Animal))
			indexDebuffTarget = indexDebuffObject(Target.(*Human))
			damageToTarget = Attacker.(*Animal).IndexAttack*indexDebuffAttacker - Target.(*Human).IndexDefense*indexDebuffTarget
			if damageToTarget < 0.0 {
				damageToTarget = 0.0
			}
			Target.(*Human).Heals -= round(damageToTarget)
			Target.(*Human).Fatigue += 5.0
			// Определение негативного статуса цели
			if damageToTarget > 21 && (Target.(*Human).Debuff.Cut || Target.(*Human).Debuff.BrokenBone) && FalseOrTrue(50) {
				Target.(*Human).Debuff.Disease = true
			}
			if Target.(*Human).Debuff.Cut && Attacker.(*Animal).Debuff.Disease && damageToTarget > 15 {
				Target.(*Human).Debuff.Disease = true
			}
			if damageToTarget > 15 && FalseOrTrue(40) {
				Target.(*Human).Debuff.Cut = true
			}
			if damageToTarget > 20 && FalseOrTrue(40) {
				Target.(*Human).Debuff.BrokenBone = true
			}
			damageToAttacker = (Target.(*Human).IndexAttack*indexDebuffTarget+Target.(*Human).Gear.IndexAttack*indexDebuffTarget)/2 - Target.(*Human).Fatigue/7 - Attacker.(*Animal).IndexDefense*indexDebuffAttacker
			if damageToAttacker < 0.0 {
				damageToAttacker = 0.0
			}
			Attacker.(*Animal).Heals -= round(damageToAttacker)
			// Определение негативного статуса атакующего
			if Attacker.(*Animal).Debuff.Cut && Target.(*Human).Debuff.Disease && damageToAttacker > 15 {
				Attacker.(*Animal).Debuff.Disease = true
			}
			if damageToAttacker > 15 && FalseOrTrue(40) {
				Attacker.(*Animal).Debuff.Cut = true
			}
			if damageToAttacker > 20 && FalseOrTrue(40) {
				Attacker.(*Animal).Debuff.BrokenBone = true
			}
			// Уменьшение прочности оружия цели
			if Target.(*Human).Gear.Durability--; Target.(*Human).Gear.Durability <= 0 {
				Target.(*Human).Gear = Gear{}
			}
		}
		checkAfterBattle(Attacker.(*Animal))
	}
}

// Определение индекса ослабления объекта
func indexDebuffObject(object interface{}) float64 {
	indexDebuff := 1.0 // Ослабляющие факторы
	switch object.(type) {
	case *Human:
		if object.(*Human).Debuff.Cut {
			indexDebuff -= 0.15
		}
		if object.(*Human).Debuff.BrokenBone {
			indexDebuff -= 0.35
		}
		if object.(*Human).Debuff.Disease {
			indexDebuff -= 0.2
		}
		return indexDebuff
	case *Animal:
		if object.(*Animal).Debuff.Cut {
			indexDebuff -= 0.15
		}
		if object.(*Animal).Debuff.BrokenBone {
			indexDebuff -= 0.35
		}
		if object.(*Animal).Debuff.Disease {
			indexDebuff -= 0.2
		}
		return indexDebuff
	default:
		return indexDebuff
	}
}

func takeReward(object interface{}) {
	switch object.(type) {
	case *Human:
		EventNow.TextAll = "<h2>Ты убил " + EventNow.EnemyHuman.Name +
			" и забрал добычу!</h2>"
		Gamer.Backpack.Wood += object.(*Human).Backpack.Wood
		if object.(*Human).Backpack.Wood != 0 {
			EventNow.TextAll += "<br>Дерево - " + strconv.Itoa(object.(*Human).Backpack.Wood) + "шт"
		}
		Gamer.Backpack.Stone += object.(*Human).Backpack.Stone
		if object.(*Human).Backpack.Stone != 0 {
			EventNow.TextAll += "<br>Камень - " + strconv.Itoa(object.(*Human).Backpack.Stone) + "шт"
		}
		Gamer.Backpack.Splint += object.(*Human).Backpack.Splint
		if object.(*Human).Backpack.Splint != 0 {
			EventNow.TextAll += "<br>Шина для перелома - " + strconv.Itoa(object.(*Human).Backpack.Splint) + "шт"
		}
		Gamer.Backpack.Skin += object.(*Human).Backpack.Skin
		if object.(*Human).Backpack.Skin != 0 {
			EventNow.TextAll += "<br>Шкура - " + strconv.Itoa(object.(*Human).Backpack.Skin) + "шт"
		}
		Gamer.Backpack.Rod += object.(*Human).Backpack.Rod
		if object.(*Human).Backpack.Rod != 0 {
			EventNow.TextAll += "<br>Прутик - " + strconv.Itoa(object.(*Human).Backpack.Rod) + "шт"
		}
		Gamer.Backpack.Medicament += object.(*Human).Backpack.Medicament
		if object.(*Human).Backpack.Medicament != 0 {
			EventNow.TextAll += "<br>Лекарство - " + strconv.Itoa(object.(*Human).Backpack.Medicament) + "шт"
		}
		Gamer.Backpack.Bandage += object.(*Human).Backpack.Bandage
		if object.(*Human).Backpack.Bandage != 0 {
			EventNow.TextAll += "<br>Бинт - " + strconv.Itoa(object.(*Human).Backpack.Bandage) + "шт"
		}
		Gamer.Backpack.Cloth += object.(*Human).Backpack.Cloth
		if object.(*Human).Backpack.Cloth != 0 {
			EventNow.TextAll += "<br>Тряпка - " + strconv.Itoa(object.(*Human).Backpack.Cloth) + "шт"
		}
		Gamer.Backpack.Food.Fruits += object.(*Human).Backpack.Food.Fruits
		if object.(*Human).Backpack.Food.Fruits != 0 {
			EventNow.TextAll += "<br>Фрукты - " + strconv.Itoa(object.(*Human).Backpack.Food.Fruits) + "шт"
		}
		Gamer.Backpack.Food.RawMeat += object.(*Human).Backpack.Food.RawMeat
		if object.(*Human).Backpack.Food.RawMeat != 0 {
			EventNow.TextAll += "<br>Сырое мясо - " + strconv.Itoa(object.(*Human).Backpack.Food.RawMeat) + "шт"
		}
		Gamer.Backpack.Food.RoastedMeat += object.(*Human).Backpack.Food.RoastedMeat
		if object.(*Human).Backpack.Food.RoastedMeat != 0 {
			EventNow.TextAll += "<br>Стейк - " + strconv.Itoa(object.(*Human).Backpack.Food.RoastedMeat) + "шт"
		}
		Gamer.Backpack.Food.Vegetables += object.(*Human).Backpack.Food.Vegetables
		if object.(*Human).Backpack.Food.Vegetables != 0 {
			EventNow.TextAll += "<br>Овощи - " + strconv.Itoa(object.(*Human).Backpack.Food.Vegetables) + "шт"
		}
		Gamer.Backpack.Bottle.Empty += object.(*Human).Backpack.Bottle.Empty
		if object.(*Human).Backpack.Bottle.Empty != 0 {
			EventNow.TextAll += "<br>Пустая бутылка - " + strconv.Itoa(object.(*Human).Backpack.Bottle.Empty) + "шт"
		}
		Gamer.Backpack.Bottle.Water += object.(*Human).Backpack.Bottle.Water
		if object.(*Human).Backpack.Bottle.Water != 0 {
			EventNow.TextAll += "<br>Бутылка с водой - " + strconv.Itoa(object.(*Human).Backpack.Bottle.Water) + "шт"
		}
		Gamer.Backpack.Bottle.Alcohol += object.(*Human).Backpack.Bottle.Alcohol
		if object.(*Human).Backpack.Bottle.Alcohol != 0 {
			EventNow.TextAll += "<br>Алкоголь - " + strconv.Itoa(object.(*Human).Backpack.Bottle.Alcohol) + "шт"
		}
		if Gamer.Gear.Name == "" && object.(*Human).Gear.Name != "" {
			Gamer.Gear = object.(*Human).Gear
			EventNow.TextAll += "<br>Оружие - " + object.(*Human).Gear.Name
		}
	case *Animal:
		EventNow.TextAll = "<h2>Ты убил " + EventNow.EnemyAnimal.Name +
			" и забрал добычу!</h2>"
		Gamer.Backpack.Wood += object.(*Animal).Reward.Wood
		if object.(*Animal).Reward.Wood != 0 {
			EventNow.TextAll += "<br>Дерево - " + strconv.Itoa(object.(*Animal).Reward.Wood) + "шт"
		}
		Gamer.Backpack.Stone += object.(*Animal).Reward.Stone
		if object.(*Animal).Reward.Stone != 0 {
			EventNow.TextAll += "<br>Камень - " + strconv.Itoa(object.(*Animal).Reward.Stone) + "шт"
		}
		Gamer.Backpack.Splint += object.(*Animal).Reward.Splint
		if object.(*Animal).Reward.Splint != 0 {
			EventNow.TextAll += "<br>Шина для перелома - " + strconv.Itoa(object.(*Animal).Reward.Splint) + "шт"
		}
		Gamer.Backpack.Skin += object.(*Animal).Reward.Skin
		if object.(*Animal).Reward.Skin != 0 {
			EventNow.TextAll += "<br>Шкура - " + strconv.Itoa(object.(*Animal).Reward.Skin) + "шт"
		}
		Gamer.Backpack.Rod += object.(*Animal).Reward.Rod
		if object.(*Animal).Reward.Rod != 0 {
			EventNow.TextAll += "<br>Прутик - " + strconv.Itoa(object.(*Animal).Reward.Rod) + "шт"
		}
		Gamer.Backpack.Medicament += object.(*Animal).Reward.Medicament
		if object.(*Animal).Reward.Medicament != 0 {
			EventNow.TextAll += "<br>Лекарство - " + strconv.Itoa(object.(*Animal).Reward.Medicament) + "шт"
		}
		Gamer.Backpack.Bandage += object.(*Animal).Reward.Bandage
		if object.(*Animal).Reward.Bandage != 0 {
			EventNow.TextAll += "<br>Бинт - " + strconv.Itoa(object.(*Animal).Reward.Bandage) + "шт"
		}
		Gamer.Backpack.Cloth += object.(*Animal).Reward.Cloth
		if object.(*Animal).Reward.Cloth != 0 {
			EventNow.TextAll += "<br>Тряпка - " + strconv.Itoa(object.(*Animal).Reward.Cloth) + "шт"
		}
		Gamer.Backpack.Food.Fruits += object.(*Animal).Reward.Food.Fruits
		if object.(*Animal).Reward.Food.Fruits != 0 {
			EventNow.TextAll += "<br>Фрукты - " + strconv.Itoa(object.(*Animal).Reward.Food.Fruits) + "шт"
		}
		Gamer.Backpack.Food.RawMeat += object.(*Animal).Reward.Food.RawMeat
		if object.(*Animal).Reward.Food.RawMeat != 0 {
			EventNow.TextAll += "<br>Сырое мясо - " + strconv.Itoa(object.(*Animal).Reward.Food.RawMeat) + "шт"
		}
		Gamer.Backpack.Food.RoastedMeat += object.(*Animal).Reward.Food.RoastedMeat
		if object.(*Animal).Reward.Food.RoastedMeat != 0 {
			EventNow.TextAll += "<br>Стейк - " + strconv.Itoa(object.(*Animal).Reward.Food.RoastedMeat) + "шт"
		}
		Gamer.Backpack.Food.Vegetables += object.(*Animal).Reward.Food.Vegetables
		if object.(*Animal).Reward.Food.Vegetables != 0 {
			EventNow.TextAll += "<br>Овощи - " + strconv.Itoa(object.(*Animal).Reward.Food.Vegetables) + "шт"
		}
		Gamer.Backpack.Bottle.Empty += object.(*Animal).Reward.Bottle.Empty
		if object.(*Animal).Reward.Bottle.Empty != 0 {
			EventNow.TextAll += "<br>Пустая бутылка - " + strconv.Itoa(object.(*Animal).Reward.Bottle.Empty) + "шт"
		}
		Gamer.Backpack.Bottle.Water += object.(*Animal).Reward.Bottle.Water
		if object.(*Animal).Reward.Bottle.Water != 0 {
			EventNow.TextAll += "<br>Бутылка с водой - " + strconv.Itoa(object.(*Animal).Reward.Bottle.Water) + "шт"
		}
		Gamer.Backpack.Bottle.Alcohol += object.(*Animal).Reward.Bottle.Alcohol
		if object.(*Animal).Reward.Bottle.Alcohol != 0 {
			EventNow.TextAll += "<br>Алкоголь - " + strconv.Itoa(object.(*Animal).Reward.Bottle.Alcohol) + "шт"
		}
	}
}

// CheckAfterBattle - Проверка после боя
func checkAfterBattle(object interface{}) {
	if Gamer.Heals == 0 {
		GameOver = true
	} else {
		switch object.(type) {
		case *Human:
			if object.(*Human).index >= 0 { // Если object это не Игрок
				if object.(*Human).Heals <= 0 {
					takeReward(object.(*Human))
					DeleteHuman(object.(*Human).index)
				} else {
					EventNow.TextAll = "<br>Враг " + EventNow.EnemyHuman.Name +
						" скрылся.<br>Здоровье - " + strconv.FormatFloat(EventNow.EnemyHuman.Heals, 'f', 0, 64)
					if EventNow.EnemyHuman.Debuff.Cut || EventNow.EnemyHuman.Debuff.BrokenBone || EventNow.EnemyHuman.Debuff.Disease {
						EventNow.TextAll += "<br>Но получил дополнительные увечья:"
						if EventNow.EnemyHuman.Debuff.Cut {
							EventNow.TextAll += "<br>Порез"
						}
						if EventNow.EnemyHuman.Debuff.BrokenBone {
							EventNow.TextAll += "<br>Перелом"
						}
						if EventNow.EnemyHuman.Debuff.Disease {
							EventNow.TextAll += "<br>Болезнь"
						}
					}
					EventNow.TextAll += "<br><br>После боя ты находишься там же.<br>" + EventNow.Place
					if EventNow.Place == Building {
						EventNow.TextAll += " - " + EventNow.Building
					}
				}
			}
		case *Animal:
			if object.(*Animal).Heals <= 0 {
				takeReward(object.(*Animal))
				DeleteAnimal(object.(*Animal).index)
			} else {
				EventNow.TextAll = "<br>Враг " + EventNow.EnemyAnimal.Name +
					" скрылся.<br>Здоровье - " + strconv.FormatFloat(EventNow.EnemyAnimal.Heals, 'f', 0, 64)
				if EventNow.EnemyAnimal.Debuff.Cut || EventNow.EnemyAnimal.Debuff.BrokenBone || EventNow.EnemyAnimal.Debuff.Disease {
					EventNow.TextAll += "<br>Но получил дополнительные увечья:"
					if EventNow.EnemyAnimal.Debuff.Cut {
						EventNow.TextAll += "<br>Порез"
					}
					if EventNow.EnemyAnimal.Debuff.BrokenBone {
						EventNow.TextAll += "<br>Перелом"
					}
					if EventNow.EnemyAnimal.Debuff.Disease {
						EventNow.TextAll += "<br>Болезнь"
					}
				}
			}
			EventNow.TextAll += "<br><br>После боя ты находишься там же.<br>" + EventNow.Place
			if EventNow.Place == Building {
				EventNow.TextAll += " - " + EventNow.Building
			}
		}
	}
	printGamerParams()
}

// DeleteAnimal - Удаление убитого животного из списка врагов
func DeleteAnimal(index int) {
	EnemyAnimal = append(EnemyAnimal[:index], EnemyAnimal[index+1:]...)
	for i := 0; i < len(EnemyAnimal); i++ {
		EnemyAnimal[i].index = i
	}
	AmountAnimal--
}

// DeleteHuman - Удаление убитого человека из списка врагов
func DeleteHuman(index int) {
	EnemyHuman = append(EnemyHuman[:index], EnemyHuman[index+1:]...)
	for i := 0; i < len(EnemyHuman); i++ {
		EnemyHuman[i].index = i
	}
	AmountHuman--
}
