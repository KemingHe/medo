# MEDO Prototype

This prototpy is:

- Built with Google Agent Development Kit
- Contiainerized with Docker
- Deployed to Google Cloud Run via Terraform
- Custom domain managed by Firebase Hosting

## 🔍 Implementation Decisions

After evaluating several options for building the NL2SQL solution for our multi-cloud data orchestration tool, we decided to implement a native LangGraph-based approach. This decision was based on the following considerations:

- Initial attempts with `vanna.ai` encountered multiple compatibility issues:
  - Main dependency incompatibility with Mac ARM64 architecture
  - Sub-dependency incompatibility with Linux when using `devcontainer`
  - Google Colab workaround wouldn't transfer well to web app demonstration (streamlit)

- LangChain/LangGraph was selected as the preferred solution because:
  - More mature and platform-agnostic ecosystem
  - Allows us to simplify by removing non-essential `vanna`-specific NL2SQL features
  - Provides better long-term maintainability without future migration concerns

## ✨ Features

[Brief description of key features]

## 🏗️ Architecture

[Overview of the system architecture]

### Components

[Description of main components]

## 🚀 Setup

[Installation and configuration instructions]

### Prerequisites

[Required dependencies and environment]

## 📋 Usage

[Examples of how to use the tool]

### API Reference

[API documentation]

## 🛠️ Development

[Development guidelines]

### Testing

[Testing procedures]

## 🗺️ Roadmap

[Future development plans]
