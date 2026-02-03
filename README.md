# Secure Go API with GoSec: Automated Vulnerability Remediation

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)
![Scanner](https://img.shields.io/badge/Scanner-Gosec-red)
![Status](https://img.shields.io/badge/Status-Remediated-success)

## Project Overview
This repository serves as a **Proof of Concept (PoC)** for automated Application Security in backend microservices. 

It demonstrates the lifecycle of a vulnerability—from introduction to detection, and finally, remediation. Unlike standard penetration testing reports, this project focuses on **Engineering Remediation**: fixing the root cause in the code.

**The Problem Solved:** Eliminates high-risk vulnerabilities (like SQL Injection) at the build stage, ensuring no insecure code ever reaches production.

---

## Technology Stack
* **Language:** Go (Golang)
* **Database:** SQLite (SQL Driver)
* **Scanner:** [Gosec](https://github.com/securego/gosec) (Golang Security Checker)
* **CI/CD:** GitHub Actions

---

## The Vulnerability: SQL Injection (SQLi)

### The Flaw (Rule G201)
The application initially used **String Concatenation** to build SQL queries. This is a critical flaw that allows attackers to manipulate database logic.

```go
// INSECURE CODE
query := fmt.Sprintf("SELECT apikey FROM users WHERE id = %s", id)
// Result: Attackers can inject "1 OR 1=1" to dump the database.
```

### Main Fix
```go
// SECURE CODE
query := "SELECT apikey FROM users WHERE id = ?"
// Result: Input is sanitized by the driver. Gosec scan passes.
```

### Automated Governance

The repository includes a GitHub Action workflow .github/workflows/go-security.yaml that enforces the following rules:

    Scan: Runs gosec on every Pull Request against the main branch.

    Report: Parses the JSON output and comments on the PR with a table of vulnerabilities.

    Block: If Severity >= HIGH, the merge button is disabled.

 Usage Guide
1. Run the Vulnerable Server
Bash

go mod tidy
go run main.go

## Server listens on :8080

2. Run Security Scan Locally
Bash

## Install Gosec
go install [github.com/securego/gosec/v2/cmd/gosec@latest](https://github.com/securego/gosec/v2/cmd/gosec@latest)

## Run Scan
gosec ./...