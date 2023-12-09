package main

import (
	"fmt"
	"time"
)

var database = make(map[string]int)

func AddClient() {
	var (
		name    string
		balance int
	)

	fmt.Scan(&name)
	fmt.Scan(&balance)
	database[name] = balance
	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")

}

func AddMoney(name string) {
	var (
		refill int
	)

	fmt.Scan(&refill)

	balance, exists := database[name]
	if !exists {
		fmt.Println("Клиент не найден")
		return
	}
	database[name] = balance + refill

	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("Итоговая сумма", name, balance+refill)
	fmt.Println("________________")
}

func GiveClientsBalance() {
	for key, val := range database {
		fmt.Println(key, val)
	}

	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")
}

func WithdrawFromBalance(name string) {
	var (
		withdraw int
	)
	fmt.Scan(&withdraw)

	balance, exists := database[name]
	if !exists {
		fmt.Println("Клиент не найден")
		return
	}
	if withdraw <= balance {
		database[name] = balance - withdraw
		fmt.Println("________________")
		fmt.Println("Готово")
		fmt.Println("Итоговая сумма", name, balance-withdraw)
	} else {
		fmt.Println("Извините у вас недостаточно средств для снятия!")
		fmt.Println("________________")
		fmt.Println("Готово")
	}
	fmt.Println("________________")
}

func main() {
	var (
		name string
	)
	for {
		var choice int
		fmt.Println("1. Добавить клиента")
		fmt.Println("2. Пополнить счёт клиента")
		fmt.Println("3. Показать баланс клиента")
		fmt.Println("4. Снять деньги клиента")
		fmt.Println("5. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			AddClient()
		} else if choice == 2 {
			fmt.Println("Имя клиента и сумму пополнения:")
			fmt.Scan(&name)
			AddMoney(name)
		} else if choice == 3 {
			GiveClientsBalance()
		} else if choice == 4 {
			fmt.Println("Имя клиента и сумму снятия:")
			fmt.Scan(&name)
			WithdrawFromBalance(name)
		} else if choice == 5 {
			break
		}
		time.Sleep(2 * time.Second)
	}

}
