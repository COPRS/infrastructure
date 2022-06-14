FROM ubuntu:20.04
ENV TZ=Europe/Paris
RUN groupadd -g 1010 safescale && useradd -u 1010 -g 1010 -s /bin/bash -d /home/safescale -m safescale
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update && apt-get install -y --no-install-recommends \
  python3 \
  python3-pip \
  jq \
  git \
  openssh-client \
  && rm -rf /var/lib/apt/lists/*
WORKDIR /opt/rs-infra-scaler
RUN chown -R safescale:safescale /opt/rs-infra-scaler
USER safescale
ENV PATH="/home/safescale/.local/bin:${PATH}"
RUN ls -ll
RUN git clone --single-branch --branch release/0.8.0 https://github.com/COPRS/infrastructure.git
WORKDIR /opt/rs-infra-scaler/infrastructure
RUN git submodule update --init --recursive
RUN pip3 install --user pyOpenSSL ecdsa -r collections/kubespray/requirements.txt
RUN ansible-galaxy collection install kubernetes.core openstack.cloud
RUN mkdir /opt/rs-infra-scaler/inventory
RUN cp inventory/sample/hosts.yaml ../inventory/hosts.yaml
COPY ./safescale /opt/rs-infra-scaler/safescale
COPY ./rs-infra-scaler /opt/rs-infra-scaler/rs-infra-scaler
RUN ln -s /opt/rs-infra-scaler/safescale /home/safescale/.local/bin/safescale
CMD ["/opt/rs-infra-scaler/rs-infra-scaler"]
