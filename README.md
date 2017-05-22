# s3cat

Does one thing: write AWS S3 objects to stdout. Explodes on failure. That's all.

## Installation

### Docker image

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo s3cat.go
docker build -t rlister/s3cat .
docker push rlister/s3cat
```

## Usage

You can concatenate multiple objects.

```
go run ./s3cat.go s3://bucket/key ...
```

Running with docker:

```
docker pull rlister/s3cat
docker run rlister/s3cat s3://bucket/key ...
```

## License

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
