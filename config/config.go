package config

import (
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Flags struct {
	DryRun  bool
	Verbose bool
	Info    bool
}

type Config struct {
	DryRun  bool `json:"-"`
	Verbose bool `json:"-"`
	Info    bool `json:"-"`
}

func LoadConfig(flags Flags) *Config {
	_ = godotenv.Load(".envrc")
	cfg := &Config{
		DryRun:  flags.DryRun,
		Verbose: flags.Verbose,
		Info:    flags.Info,
	}

	cfg.FillEnvs(".")

	return cfg
}

func (c *Config) FillEnvs(dir string) {
	_ = godotenv.Load(fmt.Sprintf("%s/.envrc", dir))

	// prefix := os.Getenv("GUS_")
	// if prefix != "" {
	// 	c.Commit.Output.Prefix = prefix
	// }
}
