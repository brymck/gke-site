gke-site
========

Google Kubernetes Engine site

Prerequisites
-------------

* Docker
* [Helm][install-helm]
* Skaffold (`brew install skaffold`)

Usage
-----

```sh
git clone https://github.com/brymck/gke-site.git
cd gke-site
make
```

### With Istio

Firstly, [install Istio][install-istio], in particular following the steps without mutual TLS authentication between sidecars.
Once that's complete, run:

```sh
make apply-istio
```

Make commands
-------------

* `make` - run `init` then `dev`
* `make init` - run any required one-time commands
* `make dev` - inject proxies if necessary and run via Skaffold
* `make install-postgresql` - install PostgreSQL via Helm
* `make apply-istio` - apply Istio manifests (requires Istio)
* `make delete-istio` - delete Istio manifests
* `make show-grafana` - open the Grafana dashboard in your browser (requires Istio)
* `make show-dashboard` - open the Kubernetes dashboard in your browser
* `make proto` - generate language-specific Protobuf files

Development
-----------

If you make any changes to Protobuf files, ensure downstream libraries are kept in sync with

```sh
make proto
```

[install-helm]: https://docs.helm.sh/using_helm/#installing-helm
[install-istio]: https://istio.io/docs/setup/kubernetes/quick-start/#installation-steps
