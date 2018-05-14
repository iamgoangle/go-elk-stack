# go-elk-stack
:fries: This is simple project to demonstrate **Elasticsearch-Logstash-Kibina (ELK)** integration with **Golang**

# Go Stacks
1. No web framework, using native go
2. **dep** for package manager `brew install dep`
3. **gin** for live-reload `go get github.com/codegangsta/gin`

# Run ELK via docker-compose
```sh
docker-compose up
```

# Ports
```
:9200 = elasticsearch
:5000 = logstash
:5601 = kibana
```

You can open kibana dashboard via [http://localhost:5601](http://localhost:5601)

# Install Dependencies
```sh
dep ensure
```

# Add new dependency
```sh
go get -u -v <dependency> or
dep ensure --add <dependency>
```

# Run App
```sh
go run main.go
```

# Run app with live-reload
```sh
gin main.go
```

# Todo
- create Makefile