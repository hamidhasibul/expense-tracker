package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
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

	if len(os.Args) < 2 {
		panic("commands required. Use \"help\" to get the list of commands")
	}
	cmd := os.Args[1]

	switch cmd {
	case "help":
		println("help - list of commands")
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

		// TODO: List Expenses
		// TODO: Delete expense by Id
		// TODO: Summerise total Expense
		// TODO: Summerise total Expense (by Month)

	}
}
