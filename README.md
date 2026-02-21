# 💸 SplitLedger – Expense Settlement API

Go · Gin · GORM · SQLite · Docker Ready

## 📋 Table of Contents

1) Project Overview

2) Features

3) Tech Stack

4) System Architecture

5) Database Schema

6) API Endpoints

7) Getting Started

8) Environment Variables

9) Settlement Algorithm

10) Money Handling Strategy

11) Example curl Requests

12) Future Improvements


## 🎯 Project Overview

SplitLedger – Expense Settlement API is a backend system built in Go that allows users to manage shared group expenses and automatically calculate optimized settlements (like Splitwise).

Users can:

Create groups (trip, roommates, events)

Add shared expenses

Split bills equally

Compute minimized transactions required to settle debts

The system focuses on financial correctness and optimized debt simplification using a greedy settlement algorithm.

## 📌 Key Highlights

• RESTful API design using Gin
• Optimized settlement algorithm (minimal transactions)
• Money stored safely using int64 (no floating point errors)
• SQLite-based persistent storage
• Group-based expense isolation
• Clean separation of models and database logic
• CORS-enabled backend

## ✨ Features
Feature	Description
👥 User Management	Create and list users
👪 Group Creation	Create expense groups
💰 Add Expense	Add shared expenses
⚖️ Equal Split	Automatically divides amount among members
🤝 Settlement	Calculates optimized minimal transactions
💾 Persistent Storage	SQLite database with GORM
🔄 Restart Safe	Data persists across server restarts

## 🛠 Tech Stack
Layer	Technology	Purpose
Language	Go 1.22	High-performance backend
Framework	Gin	HTTP routing
ORM	GORM	Database abstraction
Database	SQLite	Persistent data store
Money Type	int64 (cents)	Precision-safe finance
Middleware	CORS	Frontend-backend communication
🏗 System Architecture

## High-Level Flow

Client → HTTP Request → Gin Router → Handler
                               ↓
                          Business Logic
                               ↓
                           GORM ORM
                               ↓
                           SQLite DB

## Core Components:

models/ → Entity definitions

database/ → DB connection logic

main.go → Routing + business logic

expense.db → Persistent database file

## 🗄 Database Schema
Users
Field	Type
ID	uint
Name	string
Email	string
CreatedAt	timestamp
Groups
Field	Type
ID	uint
Name	string
Expenses
Field	Type
ID	uint
GroupID	uint
PaidBy	uint
Amount	int64 (cents)
ExpenseSplits
Field	Type
ID	uint
ExpenseID	uint
UserID	uint
Share	int64 (cents)

## 📡 API Endpoints
Users

POST /users
Create new user

GET /users
List all users

Groups

POST /groups
Create new group

Expenses

POST /expenses
Add new expense

## Example Body:

{
  "group_id": 1,
  "paid_by": 1,
  "amount": 900,
  "users": [1,2,3]
}
Settlement

GET /settle/{groupID}

Returns optimized settlement transactions.

## 🚀 Getting Started
Prerequisites

Go 1.22+

Git

Step 1: Clone Repository
git clone https://github.com/YOUR_USERNAME/GO-Project.git
cd GO-Project

Step 2: Install Dependencies
go mod tidy

Step 3: Run Server
go run main.go

Server runs on:

http://localhost:8080

## ⚙️ Environment Variables

No required environment variables.
SQLite database file (expense.db) is created automatically.

## 🧠 Settlement Algorithm

The system minimizes the number of transactions required to settle debts.

Algorithm Steps

Calculate total amount paid by each user.

Calculate total amount owed by each user.
Compute net balance = paid − owed.

Classify users:

Positive balance → Creditors
Negative balance → Debtors

Apply greedy matching:

Match highest debtor with highest creditor.
Transfer minimum amount possible.
Repeat until all balances are zero.

Time Complexity: O(n)

## 💰 Money Handling Strategy

All monetary values are stored using:

int64 (in cents)

Example:

₹900 → stored as 900

Why?

• Avoid floating-point precision issues
• Ensure financial correctness
• No rounding errors

No float or decimal type is used anywhere.

📝 Example curl Requests
Create User
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name":"A","email":"a@test.com"}'
Create Group
curl -X POST http://localhost:8080/groups \
-H "Content-Type: application/json" \
-d '{"name":"Trip Group"}'
Add Expense
curl -X POST http://localhost:8080/expenses \
-H "Content-Type: application/json" \
-d '{"group_id":1,"paid_by":1,"amount":900,"users":[1,2,3]}'
Settlement
curl http://localhost:8080/settle/1

Response:

[
  {
    "from_user_id": 2,
    "to_user_id": 1,
    "amount": 300
  },
  {
    "from_user_id": 3,
    "to_user_id": 1,
    "amount": 300
  }
]
## 🔮 Future Improvements
Category	Enhancement
🔐 Auth	JWT-based authentication
📊 Dashboard	User expense summaries
🧠 Algorithm	Support unequal splits
🧠 Algorithm	Advanced debt graph simplification
💳 Finance	Currency support
📄 Docs	Swagger API documentation
🐳 DevOps	Dockerized deployment
🧪 Testing	Unit & integration tests
