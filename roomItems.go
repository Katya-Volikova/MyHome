package main

type Item struct {
	name     string
	size     string
	color    string
	material string
	amount   string
}

func createItem(name string,
	size string,
	color string,
	material string,
	amount string) Item {
	return Item{
		name:     name,
		size:     size,
		color:    color,
		material: material,
		amount:   amount,
	}
}

func initializeItems() []Item {
	var items []Item

	name := [6]string{"Личный стол", "Кровать", "Обеденный стол", "Холодильник", "Стул", "Шкаф"}
	size := [6]string{"1.5×0.5×1", "2.3×1.5×1.8", "1.2×0.5×1", "0.8×0.8×2.1", "1×1×1", "2×1×1.3"}
	color := [6]string{"white", "white", "white", "white", "yellow", "white"}
	material := [6]string{"metal", "metal", "metal", "metal", "plastic", "tree"}
	amount := [6]string{"3", "3", "3", "1", "6", "6"}
	itemsMock := [][6]string{name, size, color, material, amount}

	for i := 0; i < 6; i++ {
		items = append(items, createItem(itemsMock[0][i], itemsMock[1][i], itemsMock[2][i], itemsMock[3][i], itemsMock[4][i]))
	}

	return items
}
