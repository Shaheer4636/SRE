
````markdown
##  Rollback Procedure (if upgrade fails)

If the Moodle upgrade fails or the system becomes unstable, follow this rollback procedure to restore your environment.

---

### Prerequisite

Ensure the backups from **Step 1: Pre-Upgrade Backups** were completed successfully before the upgrade.

---

###  1. Restore Moodle Code

```bash
rm -rf /var/www/html/moodle
cp -a /var/backups/moodle-backup-YYYY-MM-DD_HH-MM/moodle-code /var/www/html/moodle
````

> Replace `YYYY-MM-DD_HH-MM` with the actual backup timestamp.

---

###  2. Restore Moodle Database

```bash
mysql -u moodleadmin -p'MoodleDev!2025' \
  -h rds-techops-moodle01.ck3dgulnkqbo.us-west-2.rds.amazonaws.com \
  moodle < /var/backups/moodle-backup-YYYY-MM-DD_HH-MM/moodle-db.sql
```

---

### 3. Reset File Permissions

```bash
chown -R apache:apache /var/www/html/moodle
chmod -R 755 /var/www/html/moodle
```

---

###  4. Restart Apache

```bash
sudo systemctl restart httpd
```

---

### Result

Your Moodle instance should now be restored to its original **pre-upgrade** state, both in code and database.

```

Let me know if you’d like this as a downloadable `.md` file or embedded in a larger guide.
```
