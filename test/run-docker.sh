# run test by docker on Macos
# build docker image
docker build -t graphql-image:v1.0.0 .
# run service by docker image
if [[ $(docker images -q graphql-image:v1.0.0) != "" ]]; then
	docker run \
		--publish 8080:8080 \
		--env "MYSQL_ADDRS=root@tcp(docker.for.mac.localhost:3306)/kumin_store?charset=utf8&parseTime=True&loc=Local&multiStatements=true" \
		--entrypoint /bin/server graphql-image:v1.0.0
fi
