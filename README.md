gke-site
========

Google Kubernetes Engine site

Prerequisites
-------------

* Docker
* Skaffold (`brew install skaffold`)

Usage
-----

```sh
skaffold dev
```

With Istio
----------

Firstly, [install Istio][install-istio], in particular following the steps without mutual TLS authentication between sidecars.
Once that's complete, run:

```sh
kubectl label namespace default istio-injection=enabled
kubectl apply -f ./istio
```

[install-istio]: https://istio.io/docs/setup/kubernetes/quick-start/#installation-steps
