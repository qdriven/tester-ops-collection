# Windows Docker环境

## Windows Docker 瘦身

一般Windows使用WSL作为Docker的资源

```sh
$wsl -l -v --all
  NAME                   STATE           VERSION
* docker-desktop         Running         2
  docker-desktop-data    Running         2
  Ubuntu-20.04           Stopped         2

```

备份Docker

```sh
wsl --export docker-desktop-data D:\docker\docker-desktop-data.tar
wsl --export docker-desktop D:\docker\docker-desktop.tar
```

## commands

```sh
wsl --unregister docker-desktop-data
wsl --unregister docker-desktop
```

```sh
wsl --import docker-desktop-data D:\docker\docker-desktop-data .\docker-desktop-data.tar --version 2
wsl --import docker-desktop D:\docker\docker-desktop .\docker-desktop.tar --version 2
```