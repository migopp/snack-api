package api

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	DB *sql.DB
}

func CreatePostgresStore() (*PostgresStore, error) {
	connStr := "postgres://postgres:test@localhost:5432/snackapi?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		DB: db,
	}, nil
}

func (p *PostgresStore) Close() error {
	return p.DB.Close()
}

func (p *PostgresStore) CreateSnackerTable() error {
	query := `CREATE TABLE IF NOT EXISTS snackers (
        snacker_id SERIAL PRIMARY KEY,
        first_name varchar(50),
        last_name varchar(50),
        snack varchar(50),
        hearts integer
    );`
	_, err := p.DB.Exec(query)
	return err
}

func (p *PostgresStore) CreateSnacker(s *SnackerRegistration) (uint32, error) {
	query := `INSERT INTO snackers (first_name, last_name, snack, hearts)
        VALUES ($1, $2, $3, 0)
        RETURNING snacker_id;`
	var id uint32
	err := p.DB.QueryRow(query, s.FirstName, s.LastName, s.Snack).Scan(&id)
	return id, err
}

func (p *PostgresStore) DeleteSnacker(id uint32) error {
	query := `DELETE FROM snackers WHERE snacker_id = $1;`
	_, err := p.DB.Exec(query, id)
	return err
}

func (p *PostgresStore) UpdateSnacker(s *Snacker) error {
	query := `UPDATE snackers
        SET first_name = $1,
            last_name = $2,
            snack = $3,
            hearts = $4
        WHERE snacker_id = $5;`
	_, err := p.DB.Exec(query,
		s.FirstName,
		s.LastName,
		s.Snack,
		s.Hearts,
		s.ID)
	return err
}

func (p *PostgresStore) FindSnacker(id uint32) (*Snacker, error) {
	query := `SELECT * FROM snackers WHERE snacker_id = $1;`
	snacker := new(Snacker)
	err := p.DB.QueryRow(query, id).Scan(
		&snacker.ID,
		&snacker.FirstName,
		&snacker.LastName,
		&snacker.Snack,
		&snacker.Hearts)
	return snacker, err
}

func (p *PostgresStore) FetchSnackers() ([]*Snacker, error) {
	query := `SELECT * FROM snackers;`
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var snackers []*Snacker
	for rows.Next() {
		snacker := new(Snacker)
		rows.Scan(
			&snacker.ID,
			&snacker.FirstName,
			&snacker.LastName,
			&snacker.Snack,
			&snacker.Hearts)
		snackers = append(snackers, snacker)
	}

	return snackers, nil
}
