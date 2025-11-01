# ASXY

```
    ___    ____  __  __
   /   |  / __ \/ / / /
  / /| | / /_/ / /_/ / 
 / ___ |/ _, _/ __  /  
/_/  |_/_/ |_/_/ /_/   
```

## Go Log Aggregator

Collect logs from multiple files, filter by severity, and output to console or JSON.

---

## Folder Structure

```
log-aggregator/
├─ main.go           # Entry point of the program
├─ aggregator/
│  ├─ aggregator.go  # Collects and filters logs
├─ utils/
│  ├─ file_reader.go # Reads log files
├─ config/
│  ├─ config.json    # Stores log sources, severity, and output file
```

---

## Example config/config.json

```json
{
  "log_sources": ["app1.log", "app2.log"],
  "severity": ["INFO", "WARN", "ERROR"],
  "output_file": "aggregated_logs.json"
}
```

---

## Terminal Commands

### Go to project folder

```bash
cd log-aggregator
```

### Run project

```bash
go run main.go
```

### Build binary (optional)

```bash
go build -o log-aggregator main.go
./log-aggregator
```

---

## Output

* Logs printed in console
* Aggregated logs saved in `aggregated_logs.json`

---

## Contact

[contact.amish@yahoo.com](mailto:contact.amish@yahoo.com)
