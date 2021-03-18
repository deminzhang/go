# protoc
#wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-win64.zip

# go postgresql
go get -u github.com/lib/pq

# go mysql
go get -u github.com/go-sql-driver/mysql

# go xorm
go get -u github.com/go-xorm/xorm
go get -u github.com/go-xorm/cmd/xorm

# sqlite3
go get -u github.com/mattn/go-sqlite3

# go redis
go get -u github.com/garyburd/redigo/redis

# gt protobuf
# protoc @ https://github.com/protocolbuffers/protobuf/releases
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go build github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go

# toml
go get -u github.com/BurntSushi/toml

# excelize
go get -u github.com/Luxurioust/excelize

