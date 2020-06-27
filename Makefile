
SRV_IMG := alextanhongpin/go-fluent
FLU_IMG := alextanhongpin/fluentd-mongo

build-server:
	@docker build -f Dockerfile.server -t ${SRV_IMG} .

build-fluentd:
	@docker build -f Dockerfile.fluentd -t ${FLU_IMG} .

up:
	@docker-compose up -d

down:
	@docker-compose down
