// Orchestrator Sistem 60 Agen — MVP.
// Alur Faisal: PRD dulu -> drop ke Todo -> analisa -> approve -> fix -> done.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"sistem-60-agen/internal/agent"
	"sistem-60-agen/internal/kanban"
	"sistem-60-agen/internal/prd"
)

func main() {
	prdStore := prd.NewStore()
	taskStore := kanban.NewStore()
	agents := agent.DefaultAgents()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "ok", "service": "sistem-60-agen"})
	})

	mux.HandleFunc("/api/prds", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, prdStore.List())
		case http.MethodPost:
			var in struct {
				Title   string `json:"title"`
				Goal    string `json:"goal"`
				Context string `json:"context"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Title == "" || in.Goal == "" {
				http.Error(w, "title dan goal wajib (tanpa PRD agen liar)", http.StatusBadRequest)
				return
			}
			writeJSON(w, prdStore.Create(in.Title, in.Goal, in.Context))
		default:
			http.Error(w, "method tidak didukung", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, taskStore.List())
		case http.MethodPost:
			var in struct {
				PrdID    int    `json:"prd_id"`
				Title    string `json:"title"`
				Severity string `json:"severity"`
			}
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Title == "" {
				http.Error(w, "title wajib", http.StatusBadRequest)
				return
			}
			t := taskStore.Create(in.PrdID, in.Title, in.Severity)
			t.AgentName = agent.Route(in.Title)
			writeJSON(w, t)
		default:
			http.Error(w, "method tidak didukung", http.StatusMethodNotAllowed)
		}
	})

	// POST /api/tasks/{id}/move {"status":"analisa"}
	mux.HandleFunc("/api/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/move") || r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		idStr := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/api/tasks/"), "/move")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id tidak valid", http.StatusBadRequest)
			return
		}
		var in struct {
			Status string `json:"status"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Status == "" {
			http.Error(w, "status wajib: analisa/approve/fix/done/todo", http.StatusBadRequest)
			return
		}
		t, ok := taskStore.Move(id, in.Status)
		if !ok {
			http.Error(w, "transisi tidak valid atau task tidak ada", http.StatusBadRequest)
			return
		}
		writeJSON(w, t)
	})

	mux.HandleFunc("/api/agents", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{"total": len(agents), "target": 60, "agents": agents})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("sistem-60-agen jalan di :%s (%d agen terdaftar, target 60)", port, len(agents))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
