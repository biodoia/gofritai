package providers

// Provider represents a free tier service provider
type Provider struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Category    Category          `json:"category"`
	FreeTier    FreeTier          `json:"free_tier"`
	RequiresCC  bool              `json:"requires_cc"`
	URL         string            `json:"url"`
	APIEndpoint string            `json:"api_endpoint,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type Category string

const (
	CategoryCompute   Category = "compute"
	CategoryDatabase  Category = "database"
	CategoryStorage   Category = "storage"
	CategoryLLM       Category = "llm"
	CategoryAuth      Category = "auth"
	CategoryMonitor   Category = "monitoring"
	CategoryServerless Category = "serverless"
)

type FreeTier struct {
	Description string        `json:"description"`
	Limits      []Limit       `json:"limits"`
	Duration    string        `json:"duration"` // "forever", "12mo", "trial"
}

type Limit struct {
	Resource string  `json:"resource"`
	Amount   float64 `json:"amount"`
	Unit     string  `json:"unit"`
	Period   string  `json:"period,omitempty"` // "day", "month", "forever"
}

// Registry holds all known providers
var Registry = []Provider{
	// Compute
	{
		ID:       "oracle-arm",
		Name:     "Oracle Cloud ARM",
		Category: CategoryCompute,
		FreeTier: FreeTier{
			Description: "4 OCPU + 24GB RAM forever",
			Limits: []Limit{
				{Resource: "vcpu", Amount: 4, Unit: "cores"},
				{Resource: "memory", Amount: 24, Unit: "GB"},
				{Resource: "storage", Amount: 200, Unit: "GB"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://cloud.oracle.com",
	},
	{
		ID:       "fly-io",
		Name:     "Fly.io",
		Category: CategoryCompute,
		FreeTier: FreeTier{
			Description: "3 shared VMs + 160GB transfer",
			Limits: []Limit{
				{Resource: "vms", Amount: 3, Unit: "instances"},
				{Resource: "transfer", Amount: 160, Unit: "GB", Period: "month"},
			},
			Duration: "forever",
		},
		RequiresCC: true,
		URL:        "https://fly.io",
	},
	{
		ID:       "render",
		Name:     "Render",
		Category: CategoryCompute,
		FreeTier: FreeTier{
			Description: "750h/mo, spins down",
			Limits: []Limit{
				{Resource: "hours", Amount: 750, Unit: "hours", Period: "month"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://render.com",
	},
	{
		ID:       "cloudflare-workers",
		Name:     "Cloudflare Workers",
		Category: CategoryServerless,
		FreeTier: FreeTier{
			Description: "100k req/day + 10ms CPU",
			Limits: []Limit{
				{Resource: "requests", Amount: 100000, Unit: "requests", Period: "day"},
				{Resource: "cpu", Amount: 10, Unit: "ms"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://workers.cloudflare.com",
	},

	// LLM APIs
	{
		ID:       "groq",
		Name:     "Groq",
		Category: CategoryLLM,
		FreeTier: FreeTier{
			Description: "30 req/min, Llama/Mixtral/Whisper",
			Limits: []Limit{
				{Resource: "requests", Amount: 30, Unit: "requests", Period: "minute"},
			},
			Duration: "forever",
		},
		RequiresCC:  false,
		URL:         "https://groq.com",
		APIEndpoint: "https://api.groq.com/openai/v1",
	},
	{
		ID:       "google-ai-studio",
		Name:     "Google AI Studio",
		Category: CategoryLLM,
		FreeTier: FreeTier{
			Description: "60 req/min, Gemini 2.5 Pro/Flash",
			Limits: []Limit{
				{Resource: "requests", Amount: 60, Unit: "requests", Period: "minute"},
			},
			Duration: "forever",
		},
		RequiresCC:  false,
		URL:         "https://aistudio.google.com",
		APIEndpoint: "https://generativelanguage.googleapis.com/v1beta",
	},
	{
		ID:       "openrouter",
		Name:     "OpenRouter",
		Category: CategoryLLM,
		FreeTier: FreeTier{
			Description: "Free models available",
			Limits:      []Limit{},
			Duration:    "forever",
		},
		RequiresCC:  false,
		URL:         "https://openrouter.ai",
		APIEndpoint: "https://openrouter.ai/api/v1",
	},

	// Databases
	{
		ID:       "supabase",
		Name:     "Supabase",
		Category: CategoryDatabase,
		FreeTier: FreeTier{
			Description: "500MB Postgres + 50k auth + realtime",
			Limits: []Limit{
				{Resource: "storage", Amount: 500, Unit: "MB"},
				{Resource: "auth_users", Amount: 50000, Unit: "users"},
				{Resource: "bandwidth", Amount: 2, Unit: "GB", Period: "month"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://supabase.com",
	},
	{
		ID:       "turso",
		Name:     "Turso",
		Category: CategoryDatabase,
		FreeTier: FreeTier{
			Description: "9GB + 500M reads",
			Limits: []Limit{
				{Resource: "storage", Amount: 9, Unit: "GB"},
				{Resource: "reads", Amount: 500000000, Unit: "reads", Period: "month"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://turso.tech",
	},
	{
		ID:       "upstash",
		Name:     "Upstash",
		Category: CategoryDatabase,
		FreeTier: FreeTier{
			Description: "10k cmd/day Redis",
			Limits: []Limit{
				{Resource: "commands", Amount: 10000, Unit: "commands", Period: "day"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://upstash.com",
	},

	// Storage
	{
		ID:       "cloudflare-r2",
		Name:     "Cloudflare R2",
		Category: CategoryStorage,
		FreeTier: FreeTier{
			Description: "10GB storage + ZERO egress",
			Limits: []Limit{
				{Resource: "storage", Amount: 10, Unit: "GB"},
				{Resource: "egress", Amount: 0, Unit: "cost"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://www.cloudflare.com/r2",
	},
	{
		ID:       "backblaze-b2",
		Name:     "Backblaze B2",
		Category: CategoryStorage,
		FreeTier: FreeTier{
			Description: "10GB storage",
			Limits: []Limit{
				{Resource: "storage", Amount: 10, Unit: "GB"},
			},
			Duration: "forever",
		},
		RequiresCC: false,
		URL:        "https://www.backblaze.com/b2",
	},
}

// GetByID returns a provider by ID
func GetByID(id string) *Provider {
	for _, p := range Registry {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

// GetByCategory returns all providers in a category
func GetByCategory(cat Category) []Provider {
	var result []Provider
	for _, p := range Registry {
		if p.Category == cat {
			result = append(result, p)
		}
	}
	return result
}

// GetNoCC returns providers that don't require credit card
func GetNoCC() []Provider {
	var result []Provider
	for _, p := range Registry {
		if !p.RequiresCC {
			result = append(result, p)
		}
	}
	return result
}
