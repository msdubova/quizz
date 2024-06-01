package story

import (
	"fmt"
	"math/rand"
	"quizz/character"
)

func Wakeup(c character.Character) {
	items := ""
	if c.Bag.Matches {
		items += "Сірники "
	}
	if c.Bag.Torch {
		items += "Ліхтарик "
	}
	if c.Bag.Knife {
		items += "Ніж "
	} else {
		items += "Нічого"
	}

	fmt.Printf("%s прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить: %s\n", c.Name, items)
	fmt.Println("---")

	fmt.Printf("У печері темно, тому %s іде стежкою, яка веде від печери в ліс.\n", c.Name)

	fmt.Println("---")
}

func Meetanimal(c character.Character) {
	fmt.Printf("У лісі %s натикається на мертве тіло дивної тварини. \n Що робити  з твариною❓❓❓➡️ ➡️ ➡️ ➡️ ➡️  \n 1️⃣ Щоб оглянути тварину натисніть 1 на клавуатурі\n 2️⃣ Щоб просто уйти та уникнути неприємностей натисніть 2 \n 3️⃣ Щоб розбудити тварину натисніть 3 \n 4️⃣ Натисніть будь яку іншу клавішу щоб дати персонажу виріщити що робити \n", c.Name)
	var choice int

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Printf("%s підходить до тіла та оглядає його. Він помічає дивний символ на шиї тварини, але нічого більше не помічає.\n", c.Name)

	case 2:
		fmt.Printf("%s обходить тіло, уникаючи зайвих неприємностей.\n", c.Name)
	case 3:
		fmt.Printf("%s спробував роздратувати тіло, але воно ні на що не реагує.\n", c.Name)
	default:

		choice := rand.Intn(3) + 1
		switch choice {
		case 1:
			fmt.Printf("%s підходить до тіла та оглядає його. Він помічає дивний символ на шиї тварини, але нічого більше не помічає.\n", c.Name)

		case 2:
			fmt.Printf("%s обходить тіло, уникаючи зайвих неприємностей.\n", c.Name)
		case 3:
			fmt.Printf("%s спробував роздратувати тіло, але воно ні на що не реагує.\n", c.Name)

		}
	}

	fmt.Printf("Він обирає нічого з цим не робити й іти далі.\n")
	fmt.Println("---")
}

func Findbox(c character.Character) {
	fmt.Printf("Через деякий час %s приходить до безлюдного табору. Він вже втомлений і вирішує відпочити, а не йти далі. У найближчому наметі він знаходить сейф з кодовим замком з одннєй цифри.\n Введіть 1 цифру з клавіатури, щоб відкрити сейф ➡️ ➡️ ➡️ ➡️ ➡️ \n", c.Name)

	for {
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Помилка вводу. Будь ласка, введіть лише цифру.")
			continue
		}

		var code = 5

		if choice != code {
			if choice > code {
				fmt.Println("Код не підходить. Спробуй ще. \n Підказка: код менше, ніж той, що ти вводиш")
				continue

			} else if choice < code {
				fmt.Println("Код не підходить. Спробуй ще. \n Підказка: код більший, ніж той, що ти вводиш")
				continue
			}
		}
		if choice == code {
			fmt.Println("Код підходить. Сейф відкрито.")
			break
		}

	}
	fmt.Printf("---\n")
	fmt.Printf("Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає. %s непритомніє.\n А все могло бути зовсім інакше.\n 🔚", c.Name)
}
