# Description
- Protobuf
- Rest api (fallback)
- Swagger

## Makefile
- make

## Swagger
docker run --rm \
--name swagger_server \
-d \
-p 9000:8080 \
-e API_URL=swagger.json \
swaggerapi/swagger-ui && \
\
FIND_SWAGGER_FILES="$(find . -type f -name '*.swagger.json')" && \
for f in $FIND_SWAGGER_FILES; do docker cp $f swagger_server:/usr/share/nginx/html; done
