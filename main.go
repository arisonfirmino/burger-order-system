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
			ui.ClearScreen()

			fmt.Println("BURGER ORDER SYSTEM --> NOVO PEDIDO\n")

			fmt.Print("Nome do cliente: ")
			client := ReadLine(reader)

			order := services.CreateOrder(client)

			for {
				ui.ClearScreen()

				fmt.Println("BURGER ORDER SYSTEM --> NOVO PEDIDO\n")

				fmt.Printf("\nCliente: %s | Status: %s\n\n", order.Client, order.Status)

				if len(order.Products) == 0 {
					fmt.Println("Nenhum produto adicionado.")
				}

				for i, product := range order.Products {
					subtotal := product.Product.Price * float64(product.Quantity)

					fmt.Printf("%d. %dx %s | R$ %.2f\n", i+1, product.Quantity, product.Product.Name, subtotal)
				}

				fmt.Printf("\nTotal: R$ %.2f\n", services.CalculateTotal(&order))

				fmt.Println("\nSelecione um produto:\n")

				for i, product := range products {
					fmt.Printf("%d. %s | R$ %.2f\n", i+1, product.Name, product.Price)
				}

				fmt.Println("0. Finalizar pedido")

				fmt.Print("\nProduto: ")
				choice, err := ReadInt(reader)

				if err != nil || choice < 0 || choice > len(products) {
					fmt.Println("\nOpção inválida!")
					fmt.Println("\nPressione ENTER para continuar...")
					ReadLine(reader)
					continue
				}

				if choice == 0 {
					break
				}

				fmt.Print("Quantidade: ")
				quantity, err := ReadInt(reader)

				if err != nil {
					fmt.Println("\nQuantidade inválida!")
					fmt.Println("\nPressione ENTER para continuar...")
					ReadLine(reader)
					continue
				}

				product := products[choice-1]

				if err := services.AddProduct(&order, &product, quantity); err != nil {
					fmt.Printf("\nErro: %s\n", err)
				}
			}

			ui.ClearScreen()

			order.Status = "Finalizado"

			services.SaveOrder(order)
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
			for {
				ui.ClearScreen()

				orders := services.GetOrders()

				fmt.Println("BURGER ORDER SYSTEM --> PEDIDOS\n")

				if len(orders) == 0 {
					fmt.Println("Nenhum pedido encontrado.")
					fmt.Println("\nPressione ENTER para voltar...")
					ReadLine(reader)
					break
				}

				for i, order := range orders {
					fmt.Printf("%d. %s | %s | R$ %.2f\n", i+1, order.Client, order.Status, services.CalculateTotal(&order))
				}

				fmt.Println("0. Voltar")

				fmt.Print("\nSelecione um pedido: ")
				choice, err := ReadInt(reader)

				if err != nil || choice < 0 || choice > len(orders) {
					fmt.Println("\nOpção inválida!")
					fmt.Println("\nPressione ENTER para continuar...")
					ReadLine(reader)
					continue
				}

				if choice == 0 {
					break
				}

				selectedOrder := orders[choice-1]

				ui.ClearScreen()

				fmt.Println("BURGER ORDER SYSTEM --> PEDIDOS --> NOTA FISCAL\n")

				fmt.Printf("Cliente: %s\n", selectedOrder.Client)
				fmt.Printf("Status: %s\n\n", selectedOrder.Status)

				fmt.Println("---------------------------------------------")

				for _, product := range selectedOrder.Products {
					subtotal := product.Product.Price * float64(product.Quantity)

					fmt.Printf("%dx %-20s R$ %.2f\n", product.Quantity, product.Product.Name, subtotal)
				}

				fmt.Println("---------------------------------------------")

				fmt.Printf("\nTOTAL: R$ %.2f\n", services.CalculateTotal(&selectedOrder))

				fmt.Println("\n---------------------------------------------")

				fmt.Println("\nPressione ENTER para voltar aos pedidos...")
				ReadLine(reader)
			}
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
