# Copyright 2023 CS Group
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM ubuntu:20.04
ENV TZ=Europe/Paris
ARG SAFESCALE_TAG
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
