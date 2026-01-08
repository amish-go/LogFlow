gfedcb//code written by amish 
a
package aggregator

import (
	"encoding/json"
	"log-aggregator/utils"
	"strings"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Severity  string `json:"severity"`
	Message   string `json:"message"`
	Source    string `json:"source"`
}

type Aggregator struct {
	logs           []LogEntry
	severityFilter map[string]bool
}

func NewAggregator(severities []string) *Aggregator {
	filter := make(map[string]bool)
	for _, s := range severities {
		filter[strings.ToUpper(s)] = true
	}
	return &Aggregator{
		logs:           []LogEntry{},
		severityFilter: filter,
	}
}

func (a *Aggregator) CollectFromFile(filepath string) error {
	lines, err := utils.ReadLogFile(filepath)
	if err != nil {
		return err
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		var entry LogEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			continue
		}

		entry.Source = filepath

		if len(a.severityFilter) == 0 || a.severityFilter[strings.ToUpper(entry.Severity)] {
			a.logs = append(a.logs, entry)
		}
	}

	return nil
}

func (a *Aggregator) GetLogs() []LogEntry {
	return a.logs
}