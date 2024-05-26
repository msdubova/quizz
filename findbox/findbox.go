package findbox

import (
	"fmt"
	"quizz/wakeup"
)

func Findbox(c wakeup.Character) {
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
