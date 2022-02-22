# Certificates management in RS

## TLS configuration

Reference System exploits [APISIX Ingress Controller](https://apisix.apache.org/) and [Cert Manager](https://cert-manager.io/) for TLS configuration.

You need to create an [issuer](https://cert-manager.io/docs/concepts/issuer/) and a [certificate](https://cert-manager.io/docs/concepts/certificate/) for your domain name with Cert Manager.

APISIX does not work with Cert Manager for ACME HTTP01 challenges ([#781](https://github.com/apache/apisix-ingress-controller/issues/781)).  
You must use the DNS01 challenge to generate a Let's encrypt certificate. The configuration is detailled on [Cert Manager documentation](https://cert-manager.io/docs/configuration/acme/dns01).

## Stash community license

A **Stash community** licence is mandatory for the *stash operator* application, get one [here](https://license-issuer.appscode.com/?p=stash-community) and write it in the `main.yaml` inventory file.
