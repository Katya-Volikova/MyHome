package main

type Equipment struct {
	name     string
	color    string
	material string
}

func createEquipment(name string,
	color string,
	material string,
) Equipment {
	return Equipment{
		name:     name,
		color:    color,
		material: material,
	}
}

func initializeEquipment() []Equipment {
	var equipment []Equipment

	name := [6]string{"Чайник", "Тостер", "Мультиварка", "Микроволновка", "Ноутбук"}
	color := [6]string{"white", "black", "black", "black", "gray"}
	material := [6]string{"glass", "metal", "metal", "metal", "metal"}
	itemsMock := [][6]string{name, color, material}

	for i := 0; i < 5; i++ {
		equipment = append(equipment, createEquipment(itemsMock[0][i], itemsMock[1][i], itemsMock[2][i]))
	}

	return equipment
}
