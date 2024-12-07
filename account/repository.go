package account

import (
	"context"
	"database/sql"
)

type Repository interface {
	Close()
	PutAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type postgresRepositories struct {
	db *sql.DB
}

func NewPostgresRepositories(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepositories{db}, nil

}

func (r *postgresRepositories) Close() {
	r.db.Close()
}

func (r *postgresRepositories) Ping() {
	r.db.Ping()
}

func (r *postgresRepositories) PutAccount(ctx context.Context, a Account) error {
	query := `
        INSERT INTO accounts (id, name)
        VALUES ($1, $2)
        ON CONFLICT (id) DO UPDATE SET name=$2
    `
	_, err := r.db.ExecContext(ctx, query, a.ID, a.Name)
	return err
}

func (r *postgresRepositories) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	query := `
        SELECT id, name
        FROM accounts
        WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)

	a := &Account{}
	err := row.Scan(&a.ID, &a.Name)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *postgresRepositories) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	query := `
        SELECT id, name
        FROM accounts 
		ORDER BY id DESC
        OFFSET $1 LIMIT $2
    `
	rows, err := r.db.QueryContext(ctx, query, skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {

		a := &Account{}
		err := rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, *a)
	}

	return accounts, nil
}
