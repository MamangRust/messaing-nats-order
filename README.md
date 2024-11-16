## Messaging NATS Order
NATS Order is a project that demonstrates the use of NATS for messaging and processing orders. It consists of several components, including the order producer, order processor, email service, and NATS server.


## Features
- NATS order processing system
- Order messaging and processing
- Email notification on order completion

## Running the project

### Setup Email in Ethereal

`https://ethereal.email/`

```
emailUser := ""
emailPassword := ""
emailServer := "smtp.ethereal.email"
emailPort := "587"
```

### Build and run

```bash
docker-compose up -d --build
```

### Test Curl

```
    curl -X POST -H "Content-Type: application/json
" -d '{"id": 235, "status": "processed"}' http://172.20.0.4:5000/placeOrder
```
