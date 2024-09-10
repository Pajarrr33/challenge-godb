package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
    "runtime"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	var selected_menu string
	main_menu := true 
	clearScreen()

	for main_menu {
		fmt.Println("============ Main Menu ========")
		fmt.Println("1.Customer")
		fmt.Println("2.Service")
		fmt.Println("3.Order")
		fmt.Println("4.Exit")
		fmt.Println("===============================")
		fmt.Print("Masukan menu yang anda inginkan :")

		scanner.Scan()
		selected_menu = scanner.Text()

		if selected_menu == "1" {
			customer()
		} else if selected_menu == "2" {
			service()
		} else if selected_menu == "3" {
			order()
		} else if selected_menu == "4" {
			os.Exit(0)
		} else {
			fmt.Println("===========================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("===========================")
		}
	}
}

func customer() {
	var selected_menu string
	customer_menu := true 
	clearScreen()

	for customer_menu {
		fmt.Println("========== Customer Menu ======")
		fmt.Println("1.Create Customer")
		fmt.Println("2.View Of List Customer")
		fmt.Println("3.View Details Customer By ID")
		fmt.Println("4.Update Customer")
		fmt.Println("5.Delete Customer")
		fmt.Println("6.Back to Main Menu")
		fmt.Println("===============================")
		fmt.Print("Masukan menu yang anda inginkan :")

		scanner.Scan()
		selected_menu = scanner.Text()

		if selected_menu == "1" {
			create_customer()
		} else if selected_menu == "2" {
			view_of_list_customer()
		} else if selected_menu == "3" {
			view_details_customer_by_id()
		} else if selected_menu == "4" {
			update_customer()
		} else if selected_menu == "5" {
			delete_customer()
		} else if selected_menu == "6" {
			main()
		}else {
			fmt.Println("===========================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("===========================")
		}
	}
}

func create_customer() {

}

func view_of_list_customer() {

}

func view_details_customer_by_id() {

}

func update_customer() {

}

func delete_customer() {

}

func service() {
	var selected_menu string
	service_menu := true 
	clearScreen()

	for service_menu {
		fmt.Println("========== Service Menu =======")
		fmt.Println("1.Create Service")
		fmt.Println("2.View Of List Service")
		fmt.Println("3.View Details Service By ID")
		fmt.Println("4.Update Service")
		fmt.Println("5.Delete Service")
		fmt.Println("6.Back to Main Menu")
		fmt.Println("===============================")
		fmt.Print("Masukan menu yang anda inginkan :")

		scanner.Scan()
		selected_menu = scanner.Text()

		if selected_menu == "1" {
			create_service()
		} else if selected_menu == "2" {
			view_of_list_service()
		} else if selected_menu == "3" {
			view_details_service_by_id()
		} else if selected_menu == "4" {
			update_service()
		} else if selected_menu == "5" {
			delete_service()
		} else if selected_menu == "6" {
			main()
		}else {
			fmt.Println("=================================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("=================================")
		}
	}
}

func create_service() {

}

func view_of_list_service() {

}

func view_details_service_by_id() {

}

func update_service() {

}

func delete_service() {

}

func order() {
	var selected_menu string
	order_menu := true 
	clearScreen()

	for order_menu {
		fmt.Println("=========== Order Menu ========")
		fmt.Println("1.Create Order")
		fmt.Println("2.Complete Order")
		fmt.Println("3.View Of List Order")
		fmt.Println("4.View Order Details By ID")
		fmt.Println("5.Back to Main Menu")
		fmt.Println("===============================")
		fmt.Print("Masukan menu yang anda inginkan :")

		scanner.Scan()
		selected_menu = scanner.Text()

		if selected_menu == "1" {
			create_order()
		} else if selected_menu == "2" {
			complete_order()
		} else if selected_menu == "3" {
			view_of_list_order()
		} else if selected_menu == "4" {
			view_order_detail_by_id()
		} else if selected_menu == "5" {
			main()
		} else {
			fmt.Println("=================================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("=================================")
		}
	}
}

func create_order() {

}

func complete_order() {

}

func view_of_list_order() {

}

func view_order_detail_by_id() {

}

// clearScreen clears the terminal based on the OS
func clearScreen() {
    var clearCmd *exec.Cmd
    switch runtime.GOOS {
    case "linux", "darwin": // Linux/macOS
        clearCmd = exec.Command("clear")
    case "windows": // Windows
        clearCmd = exec.Command("cmd", "/c", "cls")
    default:
        fmt.Println("Unsupported platform")
        return
    }
    clearCmd.Stdout = os.Stdout
    clearCmd.Run()
}