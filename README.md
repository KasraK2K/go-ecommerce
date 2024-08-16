# Useful Documentation
---
<br />

This is not full documentation but in here I explain most useful things that you cant see in the code, and it's necessary.
<br />

I hope it will be useful for you.
<br />
<br />

## Deploy Go on linux server
<br />

For deploy go on linux server do these steps:
<br />

1. Build your app with necessary environments:
```bash
env GOOS=linux CGO_ENABLED=0 go build -o blog ./cmd/api
```
<br />

1. Create a systemd job file:
```bash
touch /etc/systemd/system/blog.service
```
<br />

1. Copy these code into `blog.service` file:
```bash
Description= instance to serve blog api
After=network.target

[Service]
Environment="MODE=production"
User=root
Group=www-data
ExecStart=<Path of your builded go file>
[Install]
WantedBy=multi-users.target
```
NOTE: As you can see this file created with production mode environment
<br />

<br />

4. Start service:
```bash
sudo systemctl start blog
```
<br />

5. Check service status:
```bash
sudo systemctl enable blog
```
<br />

6. Create blog config for nginx:
```bash
touch /etc/nginx/sites-available/blog
```
<br />

7. Copy these codes into blog config file:
```bash
server {
    server_name .com;

    location / {
        proxy_set_header  X-Real-IP  $remote_addr;
        proxy_pass http://localhost:<Port Number>;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }

    listen 80;
    listen [::]:80;
}
```
<br />

9. Create soft link for sites-enables:
```bash
sudo ln -s /etc/nginx/sites-available/blog /etc/nginx/sites-enabled
```
<br />

10. Restart nginx:
```bash
sudo systemctl restart nginx
```
