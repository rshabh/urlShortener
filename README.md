This is a url shortener.

# 📌 Database Schema - URL Shortener

This project uses **PostgreSQL** with two main tables:

---

## **1. users table**

Stores user credentials and metadata.

| Column      | Type      | Description                              |
| ----------- | --------- | ---------------------------------------- |
| `id`        | UUID      | Primary key, unique identifier for user. |
| `uname`     | TEXT      | Username (unique).                       |
| `password`  | TEXT      | Hashed password using bcrypt.            |
| `createdAt` | TIMESTAMP | Timestamp when user was created.         |

🔹 Example row:

```
id:        74043f0e-2a5e-4b28-80e5-9f50d742f447
uname:     rishabh
password:  $2a$10$pdooBlaim0pbGqfjzCLLPuI3kY2b2b/wsFWkvm01aq69kghDBjO9a
createdAt: 2025-08-29 07:00:43.689086+00:00
```

---

## **2. url table**

Stores shortened URLs mapped to users.

| Column       | Type | Description                                                                |
| ------------ | ---- | -------------------------------------------------------------------------- |
| `short`      | TEXT | Generated short code for the URL. (Primary Key)                            |
| `long`       | TEXT | Original long URL.                                                         |
| `fk_user_id` | UUID | Foreign key referencing `users.id`. Identifies which user created the URL. |

🔹 Example rows:

```
short:     wQCsO
long:      https://chatgpt.com/
fk_user_id: 74043f0e-2a5e-4b28-80e5-9f50d742f447

short:     G4Z4a
long:      https://chatgpt.com/
fk_user_id: 8d8f5729-6c0e-4722-99de-20239c20ef46
```

---

## 🔗 Relationships

* **1 User → Many URLs**
* `url.fk_user_id` references `users.id` (foreign key).
* Ensures that every shortened URL belongs to a registered user.
