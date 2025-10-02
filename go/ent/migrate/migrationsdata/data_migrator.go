package migrationsdata

import (
	"context"
	dbsql "database/sql"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/currency"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
)

func AddSEK(ctx context.Context) func(ctx context.Context) error {
	// Write future migrations to return instead of
	// panicking
	return func(ctx context.Context) error {
		tx := ent.TxFromContext(ctx)
		err := tx.Currency.Create().
			SetCurrencyCode(currency.CurrencyCodeSEK).
			SetDisplay("SEK").
			Exec(ctx)
		if err != nil && !ent.IsConstraintError(err) {
			return err
		}

		return nil
	}
}

func Run(ctx context.Context, sqlDriver *dbsql.DB, client *ent.Client, version int, fn func(ctx context.Context) error) error {
	rows, err := sqlDriver.
		Query(`SELECT version FROM schema_migrations_data`)
	if err != nil {
		return err
	}
	defer rows.Close()

	currentVersion := 0
	for rows.Next() {
		var v int
		err := rows.Scan(&v)
		if err != nil {
			return err
		}
		currentVersion = v
	}
	if currentVersion >= version {
		return nil
	}

	viewCtx := viewer.NewBackgroundContext(ctx)

	tx, err := client.BeginTx(viewCtx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txCtx := ent.NewTxContext(viewCtx, tx)
	err = fn(txCtx)
	if err != nil {
		return utils.Rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return utils.Rollback(tx, err)
	}

	if currentVersion > 0 {
		_, err = sqlDriver.Exec(`UPDATE schema_migrations_data SET version = ?`, version)
		return err
	}

	_, err = sqlDriver.Exec(`INSERT INTO schema_migrations_data (version, dirty) VALUES (?, 0)`, version)
	return err
}
