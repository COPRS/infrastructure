# Duplicate databases backups

In order to mitigate security flaws concerning the database's backups, you can configure rclone to be run pediocally and clone specified S3 buckets on different locations.

 - Configure the S3 credentials in **apps/rlcone/rclone.conf**
 - In the **apps/rclone** deployment, configure the backup retention, the buckets to clone and the cron schedule in **clone-elasticsearch-backup.yaml**, **clone-ldap-backup.yaml** and **clone-postgresql-backup.yaml** at:
     - *spec.schedule*
     - *spec.jobTemplate.spec.template.spec.containers[].env*
