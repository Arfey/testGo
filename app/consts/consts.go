package consts

const (
	Version string = "v1"
	UserNotFound = "User does not exist."
	TournamentNotFound = "Tournament does not exist."
	UserBackerInclude = "User include in backers."
	LessMinDepositLimit = "LessMinDepositLimit"
)

type Error struct {
	Error string `json:"error"`
}

var (
	CreateUsers = `CREATE TABLE IF NOT EXISTS users (
id serial PRIMARY KEY,
balance int NOT NULL DEFAULT 0
)`

	CreateTournaments = `CREATE TABLE IF NOT EXISTS tournaments (
id serial PRIMARY KEY,
deposit int NOT NULL DEFAULT 0
)`

	CreateTournamentsUsers = `CREATE TABLE IF NOT EXISTS tournaments_users (
id serial PRIMARY KEY,
tournaments_id int REFERENCES tournaments (id) ON DELETE CASCADE,
members_id int REFERENCES users (id) ON DELETE CASCADE,
backers int[] NULL,
prise int NOT NULL
)`
	DeleteTournamentsUsers = "DELETE FROM tournaments_users"
	DeleteTournaments = "DELETE FROM tournaments"
	DeleteUsers = "DELETE FROM users"
	InitDbString = "user=%s dbname=%s password=%s host=db sslmode=disable"
)
