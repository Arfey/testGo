package models

import (
	"fmt"
	"../consts"
	"strings"
)

// Provide structure of tournament model
type Tournament struct {
	Deposit int `json:"deposit"`
}

type Tournaments []Tournament

type Member struct {
	User_id int `json:"userId"`
	Backers []int `json:"backers"`
}

type Winner struct {
	Id int `json:"id"`
	Prise int `json:"prise"`
	Backers string `json:"backers"`
}

type MembersData struct {
	Balance int `json:"balance"`
	Count int `json:"count"`
}
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func CreateTournament(deposit int) (error) {
	sql := fmt.Sprintf("INSERT INTO tournaments (deposit) values (%d)", deposit)
	return simpleQuery(sql)
}

func AddMember(owner_id ,tournaments_id int, backers []int, prise int) {
	sql := fmt.Sprintf(
		`INSERT INTO tournaments_users (tournaments_id, members_id, backers, prise) values (%d, %d,
		'{` + arrayToString(append(backers, owner_id), ", ") + `}', %d)`, tournaments_id, owner_id,  prise)
	db.Query(sql)
}

func GetTournamentDeposit(id int) (int, error) {
	sql := fmt.Sprintf("SELECT deposit from tournaments WHERE id =  %d", id)
	rows, err := db.Query(sql)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		tournament := new(Tournament)
		err := rows.Scan(&tournament.Deposit)

		if err != nil {
			return 0, err
		}

		return tournament.Deposit, nil
	}

	return 0, fmt.Errorf(consts.TournamentNotFound)
}

func GetMembersData(id int, backers []int) (int, int, error) {
	sql := `SELECT min(balance), count(*) from users WHERE id IN
		(` + arrayToString(append(backers, id), ", ") + `)`
	fmt.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
		return 0, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		membersData := new(MembersData)
		err := rows.Scan(&membersData.Balance, &membersData.Count)

		if err != nil {
			return 1, 0, err
		}

		return membersData.Balance, membersData.Count, nil
	}

	return 2, 0, nil

}

func SetWinner(id, tournaments_id int) (error) {
	sql := fmt.Sprintf(
		"select prise, backers from tournaments_users where members_id = %d and tournaments_id = %d limit 1",
		id, tournaments_id)

	rows, err := db.Query(sql)

	if err != nil {
		return err
	}

	defer rows.Close()

	winner := new(Winner)

	for rows.Next() {
		rows.Scan(&winner.Prise, &winner.Backers)
	}

	arr := "(" + winner.Backers[1: len(winner.Backers) - 1] + ")"

	rows, _ = db.Query(fmt.Sprintf("select count(*) from tournaments_users where tournaments_id = %d", tournaments_id))

	var count int

	for rows.Next() {
		rows.Scan(&count)
	}

	sql_new := fmt.Sprintf(
		`update users set balance = balance + %d where id in ` + arr,
		winner.Prise * count)
	fmt.Println(sql_new)
	db.Query(sql_new)

	sql_delete := fmt.Sprintf(
		"delete from tournaments_users where tournaments_id = %d", tournaments_id)
	db.Query(sql_delete)

	db.Query(fmt.Sprintf("delete from tournaments where id = %d", tournaments_id))

	return nil

}
