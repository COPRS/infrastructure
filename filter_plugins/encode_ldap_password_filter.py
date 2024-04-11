#!/usr/bin/python3
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


import os
import hashlib
import base64
from ansible.module_utils._text import to_text, to_native
from ansible.errors import AnsibleError, AnsibleFilterError, AnsibleFilterTypeError

class FilterModule(object):
  def filters(self):
    return {
      'encode_ldap_password': self.encode_ldap_password
    }

  def encode_ldap_password(self, password):
    try:
      salt = os.urandom(4)
      sha = hashlib.sha1(password.encode('utf-8') + salt)
      digest = sha.digest()
      b64_digest = base64.b64encode(digest + salt)
      byte_res = b"".join([bytes(i, encoding='utf-8') for i in ("{SSHA}", str(b64_digest, 'utf-8'))])
      result = str(byte_res, 'utf-8')

    except Exception as e:
      raise AnsibleFilterError("encode_ldap_password - {}".format(to_native(e)), orig_exc=e)

    return to_text(result)
