package main

import (
	"expense-tracker/database"
	"expense-tracker/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Connect Database
	database.Connect()

	// Auto migrate tables
	database.DB.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.Expense{},
		&models.ExpenseSplit{},
	)

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Expense Tracker API Running",
		})
	})

	// ---------------- USERS ----------------

	r.POST("/users", func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		database.DB.Create(&user)
		c.JSON(200, user)
	})

	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		database.DB.Find(&users)
		c.JSON(200, users)
	})

	// ---------------- GROUPS ----------------

	r.POST("/groups", func(c *gin.Context) {
		var group models.Group

		if err := c.ShouldBindJSON(&group); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		database.DB.Create(&group)
		c.JSON(200, group)
	})

	// ---------------- ADD EXPENSE ----------------

	r.POST("/expenses", func(c *gin.Context) {

		var input struct {
			GroupID uint   `json:"group_id"`
			PaidBy  uint   `json:"paid_by"`
			Amount  int64  `json:"amount"` // stored in cents
			Users   []uint `json:"users"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(input.Users) == 0 {
			c.JSON(400, gin.H{"error": "Users list cannot be empty"})
			return
		}

		expense := models.Expense{
			GroupID: input.GroupID,
			PaidBy:  input.PaidBy,
			Amount:  input.Amount,
		}

		database.DB.Create(&expense)

		splitAmount := input.Amount / int64(len(input.Users))

		for _, userID := range input.Users {
			split := models.ExpenseSplit{
				ExpenseID: expense.ID,
				UserID:    userID,
				Share:     splitAmount,
			}
			database.DB.Create(&split)
		}

		c.JSON(200, gin.H{"message": "Expense added successfully"})
	})

	// ---------------- SETTLEMENT ----------------

	r.GET("/settle/:groupID", func(c *gin.Context) {

		groupIDParam := c.Param("groupID")
		groupIDInt, err := strconv.Atoi(groupIDParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid group ID"})
			return
		}
		groupID := uint(groupIDInt)

		var expenses []models.Expense
		database.DB.Where("group_id = ?", groupID).Find(&expenses)

		var splits []models.ExpenseSplit
		database.DB.
			Joins("JOIN expenses ON expenses.id = expense_splits.expense_id").
			Where("expenses.group_id = ?", groupID).
			Find(&splits)

		balance := make(map[uint]int64)

		// Add amounts paid
		for _, e := range expenses {
			balance[e.PaidBy] += e.Amount
		}

		// Subtract shares
		for _, s := range splits {
			balance[s.UserID] -= s.Share
		}

		type Person struct {
			UserID uint
			Amount int64
		}

		var debtors []Person
		var creditors []Person

		for id, amt := range balance {
			if amt < 0 {
				debtors = append(debtors, Person{id, -amt})
			} else if amt > 0 {
				creditors = append(creditors, Person{id, amt})
			}
		}

		var settlements []gin.H
		i, j := 0, 0

		for i < len(debtors) && j < len(creditors) {

			min := debtors[i].Amount
			if creditors[j].Amount < min {
				min = creditors[j].Amount
			}

			settlements = append(settlements, gin.H{
				"from_user_id": debtors[i].UserID,
				"to_user_id":   creditors[j].UserID,
				"amount":       min,
			})

			debtors[i].Amount -= min
			creditors[j].Amount -= min

			if debtors[i].Amount == 0 {
				i++
			}
			if creditors[j].Amount == 0 {
				j++
			}
		}

		c.JSON(200, settlements)
	})

	r.Run(":8080")
}
