# api


### Getting started

``` go
go mod download
go run main.go start
```

### deploy

``` bash
make deploy
```

### swagger

- create swagger yml

``` bash
swagger generate spec -o ./swagger.yml --scan-models
aws s3 cp ./swagger.yml s3://squaaat-lambda/serverless/jeonong-api/alpha/swagger.yml --acl public-read
```

- validate swagger yml

``` bash
swagger validate ./swagger.yml
```

### gorm

- DB 초기화
``` bash
go run main.go gorm create
```

- DB 삭제
``` bash
go run main.go gorm clean
```

- DB 초기화 & 삭제
``` bash
go run main.go gorm re-create
```

- Migration 코드 생성
``` bash
go run main.go gorm migrate create -v (default: yyyymmddHHMM)
```

- Migration 코드 실행
``` bash
go run main.go gorm migrate sync
``` 

