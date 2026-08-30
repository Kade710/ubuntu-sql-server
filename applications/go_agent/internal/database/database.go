package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/config"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/hardware"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/health"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/inventory"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/maintenance"
	agentnetwork "github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/network"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/osinfo"

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

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

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
			cpu,
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
			cpu = EXCLUDED.cpu,
			ram_gb = EXCLUDED.ram_gb,
			storage_gb = EXCLUDED.storage_gb,
			gpu = EXCLUDED.gpu,
			motherboard = EXCLUDED.motherboard
		RETURNING id
	`

	var serverID int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.QueryRowContext(
		ctx,
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

// UpsertNetworkInterfaces inserts or updates network interface records.
func UpsertNetworkInterfaces(
	db *sql.DB,
	serverID int,
	interfaces []agentnetwork.Interface,
) error {
	const query = `
		INSERT INTO server_management.network_interfaces (
			server_id,
			interface_name,
			mac_address,
			ip_address,
			network_type,
			speed_mbps
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (server_id, interface_name)
		DO UPDATE SET
			mac_address = EXCLUDED.mac_address,
			ip_address = EXCLUDED.ip_address,
			network_type = EXCLUDED.network_type,
			speed_mbps = EXCLUDED.speed_mbps
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin network transaction: %w", err)
	}
	defer tx.Rollback()

	for _, iface := range interfaces {
		_, err := tx.ExecContext(
			ctx,
			query,
			serverID,
			iface.Name,
			iface.MACAddress,
			iface.IPAddress,
			iface.NetworkType,
			iface.SpeedMbps,
		)
		if err != nil {
			return fmt.Errorf(
				"save network interface %s: %w",
				iface.Name,
				err,
			)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit network transaction: %w", err)
	}

	return nil
}

// AddMaintenanceLog saves a maintenance entry and returns its database ID.
func AddMaintenanceLog(
	db *sql.DB,
	serverID int,
	logEntry maintenance.Log,
) (int, error) {
	const query = `
		INSERT INTO server_management.maintenance_logs (
			server_id,
			action,
			description,
			performed_by
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var logID int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.QueryRowContext(
		ctx,
		query,
		serverID,
		logEntry.Action,
		logEntry.Description,
		logEntry.PerformedBy,
	).Scan(&logID)

	if err != nil {
		return 0, fmt.Errorf("add maintenance log: %w", err)
	}

	return logID, nil
}

func UpsertOperatingSystem(
	db *sql.DB,
	serverID int,
	info osinfo.Info,
) (int, error) {
	const query = `
		INSERT INTO server_management.operating_systems (
			server_id,
			distribution,
			version,
			kernel,
			architecture
		)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (server_id)
		DO UPDATE SET
			distribution = EXCLUDED.distribution,
			version = EXCLUDED.version,
			kernel = EXCLUDED.kernel,
			architecture = EXCLUDED.architecture
		RETURNING id
	`

	var osID int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.QueryRowContext(
		ctx,
		query,
		serverID,
		info.Distribution,
		info.Version,
		info.Kernel,
		info.Architecture,
	).Scan(&osID)

	if err != nil {
		return 0, fmt.Errorf("upsert operating system: %w", err)
	}

	return osID, nil
}

// UpsertHardwareComponents inserts or updates hardware component records.
func UpsertHardwareComponents(
	db *sql.DB,
	serverID int,
	components []hardware.Component,
) error {
	const query = `
		INSERT INTO server_management.hardware_components (
			server_id,
			component_type,
			manufacturer,
			model,
			specification
		)
		VALUES ($1, $2, NULLIF($3, ''), NULLIF($4, ''), NULLIF($5, ''))
		ON CONFLICT (server_id, component_type)
		DO UPDATE SET
			manufacturer = COALESCE(EXCLUDED.manufacturer, hardware_components.manufacturer),
			model = COALESCE(EXCLUDED.model, hardware_components.model),
			specification = COALESCE(EXCLUDED.specification, hardware_components.specification)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin hardware transaction: %w", err)
	}
	defer tx.Rollback()

	for _, component := range components {
		_, err := tx.ExecContext(
			ctx,
			query,
			serverID,
			component.Type,
			component.Manufacturer,
			component.Model,
			component.Specification,
		)
		if err != nil {
			return fmt.Errorf(
				"save hardware component %s: %w",
				component.Type,
				err,
			)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit hardware transaction: %w", err)
	}

	return nil
}

// AddHealthCheck saves a server health snapshot and returns its database ID.
func AddHealthCheck(
	db *sql.DB,
	serverID int,
	status health.Status,
) (int, error) {
	const query = `
		INSERT INTO server_management.health_checks (
			server_id,
			load_average,
			memory_percent,
			disk_percent,
			uptime_hours,
			overall_status
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var healthCheckID int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.QueryRowContext(
		ctx,
		query,
		serverID,
		status.Load1,
		status.MemoryPercent,
		status.DiskPercent,
		status.UptimeHours,
		status.Overall,
	).Scan(&healthCheckID)

	if err != nil {
		return 0, fmt.Errorf("add health check: %w", err)
	}

	return healthCheckID, nil
}

type HealthCheckRecord struct {
	ID            int
	LoadAverage   float64
	MemoryPercent float64
	DiskPercent   float64
	UptimeHours   float64
	OverallStatus string
	CreatedAt     string
}

func GetRecentHealthChecks(
	db *sql.DB,
	serverID int,
	limit int,
) ([]HealthCheckRecord, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("health check limit must be greater than zero")
	}

	if limit > 100 {
		limit = 100
	}
	const query = `
		SELECT
			id,
			load_average,
			memory_percent,
			disk_percent,
			uptime_hours,
			overall_status,
			created_at
		FROM server_management.health_checks
		WHERE server_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, serverID, limit)
	if err != nil {
		return nil, fmt.Errorf("get recent health checks: %w", err)
	}
	defer rows.Close()

	var records []HealthCheckRecord

	for rows.Next() {
		var record HealthCheckRecord

		err := rows.Scan(
			&record.ID,
			&record.LoadAverage,
			&record.MemoryPercent,
			&record.DiskPercent,
			&record.UptimeHours,
			&record.OverallStatus,
			&record.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan health checks: %w", err)
		}

		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read health checks: %w", err)
	}

	return records, nil
}

type ServerRecord struct {
	ID              int		`json:"id"`
	Hostname        string  `json:"hostname"`
	IPAddress       string	`json:"ip_address"`
	OperatingSystem string 	`json:"operating_system"`
	CPU             string	`json:"cpu"`
	RAMGB           int		`json:"ram_gb"`
	StorageGB       int		`json:"storage_gb"`
	GPU             string	`json:"gpu"`
	Motherboard     string	`json:"motherboard"`
}

func GetServers(db *sql.DB) ([]ServerRecord, error) {
	const query = `
		SELECT
			id,
			hostname,
			ip_address,
			operating_system,
			cpu,
			ram_gb,
			storage_gb,
			gpu,
			motherboard
		FROM server_management.server_inventory
		ORDER BY hostname, id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get servers: %w", err)
	}
	defer rows.Close()

	var servers []ServerRecord

	for rows.Next() {
		var server ServerRecord

		if err := rows.Scan(
			&server.ID,
			&server.Hostname,
			&server.IPAddress,
			&server.OperatingSystem,
			&server.CPU,
			&server.RAMGB,
			&server.StorageGB,
			&server.GPU,
			&server.Motherboard,
		); err != nil {
			return nil, fmt.Errorf("scan server inventory: %w", err)
		}

		servers = append(servers, server)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read server inventory: %w", err)
	}

	return servers, nil
}

func AddAlertEvent(
	db *sql.DB,
	serverID int,
	previousStatus string,
	newStatus string,
	title string,
	message string,
) (int, error) {
	const query = `
		INSERT INTO server_management.alert_events (
			server_id,
			previous_status,
			new_status,
			title,
			message
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var alertID int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err :=
		db.QueryRowContext(
			ctx,
			query,
			serverID,
			previousStatus,
			newStatus,
			title,
			message,
		).Scan(&alertID)

	if err != nil {
		return 0, fmt.Errorf("add alert event: %w", err)
	}

	return alertID, nil
}
