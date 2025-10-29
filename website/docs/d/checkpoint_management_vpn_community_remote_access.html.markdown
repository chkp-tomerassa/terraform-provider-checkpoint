---
layout: "checkpoint"
page_title: "checkpoint_management_vpn_community_remote_access"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-vpn-community-remote-access"
description: |-
This resource allows you to execute Check Point VPN Community Remote Access.
---

# Data Source: checkpoint_management_vpn_community_remote_access

This resource allows you to execute Check Point VPN Community Remote Access.

## Example Usage


```hcl
resource "checkpoint_management_vpn_community_remote_access" "vpn_community_remote_access" {
    name = "RemoteAccess"
	user_groups = ["All Users"]
}

data "checkpoint_management_vpn_community_remote_access" "data_vpn_community_remote_access" {
    name = "${checkpoint_management_vpn_community_remote_access.vpn_community_remote_access.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.  
* `gateways` - Collection of Gateway objects identified by the name or UID.
* `user_groups` - Collection of User group objects identified by the name or UID.
* `override_vpn_domains` - The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. 
* `comments` - Comments string.

`override_vpn_domains` supports the following:

* `gateway` - Participant gateway in override VPN domain identified by the name or UID.
* `vpn_domain` - VPN domain network identified by the name or UID. 
