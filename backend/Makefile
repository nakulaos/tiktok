download:
	GOPROXY=https://goproxy.cn go mod tidy

build:
	mkdir -p output
	#go env -w  GOPROXY=https://goproxy.cn,direct
	go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
	go mod tidy
	#user
	cd user/rpc && CGO_ENABLED=0 go build  -ldflags="-s -w" -o "../../output/user_rpc" user.go
	cd user/api && CGO_ENABLED=0  go build  -ldflags="-s -w" -o "../../output/user_api" user.go

	cd relation/rpc && CGO_ENABLED=0 go build  -ldflags="-s -w" -o "../../output/relation_rpc" relation.go
	cd relation/api && CGO_ENABLED=0  go build  -ldflags="-s -w" -o "../../output/relation_api" relation.go

	cd favorite/rpc && CGO_ENABLED=0 go build  -ldflags="-s -w" -o "../../output/favorite_rpc" favorite.go
	cd favorite/api && CGO_ENABLED=0  go build  -ldflags="-s -w" -o "../../output/favorite_api" favorite.go

	cd feed/rpc && CGO_ENABLED=0 go build  -ldflags="-s -w" -o "../../output/feed_rpc" feed.go
	cd feed/api && CGO_ENABLED=0  go build  -ldflags="-s -w" -o "../../output/feed_api" feed.go

	cd mq && CGO_ENABLED=0  go build  -ldflags="-s -w" -o "../output/mq_api" mq.go

	cd comment/rpc && CGO_ENABLED=0 go build  -ldflags="-s -w" -o "../../output/comment_rpc" comment.go
	cd comment/api && CGO_ENABLED=0  go build  -ldflags="-s -w" -o "../../output/comment_api" comment.go


	cd live/rpc && CGO_ENABLED=0 go build  -ldflags="-s -w" -o "../../output/live_rpc" live.go
	cd live/api && CGO_ENABLED=0  go build  -ldflags="-s -w" -o "../../output/live_api" live.go

generate:
	cd  user/api && goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir ../../script/swagger
	cd  relation/api && goctl api plugin -plugin goctl-swagger="swagger -filename relation.json" -api relation.api -dir ../../script/swagger
	cd  favorite/api && goctl api plugin -plugin goctl-swagger="swagger -filename favorite.json" -api favorite.api -dir ../../script/swagger
	cd  feed/api && goctl api plugin -plugin goctl-swagger="swagger -filename feed.json" -api feed.api -dir ../../script/swagger
	cd  comment/api && goctl api plugin -plugin goctl-swagger="swagger -filename comment.json" -api comment.api -dir ../../script/swagger
	cd  live/api && goctl api plugin -plugin goctl-swagger="swagger -filename live.json" -api live.api -dir ../../script/swagger
