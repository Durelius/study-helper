# Server setup

One binary serves every course. Each course runs as its own instance of a systemd
template unit, with its own port, database and vhost — so a broken question file fails
that course's startup and leaves the others serving.

After the one-time steps below, shipping is `./deploy/deploy.sh [course…]`.

## 1. One-time, for the machine

```sh
sudo install -d -o wilhelm -g wilhelm /var/www/chulastudy /var/www/chulastudy/courses
sudo install -d -o root -g root -m 755 /etc/chulastudy
sudo install -o root -g root -m 644 deploy/chulastudy@.service \
    /etc/systemd/system/chulastudy@.service
sudo restorecon -v /etc/systemd/system/chulastudy@.service
sudo systemctl daemon-reload
```

`StateDirectory=chulastudy` creates `/var/lib/chulastudy` on first start and keeps it
writable while the rest of the filesystem is read-only to the process. **Databases live
there and are never touched by a deploy** — the binary is replaced, the leaderboard is
not.

### SELinux

This host enforces SELinux, and it is the one thing that silently stops a deploy.

- **The unit file**: copying it through `/tmp` leaves it labelled `user_tmp_t` and
  systemd answers `Failed to open …: Permission denied`. Use `install` from the repo,
  or `restorecon` after moving it.
- **The binary**: `/var/www` defaults to `httpd_sys_content_t`, which systemd will not
  execute — `status=203/EXEC`, again "Permission denied". It needs `bin_t`:

  ```sh
  sudo semanage fcontext -a -t bin_t "/var/www/chulastudy/chulastudy"
  sudo restorecon -v /var/www/chulastudy/chulastudy
  ```

  `deploy.sh` runs `restorecon` after every upload, which is what makes that rule take
  effect on the replacement binary.

## 2. Per course

Pick a free loopback port (valuechain 8093, genai-literacy 8094) and a subdomain.

```sh
# a) port and database
sudo install -o root -g root -m 644 deploy/<course>.env /etc/chulastudy/<course>.env

# b) vhost
sed -e "s/COURSE/<course>/g" -e "s/PORT/<port>/" deploy/vhost.conf.template > /tmp/v.conf
sudo install -o root -g root -m 644 /tmp/v.conf /etc/httpd/conf.d/<course>.conf
sudo restorecon -v /etc/httpd/conf.d/<course>.conf
sudo apachectl configtest && sudo systemctl reload httpd
sudo certbot --apache -d <course>.wilhelm.my

# c) start it
sudo systemctl enable --now chulastudy@<course>
```

Apache needs to be allowed to reach the loopback port. Already on for the other sites:

```sh
sudo setsebool -P httpd_can_network_connect 1
```

### Compression

`AddOutputFilterByType DEFLATE` in the vhost is not optional. Without it the frontend
bundle goes out at ~460 KB instead of ~140 KB, and on a weak mobile connection the
transfer is cut off part-way — which presents as a blank page, not an error, because the
HTML and CSS arrive and only the script is truncated.

**certbot copies the vhost when it creates the TLS one**, so directives added to
`<course>.conf` after running certbot must be added to `<course>-le-ssl.conf` as well —
that is the file actually serving traffic. Check it with:

```sh
curl -s -o /dev/null -H 'Accept-Encoding: gzip' -w '%{size_download}\n' \
  https://<course>.wilhelm.my/assets/<hashed>.js
```

## 3. Deploying

```sh
./deploy/deploy.sh                      # binary + every course
./deploy/deploy.sh valuechain           # one course
./deploy/sync-audio.sh valuechain ~/Desktop/valuechain-audiobook
```

## Backing up a leaderboard

```sh
ssh wilhelm@<host> 'sqlite3 /var/lib/chulastudy/<course>.db ".backup /tmp/backup.db"'
scp wilhelm@<host>:/tmp/backup.db .
```

Use `.backup` rather than copying the file: a plain `cp` while the server is writing can
catch it mid-transaction and miss whatever is still in the WAL.

## Live

| Course | Subdomain | Port | Database |
|---|---|---|---|
| valuechain | valuechain.wilhelm.my | 8093 | /var/lib/chulastudy/valuechain.db |
| genai-literacy | *not yet* | 8094 | /var/lib/chulastudy/genai-literacy.db |
