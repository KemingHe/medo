# MEDO - Multi-cloud Data Orchestrator Prototype

This prototype is:

- Built with [Google Agent Development Kit](https://google.github.io/adk-docs/) (ADK)
- [Containerized with Docker](Dockerfile) with optimized [build settings](.dockerignore)
- Full [CI/CD pipeline on GCP](cloudbuild.yaml): Cloud Build → Artifact Registry → Cloud Run
- [Custom domain managed by Firebase Hosting](firebase.json)

## ✨ Features

- [Google Search grounding](google_search_agent/agent.py) for enhanced responses
- Containerized deployment with health checks
- Continuous deployment via GCP pipeline

## 🏗️ Architecture

The system is built using a modular agent-based architecture powered by Google's ADK framework:

- ADK agents handle different aspects of data orchestration
- Docker containerization ensures consistent deployment
- GCP Cloud Run provides scalable serverless hosting
- Firebase Hosting manages custom domain and routing

## 🚀 Setup

### Prerequisites

- Python 3.12+ (specified in [.python-version](.python-version))
- [Google ADK](https://google.github.io/adk-docs/) SDK
- Docker (for container builds)
- GCP account (for deployment)

### Local Development

1. Clone this repository
2. Set up Python environment:

    ```shell
    uv venv  # Creates a virtual environment using uv (https://github.com/astral-sh/uv)
    source .venv/bin/activate  # On Windows: .venv\Scripts\activate
    uv pip install -r pyproject.toml  # Installs dependencies in editable mode
    ```

3. Run locally:

    ```shell
    uv run adk web
    ```

### Deployment

The project uses [Cloud Build](cloudbuild.yaml) for automated CI/CD:

1. Push to main branch triggers build
2. Image is stored in Artifact Registry
3. Cloud Run service is updated
4. Firebase routes traffic to the new deployment

## 🛠️ Development

- [Dev container](.devcontainer/devcontainer.json) configuration available for consistent development environments
- Project uses uv for dependency management

## 🗺️ Roadmap

- Expand SQL database connectivity options
- Add support for more cloud providers
- Enhance agent capabilities with Google ADK

## 📋 Technical Notes

### Implementation Decisions

After evaluating several options for building the NL2SQL solution for our multi-cloud data orchestration tool, we decided to implement a native Google ADK-based approach. This decision was based on the following considerations:

- Initial attempts with `vanna.ai` encountered multiple compatibility issues:
  - Main dependency incompatibility with Mac ARM64 architecture
  - Sub-dependency incompatibility with Linux when using `devcontainer`
  - Google Colab workaround wouldn't transfer well to web app demonstration

- Google ADK was selected as the preferred solution because:
  - More debug and visual tracing features
  - Well-integrates with GCP's Vertex AI ecosystem
  - Allows us to simplify by removing non-essential `vanna`-specific NL2SQL features
  - Provides better long-term maintainability without future migration concerns
