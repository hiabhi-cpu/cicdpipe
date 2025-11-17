# Mini CI/CD Pipeline in Go

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)

A lightweight, self-hosted CI/CD pipeline written in Go. This tool automates the process of building and publishing Docker images directly from a `git push` to a GitHub repository.

## ✨ Core Features

* **Automated GitHub Webhook Management**: Automatically checks for and creates a repository webhook using the GitHub API on startup.
* **Git Integration**: Clones the repository on the first push and performs a `git pull` on all subsequent pushes.
* **Automated Docker Workflow**: Automatically builds a new Docker image from the repo's `Dockerfile` and publishes it to Docker Hub.
* **Robust Notifications**: Sends detailed email notifications for successful actions, new webhook creation, and critical errors (including full Docker build logs).
* **Concurrent & Graceful Shutdown**: Runs a concurrent webserver to listen for webhooks and uses channels to manage state and handle graceful `ctrl+c` shutdowns.
* **Local Environment Validation**: Checks for necessary local dependencies (Docker, Git, Ngrok) on startup.

## ⚙️ How It Works

1.  **Initialization**: You run the application locally.
2.  **Validation**: The app checks that Docker, Git, and Ngrok are installed.
3.  **Tunneling**: An [Ngrok](https://ngrok.com/) tunnel is started to get a public-facing URL.
4.  **Webhook Setup**: The app uses the GitHub API and your PAT to create a new webhook on your target repository, pointing to the new Ngrok URL.
5.  **Listening**: The Go server listens for `POST` requests from the GitHub webhook.
6.  **Trigger**: You `git push` a new commit to your repository.
7.  **Action**: GitHub sends a webhook payload to the Go server. The server then:
    * Pulls the latest code.
    * Builds the Docker image.
    * Publishes the image to Docker Hub.
8.  **Notification**: An email is sent to you with the status (success or failure, including error logs).

## 🔧 Setup & Installation

### 1. Prerequisites

Before you begin, you will need:
* [Go](https://go.dev/doc/install) (1.18+ recommended)
* [Docker](https://www.docker.com/get-started)
* [Git](https://git-scm.com/downloads)
* An [Ngrok Account](https://dashboard.ngrok.com/signup) and Auth Token
* A [GitHub Personal Access Token (PAT)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) with `repo` (full control) scopes.
* A [Docker Hub](https://hub.docker.com/) account.
* An email account (and app password if using Gmail/Google) for SMTP notifications.

### 2. Configuration

1.  Clone this repository:
    ```sh
    git clone [https://github.com/hiabhi-cpu/MiniCI_pipeline.git](https://github.com/hiabhi-cpu/MiniCI_pipeline.git)
    cd MiniCI_pipeline
    ```

2.  Create an `.env` file. You can copy the example:
    ```sh
    cp .env.example .env
    ```

3.  Edit the `.env` file with your credentials:
    ```ini
    # GitHub Credentials
    GITHUB_PAT=YOUR_GITHUB_PERSONAL_ACCESS_TOKEN
    GITHUB_USER=your-github-username
    GITHUB_REPO=your-target-repo-name

    # Docker Hub Credentials
    DOCKER_USER=your-dockerhub-username
    DOCKER_PASS=your-dockerhub-password-or-token

    # Ngrok Credentials
    NGROK_AUTH_TOKEN=your-ngrok-auth-token

    # Email (SMTP) Credentials for Notifications
    SMTP_HOST=smtp.gmail.com
    SMTP_PORT=587
    SMTP_USER=your-email@gmail.com
    SMTP_PASS=your-google-app-password

    # Notification Recipient
    EMAIL_TO=recipient-email@example.com
    ```

### 3. Run the Application

Once your `.env` file is configured, you can run the pipeline server:

```sh
go run .
