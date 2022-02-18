module ihomeMs

go 1.14

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/gin-contrib/sessions v0.0.4
	github.com/gin-gonic/gin v1.7.7
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/nats-io/nats-server/v2 v2.7.2 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	google.golang.org/protobuf v1.26.0
)

replace github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.7.0

replace github.com/nats-io/nats-server/v2 v2.7.2 => github.com/nats-io/nats-server/v2 v2.0.2
