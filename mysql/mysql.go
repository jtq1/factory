package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Configura tu DSN (Data Source Name)
	dsn := "lloretab_admin:aragontalleres2023!@tcp(185.162.171.114:3306)/lloretab_talleres_gestion"

	// Abre la conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("no se pudo conectar: %w", err)
	}

	fmt.Println("Conexión exitosa a la base de datos.")
	return db, nil
}
