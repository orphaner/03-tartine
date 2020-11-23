module github.com/orphaner/03-tartine

go 1.14

require (
	github.com/orphaner/03-tartine/pkg/core v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/spf13/pflag v1.0.5
)

replace (
	github.com/orphaner/03-tartine/pkg/core => ./pkg/core
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.6.3
	github.com/spf13/pflag => github.com/spf13/pflag v1.0.5
)
