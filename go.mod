module github.com/scjtqs2/ddns-go

go 1.17

require (
	github.com/guonaihong/gout v0.2.11
	github.com/kardianos/service v1.2.1
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/scjtqs2/utils v0.0.0-20211110033646-3f01f3014931
	github.com/sirupsen/logrus v1.8.1
	github.com/t-tomalak/logrus-easy-formatter v0.0.0-20190827215021-c074f06c5816
)

require (
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.7.7 // 修复安全漏洞
)