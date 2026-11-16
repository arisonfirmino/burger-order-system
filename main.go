package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	ui "burger-order-system/helpers"
	"burger-order-system/services"
)

func ReadLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInt(reader *bufio.Reader) (int, error) {
	text := ReadLine(reader)
	return strconv.Atoi(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	products := services.GetProducts()

	for {
		ui.ClearScreen()

		fmt.Println("--> BURGER ORDER SYSTEM\n")

		fmt.Println("1. Novo pedido")
		fmt.Println("2. Ver cardápio")
		fmt.Println("3. Consultar pedidos")
		fmt.Println("0. Sair")

		fmt.Print("\nEscolha uma opção: ")
		option, err := ReadInt(reader)

		if err != nil {
			fmt.Println("\nOpção inválida!")
			fmt.Println("\nPressione ENTER para continuar...")
			ReadLine(reader)
			continue
		}

		switch option {
		case 1:
			// Novo pedido
		case 2:
			ui.ClearScreen()

			fmt.Println("BURGER ORDER SYSTEM --> CARDÁPIO\n")

			for i, product := range products {
				fmt.Printf("%d. %s | R$ %.2f\n", i+1, product.Name, product.Price)
				fmt.Printf("%s\n\n", product.Description)
			}

			fmt.Println("\nPressione ENTER para voltar ao menu...")
			ReadLine(reader)
		case 3:
			// Consultar pedidos
		case 0:
			fmt.Println("\nEncerrando o sistema...")
			return
		default:
			fmt.Println("\nOpção inválida!")
			fmt.Println("\nPressione ENTER para continuar...")
			ReadLine(reader)
		}
	}
}
