# Automatic-Doodle-Backend
Automatic-Doodle is a Backend project for People who wants to find daily simple jobs with guaranteed payment
## Installation


```bash
git clone git@github.com:Wodemy-Labs/automatic-doodle.git
```

## Run Project

```bash
docker compose -f ./dev/docker-compose.yaml up -d && go run main.go
```

## Run Project Test

```bash
go run test ./... -v
```
## Important

Please Set  ".env"  file according to [crawl](https://github.com/Wodemy-Labs/crawl).

[Crawl](https://github.com/Wodemy-Labs/crawl) is a logger written in go.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

