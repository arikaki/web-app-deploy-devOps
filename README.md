# Go Web Application

This is a simple website written in Golang. It uses the `net/http` package to serve HTTP requests.

## Running the server

To run the server, execute the following command:

```bash
go run main.go
```

The server will start on port 8080. You can access it by navigating to `http://localhost:8080` in web browser.

# DevOps Implementation for Go Web Application

This project focuses on integrating DevOps practices into a Go-based web application. The application, which is a simple website built using the `net/http` package, serves HTTP requests.

Key DevOps practices covered in this project include:

- **Creating a Dockerfile (Multi-stage build)**
- **Containerization**
- **Continuous Integration (CI)**
- **Continuous Deployment (CD)**

## Creating a Dockerfile (Multi-stage build)

The Dockerfile is essential for building a Docker image that encapsulates the Go web application along with its dependencies. The resulting Docker image is then used to instantiate a Docker container.

We utilize a Multi-stage build approach to create the Docker image. This method leverages multiple build stages within a single Dockerfile, which helps minimize the final image size and enhance security by excluding unnecessary files and packages.

## Containerization

Containerization involves packaging an application and its dependencies into a container, which can then run on a container platform like Docker. This technique ensures the application operates consistently across various environments.

For this project, Docker is used to containerize the Go web application. Docker facilitates the building, shipping, and running of containers.

Commands to build and manage the Docker container:

```bash
docker build -t <your-docker-username>/go-web-app .
```

To run the Docker container:

```bash
docker run -p 8080:8080 <your-docker-username>/go-web-app
```

To push the Docker container to Docker Hub:
```bash
docker push <your-docker-username>/go-web-app
```

## Continuous Integration (CI)

Continuous Integration (CI) is the practice of automating the integration of code changes into a shared repository, helping to identify bugs early and ensuring the code is always ready for deployment.

GitHub Actions will be used to set up CI for the Go web application. GitHub Actions allows for the automation of workflows, including building, testing, and deploying code.

The CI workflow in GitHub Actions will include:

- Checking out the code from the repository
- Building the Docker image
- Running the Docker container
- Executing tests

## Continuous Deployment (CD)

Continuous Deployment (CD) involves automatically deploying code changes to a production environment, which helps in quickly delivering new features and fixes to users.

Argo CD will be used to implement CD for the Go web application. Argo CD is a GitOps continuous delivery tool for Kubernetes, enabling the deployment of applications to Kubernetes clusters by using Git as the source of truth.

The Argo CD setup will automatically deploy the Go web application to a Kubernetes cluster and ensure that the deployment remains synchronized with the Git repository, keeping the application up-to-date.