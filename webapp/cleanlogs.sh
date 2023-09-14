#!/bin/zsh

rm -f public/image/*.jpg
rm -f public/image/*.gif
rm -f public/image/*.png

docker compose exec nginx /bin/sh -c "rm -f /var/log/nginx/*.log && nginx -s reopen"

docker compose exec mysql /bin/sh -c "rm -f /var/log/mysql/mysql-slow.log && mysqladmin flush-logs -u root -proot"
