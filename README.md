# REST API SERVER - USES GIN-GONIC FRAMEWORK #

# (first step - get acquainted with REST API") #
server with gin-gonic framework (level 1)

## Welcome to GitHub Pages

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
  $ CGO_ENABLE=0 go build --ldflags "-extldflags \"-static\"-s -w" -o bin/application -trimpath ./*.go
```

The program after being built will be saved in `-o bin/application`. Request the system to execute with the command:
```
  $ /bin/bash -c bin/application
```

Documents:
- https://gin-gonic.com/
- https://gin-gonic.com/docs/
- https://godoc.org/github.com/gin-gonic/gin
- https://github.com/gin-gonic/gin
- https://github.com/gin-gonic/gin#api-examples
