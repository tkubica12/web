cp web /opt/web
cp web.service /etc/systemd/system/web.service
chmod +x /opt/web
systemctl enable web.service
systemctl start web.service
