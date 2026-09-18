// Package prd: dokumen planning wajib sebelum agen jalan.
// Ala Faisal: tanpa PRD agen liar (e-commerce jadi Tokopedia).
package prd

import "sync"

type PRD struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Goal    string `json:"goal"`
	Context string `json:"context"`
	Status  string `json:"status"`
}

type Store struct {
	mu   sync.Mutex
	seq  int
	data map[int]PRD
}

func NewStore() *Store {
	return &Store{data: make(map[int]PRD)}
}

func (s *Store) Create(title, goal, ctx string) PRD {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	p := PRD{ID: s.seq, Title: title, Goal: goal, Context: ctx, Status: "draft"}
	s.data[p.ID] = p
	return p
}

func (s *Store) List() []PRD {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]PRD, 0, len(s.data))
	for _, v := range s.data {
		out = append(out, v)
	}
	return out
}
