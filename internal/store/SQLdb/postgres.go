package SQLdb

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"task1/internal/models"
)

type PostgresDB struct{
	DB *sql.DB
	Link string
}
func (ps *PostgresDB) Open() error{
	db, err := sql.Open("postgres", ps.Link)
	if err != nil {
		return err
	}
	ps.DB = db
	if err := ps.DB.Ping(); err != nil{
		return err
	}
	return nil
}
func (ps *PostgresDB) Close() error{
	if err := ps.DB.Close(); err != nil{
		return err
	}
	return nil
}

func (ps *PostgresDB) DeleteByID(ID int) (error){
	if _, err := ps.DB.Exec("DELETE FROM users WHERE user_id = $1", ID); err != nil{
		return err
	}
	return nil
}

func (ps *PostgresDB) AddUser(user models.User) error{
	query := `INSERT INTO users (name, surname, phone_number) VALUES($1, $2, $3)`
	if _, err := ps.DB.Exec(query, user.Name, user.Surname, user.PhoneNum); err != nil{
		return err
	}
	return nil
}

// GetUsers returns a pretty json of all users
func (ps *PostgresDB) GetUsers() (string, error){
	query := `SELECT user_id, name, surname, phone_number FROM users`

	rows, err := ps.DB.Query(query)
	if err != nil{
		return "", err
	}

	defer rows.Close()
	var users []models.User

	for rows.Next(){
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.PhoneNum); err != nil{
			return "", err
		}
		users = append(users, user)
	}
	res, err := json.Marshal(users)
	if err != nil{
		return "", err
	}
	return string(res), nil
}
