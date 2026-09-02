package main

import "log"

func main() {
	db := connectDB()
	defer db.Close()

	log.Println("Accounting service started")

	consumeOrders(db)
}
