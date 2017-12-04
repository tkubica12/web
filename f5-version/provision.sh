cp web /opt/f5web
cp web.service /etc/systemd/system/web.service
chmod +x /opt/f5web
systemctl enable web.service
systemctl start web.service
