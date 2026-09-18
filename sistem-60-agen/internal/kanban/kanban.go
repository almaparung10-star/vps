// Package kanban: Todo -> Analisa -> Approve -> Fix -> Done.
// Ala Faisal: 80% coding via Kanban board, agen lapor per-task.
package kanban

import "sync"

const (
	StatusTodo    = "todo"
	StatusAnalisa = "analisa"
	StatusApprove = "approve"
	StatusFix     = "fix"
	StatusDone    = "done"
)

var AllowedNext = map[string][]string{
	StatusTodo:    {StatusAnalisa},
	StatusAnalisa: {StatusApprove, StatusTodo},
	StatusApprove: {StatusFix, StatusTodo},
	StatusFix:     {StatusDone, StatusTodo},
	StatusDone:    {},
}

type Task struct {
	ID        int    `json:"id"`
	PrdID     int    `json:"prd_id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	Severity  string `json:"severity"`
	AgentName string `json:"agent_name"`
}

type Store struct {
	mu   sync.Mutex
	seq  int
	data map[int]Task
}

func NewStore() *Store {
	return &Store{data: make(map[int]Task)}
}

func (s *Store) Create(prdID int, title, severity string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	t := Task{ID: s.seq, PrdID: prdID, Title: title, Status: StatusTodo, Severity: severity}
	if severity == "" {
		t.Severity = "normal"
	}
	s.data[t.ID] = t
	return t
}

func (s *Store) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Task, 0, len(s.data))
	for _, v := range s.data {
		out = append(out, v)
	}
	return out
}

// Move validasi transisi status ala Kanban Faisal.
func (s *Store) Move(id int, next string) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.data[id]
	if !ok {
		return Task{}, false
	}
	for _, n := range AllowedNext[t.Status] {
		if n == next {
			t.Status = next
			s.data[id] = t
			return t, true
		}
	}
	return Task{}, false
}
