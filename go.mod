module github.com/orphaner/03-tartine

go 1.15

require (
	github.com/gin-contrib/gzip v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/pflag v1.0.5
)

replace (
	github.com/gin-contrib/gzip => github.com/gin-contrib/gzip v0.0.3
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.6.3
	github.com/sirupsen/logrus => github.com/sirupsen/logrus v1.7.0
	github.com/spf13/pflag => github.com/spf13/pflag v1.0.5
)
