NAME=expert_system

SRC_PATH=src/*.go

all: $(NAME)

$(NAME): $(SRC_PATH)
	@go build -o $(NAME) $(SRC_PATH)
	@echo "Exec $(NAME) OK"

fclean:
	@rm -rf $(NAME)
	@echo "Exec $(NAME) removed"

clean:
	@rm -rf $(NAME)
	@echo "Exec $(NAME) removed"

clean_deploy: clean_jaeger clean_dgraph

clean_jaeger:
	@docker stop jaeger
	@docker rm -f jaeger

clean_dgraph:
	@docker stop dgraph
	@docker rm -f dgraph
	@docker volume rm vgraph

re: clean all

goget:
	@go get github.com/opentracing/opentracing-go
	@go get github.com/uber/jaeger-client-go
	@go get github.com/dgraph-io/dgo
	
deploy: jaeger dgraph

jaeger:
	@docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.6 >> /dev/null
	@echo "Jaeger has been deployed in docker"

dgraph:
	@docker volume create vgraph >> /dev/null
	@docker run -d -t --privileged -p 5080:5080 -p 6080:6080 -p 8080:8080 -p 9080:9080 -p 8000:8000 -v vgraph:/dgraph --name dgraph dgraph/dgraph:v1.0.10 dgraph zero &> /dev/null
	@docker exec -d dgraph dgraph alpha --lru_mb 2048 --zero 127.0.0.1:5080 &> /dev/null
	@docker exec -d dgraph dgraph-ratel &> /dev/null
	@echo "Dgraph has been deployed in docker"

clean_db:
	@curl -X POST 127.0.0.1:8080/alter -d '{"drop_all": true}'

.PHONY: all fclean clean re deploy jaeger dgraph clean_deploy clean_jaeger clean_dgraph goget clean_db
