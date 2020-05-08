
docker:
	docker build -f eyedeekay/echo .

run: docker
	docker run --restart=always --publish 7777:7777 --name echobot eyedekay/echo