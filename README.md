# E-Commerce Backend API (Go + Fiber + GORM)

## 🚀 Introduction
This project is a **high-performance** and **scalable** E-Commerce backend API built using **Go (Golang)** and the **Fiber** web framework. It is designed to efficiently handle user authentication, product management, shopping carts, orders, and more. The system ensures **speed, security, and scalability**, making it ideal for production-grade applications.

![DALL·E 2025-02-14 11 54 30 - A simple and clean social preview image for an e-commerce backend API built with Go (Golang)  The design should include the Go Gopher mascot holding a](https://github.com/user-attachments/assets/bba5e560-1599-400f-a2c3-c05dd1f6c8ef)


## 🏗️ Tech Stack & Why I Chose It

| Technology  | Purpose | Why? |
|------------|---------|------|
| **Go (Golang)** | Backend Language | High performance, concurrency support, efficient memory management |
| **Fiber** | Web Framework | Fast and lightweight, optimized for high-speed APIs |
| **GORM** | ORM for Database | Simplifies database operations, provides robust query capabilities |
| **PostgreSQL/MySQL** | Database | Reliable, scalable, and supports complex queries |
| **Bcrypt** | Password Hashing | Ensures secure storage of user credentials |
| **JWT (JSON Web Token)** | Authentication | Secure and scalable authentication mechanism |
| **Docker** (Planned) | Containerization | Ensures consistency across environments |

## 📂 Project Structure
```
C:\Go-Lang\ecommerce-backend
│── internal/
│   ├── config/          # Database & App Configurations
│   ├── models/          # Data Models (User, Product, Order, etc.)
│   ├── handlers/        # API Handlers (Business Logic)
│   ├── routes/          # Route Definitions
│── pkg/                 # Utility Functions (Helper Methods, Middleware, etc.)
│── main.go              # Entry Point (Main.go)
│── .env                 # Environment Variables (Database URL, Secrets, etc.)
│── go.mod               # Go Module File
│── TODO.md              # Pending Tasks & Features
│── README.md            # Project Documentation
```

## ✅ Features Implemented So Far
### 1️⃣ **User Authentication & Authorization**
- Secure user registration with **Bcrypt password hashing**.
- JWT-based authentication for **secure API access**.
- User roles (**Admin & Customer**) to **restrict access** to specific actions.

### 2️⃣ **Product Management**
- CRUD operations for **products** (Add, Update, Delete, Fetch).
- Admin-only access for adding and updating products.
- Users can browse the product catalog.

### 3️⃣ **Shopping Cart System**
- Users can **add, remove, and update** cart items.
- The cart persists per user and is linked to their account.

### 4️⃣ **Order Processing**
- Users can checkout their cart, converting it into an **order**.
- Orders contain **order items, total price, and status**.
- **Admin can update order status** (e.g., Pending → Shipped → Delivered).

### 5️⃣ **Fetching Order History**
- Users can **view their past orders** with details.
- Each order includes **ordered items, prices, and status updates**.

## 🚀 Upcoming Features (Planned)
🔜 **Payment Integration** (Stripe/PayPal for secure transactions)  
🔜 **Admin Dashboard** (For managing products & orders)  
🔜 **Docker Support** (For easy deployment & environment consistency)  
🔜 **Unit Testing & API Testing** (Ensuring reliability)  

## ⚡ How to Run the Project
### 1️⃣ Clone the Repository
```bash
git clone https://github.com/your-username/ecommerce-backend.git
cd ecommerce-backend
```
### 2️⃣ Install Dependencies
```bash
go mod tidy
```
### 3️⃣ Set Up Environment Variables
Create a `.env` file and configure the database connection:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=ecommerce_db
JWT_SECRET=your-secret-key
```
### 4️⃣ Run Migrations (Initialize Database)
```bash
go run migrate.go
```
### 5️⃣ Start the Server
```bash
go run main.go
```
API will be running at: `http://localhost:8000`

## 📌 API Endpoints
### User Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| **POST** | `/ecom/auth/register` | Register a new user |
| **POST** | `/ecom/auth/login` | Authenticate user & get JWT token |

### Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| **POST** | `/ecom/products` | Add a new product (Admin only) |
| **GET** | `/ecom/products` | Get all products |
| **PUT** | `/ecom/products/:id` | Update product (Admin only) |
| **DELETE** | `/ecom/products/:id` | Delete product (Admin only) |

### Cart
| Method | Endpoint | Description |
|--------|----------|-------------|
| **POST** | `/ecom/cart` | Add item to cart |
| **GET** | `/ecom/cart` | View cart items |
| **DELETE** | `/ecom/cart/:id` | Remove item from cart |

### Orders
| Method | Endpoint | Description |
|--------|----------|-------------|
| **POST** | `/ecom/orders` | Checkout cart & place order |
| **GET** | `/ecom/orders` | Get user order history |
| **PUT** | `/ecom/orders/:id/status` | Update order status (Admin only) |

## 👥 Contributors
- **Your Name** (Lead Developer)
- Open to contributors! Feel free to submit PRs.

## 📜 License
This project is **open-source** and licensed under the MIT License.

---
📢 **Have suggestions or improvements?** Feel free to open an issue or contribute! 🚀

