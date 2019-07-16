echo "* * * * * /usr/local/bin/php /var/www/html/artisan schedule:run >> /var/log/laravel-scheduler.log 2>&1"  >> /var/spool/cron/crontabs/root
/usr/sbin/crond