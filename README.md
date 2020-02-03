# dev


wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz

#go mysql
go get -u github.com/go-sql-driver/mysql

#go postgresql
go get -u github.com/lib/pq

#go xorm
go get -u github.com/go-xorm/xorm
go get github.com/go-xorm/cmd/xorm

#go redis
go get -u github.com/garyburd/redigo/redis

#
go get -u github.com/mattn/go-sqlite3

#gt protobuf
go get github.com/golang/protobuf/protoc-gen-go
cd github.com/golang/protobuf/protoc-gen-go
go build
go install

go get github.com/golang/protobuf/proto
cd github.com/golang/protobuf/proto
go build
go install

#ttest