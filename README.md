# 🌐 URL Monitor

A simple yet powerful URL monitoring service built with Go. It regularly checks the availability of specified URLs and sends real-time alerts via Telegram when any URL becomes unavailable.

---

## 🚀 Getting Started

### 🐳 Run with Docker Compose

To get started quickly, simply run:

```
docker-compose up -d --build
```

Make sure you have Docker and Docker Compose installed.

---

### ⚙️ Environment Configuration

Create a `.env` file in the root directory. You can use `.env.example` as a reference:

cp .env.example .env

Configure the necessary values such as:

- PostgreSQL connection details  
- Telegram Bot Token  
- Telegram Chat ID  
- Application port  
- JWT secret key  

---

## 📦 Features

- ✅ Monitor any public URL  
- 🕒 Automatic health checks every 1 minute  
- 🔔 Instant **Telegram notifications** on failure  
- 📊 Log history per URL (viewable via REST API)  
- 🧹 Old logs cleaned up automatically every hour  
- 🔒 JWT-based authentication  

---

## 🔐 Authentication

All protected endpoints require JWT authentication.

### 👤 Register and Login

Before using the system, you must **register** and then **login** to receive a JWT token:

**Register**  
POST /auth/register  
```
Body:  
{  
  "email": "user@example.com",  
  "password": "yourpassword"  
}
```
**Login**  
POST /auth/login  
```
Body:  
{  
  "email": "user@example.com",  
  "password": "yourpassword"  
}
```

The response will contain a JWT token.

### 🔑 Using the Token

Include the token in each protected request:

Authorization: Bearer <your_token>

---

## 📌 API Endpoints

### 📁 URL Management

**Create URL**  
POST /url  
```
Headers:  
Authorization: Bearer <your_token>  
Body:  
{  
  "address": "https://example.com"  
}
```

**Get All URLs**  
GET /url
```  
Headers:  
Authorization: Bearer <your_token>  
```

**Delete URL**  
DELETE /url/{id}
```  
Headers:  
Authorization: Bearer <your_token>  
```

---

### 📜 Monitor Logs

**Get Logs for Specific URL**  
GET /url/{id}/logs  
```
Headers:  
Authorization: Bearer <your_token>  
```

Returns a list of recent check results, including timestamp, status, HTTP code, and any error messages.

---

## 📣 Telegram Alerts

When a monitored URL becomes unavailable (connection error or HTTP status >= 400), an alert is sent to your Telegram chat:

```
🚨 URL https://example.com is down!  
Code: 500  
Error: Internal Server Error
```

To configure Telegram alerts, make sure you set the following environment variables:

- TELEGRAM_BOT_TOKEN  
- TELEGRAM_CHAT_ID  

You can get these by creating a bot with @BotFather and adding it to a chat.

---

## 🧪 Health Checks

- Every 1 minute, all active URLs are checked.  
- Logs older than 1 hour are automatically deleted.  
- Results are stored in the database and accessible via the API.  

---

## 💬 Contact

Feel free to contribute or open an issue. For questions or suggestions, ping me on Telegram or GitHub.

---

## 📜 License

MIT License – use it freely!
