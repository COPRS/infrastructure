#!/bin/bash
MD5="$(echo $BUCKET | grep -E -i -o '.[0-9a-f]{32}$')"
BUCKET_DEST="${BUCKET//$MD5}.rclone"
rclone copy source:$BUCKET dest:$BUCKET_DEST/current --backup-dir dest:$BUCKET_DEST/archive/$(date +%s)
limit=$(( `date +%s` - $RETENTION_DAYS*86400 ))
for backup in $(rclone lsf dest:$BUCKET_DEST/archive); do 
    echo $backup
    if [[ ${backup%%/} -lt $limit ]]; then         
    rclone purge dest:$BUCKET_DEST/archive/$backup
    fi
done
