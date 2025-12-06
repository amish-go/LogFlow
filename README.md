# LogFlow

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

Collect logs from multiple files, filter by severity, and output to console or JSON.

</div>

---

## Features

- Collect logs from multiple sources
- Filter by severity level (INFO, WARN, ERROR)
- Console and JSON output
- Simple configuration

## Installation

```bash
cd logflow
go run main.go
```

## Configuration

Edit `config/config.json`:

```json
{
  "log_sources": ["app1.log", "app2.log"],
  "severity": ["INFO", "WARN", "ERROR"],
  "output_file": "aggregated_logs.json"
}
```

## Usage

```bash
# Run
go run main.go

# Build binary
go build -o logflow main.go
./logflow
```

## Output

- Logs printed in console
- Aggregated logs saved in `aggregated_logs.json`

## License

MIT License

## Contact

Email: contact.amish@yahoo.com
