#!/bin/bash
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


# Install necessary ldap packages
sudo DEBIAN_FRONTEND=nointerractive \
apt install -y ldap-auth-client=0.5.4 \
ldap-auth-config=0.5.4 \
libnss-ldap=265-5ubuntu1 \
libpam-ldap=186-4ubuntu1 \
ldap-utils
