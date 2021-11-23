#!/bin/bash

# Apply network-related security rules

cat <<-'EOF' | sudo tee -a /etc/sysctl.conf > /dev/null
# IP Spoofing protection
net.ipv4.conf.default.rp_filter = 1
net.ipv4.conf.all.rp_filter = 1
# Block SYN attacks
net.ipv4.tcp_syncookies = 1
# Controls IP packet forwarding
net.ipv4.conf.all.forwarding = 0
# Ignore ICMP redirects
net.ipv4.conf.all.accept_redirects = 0
net.ipv6.conf.all.accept_redirects = 0
net.ipv4.conf.default.accept_redirects = 0
net.ipv6.conf.default.accept_redirects = 0
# Ignore send redirects
net.ipv4.conf.all.send_redirects = 0
net.ipv4.conf.default.send_redirects = 0
# Disable source packet routing
net.ipv4.conf.all.accept_source_route = 0
net.ipv6.conf.all.accept_source_route = 0
net.ipv4.conf.default.accept_source_route = 0
net.ipv6.conf.default.accept_source_route = 0
# Log Martians
net.ipv4.conf.all.log_martians = 1
# Block SYN attacks
net.ipv4.tcp_max_syn_backlog = 2048
net.ipv4.tcp_synack_retries = 2
net.ipv4.tcp_syn_retries = 5
# Log Martians
net.ipv4.icmp_ignore_bogus_error_responses = 1
# Ignore ICMP broadcast requests
net.ipv4.icmp_echo_ignore_broadcasts = 1
# Ignore Directed pings
net.ipv4.icmp_echo_ignore_all = 1
kernel.randomize_va_space = 1
# disable IPv6 if required (IPv6 might cause issues with the Internet connection being slow)
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6 = 1
net.ipv6.conf.lo.disable_ipv6 = 1
# Accept Redirects? No, this is not router
net.ipv4.conf.all.secure_redirects = 0
# Log packets with impossible addresses to kernel log? yes
net.ipv4.conf.default.secure_redirects = 0

# [IPv6] Number of Router Solicitations to send until assuming no routers are present.
# This is host and not router.
net.ipv6.conf.default.router_solicitations = 0
# Accept Router Preference in RA?
net.ipv6.conf.default.accept_ra_rtr_pref = 0
# Learn prefix information in router advertisement.
net.ipv6.conf.default.accept_ra_pinfo = 0
# Setting controls whether the system will accept Hop Limit settings from a router advertisement.
net.ipv6.conf.default.accept_ra_defrtr = 0
# Router advertisements can cause the system to assign a global unicast address to an interface.
net.ipv6.conf.default.autoconf = 0
# How many neighbor solicitations to send out per address?
net.ipv6.conf.default.dad_transmits = 0
# How many global unicast IPv6 addresses can be assigned to each interface?
net.ipv6.conf.default.max_addresses = 1
EOF

sudo sysctl -p

# Secure /tmp and /var/tmp
## Creating a 1GB filesystem file for the /tmp parition.
sudo fallocate -l 1G /tmpdisk
sudo fallocate -l 1G /vartmpdisk

sudo mkfs.ext4 /tmpdisk
sudo mkfs.ext4 /vartmpdisk

sudo chmod 0600 /tmpdisk
sudo chmod 0600 /vartmpdisk

## Mounting the new /tmp partition and setting the right permissions.
sudo mount -o loop,noexec,nosuid,rw /tmpdisk /tmp
sudo mv /var/tmp /var/tmpold
sudo mkdir /var/tmp
sudo mount -o loop,noexec,nosuid,rw /vartmpdisk /var/tmp
sudo chmod 1777 /tmp
sudo chmod 1777 /var/tmp

## Setting the /tmp in the fstab.
cat <<-'EOF' | sudo tee -a /etc/fstab >/dev/null
/tmpdisk	/tmp	ext4	loop,nosuid,noexec,rw	0 0
/vartmpdisk  /var/tmp    ext4    loop,nosuid,noexec,rw   0 0
EOF
sudo mount -o remount /tmp
sudo mount -o remount /var/tmp

sudo cp -prf /var/tmpold/* /var/tmp/
sudo rm -rf /var/tmpold

# Secure Shared Memory
cat <<-'EOF' | sudo tee -a /etc/fstab >/dev/null
tmpfs	/run/shm	tmpfs	ro,noexec,nosuid	0 0
EOF

# Limit number of processes for the users
cat <<-'EOF' | sudo tee -a /etc/security.limits.conf > /dev/null
@users hard nproc 100
EOF

# SSH hardening
cat <<-'EOF' | sudo tee /etc/ssh/sshd_config > /dev/null
Port 22
Protocol 2
HostKey /etc/ssh/ssh_host_rsa_key
HostKey /etc/ssh/ssh_host_dsa_key
HostKey /etc/ssh/ssh_host_ecdsa_key
HostKey /etc/ssh/ssh_host_ed25519_key
UsePrivilegeSeparation yes
KeyRegenerationInterval 3600
ServerKeyBits 1024
SyslogFacility AUTH
LogLevel VERBOSE
LoginGraceTime 120
PermitRootLogin no
StrictModes yes
RSAAuthentication yes
PubkeyAuthentication yes
IgnoreRhosts yes
RhostsAuthentication no
RhostsRSAAuthentication no
HostbasedAuthentication no
PermitEmptyPasswords no
ChallengeResponseAuthentication no
PasswordAuthentication no
ClientAliveInterval 300
ClientAliveCountMax 0
AllowTcpForwarding no
X11Forwarding no
X11DisplayOffset 10
UseDNS no
AllowAgentForwarding no
AllowTcpForwarding no
PrintMotd no
AcceptEnv LANG LC_*
PrintLastLog yes
TCPKeepAlive no
MaxSessions 2
MaxAuthTries 3
Compression no
UsePAM yes
Subsystem sftp /usr/lib/openssh/sftp-server
EOF

sudo sed -i -e "/^.*.pam_motd.so.*$/s/^/#/" /etc/pam.d/sshd

sudo systemctl restart sshd

# Install fail2ban

sudo apt-get install -y  fail2ban
cat <<-'EOF' | tee -a /etc/fail2ban/jail.local > /dev/null
[sshd]
enabled = true
port = 22
filter = sshd
logpath = /var/log/auth.log
maxretry = 3
EOF

sudo systemctl enable fail2ban --now