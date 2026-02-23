# 📡 MQTT Chat Simulation (Go + EMQX + Docker)

A fully simulated real-time chat system using:

- Go
- MQTT (Pub/Sub)
- EMQX Broker (Docker)
- Retained Messages
- Ephemeral Messages (TTL)
- Typing Indicators
- Dynamic Topics
- One-to-One Chat
- Group Chat
- Rule Engine
- YAML Configuration

---

# 🧠 Architecture

Publisher  --->  EMQX Broker  --->  Subscriber

Publisher and Subscriber run independently.

---




# 🚀 How To Run

### 1️⃣ Start EMQX

```
docker-compose up -d
```

### 2️⃣ Install deps

```
go mod tidy
```

### 3️⃣ Run Subscriber

```
go run main.go sub --user=alice --topic=group/dev
```

### 4️⃣ Run Publisher

```
go run main.go pub --user=bob --topic=group/dev
```

---

# 🧪 Example

Publisher:
```
Hello
/typing
```

Subscriber receives:
```
[group/dev][normal] bob: Hello
[group/dev] bob is typing...
```

---

# 🛑 Stop Broker

```
docker-compose down
```

---


