# GoLang AWS Deployment Tool

This GoLang application is designed to provide a streamlined deployment solution for PERN/MERN stack applications on AWS. It facilitates rapid deployment with a single command, abstracting away complexity and reducing deployment time from hours to minutes.

## Features

- **One-Click Deployment**: Users can simply paste their GitHub link and deploy their application to AWS with a single command.
- **Automated Infrastructure Provisioning**: The application automates infrastructure provisioning on AWS using services like EC2, RDS, Route 53, and IAM, ensuring scalability and security.
- **Containerization**: Docker is employed for containerization and orchestration, ensuring consistency and portability across environments.
- **CI/CD Pipelines**: Configured CI/CD pipelines using Jenkins/CircleCI integrated with version control systems for automated testing and deployment, ensuring efficient deployment processes.
- **User-Friendly CLI Tool**: Designed a user-friendly CLI tool for easy interaction, abstracting away complexity for a seamless deployment experience.
- **Error Handling and Rollback Mechanisms**: Incorporated error handling and rollback mechanisms to ensure smooth deployment processes and minimize downtime.

## Installation

To install the GoLang AWS Deployment Tool, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/your-username/aws-deployment-tool.git
cd aws-deployment-tool

go run main.go
