NAME=tomaskral/kompose-demo-frontend
VERSION=kubecon2016

.PHONY: image push

image:
	docker build -t $(NAME):$(VERSION) .

push:
	docker push $(NAME):$(VERSION)
