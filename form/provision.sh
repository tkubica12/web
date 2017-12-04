cp web /opt/form
cp web.service /etc/systemd/system/web.service
chmod +x /opt/form
systemctl enable web.service
systemctl start web.service
