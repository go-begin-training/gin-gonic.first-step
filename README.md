# server gin-gonic (first step - get acquainted with REST API")
server with gin-gonic framework (level 1)

### How can this project run? ###

We use the files "go.mod" and "go.sum" to contain the necessary configuration packages.
If the project you pull requires importing packages, open a terminal at the project root directory and enter the following command:
```
  $ go mod download
```

Next, run directly with the command:
```
  $ go run * .go
```
Or build into a program with the command:
```
  $ CGO_ENABLE = 0 go build --ldflags "-extldflags \" - static \ "-s -w" -o bin/application -trimpath ./*.go
```

The program after being built will be saved in `-o bin/application`. Request the system to execute with the command:
```
  $ /bin/bash -c bin/application
```

Documents:
  - https://www.mongodb.com/
  - https://docs.mongodb.com/drivers/go/
  - https://godoc.org/go.mongodb.org/mongo-driver/mongo
  - https://github.com/mongodb/mongo-go-driver
  - https://github.com/mongodb/mongo-go-driver#usage