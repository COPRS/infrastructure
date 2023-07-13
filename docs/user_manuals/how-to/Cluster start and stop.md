# Start and stop the cluster

## Stop the cluster

1. Drain the Kubernetes cluster.

    ```bash
    nodes=$(kubectl get nodes --no-headers -ocustom-columns=":.metadata.name" --selector='!node-role.kubernetes.io/master')
    # Drain all node but keep linkerd Pod. Tt is necessary as we delete pod so the newly created one get linkerd-proxy injected.
    # --disable-eviction is used to bypass podDisruptionPolicy set by the different operators.
    for node in $nodes; do
        kubectl drain \
            --ignore-daemonsets \
            --delete-emptydir-data \
            --disable-eviction \
            --force \
            --grace-period=0 \
            --pod-selector app.kubernetes.io/instance!=linkerd \
        $node   
    done

    # When all pods are pending except linkerd and daemonsets, delete linkerd pods too.
    for node in $nodes; do
        kubectl drain \
            --ignore-daemonsets \
            --delete-emptydir-data \
            --disable-eviction \
            --force \
            --grace-period=0 \
            --pod-selector app.kubernetes.io/instance=linkerd \
        $node
    done
    ```

2. Turn off the machines.

    ```Bash
    safescale cluster stop CLUSTER_NAME
    ```

## Start the cluster

1. Start the machines.

    ```Bash
    safescale cluster start CLUSTER_NAME
    ```

2. Rollout kube-system daemonsets and deployments

    We first need to restore networking in the cluster. Some pods may be in terminated state.
    This is because the master are turned off before pod eviction is triggered after node is down.
    <!-- 
    /!\ SafeScale does not support a stop delay between masters and nodes.
    Configuration can be made of kube-controller-manager. 
    Cf: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager
    - --pod-eviction-timeout
    - --node-monitor-period
    - --node-monitor-grace-period
    Kubespray configuration detailled in `collections/kubespray/docs/large-deployments.md`.
    -->

    ```Bash
    kubectl --namespace kube-sytem annotate daemonset --all kubernetes.io/change-cause="node restart" --overwrite=true
    kubectl --namespace kube-system rollout restart daemonset
    ```

3. Uncordon the Kubernetes cluster

    Once, all the pods in kube-system namespace are in a ready state, uncordon the nodes.

    ```Bash
    nodes=$(kubectl get nodes --no-headers -ocustom-columns=":.metadata.name" --selector='!node-role.kubernetes.io/master')

    for node in $nodes; do
        kubectl uncordon $node;
    done
    ```

    Linkerd must be ready before starting the next step. Check Linkerd control-plane status with the following command.

    ```Bash
    kubectl --namespace networking get pod --selector=app.kubernetes.io/instance=linkerd --watch
    ```

4. Rollout other daemonsets

    Finally, we can rollout other daemonsets.

    ```Bash
    namespaces=$(kubectl get namespaces --no-headers -ocustom-columns=":.metadata.name" --selector=kubernetes.io/metadata.name!=kube-system)

    for namespace in $namespaces; do
        kubectl -n $namespace annotate daemonset --all kubernetes.io/change-cause="node restart" --overwrite=true
        kubectl -n $namespace rollout restart daemonset
    done
    ```
