# DevSelfie

A highly optimized CLI utility that algorithmically parses VCS logs to synthesize a machine-readable coding productivity profile. 

**Developed by: Gyandeep09**

## Arch & Cap
DevSelfie dynamically computes terminal-based TS visualizations by executing lexical analysis against underlying VCS object models and remote event streams.

- **Local I/O Parsing**: Executes low-level FS traversal against the hidden `.git/logs/HEAD` construct for zero-latency heuristic extraction and TS aggregation.
- **Remote API Sync**: Accepts the `-github <username>` flag to invoke stateless REST API sync, retrieving high-fidelity `PushEvent` datagrams from the GitHub ecosystem.
- **Chrono-Norm**: Programmatically translates UTC-based remote timestamps into the host machine's local TZ to ensure strict chronological accuracy.
- **Data Serialization**: Emits strict standard JSON payloads via the `-json` flag, engineered for seamless integration into downstream CI/CD pipelines, automated dashboards, or secondary runtime envs.

## Install

Execute the following directive to compile and install the executable binary via the standard Go toolchain:
```bash
go install github.com/Gyandeep09/devselfie@latest
```

## Runtime Cmds

**Invoke local parsing engine:**
```bash
devselfie
```

**Target isolated repo structures:**
```bash
devselfie -dir ../target-system-dir
```

**Sync remote telemetry:**
```bash
devselfie -github torvalds
```

**Emit serialized JSON payloads:**
```bash
devselfie -github torvalds -json
```
