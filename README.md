# Host-DNS-Updater

## Usage

### Install
```go
    git clone github.com/lilpipidron/Host-DNS-Updater
```

### Build service
```
    cd Host-DNS-Updateer/service/cmd
    go build -o service
```

### Start service 
``` 
    sudo ./service
```

### Build client 
```
    cd ../../client
    go build -o client
```

### Start client
```
    go install
    client [command]
```