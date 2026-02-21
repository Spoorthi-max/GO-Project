📌 Project Title

Splitwise-Style Expense Tracker REST API (Go + Gin + GORM)

📖 Project Description

This project is a REST API built using Go (Golang) to track shared expenses and split bills among friends.

Users can:

Create groups (e.g., roommates, trip groups)

Add shared expenses

Split expenses equally among members

Compute optimized settlements minimizing the number of transactions

The system implements a settlement algorithm to reduce unnecessary payments between users.

🛠️ Tech Stack

Language: Go (Golang)

Framework: Gin

ORM: GORM

Database: SQLite

Money Handling: int64 (stored in cents)

🗂️ Database Schema
User
Field	Type
ID	uint
Name	string
Email	string
CreatedAt	timestamp
Group
Field	Type
ID	uint
Name	string
Expense
Field	Type
ID	uint
GroupID	uint
PaidBy	uint
Amount	int64 (cents)
ExpenseSplit
Field	Type
ID	uint
ExpenseID	uint
UserID	uint
Share	int64 (cents)
💰 Money Handling Approach

All monetary values are stored using int64 in cents instead of float.

Example:

₹900 stored as 900 (cents representation)

Prevents floating-point precision errors

Ensures financial accuracy

No float or decimal type is used.

🔗 API Endpoints
Create User
POST /users
Get All Users
GET /users
Create Group
POST /groups
Add Expense
POST /expenses

Body:

{
  "group_id": 1,
  "paid_by": 1,
  "amount": 900,
  "users": [1,2,3]
}
Settlement
GET /settle/{groupID}
⚙️ Settlement Algorithm

The system calculates the net balance of each user in a group:

Add total amount paid by each user.

Subtract their share from total expenses.

Users with positive balance → Creditors.

Users with negative balance → Debtors.

Apply greedy matching:

Match largest debtor with largest creditor.

Transfer minimum of the two balances.

Repeat until balances are zero.

This minimizes the number of transactions required.

Time Complexity: O(n)

📊 Example Scenario

Users: A, B, C
Expense: ₹900
Paid by A
Split equally

Each share = ₹300

Net balances:

A → +600

B → -300

C → -300

Settlement:

B pays A ₹300

C pays A ₹300

Minimum transactions = 2
