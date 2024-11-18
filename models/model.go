package models

type Route struct {
	Path           []string `yaml:"path"`
	Method         []string `yaml:"method"`
	Description    string   `yaml:"description"`
	Timeout        string   `yaml:"timeout"`
	Retries        int      `yaml:"retries"`
	GeneratedRoute []string `yaml:"generatedRoute"`
	RealUrl        []string `yaml:"realUrl"`
}

type Service struct {
	Name     string  `yaml:"name"`
	URL      string  `yaml:"url"`
	Leader   string  `yaml:"leader"`
	Instance []int   `yaml:"instance"`
	Routes   []Route `yaml:"routes"`
}

type APIConfig struct {
	Name            string    `yaml:"name"`
	Version         string    `yaml:"version"`
	Description     string    `yaml:"description"`
	DefaultRoute    string    `yaml:"defaultRoute"`
	GatewayInfo     string    `yaml:"gatewayInfo"`
	Username        string    `yaml:"username"`
	Password        string    `yaml:"password"`
	LoadBalancing   string    `yaml:"loadBalancing"`
	MainApplication string    `yaml:"mainApplication"`
	Services        []Service `yaml:"services"`
}
type Gateway struct {
	GatewayTag string    `yaml:"gateway"`
	APIConfig  APIConfig `yaml:"api-config"`
}
