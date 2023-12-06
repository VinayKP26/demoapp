# demoapp
### Prerequisites:
1. K8s cluster

2. Istio installed

3. Docker install on build env.

4. Container registry access to push build image (or publically avilable image vinaykp26/demoapp:1.0.2)

5. Openssl utility if using self signed certificate

## Installation:

Generate the certificate and key using openssl.
```
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -sha256 -days 365 -nodes
```

Create the tls secret in istio-namespace.
```
kubectl create -n istio-system secret tls istio-ingressgateway-certs --key key.pem --cert cert.pem
```

Create namespace to app deployment
```
kubectl create ns demo
```

Enable the istio envoy proxy auto injection to namespace pods.
```
kubectl label namespace demo istio-injection=enabled
```

Go to root path of source code (cd ./demoapp/)

Update the below required value as per your domain or ip in values.yaml file of ops/helm

```

istio:
  enabled: true
  istioIngressNamespace: istio-system
  gatewayName: demo-gateway
  host: "35.223.164.29.nip.io"
  uri:
    prefix: "/api/v1/"

```

Install the app in demo namespace.
```
helm upgrade --install demo ops/helm -n demo
```
## Application DEMOAPP Routes:
1. GET /api/v1/info

2. POST /api/v1/data

sample data:
```
{
  "_id" : 1,
  "name": "Demo",
  "message" : "Hello Mr. Demo, Welcome"
}

```