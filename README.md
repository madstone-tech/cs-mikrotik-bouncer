# CrowdSec Mikrotik Bouncer

This is a very simple and basic Mikrotik RouterOS Bouncer for CrowdSec. Redimentary would be an understatement. I am just trying to wrap my head around how the REST API works vs the Oldschool API. You will need an `.env` file in the root directory of the project which contains the following variables:

* `ADDR`: The IP address of the RouterOS device
* `USERNAME`: The username of the RouterOS device
* `PASSWORD`: The password of the RouterOS device

I am currently only supporting REST API / RouterOS 7.X but hope to add support for RouterOS 6.X in the future.