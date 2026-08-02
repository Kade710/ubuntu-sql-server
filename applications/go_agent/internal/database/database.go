package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/config"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/inventory"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect opens and verifies a PostgreSQL database connection.
func Connect(cfg config.Config) (*sql.DB, error) {
	connectionURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}

	query := connectionURL.Query()
	query.Set("sslmode", "disable")
	connectionURL.RawQuery = query.Encode()

	db, err := sql.Open("pgx", connectionURL.String())
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

// RegisterServer inserts a new server or updates the existing hostname record.
func RegisterServer(db *sql.DB, server inventory.Server) (int, error) {
	const query = `
	    INSERT INTO server_management.server_inventory (
			hostname,
			ip_address,
			operating_system,
			ram_gb,
			storage_gb,
			gpu,
			motherboard
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (hostname)
		DO UPDATE SET
			ip_address = EXCLUDED.ip_address,
			operating_system = EXCLUDED.operating_system,
			ram_gb = EXCLUDED.ram_gb
		RETURNING id
	`

	var serverID int

	err := db.QueryRow(
		query,
		server.Hostname,
		server.IPAddress,
		server.OperatingSystem,
		server.CPU,
		server.RAMGB,
		server.StorageGB,
		server.GPU,
		server.Motherboard,
	).Scan(&serverID)

	if err != nil {
		return 0, fmt.Errorf("register server: %w", err)
	}

	return serverID, nil
}
