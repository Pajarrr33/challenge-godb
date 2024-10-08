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
	"github.com/joho/godotenv"
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
		fmt.Print("Insert a menu you want :")

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
			fmt.Println("A menu that you insert is not found")
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
		fmt.Print("Insert a menu you want :")

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
			fmt.Println("A menu that you insert is not found")
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

	fmt.Print("Insert customer id     : ")
	scanner.Scan()
	customer.Customer_id, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Print("Insert customer name   : ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Insert customer Phone  : ")
	scanner.Scan()
	customer.Phone = scanner.Text()

	fmt.Print("Insert customer Adress : ")
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
			fmt.Println("Customer Id Not Found Please enter a different Id")
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
		fmt.Print("Insert a menu you want :")

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
			fmt.Println("A menu that you insert is not found")
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
	db := connectDb()        
	defer db.Close()   

	service := entity.Service{}

	fmt.Print("Insert a service id :")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	select_by_id := "SELECT service_id,service_name,unit,price FROM service WHERE service_id = $1"
	err = db.QueryRow(select_by_id,id).Scan(&service.Service_id,&service.Service_name,&service.Unit,&service.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("===============================")
			fmt.Println("Service Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}

	fmt.Println("======== Update Service ==========")

	fmt.Print("Enter Service Name   : ")
	scanner.Scan()
	name := scanner.Text()
	if name != "" {
		service.Service_name = name
	}

	fmt.Print("Enter Service Unit   : ")
	scanner.Scan()
	unit := scanner.Text()
	if unit != "" {
		service.Unit = unit
	}

	fmt.Print("Enter Service Price  : ")
	scanner.Scan()
	input := scanner.Text()

	// Check if the input is empty
	if input != "" {
		price, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		service.Price = price
	} 

	// fill the updated_at value with time now
	service.Updated_at = time.Now()

	fmt.Println("=================================")

	update := "UPDATE service SET service_name = $2, unit = $3, price = $4, updated_at = $5 WHERE service_id = $1;"
	_, err = db.Exec(update, id, service.Service_name,service.Unit, service.Price,service.Updated_at)
	if err != nil {
		panic(err)  // Handle error if the query fails
	} else {
		fmt.Println("Successfully updated")  // Log success
	}
}

func delete_service() {
	db := connectDb()        
	defer db.Close()   

	service := entity.Service{}

	fmt.Print("Insert a service id :")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	// check if customer is exist
	select_by_id := "SELECT service_id FROM service WHERE service_id = $1"
	err = db.QueryRow(select_by_id,id).Scan(&service.Service_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("===============================")
			fmt.Println("Service Not Found. Please enter a different ID.")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}

	order_detail := entity.Order_detail{}

	service_in_order := `SELECT service_id FROM order_detail WHERE service_id = $1`

	err = db.QueryRow(service_in_order,id).Scan(&order_detail.Service_id)

	if err != nil {
		if err == sql.ErrNoRows {
			delete := "DELETE FROM service WHERE service_id = $1"

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
		fmt.Println("Service ID is being used in orders. Please delete the order first.")
	}
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
		fmt.Print("Insert a menu you want :")

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
			fmt.Println("A menu that you insert is not found")
			fmt.Println("=========================")
		}
	}
}

func create_order() {
	order := entity.Order{}
	order_detail := entity.Order_detail{}

	db := connectDb()        
	defer db.Close()   
	var err error 

	fmt.Println("========== Create Order ======")

	fmt.Print("Insert Order Id     : ")
	scanner.Scan()
	order.Order_id, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	// check if id order is already exist 
	check_id_order := `SELECT order_id FROM "order" WHERE order_id = $1;`

	err = db.QueryRow(check_id_order,order.Order_id).Scan(&order.Order_id)
	if err == nil {
		fmt.Println("===============================")
		fmt.Println("Order ID already exists. Please enter a different ID.")
		fmt.Println("===============================")
		return
	}

	fmt.Print("Insert Customer Id   : ")
	scanner.Scan()
	order.Customer_id, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	customer_exist := "SELECT customer_id FROM customer WHERE customer_id = $1"
	err = db.QueryRow(customer_exist,order.Customer_id).Scan(&order.Customer_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("===============================")
			fmt.Println("Customer Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}

	fmt.Print("Insert Received By  : ")
	scanner.Scan()
	order.Received_by = scanner.Text()

	err = db.QueryRow("SELECT MAX(order_detail_id) + 1 FROM order_detail;").Scan(&order_detail.Order_detail_id)
	if err != nil {
		panic(err)
	} 

	fmt.Print("Insert Service Id : ")
	scanner.Scan()
	order_detail.Service_id, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Print("Insert Quantity : ")
	scanner.Scan()
	order_detail.Qty, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	// fill the created_at and updated_at value with time now
	order.Order_date = time.Now()
	order.Created_at = time.Now()
	order.Updated_at = time.Now()

	fmt.Println("=================================")

	tx,err := db.Begin()
	if err != nil {
		panic(err)
	}

	insert_order := `INSERT INTO "order" (order_id,customer_id,order_date,received_by,created_at,updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	
	_,err = tx.Exec(insert_order,order.Order_id,order.Customer_id,order.Order_date,order.Received_by,order.Created_at,order.Updated_at)
	if err != nil {
		panic(err)  // Handle error if the query fails
	} 

	insert_order_detail := "INSERT INTO order_detail (order_detail_id,order_id,service_id,qty) VALUES ($1,$2,$3,$4)"
	_,err = tx.Exec(insert_order_detail,order_detail.Order_detail_id,order.Order_id,order_detail.Service_id,order_detail.Qty)
	if err != nil {
		panic(err)  // Handle error if the query fails
	} 

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully added")
	}
}

func complete_order() {
	order := entity.Order{}

	db := connectDb()        
	defer db.Close()   
	var err error 

	fmt.Println("========== Complete Order ======")

	fmt.Print("Insert Order Id          : ")
	scanner.Scan()
	order.Order_id, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	// check if id order is exist 
	check_id_order := `SELECT order_id FROM "order" WHERE order_id = $1;`

	err = db.QueryRow(check_id_order,order.Order_id).Scan(&order.Order_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Order Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}

	fmt.Print("Insert Completion Date(YYYY-MM-DD H:M:S)  : ")
	scanner.Scan()
	order.Completion_date, err = time.Parse("2006-01-02 15:04:05", scanner.Text())
	if err != nil {
		panic(err)
	}

	order.Updated_at = time.Now()

	fmt.Println("=================================")

	update := `UPDATE "order" SET completion_date = $2, updated_at = $3 WHERE order_id = $1;`
	_, err = db.Exec(update, order.Order_id, order.Completion_date,order.Updated_at)
	if err != nil {
		panic(err)  // Handle error if the query fails
	} else {
		fmt.Println("Successfully updated")  // Log success
	}
}

func view_of_list_order() {
	db := connectDb()        
	defer db.Close()   
	
	orders := []entity.Order{}

	select_all := `SELECT order_id,customer_id,order_date,COALESCE(completion_date, '0001-01-01 00:00:00') AS completion_date,received_by,created_at,updated_at FROM "order";`

	rows, err := db.Query(select_all)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		order := entity.Order{}
		err := rows.Scan(&order.Order_id,&order.Customer_id,&order.Order_date,&order.Completion_date,&order.Received_by,&order.Created_at,&order.Updated_at)
		if err != nil {
			panic(err)  // Handle error during scanning
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("==== List of all order =====")
	for _, order := range orders {
		fmt.Println(order.Order_id,order.Customer_id,order.Order_date,order.Completion_date,order.Received_by,order.Created_at,order.Updated_at)
	}
	fmt.Println("===============================")
}

func view_order_detail_by_id() {
	order := entity.Order{}
	order_detail := entity.Order_detail{}

	db := connectDb()        
	defer db.Close()   

	fmt.Println("========== View Order Details By Id ======")

	fmt.Print("Insert Order Id   : ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	select_by_id := `SELECT "order".order_id,customer_id,order_detail.service_id,order_detail.qty,order_date,COALESCE(completion_date, '0001-01-01 00:00:00') AS completion_date,received_by,created_at,updated_at 
	FROM "order" 
	INNER JOIN order_detail ON "order".order_id = order_detail.order_id
	WHERE "order".order_id = $1;`

	err = db.QueryRow(select_by_id,id).Scan(&order.Order_id, &order.Customer_id, &order_detail.Service_id, &order_detail.Qty, &order.Order_date, &order.Completion_date, &order.Received_by, &order.Created_at, &order.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Order Not Found")
			fmt.Println("===============================")
			return
		}

		panic(err)
	}

	fmt.Println(order.Order_id, order.Customer_id, order_detail.Service_id, order_detail.Qty, order.Order_date, order.Completion_date, order.Received_by, order.Created_at, order.Updated_at)
	fmt.Println("===============================")
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

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Database connection constants
	
		host     := os.Getenv("DB_HOST")  		    	// Host where the database is running
		port,err := strconv.Atoi(os.Getenv("DB_PORT"))  // Port where PostgreSQL is listening
		if err != nil {
			panic(err)
		}
		user     := os.Getenv("DB_USERNAME")  			// Database user
		password := os.Getenv("DB_PASSWORD")   			// Database password
		dbname   := os.Getenv("DB_DATABASE")        	// Database name
		dbconnection := os.Getenv("DB_CONNECTION")  	// Name of the database you want to connect

	// Connection string for PostgreSQL
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    				host, port, user, password, dbname)

	db, err := sql.Open(dbconnection, psqlInfo)  // Open a connection using the connection string

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