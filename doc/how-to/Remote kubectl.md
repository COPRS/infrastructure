# Remote cluster administration using kubectl

## Install kubectl and the oidc plugin

 - Install kubectl using the official documentation: https://kubernetes.io/fr/docs/tasks/tools/install-kubectl/
 - Install the kubelogin using official documentation: https://github.com/int128/kubelogin or by running:
 ```shellsession
 curl -LO https://github.com/int128/kubelogin/releases/download/v1.25.1/kubelogin_linux_amd64.zip
 unzip kubelogin_linux_amd64.zip
 sudo mv kubelogin /usr/local/bin/kubectl-oidc_login
 ```

## Configure your kubeconfig file to use the platform's authentification

Setup your `kubeconfig` file like the following example, and set manually the the variables that come from the inventory:
```yaml
apiVersion: v1
clusters:
- cluster:
    server: https://kube.{{ platform_domain_name }}
  name: {{ cluster.name }}
contexts:
- context:
    cluster: {{ cluster.name }}
    user: oidc
  name: default
current-context: default
kind: Config
preferences: {}
users:
- name: oidc
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      args:
        - oidc-login
        - get-token
        - --oidc-issuer-url=https://iam.{{ platform_domain_name }}/auth/realms/{{ keycloak.realm.name }}
        - --oidc-client-id=kubectl
        - --oidc-client-secret={{ kubectl_oidc.oidc_client_secret }}
      command: kubectl
      env: null
      provideClusterInfo: false

```

## Run your kubectl commands as usual

On first login or on token expiration, your browser will open a login page where you can use your platform credentials.

> Note: Add `--grant-type=authcode-keyboard` to the args if you want to copy-paste a link in your browser manually, it is useful if you are in a ssh session
