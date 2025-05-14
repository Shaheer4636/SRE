
```ini
listen=YES
listen_ipv6=NO
listen_port=990

ssl_enable=YES
implicit_ssl=YES
rsa_cert_file=/etc/ssl/private/vsftpd.pem
rsa_private_key_file=/etc/ssl/private/vsftpd.pem

force_local_logins_ssl=YES
force_local_data_ssl=YES

pasv_enable=YES
pasv_min_port=40000
pasv_max_port=40010
pasv_address=98.82.15.112
```

---

## Directive Explanations

### `listen=YES`
* Tells `vsftpd` to run in **standalone mode** and listen for incoming connections (not from inetd).

### `listen_ipv6=NO`
* Disables IPv6 mode.
* Required if you're using `listen=YES` and only need IPv4 (which is typical for EC2 setups).

### `listen_port=990`
* Binds the FTP service to **port 990**.
* This is the **default port for Implicit FTPS**, where SSL starts immediately without negotiation.

### `ssl_enable=YES`
* Enables **SSL/TLS support** in `vsftpd`.

### `implicit_ssl=YES`
* Configures `vsftpd` to use **Implicit SSL/TLS**, which starts **TLS immediately** after connection (as required by port 990).
* Without this, it defaults to Explicit TLS and fails for WinSCP in implicit mode.

### `rsa_cert_file=/etc/ssl/private/vsftpd.pem`
### `rsa_private_key_file=/etc/ssl/private/vsftpd.pem`
* Path to your **self-signed SSL certificate and private key**.
* Both must point to the same `.pem` file if you combined them.

### `force_local_logins_ssl=YES`
* Requires **local users** (e.g., `ec2-user`, `ftpuser`) to log in via SSL/TLS only.
* Disables plain login (improves security).

### `force_local_data_ssl=YES`
* Ensures **data transfers** (like uploads/downloads) are encrypted via SSL/TLS.

### `pasv_enable=YES`
* Enables **Passive FTP mode**, which is necessary for most FTP clients and AWS EC2 setups to avoid firewall/NAT issues.

### `pasv_min_port=40000`
### `pasv_max_port=40010`
* Defines the **range of ports** used for passive data connections.
* These must be **opened in your EC2 security group and firewall**.

### `pasv_address=98.82.15.112`
* Sets the **public IP address** the server advertises to clients in Passive mode.
* Critical for AWS EC2 servers with internal/private IPs.
