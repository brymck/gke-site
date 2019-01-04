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

### With Istio

Firstly, [install Istio][install-istio], in particular following the steps without mutual TLS authentication between sidecars.
Once that's complete, run:

```sh
make apply-istio
```

To view the Grafana dashboard, run

```sh
make forward-grafana
```

and visit [http://localhost:3000/](http://localhost:3000/).

To remove usage of Istio, run

```sh
make delete-istio
```

Development
-----------

If you make any changes to Protobuf files, ensure downstream libraries are kept in sync with

```sh
make proto
```

[install-istio]: https://istio.io/docs/setup/kubernetes/quick-start/#installation-steps
