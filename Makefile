all: image dockerpush

.PHONY: image
image:
	docker build . -t simple-bash:latest

dockerpush:
	docker tag simple-bash unterhol/simple-bash:latest
	docker push unterhol/simple-bash:latest