

## To Install **PHP 7.3** (from `amazon-linux-extras`)

```bash
# 1. Remove all PHP again just in case
sudo yum remove php* -y

# 2. Disable php8.2 extras if enabled
sudo amazon-linux-extras disable php8.2

# 3. Enable php7.3 extras
sudo amazon-linux-extras enable php7.3

# 4. Clean and make cache
sudo yum clean metadata
sudo yum makecache

# 5. Install PHP 7.3 and common extensions
sudo yum install -y php php-cli php-fpm php-common php-json php-mysqlnd php-pdo php-xml php-mbstring php-curl php-opcache

# 6. Confirm version
php -v
```
