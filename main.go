package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"time"
)

type Expense struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

func main() {

	expenses := make([]Expense, 0)

	content, err := os.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &expenses)
	if err != nil {
		panic(err)
	}

	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	if len(os.Args) < 2 {
		panic("commands required. Use \"help\" to get the list of commands")
	}
	cmd := os.Args[1]

	switch cmd {
	case "help":
		println("help - list of commands")

	case "list":
		fmt.Printf("%-3s %-11s %-12s %6s\n\n", "ID", "Date", "Description", "Amount")

		for _, expense := range expenses {
			fmt.Printf("%-3d %-11s %-12s %6d\n", expense.Id, expense.Date, expense.Description, expense.Amount)
		}

	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		descPtr := addCmd.String("description", "description", "expense description")
		amountPtr := addCmd.Int("amount", 0, "expense amount")
		addCmd.Parse(os.Args[2:])

		var newId int

		if len(expenses) > 0 {
			newId = expenses[len(expenses)-1].Id + 1
		} else {
			newId = 1
		}

		newExpense := Expense{
			Id:          newId,
			Date:        time.Now().Format("2006-01-02"),
			Description: *descPtr,
			Amount:      *amountPtr,
		}

		expenses = append(expenses, newExpense)

		updatedData, err := json.MarshalIndent(expenses, "", " ")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("./data.json", updatedData, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", newId)

	case "delete":

		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		idPtr := deleteCmd.Int("id", 0, "expense id")

		deleteCmd.Parse(os.Args[2:])

		for i, expense := range expenses {
			if expense.Id == *idPtr {
				expenses = slices.Delete(expenses, i, i+1)
				break
			}
		}

		updatedData, err := json.MarshalIndent(expenses, "", " ")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("./data.json", updatedData, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("Expense deleted successfully")

	case "summary":

		sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		monthPtr := sumCmd.Int("month", 0, "expense month")
		sumCmd.Parse(os.Args[2:])
		count := 0

		if *monthPtr < 0 || *monthPtr > 12 {
			fmt.Println("Invalid month. Please enter a value between 1 and 12.")
			return
		}

		if *monthPtr == 0 {
			for _, expense := range expenses {
				count += expense.Amount
			}
			fmt.Printf("Total expenses: $%d\n", count)
			return
		}

		for _, expense := range expenses {
			t, err := time.Parse("2006-01-02", expense.Date)
			if err != nil {
				panic(err)
			}

			expenseMonth := t.Month()

			if *monthPtr == int(expenseMonth) {
				count += expense.Amount
			}
		}

		fmt.Printf("Total expenses for %s: $%d\n", months[*monthPtr], count)

	}
}
