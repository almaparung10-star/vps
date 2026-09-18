# Sistem 60 Agen — Go + Postgres
Tiru arsitektur Mas Faisal (Builder Story 4): manusia jadi atasan via PRD + Kanban, 60 skill = 60 agen.

## Struktur
- `cmd/api` : REST API orchestrator (Kanban + PRD + Agent registry)
- `internal/kanban` : Todo -> Analisa -> Approve -> Fix -> Done
- `internal/prd` : PRD dokumen planning
- `internal/agent` : registry 60 skill + runner (goroutine)
- `migrations` : schema Postgres
- `docker-compose.yml` : postgres 16 + api

## Jalan lokal
```bash
cp .env.example .env
docker compose up -d db
go run ./cmd/api
# health: http://localhost:8080/health
```

## API MVP
- `GET /health`
- `GET /api/prds` `POST /api/prds`
- `GET /api/tasks` `POST /api/tasks` `POST /api/tasks/{id}/move`
- `GET /api/agents`

Alur ala Faisal: buat PRD dulu -> drop ke Todo -> agen analisa -> manusia Approve -> fixing -> Done + lapor Telegram (menyusul).
