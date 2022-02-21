#!/usr/bin/python3

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
