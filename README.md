# DevSelfie

A highly optimized Command-Line Interface (CLI) utility that algorithmically parses version-control logs to synthesize a machine-readable coding productivity profile. 

**Developed by: Gyandeep09**

## Architecture & Capabilities
DevSelfie dynamically computes terminal-based temporal visualizations by executing lexical analysis against underlying Git object models and remote event streams.

- **Local Log Parsing**: Executes low-level filesystem traversal against the hidden `.git/logs/HEAD` construct for zero-latency heuristic extraction and temporal aggregation.
- **Remote API Synchronization**: Accepts the `-github <username>` flag to invoke stateless REST API synchronization, retrieving high-fidelity `PushEvent` datagrams from the GitHub ecosystem.
- **Chronological Normalization**: Programmatically translates UTC-based remote timestamps into the host machine's local time zone to ensure strict chronological accuracy.
- **Data Serialization**: Emits strict standard JSON payloads via the `-json` flag, engineered for seamless integration into downstream CI/CD pipelines, automated dashboards, or secondary runtime environments.

## Installation

Execute the following directive to compile and install the executable binary via the standard Go toolchain:
```bash
go install github.com/Gyandeep09/devselfie@latest
```

## Runtime Commands

**Invoke localized parsing engine:**
```bash
devselfie
```

**Target isolated repository structures:**
```bash
devselfie -dir ../target-system-dir
```

**Synchronize remote telemetry:**
```bash
devselfie -github torvalds
```

**Emit serialized JSON payloads:**
```bash
devselfie -github torvalds -json
```
