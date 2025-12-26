# Documentación Hyperledger Fabric & ERC-20

## Overview
This is an MkDocs documentation project that provides guides for Hyperledger Fabric networks, ERC-20 tokens, and consensus algorithms. The documentation is written in Spanish.

## Project Structure
```
docs/                     # Documentation source files
  consenso/               # Consensus algorithms documentation
  red/                    # Network management guides
  stylesheets/            # Custom CSS styles
  token/                  # ERC-20 token documentation
  index.md                # Homepage
mkdocs.yml                # MkDocs configuration
```

## Technology Stack
- **MkDocs**: Static site generator for documentation
- **MkDocs Material**: Material Design theme for MkDocs
- **Python 3.11**: Runtime environment

## Development
The development server runs on port 5000 using:
```bash
mkdocs serve -a 0.0.0.0:5000
```

## Deployment
The site is deployed as static files. The build command generates the `site/` directory:
```bash
mkdocs build
```

## Recent Changes
- 2025-12-26: Initial setup in Replit environment
