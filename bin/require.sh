
# go postgresql
go get -u github.com/lib/pq

#go mysql
go get -u github.com/go-sql-driver/mysql

# go xorm
go get -u github.com/go-xorm/xorm
go get -u github.com/go-xorm/cmd/xorm

# go redis
go get -u github.com/garyburd/redigo/redis

# sqlite3
go get -u github.com/mattn/go-sqlite3

# toml
go get -u github.com/BurntSushi/toml

# gt protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
cd github.com/golang/protobuf/protoc-gen-go
go build
go install

go get -u github.com/golang/protobuf/proto
cd github.com/golang/protobuf/proto
go build
go install
# excelize
go get -u github.com/Luxurioust/excelize

