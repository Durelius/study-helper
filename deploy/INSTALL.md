# One-time setup on the Linode

Everything after this is `./deploy/deploy.sh`.

## 1. DNS

An `A` record for `valuechain.wilhelm.my` pointing at the Linode.

## 2. The service

```sh
sudo install -d -o wilhelm -g wilhelm /var/www/chulavaluechain
sudo install -o root -g root -m 644 deploy/valuechain.service \
    /etc/systemd/system/chulavaluechain.service
sudo restorecon -v /etc/systemd/system/chulavaluechain.service
sudo systemctl daemon-reload
sudo systemctl enable chulavaluechain
```

### SELinux

This host enforces SELinux, and it is the one thing that will silently stop the deploy.
Two labels have to be right, and neither is by default:

- **The unit file.** Copying it in through `/tmp` leaves it labelled `user_tmp_t`, and
  systemd answers `Failed to open …: Permission denied`. Use `install` from the repo as
  above, or `restorecon` after moving it.
- **The binary.** `/var/www` defaults to `httpd_sys_content_t`, which systemd will not
  execute — the failure is `status=203/EXEC`, again "Permission denied". The binary
  needs `bin_t`, and a persistent rule keeps it that way through future deploys and any
  full relabel:

  ```sh
  sudo semanage fcontext -a -t bin_t "/var/www/chulavaluechain/chulavaluechain"
  sudo restorecon -v /var/www/chulavaluechain/chulavaluechain
  ```

  `deploy.sh` runs `restorecon` after every upload, which is what makes that rule take
  effect on the replacement binary.

`StateDirectory=chulavaluechain` creates `/var/lib/chulavaluechain` on first start and
keeps it writable while the rest of the filesystem is read-only to the process. The
leaderboard lives in `study.db` there, so **do not** delete that directory between
deploys — the binary is replaced, the database is not.

## 3. Apache

```sh
sudo install -o root -g root -m 644 deploy/valuechain-vhost.conf \
    /etc/httpd/conf.d/valuechain.conf
sudo restorecon -v /etc/httpd/conf.d/valuechain.conf
sudo apachectl configtest && sudo systemctl reload httpd
sudo certbot --apache -d valuechain.wilhelm.my
```

The proxy also has to be allowed to reach the loopback port. It was already on for the
courts app, so this is usually a no-op:

```sh
sudo setsebool -P httpd_can_network_connect 1
```

### Compression

`AddOutputFilterByType DEFLATE` in the vhost is not optional. Without it the frontend
bundle goes out at ~460 KB instead of ~140 KB, and on a weak mobile connection the
transfer is cut off part-way — which presents as a blank page, not as an error, because
the HTML and CSS arrive and only the script is truncated.

**certbot copies the vhost when it creates the TLS one**, so if you add directives to
`valuechain.conf` after running certbot, add them to `valuechain-le-ssl.conf` as well —
that is the file actually serving traffic.

Check it with:

```sh
curl -s -o /dev/null -H 'Accept-Encoding: gzip' -w '%{size_download}\n' \
  https://valuechain.wilhelm.my/assets/<hashed>.js
```

## 4. First deploy

```sh
./deploy/deploy.sh
```

**Done on 2026-09-21.** Live at https://valuechain.wilhelm.my, on `127.0.0.1:8093`,
certificate expiring 2026-12-20 and renewing on certbot's existing timer.

## Backing up the leaderboard

```sh
ssh wilhelm@<host> 'sqlite3 /var/lib/chulavaluechain/study.db ".backup /tmp/study-backup.db"'
scp wilhelm@<host>:/tmp/study-backup.db .
```
