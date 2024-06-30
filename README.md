# host-dns-updater

## Usage

### Install
```
    git clone https://github.com/lilpipidron/Host-DNS-Updater
```

### Build service
```
    cd host-dns-updater/service/cmd
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
