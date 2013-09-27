package iglo

type API struct {
	Version        string          `json:"_version"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Metadata       Metadata        `json:"metadata"`
	ResourceGroups []ResourceGroup `json:"resourceGroups"`
}

type Host struct {
	Value string `json:"value"`
}

type Format struct {
	Value string `json:"value"`
}

type Metadata struct {
	Format Format `json:"FORMAT"`
	Host   Host   `json:"HOST"`
}

type ResourceGroup struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Resources   []Resource `json:"resources"`
}

type Header struct {
	Value string `json:"value"`
}

type Model struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Headers     map[string]Header `json:"headers"`
	Body        string            `json:"body"`
	Schema      string            `json:"schema"`
}

type Parameter struct {
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Required    bool     `json:"required"`
	Default     string   `json:"default"`
	Example     string   `json:"example"`
	Values      []string `json:"values"`
}

type Example struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Requests    []Request  `json:"requests"`
	Responses   []Response `json:"responses"`
}

type Request struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Headers     map[string]Header `json:"headers"`
	Body        string            `json:"body"`
	Schema      string            `json:"schema"`
}

type Response struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Headers     map[string]Header `json:"headers"`
	Body        string            `json:"body"`
	Schema      string            `json:"schema"`
}

type Action struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Method      string               `json:"method"`
	Parameters  map[string]Parameter `json:"parameters"`
	Headers     map[string]Header    `json:"headers"`
	Examples    []Example            `json:"examples"`
}

type Resource struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	UriTemplate string               `json:"uriTemplate"`
	Model       Model                `json:"model"`
	Parameters  map[string]Parameter `json:"parameters"`
	Headers     map[string]Header    `json:"headers"`
	Actions     []Action             `json:"actions"`
}
