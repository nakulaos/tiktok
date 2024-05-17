build-k8s-gitlabAndHarbor:
	kubectl apply -f ./deploy/k8s-gitlab.yaml
	kubectl apply -f ./deploy/k8s-harbor-pvc.yaml
	cd ./script/k8s  && helm install harbor -n harbor harbor/

build-k8s-gitlab-runner:
	# 需要查询token
	kubectl apply -f ./deploy/k8s-gitlab-runner.yaml


initial-k8s:
	minikube start --memory 8192  --cpus 6
	minikube dashboard


delete-k8s-harbor:
	kubectl delete namespace harbor


stop-k8s:
	minikube stop

run-backend-docker:
	# 先删除之前创建的镜像
	docker image rm tiktok-tiktok-userapi
	docker image rm tiktok-tiktok-userrpc
	docker image rm tiktok-tiktok-relationapi
	docker image rm tiktok-tiktok-relationrpc

	docker-compose  -f docker-compose-env.yml up -d

build-backend-docker:
	docker-compose -f docker-compose-build.yml up -d


run:
	cd backend &&  make build
	# 先删除之前创建的镜像
#
	docker-compose  -f docker-compose-env.yml down

	docker image rm tiktok-tiktok-userapi
	docker image rm tiktok-tiktok-userrpc
	docker image rm tiktok-tiktok-relationapi
	docker image rm tiktok-tiktok-relationrpc
	docker image rm tiktok-tiktok-favoriteapi
	docker image rm tiktok-tiktok-favoriterpc
	docker image rm tiktok-tiktok-feedapi
	docker image rm tiktok-tiktok-feedrpc
	docker image rm tiktok-tiktok-mqapi
	docker image rm tiktok-tiktok-commentapi
	docker image rm tiktok-tiktok-commentrpc
	docker image rm tiktok-tiktok-liveapi
	docker image rm tiktok-tiktok-liverpc

	docker-compose  -f docker-compose-env.yml up -d

