// Package agent: registry 60 skill = 60 agen.
// Ala Faisal: 1 skill = 1 agen spesifik (repo + dokumen spesifik) biar hemat token.
// 1 fitur = 1 session, jaga konteks 50-75%.
package agent

type Agent struct {
	Name  string `json:"name"`
	Skill string `json:"skill"`
	Role  string `json:"role"`
	Repo  string `json:"repo"`
}

// DefaultAgents: 17 inti ala Pantau 360, slot sampai 60.
func DefaultAgents() []Agent {
	roles := []Agent{
		{Name: "analis-1", Skill: "analisa-prd", Role: "analis"},
		{Name: "frontend-1", Skill: "developer-frontend", Role: "developer"},
		{Name: "frontend-2", Skill: "automation-qa-frontend", Role: "qa"},
		{Name: "backend-1", Skill: "developer-backend", Role: "developer"},
		{Name: "backend-2", Skill: "developer-sales-pipeline", Role: "developer"},
		{Name: "qa-1", Skill: "qa-e2e-api", Role: "qa"},
		{Name: "qa-2", Skill: "unit-test", Role: "qa"},
		{Name: "qa-3", Skill: "api-testing", Role: "qa"},
		{Name: "qa-4", Skill: "selenium-frontend", Role: "qa"},
		{Name: "qa-5", Skill: "n2n-test", Role: "qa"},
		{Name: "sec-1", Skill: "security-testing", Role: "qa"},
		{Name: "audit-1", Skill: "audit-code", Role: "audit"},
		{Name: "infra-1", Skill: "infra-deploy", Role: "infra"},
		{Name: "infra-2", Skill: "grafana-monitor", Role: "infra"},
		{Name: "bugfix-1", Skill: "bugfix-sentry", Role: "developer"},
		{Name: "docs-1", Skill: "update-dokumen-prd", Role: "docs"},
		{Name: "notulen-1", Skill: "notulen-meeting", Role: "productivity"},
	}
	// Slot cadangan sampai 60 (jaksa, intel, ekonomi, politik, sentinel, dst).
	extra := []string{
		"sentinel-sosmed", "sentinel-sentimen", "sentinel-cluster", "sentinel-report",
		"mockup-sales", "seo-tool", "osint-tool", "crm-handoff", "scheduler-pagi",
		"weekly-report", "telegram-lapor",
	}
	for _, e := range extra {
		roles = append(roles, Agent{Name: e, Skill: e, Role: "support"})
	}
	return roles
}

// Route: pilih agen berdasarkan kata kunci judul task (sederhana, nanti bisa pakai LLM).
func Route(title string) string {
	t := title
	contains := func(s, sub string) bool {
		return len(s) >= len(sub) && (func() bool {
			for i := 0; i+len(sub) <= len(s); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
			return false
		})()
	}
	lower := ""
	for _, r := range t {
		if r >= 'A' && r <= 'Z' {
			lower += string(r + 32)
		} else {
			lower += string(r)
		}
	}
	switch {
	case contains(lower, "logo") || contains(lower, "frontend") || contains(lower, "tampilan"):
		return "frontend-1"
	case contains(lower, "api") || contains(lower, "backend"):
		return "backend-1"
	case contains(lower, "bug") || contains(lower, "sentry") || contains(lower, "peang"):
		return "bugfix-1"
	case contains(lower, "test") || contains(lower, "qa"):
		return "qa-1"
	case contains(lower, "deploy") || contains(lower, "server") || contains(lower, "grafana"):
		return "infra-1"
	case contains(lower, "meeting") || contains(lower, "notulen"):
		return "notulen-1"
	default:
		return "analis-1"
	}
}
