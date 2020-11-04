# KIS2020 Demo Application

![Unit Test](https://github.com/guni1192/kis2020-demo-application/workflows/Unit%20Test/badge.svg)
![Lint](https://github.com/guni1192/kis2020-demo-application/workflows/Lint/badge.svg)
![Go](https://img.shields.io/github/go-mod/go-version/guni1192/kis2020-demo-application)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/guni1192/kis2020-demo-application)](https://pkg.go.dev/github.com/guni1192/kis2020-demo-application)


Sample Demo Application


## Getting Started

### Rename Paramerters

1. Push `use this template` button in this repository.
2. Rename variables for your application.
Here's the startup script.

```bash
(
    export APPLICATION=your-application-name
    export USER=your-user-name
    export REPOSITORY=$USER/$APPLICATION
    grep -l 'hello' Dockerfile .github/workflows/*.yaml | xargs sed -i.bak -e "s/hello/$APPLICATION/g"
    grep -l 'guni1192/kis2020-demo-application' * | xargs sed -i.bak -e "s@guni1192/kis2020-demo-application@$REPOSITORY@g"
)
```

### Setup Container Registry

1. Get your token for Github Packages.
`User > Settings > Developper Settins > Personal access tokens > Generate new token`
**Check `write:packages` and `delete:packages` after click Generate token.**
2. Copy your token and paste the token set your repository secret variables (name: CR_PAT).
`Access Your Repository > Settings > Secrets > New Secrets`


### Docker Build in Local

```bash
docker build -t ghcr.io/guni1192/kis2020-demo-application .
```

### Running on Docker

```bash
docker run --rm -p 8080:8080 ghcr.io/guni1192/kis2020-demo-application
```

### Local Build

```
make build
```

## CI

- Lint (`go vet`, `golangci-lint`)
- Unit Tests
- Publish library documents
- Build and push container image to `ghcr.io`
- Vulnerability scanning in container image

### Lint

Analysis your code staticly when each push.

tools
- `go vet`
- `golangci-lint`

### Publish library documents

Your application documents is generated by GitHub actions when your application released (ref: http://pkg.go.dev).

### Build and push container image to ghcr.io

Build and push container image when you tagged a commit.
Container image tag format is `ghcr.io/{{ owner }}/{{ application }}:{{ version }}`.

### Vulnerability scanning in container image

Scan Container image when each commits.
If vulnerabilities are found, send alert message to your PR.

## License

Put your application license.
