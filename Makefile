all: init dev

init: install-postgresql

dev:
	skaffold dev

install-postgresql:
	helm status gke-site-psql >/dev/null || helm install --name gke-site-psql stable/postgresql

apply-istio:
	kubectl label namespace default istio-injection=enabled
	kubectl apply -f ./istio

delete-istio:
	kubectl label namespace default istio-injection-
	kubectl delete -f ./istio

show-grafana:
	$(eval POD_NAME = $(shell kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}'))
	(sleep 1s; open 'http://localhost:3000/') &
	kubectl -n istio-system port-forward $(POD_NAME) 3000:3000

show-dashboard:
	$(eval TOKEN_NAME = $(shell kubectl -n kube-system get secret | grep '^admin-user' | awk '{print $$1}'))
	kubectl -n kube-system describe secret $(TOKEN_NAME) | grep --color=never '^token:'
	(sleep 1s; open 'http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/service?namespace=default') &
	kubectl proxy

proto:
	cd src/frontend/server && ./genproto.sh
	cd src/helloservice && ./genproto.sh

.PHONY: all init dev install-postgresql apply-istio delete-istio show-grafana show-dashboard proto
