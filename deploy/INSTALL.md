# One-time setup on the Linode

Everything after this is `./deploy/deploy.sh`.

## 1. DNS

An `A` record for `valuechain.wilhelm.my` pointing at the Linode.

## 2. The service

```sh
sudo install -d -o wilhelm -g wilhelm /var/www/chulavaluechain
sudo cp deploy/valuechain.service /etc/systemd/system/chulavaluechain.service
sudo systemctl daemon-reload
sudo systemctl enable chulavaluechain
```

`StateDirectory=chulavaluechain` creates `/var/lib/chulavaluechain` on first start and
keeps it writable while the rest of the filesystem is read-only to the process. The
leaderboard lives in `study.db` there, so **do not** delete that directory between
deploys — the binary is replaced, the database is not.

## 3. Apache

```sh
sudo cp deploy/valuechain-vhost.conf /etc/httpd/conf.d/valuechain.conf
sudo apachectl configtest && sudo systemctl reload httpd
sudo certbot --apache -d valuechain.wilhelm.my
```

SELinux has to allow the proxy to reach the loopback port, the same as the courts app:

```sh
sudo setsebool -P httpd_can_network_connect 1
```

## 4. First deploy

```sh
./deploy/deploy.sh
```

## Backing up the leaderboard

```sh
ssh wilhelm@<host> 'sqlite3 /var/lib/chulavaluechain/study.db ".backup /tmp/study-backup.db"'
scp wilhelm@<host>:/tmp/study-backup.db .
```
