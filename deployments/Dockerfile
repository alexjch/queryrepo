## Build stage
FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/queryrepo/

## Packaging stage
FROM centos:8

WORKDIR /app/
RUN yum install dnf-plugins-core -y
RUN rm -rf /etc/yum.repos.d/
COPY --from=builder /app/main repoquerysrv
EXPOSE 9001
RUN groupadd -r appuser && useradd -r -g appuser appuser
CMD ["/app/repoquerysrv", "-repoUrl", "http://vault.centos.org/7.8.2003/os/Source/"]
