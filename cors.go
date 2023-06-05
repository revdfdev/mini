package mini

type CorsConfig struct {
	AllowedOrigins []string
	AllowedHeaders []string
	AllowedMethods []string
}

func (mini *Mini) SetCorsConfig(config *CorsConfig) {
	mini.corsConfig = config
	mini.corConfigSet = true
}
