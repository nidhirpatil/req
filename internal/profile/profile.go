package profile

type Auth struct {
	Type   string `json:"type"`   // "bearer" or "basic"
	Token  string `json:"token"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
}

type Profile struct {
	Name    string            `json:"name"`
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Auth    Auth              `json:"auth"`
}
