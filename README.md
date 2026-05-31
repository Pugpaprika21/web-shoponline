# Web ShopOnline

ระบบขายของ Online พร้อมหน้า Admin สำหรับจัดการสินค้า, คลังสินค้า, คำสั่งซื้อ, เมนู และ Roles

## Tech Stack

- **Backend:** Go (Fiber Framework) + GORM
- **Frontend:** Tailwind CSS + HTMX + jQuery
- **Database:** PostgreSQL
- **DevOps:** Docker + Docker Compose
- **Hot Reload:** Air

## Project Structure

```
web-shoponline/
├── cmd/server/              # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── database/            # GORM connection, migration, seed
│   ├── dto/                 # Data Transfer Objects
│   ├── handler/             # HTTP handlers (Fiber)
│   ├── middleware/          # Auth middleware
│   ├── model/               # GORM models
│   ├── repository/          # Database queries (GORM)
│   ├── router/              # Route definitions
│   └── service/             # Business logic
├── migrations/              # SQL reference files
├── web/
│   ├── static/              # CSS, JS, images
│   └── templates/           # HTML templates (HTMX)
├── .air.toml                # Air hot-reload config
├── docker-compose.yml       # Production
├── docker-compose.dev.yml   # Development (with Air)
├── Dockerfile               # Production build
├── Dockerfile.dev           # Development build (Air)
└── .env
```

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.22+ (for local development)
- Air (for hot-reload): `go install github.com/air-verse/air@latest`

### Development (with Hot Reload)

```bash
# Start PostgreSQL only
docker-compose up postgres -d

# Run with Air hot-reload
air

# OR use Docker dev compose
docker-compose -f docker-compose.dev.yml up --build
```

### Production

```bash
docker-compose up --build
```

The application will be available at: http://localhost:3000

## Pages

### User (Public)
- **Home:** http://localhost:3000/
- **Products:** http://localhost:3000/products
- **Login:** http://localhost:3000/login
- **Register:** http://localhost:3000/register
- **Cart:** http://localhost:3000/cart

### Admin
- **Admin Login:** http://localhost:3000/admin/login
- **Dashboard:** http://localhost:3000/admin
- **Products:** http://localhost:3000/admin/products
- **Inventory:** http://localhost:3000/admin/inventory
- **Orders:** http://localhost:3000/admin/orders
- **Menu Settings:** http://localhost:3000/admin/settings/menus
- **Role Settings:** http://localhost:3000/admin/settings/roles

## API Endpoints

### Auth
- `POST /api/auth/login` - Login
- `POST /api/auth/register` - Register
- `POST /api/auth/logout` - Logout
- `GET /api/auth/me` - Get current user

### Public
- `GET /api/products` - List active products
- `GET /api/products/:id` - Get product details

### Cart (requires login)
- `GET /api/cart` - Get cart
- `POST /api/cart/items` - Add to cart
- `DELETE /api/cart/items/:id` - Remove from cart
- `GET /api/cart/count` - Get cart item count

### Admin (requires admin role)
- `GET/POST /api/admin/products` - List/Create products
- `PUT/DELETE /api/admin/products/:id` - Update/Delete product
- `GET/POST /api/admin/inventory` - List/Add inventory
- `GET /api/admin/orders` - List orders
- `GET /api/admin/orders/:id` - Get order details
- `PATCH /api/admin/orders/:id/status` - Update order status
- `GET/POST /api/admin/menus` - List/Create menus
- `PUT/DELETE /api/admin/menus/:id` - Update/Delete menu
- `GET/POST /api/admin/roles` - List/Create roles
- `PUT/DELETE /api/admin/roles/:id` - Update/Delete role
- `GET /api/admin/permissions` - List all permissions

## Default Admin Account

- **Email:** admin@shoponline.com
- **Password:** admin123
