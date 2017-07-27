# Golang test api

## Deploy

```console
git clone https://github.com/Arfey/testGo.git
cd testGo
docker-compose up
```
http://localhost:8081/v1/


## Players


***Create player***
```console

POST /v1/players/
data: {"balance": 300}

curl -X POST -H 'Content-Type: application/json' -d "{\"balance\": 300}" http://localhost:8081/v1/players/
```
***Create Tournament***
```console
POST /v1/tournaments/
data: {"deposit": 1000}

curl -X POST -H 'Content-Type: application/json' -d "{\"deposit\": 1000}" http://localhost:8081/v1/tournaments/
```
***Add player to tournament***
```console
POST /v1/tournaments/1/

data: {"userId": 1000}

curl -X POST -H 'Content-Type: application/json' -d "{\"userId\": 15}" http://localhost:8081/v1/tournaments/1/
```
***Add player to tournament with backers***
```console
POST /v1/tournaments/1/

data: { "userId": 14, "backers": [11, 12, 13]}

curl -X POST -H 'Content-Type: application/json' -d "{\"userId\": 15 \"backers\": [11, 12, 13]}" http://localhost:8081/v1/tournaments/1/
```
***Set result***
```console
PUT /v1/tournaments/1/
```

data: {"id": 14} // id of winner

curl -X PUT -H 'Content-Type: application/json' -d "{\"userId\": 15}" http://localhost:8081/v1/tournaments/1/

***Reset***

POST /v1/reset/

curl -X POST -H 'Content-Type: application/json' http://localhost:8081/v1/reset/

```
