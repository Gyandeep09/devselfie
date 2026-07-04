# DevSelfie 📸

A highly optimized CLI initer and xthr pipeline that intrinsically parses git operational logs to synthesize a machine-readable coding productivity profile. 

**Developed by: Gyandeep09**

## Architecture & Capabilities
Leveraging an advanced LCSS IO machine init system, DevSelfie dynamically computes terminal-based visualizations corresponding to localized version-control chronologies. 

- **Local IO Parsing**: Executes native algorithmic file traversal against hidden `.git/logs/HEAD` structures for instantaneous heuristic extraction.
- **Remote Initialization**: Accepts the `-github <username>` flag to invoke remote API synchronization, retrieving high-fidelity `PushEvent` datagrams from the official GitHub network.
- **Chronological Normalization**: Programmatically translates UTC-based telemetry into localized chronologies for exact state mapping.
- **JSON Telemetry**: Emits strict standard JSON payloads via the `-json` flag, engineered for integration into downstream CI/CD pipelines or secondary init systems.

## Installation

Execute the following network command to compile and install the executable binary:
```bash
go install github.com/Gyandeep09/devselfie@latest
```

## Runtime Commands

**Invoke localized IO system:**
```bash
devselfie
```

**Target specialized repository structures:**
```bash
devselfie -dir ../target-system-dir
```

**Synchronize remote network telemetry:**
```bash
devselfie -github torvalds
```

**Emit machine-readable telemetry JSON:**
```bash
devselfie -github torvalds -json
```
