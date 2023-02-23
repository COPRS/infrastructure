FROM ubuntu:20.04
ENV TZ=Europe/Paris
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update && apt-get install -y --no-install-recommends \
  software-properties-common \
  curl \
  wget \
  net-tools \
  jq \
  openssh-client \
  && rm -rf /var/lib/apt/lists/*
RUN wget -qO- https://github.com/CS-SI/SafeScale/releases/download/${SAFESCALE_TAG}/safescale-${SAFESCALE_TAG}-linux-amd64.tar.gz | tar -xzf - -C / ./safescaled
CMD ["/safescaled"]
