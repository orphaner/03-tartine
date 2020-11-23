module github.com/tartine

go 1.15

require (
	github.com/gin-contrib/gzip v0.0.3 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/orphaner/03-tartine/pkg/core v0.0.0
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5
)

replace github.com/orphaner/03-tartine/pkg/core => ./pkg/core
