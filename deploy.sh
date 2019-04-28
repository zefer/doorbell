# TODO: make this generic and/or replace with https://github.com/zefer/ansible.
GOOS=linux GOARM=6 GOARCH=arm go build
scp ./doorbell joe@hendrix:/home/joe/doorbell
