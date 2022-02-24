# python 3 headers, required if submitting to Ansible
from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

DOCUMENTATION = """
    lookup: openssl_certificate
    author: Aubin Lambaré <aubin.lambare@csgroup.eu>
    short_description: create openssl  certificate
    description:
        - This lookup return either a CA certificate
    options:
        _terms:
            description:
                - Common name (CN) for the certificate.
            type: string
            required: true
        notAfter:
            description:
                - Amount of seconds before certificate expirations. 
                - Default to 315360000 (10 years: 10*365*24*60*60).
            type: integer
        san:
            description:
                - Subject alternative names.
                - Expected domain name value. Ex: mysubdomain.mydomain.root   
            type: string
        size:
            description: 
                - Size of the private key. 
                - Default is 2048.
            type: integer
"""

from ansible.errors import AnsibleError
from ansible.plugins.lookup import LookupBase
from ansible.module_utils.common.text.converters import to_native
from ansible.utils.display import Display
from OpenSSL import crypto
import random

display = Display()

DefaultRSAKeySize = 2048

def generate_key(size: int) -> bytes:
    key = crypto.PKey()

    cryptoType = crypto.TYPE_RSA

    display.vvv("Create RSA key of size %i" % size)

    key.generate_key(cryptoType, size)

    return crypto.dump_privatekey(crypto.FILETYPE_PEM, key)

def generate_certificate(privateKey:bytes, cn:str, notAfter:int, san:bytes = None) -> str:
    
    try:
        key = crypto.load_privatekey(crypto.FILETYPE_PEM, privateKey)

        cert = crypto.X509()
        cert.set_serial_number(random.getrandbits(64))
        cert.get_subject().CN = cn
        cert.gmtime_adj_notBefore(0)
        cert.gmtime_adj_notAfter(notAfter)

        extensions = [
            crypto.X509Extension(
                b'basicConstraints', False, b'CA:TRUE'),
            crypto.X509Extension(
                b'keyUsage', False, b'keyCertSign,cRLSign')
        ]

        if san != None:
            extensions.append(crypto.X509Extension(b'subjectAltName', False, san))
            display.vvv("Add SAN %s" % san.decode('utf-8'))

        cert.add_extensions(extensions)

        # self-sign the certificate
        cert.set_issuer(cert.get_subject())
        cert.set_pubkey(key)
        cert.sign(key, 'sha256')

        display.vvv("Certificate signed with sha256.")

    except Exception as e:
        raise AnsibleError("Cannot create certificate :%s" % to_native(e))
    
    return crypto.dump_certificate(crypto.FILETYPE_PEM, cert)

class LookupModule(LookupBase):

    def run(self, terms, variables=None, **kwargs):

        self.set_options(var_options=variables, direct=kwargs)
        
        ret = []
        for term in terms:

            if self.has_option('size') and self.get_option('size') != None:
                size = self.get_option('size')
            else:
                size = DefaultRSAKeySize
            
            privateKey = generate_key(size)
            

            if self.has_option('notAfter') and self.get_option('notAfter') != None:
                    notAfter:int = self.get_option('notAfter')
            else:
                # default set to 10 years
                notAfter = 10*365*24*60*60

            display.vvv("Set expiration date to %i secondes" % notAfter) 

            if self.has_option('san') and self.get_option('san') != None:
                san = b'DNS:' + str.encode(self.get_option('san'))
            else:
                san = None

            certificate = generate_certificate(privateKey, term, notAfter, san)    
                    
            ret.append({
                'key': privateKey.decode('utf-8'),
                'crt': certificate.decode('utf-8')
            })

        return ret
