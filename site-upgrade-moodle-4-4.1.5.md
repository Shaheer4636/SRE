
````markdown
#  Moodle Upgrade Guide: 4.0dev+ → 4.1.5

**Step-by-step technical document** outlining how you successfully upgraded your Moodle site to **version 4.1.5** from **4.0dev+**, including PHP version alignment, plugin cleanup, and environment fixes.

---

##  Objective

Upgrade Moodle core from 4.0dev+ to Moodle 4.1.5 in a safe, testable, and recoverable manner.

---

##  1. Pre-Upgrade Backups

###  Backup Moodle Code

```bash
TIMESTAMP=$(date +%F_%H-%M)
mkdir -p /var/backups/moodle-backup-$TIMESTAMP
cp -a /var/www/html/moodle /var/backups/moodle-backup-$TIMESTAMP/moodle-code
`````

###  Backup Moodle Database (RDS)

```bash
mysqldump -u moodleadmin -p'MoodleDev!2025' \
  -h rds-techops-moodle01.ck3dgulnkqbo.us-west-2.rds.amazonaws.com \
  --ssl-ca=/tmp/rds-ca.pem \
  moodle > /var/backups/moodle-backup-$TIMESTAMP/moodle-db.sql
```

---

##  2. Download and Extract Moodle 4.1.5

```bash
cd /tmp
wget -O moodle-4.1.5.tgz https://download.moodle.org/download.php/direct/stable401/moodle-latest-401.tgz
tar -xvzf moodle-4.1.5.tgz
```

---

## 3. Replace Moodle Codebase

```bash
cp /var/www/html/moodle/config.php /tmp/config.php.bak
rm -rf /var/www/html/moodle/*
cp -a /tmp/moodle/* /var/www/html/moodle/
cp /tmp/config.php.bak /var/www/html/moodle/config.php
chown -R apache:apache /var/www/html/moodle
chmod -R 755 /var/www/html/moodle
```

---

##  4. Ensure Correct PHP Version (PHP 8.1)

###  Remove Old PHP and Enable PHP 8.1

```bash
sudo yum remove php* -y
sudo amazon-linux-extras enable php8.1
sudo yum clean metadata
sudo yum install -y php php-cli php-fpm php-mysqlnd php-gd php-intl php-sodium php-soap php-xml
```

###  Set PHP 8.1 as Default

```bash
sudo alternatives --install /usr/bin/php php /usr/bin/php-8.1 1
sudo alternatives --set php /usr/bin/php-8.1
```

###  Restart Apache

```bash
sudo systemctl restart httpd
```

---

##  5. Fix Required PHP Settings

### Increase `max_input_vars`

```bash
sudo sed -i -E 's/^\s*;?\s*max_input_vars\s*=.*/max_input_vars = 5000/' /etc/php.ini
```

```bash
sudo systemctl restart httpd
```

---

##  6. Clean Moodle Caches (Optional)

```bash
rm -rf /var/www/moodledata/cache/*
rm -rf /var/www/moodledata/localcache/*
rm -rf /var/www/moodledata/temp/*
rm -rf /var/www/moodledata/muc/*
```

---

##  7. Remove Incompatible or Missing Plugins

Manually remove plugins that block the upgrade:

```bash
rm -rf /var/www/html/moodle/question/type/ordering
rm -rf /var/www/html/moodle/filter/codehighlighter
rm -rf /var/www/html/moodle/lib/editor/tiny/plugins/html
rm -rf /var/www/html/moodle/lib/editor/tiny/plugins/noautolink
rm -rf /var/www/html/moodle/lib/editor/tiny/plugins/premium
rm -rf /var/www/html/moodle/admin/tool/mfa
rm -rf /var/www/html/moodle/report/themeusage
rm -rf /var/www/html/moodle/lib/h5p/h5plib_v127
```

Set permissions again:

```bash
chown -R apache:apache /var/www/html/moodle
```

---

##  8. Run the Upgrade

```bash
sudo -u apache php /var/www/html/moodle/admin/cli/upgrade.php
```

Upgrade successfully ran after all plugin and environment issues were resolved.

---

##  9. Post-Upgrade Verification

* Verified `version.php` shows: `4.1.5+`
* Verified PHP extensions: `gd`, `intl`, `sodium`, `soap` installed
* Verified database version:

````
