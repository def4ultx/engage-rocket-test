# Engage rocket - Score calculator

## Design

domain package - Where to store domain model. All are pure function.

handler package - Orchestrate domain logic. Taking care of json encode/decode, validation

integration package - using for running integration test

## How to run application

```bash
go mod vendor
go run main.go
```

or (using make):

```bash
make
```

## How to run test

```bash
go test -mod vendor -race -count=1 ./...
```

or (using make):

```bash
make test
```

## How to run integration test

```bash
docker build -t engagerocket/score-server:latest .
docker-compose up -d
go test -v -tags=integration -count=1 ./integration/...
docker-compose down
```

or (using make):

```bash
make integration
```

## API

### Calculate score

Calculate average score for each category.

#### Endpoint

```
POST /score
```

#### Input

```json
{
    "scores": {
        "managers": [
            { "userId": 1, "score": 1 },
            { "userId": 2, "score": 5 }
        ],
        "team": [
            { "userId": 4, "score": 1 },
            { "userId": 5, "score": 5 },
            { "userId": 6, "score": 3 },
            { "userId": 7, "score": 2 }
        ],
        "others": [
            { "userId": 8, "score": 1 },
            { "userId": 9, "score": 5 }
        ]
    }
}
```

#### Output

```json
{
    "success": true,
    "data": {
        "scores": {
            "manager": 3,
            "team": 2.75,
            "others": 0
        }
    },
    "errors": []
}
```
