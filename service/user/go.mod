module ihomeMs/service/user

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.8
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/nats-io/nats-server/v2 v2.7.2 // indirect
	github.com/tedcy/fdfs_client v0.0.0-20200106031142-21a04994525a
)

replace github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.7.0
replace github.com/nats-io/nats-server/v2 v2.7.2 => github.com/nats-io/nats-server/v2 v2.0.2
