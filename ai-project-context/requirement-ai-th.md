# Project Context & AI Rules

## 1. Project Overview
* **Name:** web-shoponline
* **Description:** ระบบขายของ online มีหน้า สำหรับ admin user
* **Tech Stack:**
    * **Backend:** Golang (Fiber Framework) (gorm) เน้นการทำ REST API, Dynamic JSON Schema Parsing และ Rendering HTMX Components
    * **Frontend:** Tailwind CSS, HTMX, jQuery
    * **Database:** PostgreSQL (Pgsql)
    * **DevOps:** Docker, Docker Compose

---

## 2. Architecture & Code Style Guidelines

### 📂 File & Directory Structure
* **Frontend:** แยกโฟลเดอร์สำหรับ Static Assets และ HTML/Components Templates ชัดเจน
* **Backend:** ใช้โครงสร้างสากล **Standard Go Project Layout** แยก Layer ชัดเจน:
    * `cmd/` (Entry point)
    * `internal/handler/` (Fiber Handlers)
    * `internal/service/` (Business Logic)
    * `internal/repository/` (Database Queries)
    * `internal/model/` (Structs & DTOs)

### 💻 Coding Standards
* **Architecture Pattern:** Handler -> Service -> Repository -> Database 
* **Data Flow:** รับส่งข้อมูลระหว่าง Layer ด้วย **Go Structs / DTOs** เสมอ ห้ามส่งโครงสร้างดิบจาก Database (Models) ออกไปที่ Handler โดยตรง
* **Fiber Context:** ทุกๆ Handler และ Service ต้องมีการรับส่งและจัดการ `fiber.Ctx` และ `context.Context` อย่างถูกต้อง
* **Type Safety & Validation:** * ใช้ฟีเจอร์ Strong Typing ของ Go (หลีกเลี่ยงการใช้ `interface{}` หรือ `any` หากไม่จำเป็น)
    * ใช้ **Struct Tags** (`json:`, `form:`, `validate:`) ร่วมกับ Library เช่น `go-playground/validator` สำหรับทำ Request Data Validation เสมอ

---

## 3. Environment & Configuration
* **Local Setup:** ทำงานผ่าน Docker / Docker Compose เป็นหลัก
* **Environment Variables:** ห้าม Hardcode ค่า Config หรือ Variable ต่างๆ ลงในไฟล์โค้ดหรือไฟล์ YAML ของ Docker โดยตรง ให้ดึงผ่าน `.env` หรือ `env_file` เข้ามาผูกกับ Config Struct ของ Go ในตอนเริ่ม Startup App เสมอ

---

## 4. AI Communication & Output Rules
* **Tone:** ตอบคำถามแบบกระชับ ตรงประเด็น (Concise) ไม่ต้องเกริ่นนำยาว
* **Code Quality:** เมื่อเขียน Code Snippet ต้องเป็นโค้ดที่พร้อมใช้งานจริง (Production-ready) มี **Explicit Error Handling** (`if err != nil`) ครบถ้วนตามสไตล์ของ Go และปิดการเชื่อมต่อ Resource (เช่น `defer rows.Close()`) เสมอ
* **Language:** สามารถอธิบายคอนเซปต์ด้วยภาษาไทย แต่ชื่อตัวแปร ฟังก์ชัน โครงสร้าง Struct, แท็ก JSON และ Code Comment ต้องเป็นภาษาอังกฤษมาตรฐาน