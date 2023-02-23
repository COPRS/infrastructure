FROM golang:1.18 as scaler-bin
# Build binarie GO
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY protos/ protos/
RUN go build -o ./build/rs-infra-scaler .

FROM alpine as external-bin
ARG SAFESCALE_TAG
# Download KubeSpray
RUN apk add wget unzip
RUN wget -q https://github.com/kubernetes-sigs/kubespray/archive/refs/heads/release-2.17.zip -O kubespray.zip && unzip kubespray.zip >/dev/null && rm kubespray.zip
# Download Safescale
RUN wget -q -O- https://github.com/CS-SI/SafeScale/releases/download/${SAFESCALE_TAG}/safescale-${SAFESCALE_TAG}-linux-amd64.tar.gz | tar -xzf - ./safescale

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
WORKDIR /opt/rs-infra-scaler/infrastructure
RUN chown -R safescale:safescale /opt/rs-infra-scaler
USER safescale
ENV PATH="/home/safescale/.local/bin:${PATH}"
COPY --from=external-bin --chown=safescale:safescale /safescale /home/safescale/.local/bin/safescale
COPY --from=external-bin --chown=safescale:safescale /kubespray-release-2.17 ./collections/kubespray/
RUN pip3 install --no-cache-dir --user pyOpenSSL ecdsa -r collections/kubespray/requirements.txt
RUN ansible-galaxy collection install kubernetes.core openstack.cloud
COPY --chown=safescale:safescale ansible_resources/*.yaml ansible_resources/ansible.cfg ./
COPY --chown=safescale:safescale ansible_resources/roles ./roles
COPY --chown=safescale:safescale ansible_resources/inventory/sample/hosts.yaml ../inventory/hosts.yaml
COPY --from=scaler-bin --chown=safescale:safescale /app/build/rs-infra-scaler /opt/rs-infra-scaler/rs-infra-scaler
CMD ["/opt/rs-infra-scaler/rs-infra-scaler"]
