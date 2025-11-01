package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"log-aggregator/aggregator"
	"log-aggregator/config"
)

type Config struct {
	LogSources     []string `json:"log_sources"`
	SeverityLevels []string `json:"severity_levels"`
	OutputFile     string   `json:"output_file"`
}

func main() {
	configData, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Fatal("Failed to read config:", err)
	}

	var cfg Config
	if err := json.Unmarshal(configData, &cfg); err != nil {
		log.Fatal("Failed to parse config:", err)
	}

	agg := aggregator.NewAggregator(cfg.SeverityLevels)

	for _, source := range cfg.LogSources {
		if err := agg.CollectFromFile(source); err != nil {
			fmt.Printf("Warning: Failed to collect from %s: %v\n", source, err)
		}
	}

	logs := agg.GetLogs()

	output, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		log.Fatal("Failed to marshal logs:", err)
	}

	fmt.Println(string(output))

	if err := os.WriteFile(cfg.OutputFile, output, 0644); err != nil {
		log.Fatal("Failed to write output file:", err)
	}

	fmt.Printf("\nAggregated %d logs to %s\n", len(logs), cfg.OutputFile)
}