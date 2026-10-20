# زیرساخت Docker

این پوشه شامل زیرساخت مورد نیاز پروژه GB1 است و سرویس‌های زیر را اجرا می‌کند:

- **PostgreSQL 18.1**
- **pgAdmin 4**
- **Redis 8.4**

تمام سرویس‌ها از طریق شبکه Docker مشترک **padegan-network** با یکدیگر در ارتباط هستند.

---

# پیش‌نیازها

قبل از اجرای پروژه، موارد زیر باید نصب باشند:

- Docker
- Docker Compose (نسخه 2 یا بالاتر)
- Make (اختیاری)

بررسی نصب:

```bash
docker --version
docker compose version
```

---

# ساختار پروژه

```text
development/
├── .env
├── postgres/
│   └── init/
├── redis/
└── README.md
```

---

# PostgreSQL

## اطلاعات اتصال

| فیلد                | مقدار       |
|---------------------|-------------|
| Host (داخل Docker)  | postgres    |
| Host (از روی سیستم) | localhost   |
| Port                | 5432        |
| Database            | padegan-db     |
| Username            | admin       |
| Password            | password123 |

### اتصال از داخل Docker

```text
postgres:5432
```

### اتصال از روی Host Machine

```text
localhost:5432
```

---

# pgAdmin

## دسترسی

```text
http://localhost:5050
```

## اطلاعات ورود

| فیلد     | مقدار             |
|----------|-------------------|
| Email    | admin@example.com |
| Password | password123       |

---

# Redis

## اطلاعات اتصال

| فیلد                | مقدار       |
|---------------------|-------------|
| Host (داخل Docker)  | redis       |
| Host (از روی سیستم) | localhost   |
| Port                | 6379        |
| Password            | password123 |

### اتصال از داخل Docker

```text
redis:5432
```

### اتصال از روی Host Machine

```text
localhost:5432
```

---

ارتباط بین سرویس‌ها از طریق **نام سرویس** انجام می‌شود.

مثال:

| سرویس      | Host     |
|------------|----------|
| PostgreSQL | postgres |
| Redis      | redis    |

---

# Volumes

داده‌های سرویس‌ها در Docker Volume ذخیره می‌شوند.

| Volume            | کاربرد              |
|-------------------|---------------------|
| padegan-postgres-data | داده‌های PostgreSQL |
| padegan-pgadmin-data  | تنظیمات pgAdmin     |
| padegan-redis-data    | داده‌های Redis      |

مشاهده Volumeها:

```bash
docker volume ls
```

---

# اطلاعات اتصال سریع

## PostgreSQL

```text
Host: localhost
Port: 5432
Database: padegan-db
User: admin
Password: password123
```

## Redis

```text
Host: localhost
Port: 6379
Password: password123
```

## pgAdmin

```text
URL: http://localhost:5050
Email: admin@example.com
Password: password123
```