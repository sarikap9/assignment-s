build:
	docker build -t ping-pong-go .

run:
	docker run -p 3000:8081 ping-pong-go

k8s-deploy:
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml

k8s-clean:
	kubectl delete -f deployment.yaml
	kubectl delete -f service.yaml
