all:
	skaffold dev

forward-grafana:
	$(eval POD_NAME = $(shell kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}'))
	kubectl -n istio-system port-forward $(POD_NAME) 3000:3000

apply-istio:
	kubectl label namespace default istio-injection=enabled
	kubectl apply -f ./istio

delete-istio:
	kubectl label namespace default istio-injection-
	kubectl delete -f ./istio

.PHONY: forward-grafana apply-istio delete-istio
