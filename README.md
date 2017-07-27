# Golang test api

## Players
```console

***Create player***

POST /v1/players/
data: {"balance": 300}

curl -X POST -H 'Content-Type: application/json' -d "{\"balance\": 300}" http://localhost:8081/v1/players/

***Create Tournament***

POST /v1/tournaments/
data: {"deposit": 1000}

curl -X POST -H 'Content-Type: application/json' -d "{\"deposit\": 1000}" http://localhost:8081/v1/tournaments/

***Add player to tournament***

POST /v1/tournaments/1/

data: {"userId": 1000}

curl -X POST -H 'Content-Type: application/json' -d "{\"userId\": 15}" http://localhost:8081/v1/tournaments/1/

***Add player to tournament with backers***

POST /v1/tournaments/1/

data: { "userId": 14, "backers": [11, 12, 13]}

curl -X POST -H 'Content-Type: application/json' -d "{\"userId\": 15 \"backers\": [11, 12, 13]}" http://localhost:8081/v1/tournaments/1/

***Set result***

PUT /v1/tournaments/1/

data: {"id": 14} // id of winner

curl -X PUT -H 'Content-Type: application/json' -d "{\"userId\": 15}" http://localhost:8081/v1/tournaments/1/

***Reset***

POST /v1/reset/

curl -X POST -H 'Content-Type: application/json' http://localhost:8081/v1/reset/

```