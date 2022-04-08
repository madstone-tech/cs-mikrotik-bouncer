# CrowdSec Mikrotik Bouncer

This is a very simple and basic Mikrotik RouterOS Bouncer for CrowdSec. Redimentary would be an understatement. I am just trying to wrap my head around how the REST API works vs the Oldschool API. You will need to add the following environment variables to your machine:

* `ADDR`: The IP address of the RouterOS device
* `USERNAME`: The username of the RouterOS device
* `PASSWORD`: The password of the RouterOS device

I am currently only supporting REST API / RouterOS 7.X but hope to add support for RouterOS 6.X in the future.

This bouncer currently leverages Crowdsec's Custom Bouncer binary, which you will need to install and configure [here](https://docs.crowdsec.net/docs/bouncers/custom/). I have included an example custom bouncer configuration example that you can modify and setup on your Crowdsec server.

Build the binary by cloning into the repo, and use `go build .` to create a binary in this repo. 

You will need to set your Environment Variables up in your `crowdsec-custom-bouncer.service` systemd file. An example can be found below of the modified `[Service]` section.

```
[Service]
Environment=ADDR=192.168.88.1
Environment=USERNAME=amulet
Environment=PASSWORD=Password
Type=notify
ExecStart=/usr/bin/crowdsec-custom-bouncer -c /etc/crowdsec/bouncers/>
ExecStartPost=/bin/sleep 0.1
```
