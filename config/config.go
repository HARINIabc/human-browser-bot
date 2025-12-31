package config

type Config struct {
	StartHour  int
	EndHour    int
	MaxActions int
	DemoText   string
}

func Load() *Config {
	return &Config{
		StartHour:  9,
		EndHour:   18,
		MaxActions: 3,
		DemoText:  "Hi! I am a human-like browser bot",
	}
}
