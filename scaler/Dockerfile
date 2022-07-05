FROM ubuntu:20.04
ENV TZ=Europe/Paris
RUN groupadd -g 1010 safescale && useradd -u 1010 -g 1010 -s /bin/bash -d /home/safescale -m safescale
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update && apt-get install -y --no-install-recommends \
  python3 \
  python3-pip \
  jq \
  openssh-client \
  && rm -rf /var/lib/apt/lists/*
USER safescale
WORKDIR /opt/rs-infra-scaler/infrastructure
ENV PATH="/home/safescale/.local/bin:${PATH}"
COPY --chown=safescale:safescale collections ./collections
RUN pip3 install --no-cache-dir --user pyOpenSSL ecdsa -r collections/kubespray/requirements.txt
RUN ansible-galaxy collection install kubernetes.core openstack.cloud
COPY --chown=safescale:safescale *.yaml ansible.cfg ssh-main-gateway.conf ./
COPY --chown=safescale:safescale roles ./roles
COPY --chown=safescale:safescale inventory/sample/hosts.yaml ../inventory/hosts.yaml
COPY --chown=safescale:safescale ./scaler/build/safescale /home/safescale/.local/bin/safescale
COPY --chown=safescale:safescale ./scaler/build/rs-infra-scaler /opt/rs-infra-scaler/rs-infra-scaler
CMD ["/opt/rs-infra-scaler/rs-infra-scaler"]
