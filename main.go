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

const filePath = "./data.json"

var months = map[int]string{
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

func main() {

	expenses := make([]Expense, 0)

	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &expenses)
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		panic("commands required. Use \"help\" to get the list of commands")
	}
	cmd := os.Args[1]

	switch cmd {
	case "help":
		displayCmds()

	case "list":
		listExpenses(expenses)

	case "add":
		err := handleAdd(expenses)
		if err != nil {
			fmt.Println(err)
		}

	case "delete":

		err := handleDelete(expenses)
		if err != nil {
			fmt.Println(err)
		}

	case "summary":
		displayExpenseSummery(expenses)

	case "update":
		err := handleUpdate(expenses)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Println("Invalid command.Use \"help\" to get the list of commands")

	}
}

func displayCmds() {
	fmt.Println("Available commands:")
	fmt.Println("  list          - List all expenses")
	fmt.Println("  add           - Add a new expense. Usage: add --description <desc> --amount <amount>")
	fmt.Println("  delete        - Delete an expense by ID. Usage: delete --id <id>")
	fmt.Println("  update        - Update an expense by ID. Usage: update --id <id> --description <desc> --amount <amount>")
	fmt.Println("  summary       - Show total expenses. Use --month <1-12> for a specific month")
	fmt.Println("  help          - Display available commands")
}

func listExpenses(expenses []Expense) {
	fmt.Printf("%-3s %-11s %-12s %6s\n\n", "ID", "Date", "Description", "Amount")

	for _, expense := range expenses {
		fmt.Printf("%-3d %-11s %-12s %6d\n", expense.Id, expense.Date, expense.Description, expense.Amount)
	}
}

func writeExpenses(expenses []Expense) error {
	updatedData, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		panic(err)
	}

	return os.WriteFile(filePath, updatedData, 0600)
}

func handleAdd(expenses []Expense) error {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	descPtr := addCmd.String("description", "description", "expense description")
	amountPtr := addCmd.Int("amount", 0, "expense amount")
	addCmd.Parse(os.Args[2:])

	if *amountPtr <= 0 {
		return fmt.Errorf("invalid amount. Please enter a positive value")
	}

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

	err := writeExpenses(expenses)
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", newId)

	return nil
}

func handleDelete(expenses []Expense) error {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	idPtr := deleteCmd.Int("id", 0, "expense id")

	deleteCmd.Parse(os.Args[2:])

	if *idPtr <= 0 {
		return fmt.Errorf("invalid ID. Please enter a positive integer")
	}

	found := false

	for i, expense := range expenses {
		if expense.Id == *idPtr {
			expenses = slices.Delete(expenses, i, i+1)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Expense not found")

	}

	err := writeExpenses(expenses)
	if err != nil {
		panic(err)
	}

	fmt.Println("Expense deleted successfully")

	return nil
}

func handleUpdate(expenses []Expense) error {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	idPtr := updateCmd.Int("id", 0, "expense id")
	descriptionPtr := updateCmd.String("description", "", "expense description")
	amountPtr := updateCmd.Int("amount", 0, "expense amount")
	updateCmd.Parse(os.Args[2:])

	if *descriptionPtr == "" || *amountPtr <= 0 {
		return fmt.Errorf("invalid description or amount. Please provide valid values")
	}

	found := false
	for i, expense := range expenses {
		if expense.Id == *idPtr {
			expenses[i].Description = *descriptionPtr
			expenses[i].Amount = *amountPtr
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("expense not found")
	}

	err := writeExpenses(expenses)
	if err != nil {
		return err
	}

	fmt.Printf("Expense updated successfully (ID: %d)\n", *idPtr)

	return nil
}

func displayExpenseSummery(expenses []Expense) {

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
