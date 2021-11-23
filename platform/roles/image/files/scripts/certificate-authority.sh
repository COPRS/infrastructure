#!/bin/bash

# Parameters

SF_ETCDIR="/opt/safescale/etc"
CountryName="FR"
StateOrProvinceName="Haute-Garonne"
LocalityName="Toulouse"
OrganizationName="CS-Group"
OrganizationalUnitName="S1PDGS"
CommonName="s1pdgs-CA.crt"

# Setup

sudo apt-get install -y openssl
sudo mkdir -p ${SF_ETCDIR}/pki/ca/{certs,crl,csr,newcerts,private}
sudo chmod 0700 ${SF_ETCDIR}/pki/ca/private
sudo touch ${SF_ETCDIR}/pki/ca/index.txt
sudo touch ${SF_ETCDIR}/pki/ca/index.txt.attr

# Generate Certificate Authority

sudo echo "1000" > ${SF_ETCDIR}/pki/ca/serial | &>/dev/null
sudo echo "1000" > ${SF_ETCDIR}/pki/ca/crlnumber | &>/dev/null

cat <<-EOF | sudo tee ${SF_ETCDIR}/pki/ca/openssl.cnf &>/dev/null
[ ca ]
default_ca = CA_default

[ CA_default ]
# Directory and file locations.
dir               = ${SF_ETCDIR}/pki/ca
certs             = \$dir/certs
crl_dir           = \$dir/crl
new_certs_dir     = \$dir/newcerts
database          = \$dir/index.txt
serial            = \$dir/serial
RANDFILE          = \$dir/private/.rand

# The root key and root certificate.
private_key       = \$dir/private/rootca.key.pem
certificate       = \$dir/certs/rootca.cert.pem

# For certificate revocation lists.
crlnumber         = \$dir/crlnumber
crl               = \$dir/crl/rootca.crl.pem
crl_extensions    = crl_ext
default_crl_days  = 30

# SHA-1 is deprecated, so use SHA-2 instead.
default_md        = sha256

name_opt          = ca_default
cert_opt          = ca_default
default_days      = 375
preserve          = no
policy            = policy_default

[ policy_default ]
# See the POLICY FORMAT section of the 'ca' man page.
countryName             = optional
stateOrProvinceName     = optional
localityName            = optional
organizationName        = optional
organizationalUnitName  = optional
commonName              = supplied
emailAddress            = optional

[ req ]
# Options for the 'req' tool (man req).
default_bits        = 4096
distinguished_name  = req_distinguished_name
string_mask         = utf8only
# SHA-1 is deprecated, so use SHA-2 instead.
default_md          = sha256

# Extension to add when the -x509 option is used.
x509_extensions     = v3_ca

[ req_distinguished_name ]
# See <https://en.wikipedia.org/wiki/Certificate_signing_request>.
countryName                     = Country Name (2 letter code)
stateOrProvinceName             = State or Province Name
localityName                    = Locality Name
0.organizationName              = Organization Name
organizationalUnitName          = Organizational Unit Name
commonName                      = Common Name
emailAddress                    = Email Address

# Optionally, specify some defaults.
countryName_default             = ${CountryName}
stateOrProvinceName_default     = ${StateOrProvinceName}
localityName_default            = ${LocalityName}
0.organizationName_default      = ${OrganizationName}
organizationalUnitName_default  = ${OrganizationalUnitName}
emailAddress_default            = 

[ v3_ca ]
# Extensions for a typical CA (man x509v3_config).
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid:always,issuer
basicConstraints = critical, CA:true
keyUsage = critical, digitalSignature, cRLSign, keyCertSign

[ usr_cert ]
# Extensions for client certificates (man x509v3_config).
basicConstraints = CA:FALSE
nsCertType = client, email
nsComment = "OpenSSL Generated Client Certificate"
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer
keyUsage = critical, nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, emailProtection

[ server_cert ]
# Extensions for server certificates (man x509v3_config).
basicConstraints = CA:FALSE
nsCertType = server
nsComment = "OpenSSL Generated Server Certificate"
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer:always
keyUsage = critical, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth

[ crl_ext ]
# Extension for CRLs (man x509v3_config).
authorityKeyIdentifier=keyid:always

[ ocsp ]
# Extension for OCSP signing certificates (man ocsp).
basicConstraints = CA:FALSE
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer
keyUsage = critical, digitalSignature
extendedKeyUsage = critical, OCSPSigning
EOF


sudo openssl genrsa -out ${SF_ETCDIR}/pki/ca/private/rootca.key.pem 4096 
sudo chmod 0400 ${SF_ETCDIR}/pki/ca/private/rootca.key.pem

SUBJ="/C=${CountryName}/ST=${StateOrProvinceName}/L=${LocalityName}/O=${OrganizationName}/OU=${OrganizationalUnitName}/CN=${CommonName}"
sudo openssl req -config ${SF_ETCDIR}/pki/ca/openssl.cnf \
            -key ${SF_ETCDIR}/pki/ca/private/rootca.key.pem \
            -new -x509 -days 10000 -sha256 \
            -out ${SF_ETCDIR}/pki/ca/certs/rootca.cert.pem \
            -extensions v3_ca \
            -subj "${SUBJ}"
sudo chmod 0444 ${SF_ETCDIR}/pki/ca/certs/rootca.cert.pem


# Enable Certificate Authority 
sudo ln -s ${SF_ETCDIR}/pki/ca/certs/rootca.cert.pem "/usr/local/share/ca-certificates/${CommonName}"

#sudo /bin/update-ca-trust${SF_ETCDIR}/pki/ca/crlnumber &>/dev/null
sudo update-ca-certificates

# Prevent safescale to create a CA (bild in check is not working)

sudo mkdir -p ${SF_ETCDIR}/pki/ca/signca/certs
sudo touch ${SF_ETCDIR}/pki/ca/signca/certs/rootca.cert.pem