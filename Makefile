
docker:
	docker build -t eyedeekay/echo .

run: docker
	docker run -itd --restart=always --publish 7777:7777 --name echobot eyedeekay/echo