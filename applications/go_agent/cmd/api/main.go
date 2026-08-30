package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/config"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/database"
)

func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	}); err != nil {
		log.Printf("failed to encode error response: %v", err)
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	log.Println("PostgreSQL connection successful.")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if _, err := fmt.Fprintln(w, `{"status":"ok"}`); err != nil {
			log.Printf("failed to write response: %v", err)
		}
	})

	mux.HandleFunc("GET /api/servers", func(w http.ResponseWriter, r *http.Request) {
		servers, err := database.GetServers(db)
		if err != nil {
			writeJSONError(w, "failed to retrieve servers", http.StatusInternalServerError)
			log.Printf("failed to retrieve servers: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(servers); err != nil {
			log.Printf("failed to encode servers response: %v", err)
		}
	})

	mux.HandleFunc("POST /api/servers", func(w http.ResponseWriter, r *http.Request) {
		var server database.ServerRecord

		if err := json.NewDecoder(r.Body).Decode(&server); err != nil {
			writeJSONError(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if server.Hostname == "" {
			writeJSONError(w, "hostname is required", http.StatusBadRequest)
			return
		}

		created, err := database.CreateServer(db, server)
		if err != nil {
			writeJSONError(w, "failed to create server", http.StatusInternalServerError)
			log.Printf("failed to create server: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(created); err != nil {
			log.Printf("failed to encode created server response: %v", err)
		}
	})

	mux.HandleFunc("GET /api/servers/{id}", func(w http.ResponseWriter, r *http.Request) {
		serverID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || serverID <= 0 {
			writeJSONError(w, "invalid server id", http.StatusBadRequest)
			return
		}

		server, err := database.GetServerByID(db, serverID)
		if err != nil {
			writeJSONError(w, "server not found", http.StatusNotFound)
			log.Printf("failed to retrieve server %d: %v", serverID, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(server); err != nil {
			log.Printf("failed to encode server response: %v", err)
		}
	})

	mux.HandleFunc("PUT /api/servers/{id}", func(w http.ResponseWriter, r *http.Request) {
		serverID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || serverID <= 0 {
			writeJSONError(w, "invalid server id", http.StatusBadRequest)
			return
		}

		var server database.ServerRecord

		if err := json.NewDecoder(r.Body).Decode(&server); err != nil {
			writeJSONError(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if server.Hostname == "" {
			writeJSONError(w, "hostname is required", http.StatusBadRequest)
			return
		}

		updated, err := database.UpdateServer(db, serverID, server)
		if err != nil {
			writeJSONError(w, "server not found", http.StatusNotFound)
			log.Printf("failed to update server %d: %v", serverID, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(updated); err != nil {
			log.Printf("failed to encode updated server response: %v", err)
		}
	})

	mux.HandleFunc("DELETE /api/servers/{id}", func(w http.ResponseWriter, r *http.Request) {
		serverID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || serverID <= 0 {
			writeJSONError(w, "invalid server id", http.StatusBadRequest)
			return
		}

		if err := database.DeleteServer(db, serverID); err != nil {
			if err == sql.ErrNoRows {
				writeJSONError(w, "server not found", http.StatusNotFound)
				return
			}

			writeJSONError(w, "failed to delete server", http.StatusInternalServerError)
			log.Printf("failed to delete server %d: %v", serverID, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Ubuntu SQL Server API listening on :8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
