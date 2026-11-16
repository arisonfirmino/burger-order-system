package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	ui "burger-order-system/helpers"
)

func ReadLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInt(reader *bufio.Reader) (int, error) {
	text := ReadLine(reader)
	return strconv.Atoi(text)
}

func Pause(reader *bufio.Reader) {
	println("\nPressione ENTER para continuar...")
	ReadLine(reader)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		ui.ClearScreen()

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
			// Ver cardápio
		case 3:
			// COnsultar pedidos
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
