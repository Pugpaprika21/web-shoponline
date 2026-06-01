# How to Run - Web ShopOnline

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop/) & Docker Compose
- [Go 1.22+](https://go.dev/dl/) (สำหรับ development)
- [Air](https://github.com/air-verse/air) (สำหรับ hot-reload, optional)

---

## 1. Production (Docker)

รันทุกอย่างผ่าน Docker ครบจบในคำสั่งเดียว:

```bash
docker-compose up --build -d
```

| Service | URL | Description |
|---------|-----|-------------|
| App | http://localhost:3000 | เว็บหลัก |
| pgAdmin | http://localhost:5050 | จัดการ Database |
| PostgreSQL | localhost:5432 | Database |

### หยุด App

```bash
docker-compose down
```

### หยุด + ลบ Database (Reset ทั้งหมด)

```bash
docker-compose down -v
```

### ดู Logs

```bash
# ดู logs ทุก service
docker-compose logs -f

# ดู logs เฉพาะ app
docker-compose logs -f app
```

### Rebuild หลังแก้โค้ด

```bash
docker-compose up --build -d
```

---

## 2. Development (Local + Docker DB)

เหมาะสำหรับพัฒนา — แก้โค้ดแล้วเห็นผลทันที

### Step 1: Start Database

```bash
docker-compose up postgres -d
```

### Step 2: Copy Environment File

```bash
cp .env.example .env
```

แก้ `DB_HOST=localhost` ใน `.env` (เพราะรัน app นอก Docker)

### Step 3: Run App

**วิธี A — ใช้ Air (Hot Reload):**

```bash
# ติดตั้ง Air (ครั้งแรก)
go install github.com/air-verse/air@latest

# รัน
air
```

แก้ไฟล์ `.go` หรือ `.html` แล้ว app จะ restart อัตโนมัติ

**วิธี B — รันตรง:**

```bash
go run ./cmd/server
```

**วิธี C — Build แล้วรัน:**

```bash
go build -o bin/server ./cmd/server
./bin/server
```

---

## 3. Development ผ่าน Docker (Air + Volume Mount)

```bash
docker-compose -f docker-compose.dev.yml up --build
```

---

## 4. Default Accounts

### Admin

| Field | Value |
|-------|-------|
| Email | `admin@shoponline.com` |
| Password | `admin123` |
| URL | http://localhost:3000/admin/login |

### pgAdmin

| Field | Value |
|-------|-------|
| Email | `admin@admin.com` |
| Password | `admin123` |
| URL | http://localhost:5050 |

### pgAdmin → Connect to Database

หลัง login pgAdmin แล้ว:
1. Right-click **Servers** → **Register** → **Server**
2. Tab **General**: Name = `ShopOnline`
3. Tab **Connection**:
   - Host: `postgres`
   - Port: `5432`
   - Database: `shoponline_db`
   - Username: `shoponline`
   - Password: `shoponline_secret`
4. กด **Save**

---

## 5. Environment Variables

ตั้งค่าใน `.env`:

```env
# Application
APP_PORT=3000
APP_ENV=development

# Database
DB_HOST=localhost          # ใช้ "postgres" ถ้ารันใน Docker
DB_PORT=5432
DB_USER=shoponline
DB_PASSWORD=shoponline_secret
DB_NAME=shoponline_db
DB_SSLMODE=disable

# JWT
JWT_SECRET=your-secret-key-change-in-production
```

> ⚠️ **Production:** เปลี่ยน `JWT_SECRET` เป็นค่าที่ปลอดภัย

---

## 6. Project Structure

```
web-shoponline/
├── cmd/server/              # Entry point
├── internal/
│   ├── config/              # Configuration
│   ├── database/            # GORM connection, migration, seed
│   ├── dto/                 # Data Transfer Objects
│   ├── handler/             # HTTP handlers (Fiber)
│   ├── middleware/          # JWT auth middleware
│   ├── model/               # GORM models
│   ├── repository/          # Database queries + interfaces
│   ├── router/              # Route definitions
│   └── service/             # Business logic + interfaces
├── migrations/              # SQL reference
├── web/
│   ├── static/              # CSS, JS
│   └── templates/           # HTML templates
├── .air.toml                # Air hot-reload config
├── .env                     # Environment variables
├── docker-compose.yml       # Production Docker
├── docker-compose.dev.yml   # Development Docker (Air)
├── Dockerfile               # Production image
└── Dockerfile.dev           # Development image
```

---

## 7. Useful Commands

```bash
# Build only (verify compilation)
go build ./cmd/server

# Run tests
go test ./...

# Tidy dependencies
go mod tidy

# Check containers status
docker-compose ps

# Enter postgres shell
docker exec -it web-shoponline-postgres-1 psql -U shoponline -d shoponline_db

# View all tables
docker exec -it web-shoponline-postgres-1 psql -U shoponline -d shoponline_db -c "\dt"
```

---

## 8. API Quick Test

```bash
# Login (get token)
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@shoponline.com","password":"admin123"}'

# Use token for protected APIs
curl http://localhost:3000/api/admin/products \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

## 9. Pages Overview

### Public (ไม่ต้อง login)
- `/` — Home
- `/products` — Products listing
- `/login` — User login
- `/register` — User registration
- `/cart` — Shopping cart (ต้อง login เพื่อใช้งาน)

### Admin (ต้อง login ด้วย admin account)
- `/admin` — Dashboard
- `/admin/products` — จัดการสินค้า
- `/admin/inventory` — จัดการคลังสินค้า
- `/admin/orders` — จัดการคำสั่งซื้อ
- `/admin/users` — จัดการผู้ใช้/ผู้ดูแล
- `/admin/settings/menus` — ตั้งค่าเมนู
- `/admin/settings/roles` — ตั้งค่า Roles & Permissions
- `/admin/settings/store` — ตบแต่งหน้าร้าน
