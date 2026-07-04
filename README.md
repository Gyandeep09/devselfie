# DevSelfie 📸

A high-performance command-line tool built in **Go** that parses your git logs and visualizes your coding productivity as beautiful ASCII art. 

**Developed by: Gyandeep09**

## What is it?
Instead of a flat green contribution graph, DevSelfie dynamically generates a terminal layout based on the times of day you commit code. 

Are you pulling all-nighters? DevSelfie will catch that and give you the **Night Owl** badge. Coding early before work? You get the **Morning Bird** badge!

## Features
- **Local Parsing**: Extremely fast, natively reads your hidden `.git/logs/HEAD` file.
- **GitHub Integration**: Don't want to parse locally? Use `-github <username>` to fetch and analyze a user's recent public events via the official GitHub API. 
- **Timezone Aware**: Automatically converts remote UTC API timestamps to your local timezone for perfect accuracy.
- **JSON Export**: Hook DevSelfie into your CI/CD pipelines or frontend apps by outputting raw data using the `-json` flag!

## Installation

If you have Go installed, you can simply run:
```bash
go install github.com/Gyandeep09/devselfie@latest
```
*(Note: Replace the URL above with your actual GitHub repository link once published!)*

## Usage

**Analyze your current local git repository:**
```bash
devselfie
```

**Analyze a specific local repository:**
```bash
devselfie -dir ../my-awesome-project
```

**Analyze any public GitHub User:**
```bash
devselfie -github torvalds
```

**Output raw data as JSON:**
```bash
devselfie -github torvalds -json
```

## The Badges
- 🌅 **Morning Bird** (5 AM - 11 AM)
- ☀️ **Afternoon Hustler** (12 PM - 5 PM)
- 🌆 **Evening Coder** (6 PM - 10 PM)
- 🦉 **Night Owl** (11 PM - 4 AM)

---
*Built with ❤️ in Go.*
