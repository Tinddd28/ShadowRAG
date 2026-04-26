package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP       HTTP       `yaml:"http"`
	Logger     Logger     `yaml:"logger"`
	Database   Database   `yaml:"database"`
	LLM        LLM        `yaml:"llm"`
	Embedding  Embedding  `yaml:"embedding"`
	Concurency Concurency `yaml:"concurency"`
}

type HTTP struct {
	Port         string `yaml:"host"`
	IdleTimeout  int    `yaml:"idle_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
}

type Logger struct {
	Level      string `yaml:"level"`
	Dir        string `yaml:"dir"`
	Format     string `yaml:"format"`
	SavingDays int    `yaml:"saving_days"`
}

type Database struct {
	Url     string `yaml:"url"`
	MaxConn int    `yaml:"max_conn"`
}

type LLM struct {
	Model  string `yaml:"model"`
	ApiKey string `yaml:"api_key"`
	Host   string `yaml:"host"`
}

type Embedding struct {
	Model         string `yaml:"model"`
	ApiKey        string `yaml:"api_key"`
	Host          string `yaml:"host"`
	Dim           int    `yaml:"dim"`
	ContextLength int    `yaml:"context_length"`
}

type Concurency struct {
	MaxAsync          int `yaml:"max_sync"`
	MaxParallelInsert int `yaml:"max_parallel_insert"`
	EmbFuncMaxAsync   int `yaml:"emb_func_max_async"`
	EmbBatchNum       int `yaml:"emb_batch_num"`
}

const configPath = "./config.yaml"

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
