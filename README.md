Connman IP Updater
---

This is a tool which allows you to use [connman vpn](https://github.com/chewitt/connman/blob/master/doc/vpn-config-format.txt) ([Wireguard](https://www.wireguard.com/)) with a DDNS.
It was originally build for [LibreELEC](https://libreelec.tv/) but may work on other platforms as well.

Installation (LibreELEC)
---

> [!NOTE]
> This guide assumes you have already set up Wireguard in your LibreELEC installation. If not click [here](https://wiki.libreelec.tv/configuration/wireguard) for a guide.
> Make sure you have enabled **SSH** and **cron** in Kodi.

1. SSH into your LibreELEC computer:
```sh
# your IP address here
# default password is libreelec
ssh root@123.123.123.123
```

2. `cd` into the wireguard directory:
```sh
cd /storage/.config/wireguard/
```

3. Download your required binary from the [releases page](https://github.com/TgZ39/DDnsIpUpdater/releases) (depending on the architecture you're using you may need to build it yourself):
```sh
wget https://github.com/TgZ39/DDnsIpUpdater/releases/download/<VERSION>/<NAME> # Download the binary
mv <NAME> ip_updater # rename the file
```

4. Create the config file:
```sh
nano ip_updater_config.toml
```
and edit it to something like this:
```toml
domain = "your-ddns-domain.duckdns.org"
wireguardconfigfile = "your-wireguard.config"
```
save the file `STRG + o` and exit the file `STRG + x`.

5. Create an update shell script:
```sh
nano update.sh
```
and edit it to this:
```sh
#! /bin/bash

cd /storage/.config/wireguard/
./ip_updater
```

6. Create a cron job to run the shell scirpt every 5 minutes:
```sh
crontab -e
```
add this line to the crontab file:
```cron
*/5 * * * * /storage/.config/wireguard/update.sh > /dev/null
```
