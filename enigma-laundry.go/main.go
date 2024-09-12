package main

import (
	"bufio"
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	_ "github.com/lib/pq"
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
			fmt.Println("===============================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("===============================")
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
			fmt.Println("===============================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("===============================")
		}
	}
}

func create_customer() {
	customer := entity.Customer{}

	db := connectDb()        
	defer db.Close()          
	var err error

	fmt.Println("========== Create Customer ======")

	fmt.Print("Masukan id customer     : ")
	scanner.Scan()
	customer.Customer_id, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Print("Masukan name customer   : ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Masukan Phone customer  : ")
	scanner.Scan()
	customer.Phone = scanner.Text()

	fmt.Print("Masukan Adress customer : ")
	scanner.Scan()
	customer.Address = scanner.Text()

	// fill the created_at and updated_at value with time now
	customer.Created_at = time.Now()
	customer.Updated_at = time.Now()

	fmt.Println("=================================")

	// check if id customer is already exist 
	check_id := "SELECT customer_id FROM customer WHERE customer_id = $1;"

	err = db.QueryRow(check_id,customer.Customer_id).Scan(&customer.Customer_id)
	if err == nil {
		fmt.Println("Customer ID already exist. Please enter a different ID")
		return
	}

	// insert customer data into db
	insert_query := "INSERT INTO customer (customer_id, name, phone, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(insert_query, customer.Customer_id, customer.Name, customer.Phone, customer.Address, customer.Created_at, customer.Updated_at)
	if err != nil {
		panic(err)  // Handle error if the query fails
	} else {
		fmt.Println("Successfully added")  // Log success
	}
}

func view_of_list_customer() {
	db := connectDb()        
	defer db.Close()   
	
	customers := []entity.Customer{}

	select_all := "SELECT customer_id,name,phone,address,created_at,updated_at FROM customer;"

	rows, err := db.Query(select_all)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Customer_id, &customer.Name, &customer.Phone, &customer.Address, &customer.Created_at, &customer.Updated_at)
		if err != nil {
			panic(err)  // Handle error during scanning
		}
		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("==== List of all customer =====")
	for _, customer := range customers {
		fmt.Println(customer.Customer_id,customer.Name,customer.Address,customer.Created_at,customer.Updated_at)
	}
	fmt.Println("===============================")
}

func view_details_customer_by_id() {
	db := connectDb()        
	defer db.Close()   
	
	customer := entity.Customer{}

	select_by_id := "SELECT customer_id,name,phone,address,created_at,updated_at FROM customer WHERE customer_id = $1"

	fmt.Print("Insert a customer id :")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	err = db.QueryRow(select_by_id,id).Scan(&customer.Customer_id, &customer.Name, &customer.Phone , &customer.Address, &customer.Created_at, &customer.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("===============================")
			fmt.Println("Customer Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}
	
	fmt.Println("===============================")
	fmt.Println(customer.Customer_id, customer.Name, customer.Phone, customer.Address, customer.Created_at, customer.Updated_at)
	fmt.Println("===============================")
}

func update_customer() {
	db := connectDb()        
	defer db.Close()   

	customer := entity.Customer{}

	fmt.Print("Insert a customer id :")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	select_by_id := "SELECT customer_id,name,phone,address,created_at,updated_at FROM customer WHERE customer_id = $1"
	err = db.QueryRow(select_by_id,id).Scan(&customer.Customer_id, &customer.Name, &customer.Phone , &customer.Address, &customer.Created_at, &customer.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("===============================")
			fmt.Println("Customer Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}

	fmt.Println("======== Update Customer ==========")

	fmt.Print("Enter Customer Name   : ")
	scanner.Scan()
	name := scanner.Text()
	if name != "" {
		customer.Name = name
	}

	fmt.Print("Enter Customer Phone  : ")
	scanner.Scan()
	phone := scanner.Text()
	if phone != "" {
		customer.Phone = phone
	}

	fmt.Print("Enter Customer Address : ")
	scanner.Scan()
	address := scanner.Text()
	if address != "" {
		customer.Address = address
	}

	// fill the updated_at value with time now
	customer.Updated_at = time.Now()

	fmt.Println("=================================")

	update := "UPDATE customer SET name = $2, phone = $3, address = $4, updated_at = $5 WHERE customer_id = $1;"
	_, err = db.Exec(update, id, customer.Name, customer.Phone, customer.Address, customer.Updated_at)
	if err != nil {
		panic(err)  // Handle error if the query fails
	} else {
		fmt.Println("Successfully updated")  // Log success
	}
}

func delete_customer() {
	db := connectDb()        
	defer db.Close()   

	customer := entity.Customer{}

	fmt.Print("Insert a customer id :")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	// check if customer is exist
	select_by_id := "SELECT customer_id FROM customer WHERE customer_id = $1"
	err = db.QueryRow(select_by_id,id).Scan(&customer.Customer_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("===============================")
			fmt.Println("Customer Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}

	// check if customer id exits in order

	order := entity.Order{}

	customer_in_order := `SELECT customer_id FROM "order" WHERE customer_id = $1`

	err = db.QueryRow(customer_in_order,id).Scan(&order.Customer_id)

	if err != nil {
		if err == sql.ErrNoRows {
			delete := "DELETE FROM customer WHERE customer_id = $1"

			_, err = db.Exec(delete, id)
			if err != nil {
				panic(err)  // Handle error if the query fails
			} else {
				fmt.Println("Successfully deleted data")  // Log success
			}
		} else {
			panic(err)
		}
	} else {
		fmt.Println("Customer ID is being used in orders. Please delete the order first.")
	}
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
			fmt.Println("=====================================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("=====================================")
		}
	}
}

func create_service() {
	service := entity.Service{}

	db := connectDb()        
	defer db.Close()   
	var err error 

	fmt.Println("========== Create Service ======")

	fmt.Print("Insert Service Id     : ")
	scanner.Scan()
	service.Service_id, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Print("Insert Service Name   : ")
	scanner.Scan()
	service.Service_name = scanner.Text()

	fmt.Print("Insert Service Unit  : ")
	scanner.Scan()
	service.Unit = scanner.Text()

	fmt.Print("Insert Service Price : ")
	scanner.Scan()
	service.Price, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	// fill the created_at and updated_at value with time now
	service.Created_at = time.Now()
	service.Updated_at = time.Now()

	fmt.Println("=================================")

	// check if id service is already exist 
	check_id := "SELECT service_id FROM service WHERE service_id = $1;"

	err = db.QueryRow(check_id,service.Service_id).Scan(&service.Service_id)
	if err == nil {
		fmt.Println(" Service ID already exists. Please enter a different ID")
		return
	}

	// insert customer data into db
	insert_query := "INSERT INTO service (service_id,service_name,unit,price,created_at,updated_at) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err = db.Exec(insert_query, service.Service_id, service.Service_name, service.Unit, service.Price, service.Created_at, service.Updated_at)
	if err != nil {
		panic(err)  // Handle error if the query fails
	} else {
		fmt.Println("Successfully added")  // Log success
	}
}

func view_of_list_service() {
	db := connectDb()        
	defer db.Close()   
	
	services := []entity.Service{}

	select_all := "SELECT service_id,service_name,unit,price,created_at,updated_at FROM service;"

	rows, err := db.Query(select_all)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		service := entity.Service{}
		err := rows.Scan(&service.Service_id,&service.Service_name,&service.Unit,&service.Price,&service.Created_at,&service.Updated_at)
		if err != nil {
			panic(err)  // Handle error during scanning
		}
		services = append(services, service)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("==== List of all service =====")
	for _, service := range services {
		fmt.Println(service.Service_id,service.Service_name,service.Unit,service.Price,service.Created_at,service.Updated_at)
	}
	fmt.Println("===============================")
}

func view_details_service_by_id() {
	db := connectDb()        
	defer db.Close()   
	
	service := entity.Service{}

	select_by_id := "SELECT service_id,service_name,unit,price,created_at,updated_at FROM service WHERE service_id = $1"

	fmt.Print("Insert a service id :")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	err = db.QueryRow(select_by_id,id).Scan(&service.Service_id,&service.Service_name,&service.Unit,&service.Price,&service.Created_at,&service.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("===============================")
			fmt.Println("Service Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}
	
	fmt.Println("===============================")
	fmt.Println(service.Service_id,service.Service_name,service.Unit,service.Price,service.Created_at,service.Updated_at)
	fmt.Println("===============================")
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
			fmt.Println("=========================")
			fmt.Println("Menu yang anda masukan tidak ada!")
			fmt.Println("=========================")
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
        return
    }
    clearCmd.Stdout = os.Stdout
    clearCmd.Run()
}


func connectDb() *sql.DB {
	// Database connection constants
	const (
		host     = "localhost"  // Host where the database is running
		port     = 5432         // Port where PostgreSQL is listening
		user     = "postgres"   // Database user
		password = "Areman44"   // Database password
		dbname   = "enigma_laundry"      // Database name
	)

	// Connection string for PostgreSQL
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    				host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)  // Open a connection using the connection string

	if err != nil {
		panic(err)  // Handle error if connection fails
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		panic(err)  // Handle error if ping fails
	}

	return db
}