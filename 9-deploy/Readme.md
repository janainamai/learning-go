## Executando projeto localmente com docker

### No terminal 1:
docker-compose up -d

docker-compose exec goapp bash

go run main.go

### No terminal 2:
curl http://localhost:8080

## Gerando imagem docker para produção

### No terminal:
docker build -t janainamai/learning-go/9-deploy:latest -f Dockerfile.prod .

docker run --rm -p 8080:8080 janainamai/learning-go/9-deploy:latest