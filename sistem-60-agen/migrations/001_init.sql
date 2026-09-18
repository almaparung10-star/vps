-- Schema awal Sistem 60 Agen (ala Faisal: PRD + Kanban + 60 skill)
CREATE TABLE IF NOT EXISTS prds (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  goal TEXT NOT NULL,
  context TEXT DEFAULT '',
  status TEXT NOT NULL DEFAULT 'draft',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS agents (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  skill TEXT NOT NULL,
  repo TEXT DEFAULT '',
  role TEXT NOT NULL DEFAULT 'developer',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  prd_id INT REFERENCES prds(id) ON DELETE SET NULL,
  title TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'todo',
  severity TEXT DEFAULT 'normal',
  agent_name TEXT DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS sessions (
  id SERIAL PRIMARY KEY,
  task_id INT REFERENCES tasks(id) ON DELETE CASCADE,
  agent_name TEXT NOT NULL,
  context_pct INT NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS logs (
  id SERIAL PRIMARY KEY,
  task_id INT REFERENCES tasks(id) ON DELETE CASCADE,
  message TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
