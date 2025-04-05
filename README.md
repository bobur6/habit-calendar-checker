# habit-calendar-checker
Simple habit tracker REST API built with Go. Create and manage habit lists and tasks.

# 🚀 Habit Tracker API

### 🔄 Полноценный REST API для трекинга привычек  
CRUD для списков и задач, авторизация, миграции, контейнеризация и всё по этапам!  
> Выполнено как часть учебного проекта: [Task 7 – Task 12]

---

## 📸 Overview

![Go](https://img.shields.io/badge/Go-1.21-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue)
![Docker](https://img.shields.io/badge/Dockerized-yes-blue)

Habit Tracker API — это современный backend-сервис, позволяющий:
- Создавать **списки привычек**  
- Добавлять и управлять **ежедневными задачами**  
- Использовать **аутентификацию и авторизацию**  
- Подключать базу данных (PostgreSQL) через **GORM**
- Работать с **миграциями** и **Docker**

---

## ⚙️ Установка и запуск

> Убедитесь, что у вас установлен **Go**, **PostgreSQL**, и (опционально) **Docker**.

### 🔧 Локальный запуск (без Docker)

```bash
git clone https://github.com/your-username/habit-tracker-api.git
cd habit-tracker-api
go run main.go
