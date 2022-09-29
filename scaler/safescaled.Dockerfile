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
COPY ./scaler/build/safescaled /safescaled
CMD ["/safescaled"]
