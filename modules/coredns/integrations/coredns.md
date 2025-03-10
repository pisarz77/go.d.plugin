<!--startmeta
custom_edit_url: "https://github.com/netdata/go.d.plugin/edit/master/modules/coredns/README.md"
meta_yaml: "https://github.com/netdata/go.d.plugin/edit/master/modules/coredns/metadata.yaml"
sidebar_label: "CoreDNS"
learn_status: "Published"
learn_rel_path: "Data Collection/DNS and DHCP Servers"
message: "DO NOT EDIT THIS FILE DIRECTLY, IT IS GENERATED BY THE COLLECTOR'S metadata.yaml FILE"
endmeta-->

# CoreDNS

Plugin: go.d.plugin
Module: coredns

<img src="https://img.shields.io/badge/maintained%20by-Netdata-%2300ab44" />

## Overview

This collector monitors CoreDNS instances.




This collector is supported on all platforms.

This collector supports collecting metrics from multiple instances of this integration, including remote instances.


### Default Behavior

#### Auto-Detection

This integration doesn't support auto-detection.

#### Limits

The default configuration for this integration does not impose any limits on data collection.

#### Performance Impact

The default configuration for this integration is not expected to impose a significant performance impact on the system.


## Metrics

Metrics grouped by *scope*.

The scope defines the instance that the metric belongs to. An instance is uniquely identified by a set of labels.



### Per CoreDNS instance

These metrics refer to the entire monitored application.

This scope has no labels.

Metrics:

| Metric | Dimensions | Unit |
|:------|:----------|:----|
| coredns.dns_request_count_total | requests | requests/s |
| coredns.dns_responses_count_total | responses | responses/s |
| coredns.dns_request_count_total_per_status | processed, dropped | requests/s |
| coredns.dns_no_matching_zone_dropped_total | dropped | requests/s |
| coredns.dns_panic_count_total | panics | panics/s |
| coredns.dns_requests_count_total_per_proto | udp, tcp | requests/s |
| coredns.dns_requests_count_total_per_ip_family | v4, v6 | requests/s |
| coredns.dns_requests_count_total_per_per_type | a, aaaa, mx, soa, cname, ptr, txt, ns, ds, dnskey, rrsig, nsec, nsec3, ixfr, any, other | requests/s |
| coredns.dns_responses_count_total_per_rcode | noerror, formerr, servfail, nxdomain, notimp, refused, yxdomain, yxrrset, nxrrset, notauth, notzone, badsig, badkey, badtime, badmode, badname, badalg, badtrunc, badcookie, other | responses/s |

### Per server

These metrics refer to the DNS server.

Labels:

| Label      | Description     |
|:-----------|:----------------|
| server_name | Server name. |

Metrics:

| Metric | Dimensions | Unit |
|:------|:----------|:----|
| coredns.server_dns_request_count_total | requests | requests/s |
| coredns.server_dns_responses_count_total | responses | responses/s |
| coredns.server_request_count_total_per_status | processed, dropped | requests/s |
| coredns.server_requests_count_total_per_proto | udp, tcp | requests/s |
| coredns.server_requests_count_total_per_ip_family | v4, v6 | requests/s |
| coredns.server_requests_count_total_per_per_type | a, aaaa, mx, soa, cname, ptr, txt, ns, ds, dnskey, rrsig, nsec, nsec3, ixfr, any, other | requests/s |
| coredns.server_responses_count_total_per_rcode | noerror, formerr, servfail, nxdomain, notimp, refused, yxdomain, yxrrset, nxrrset, notauth, notzone, badsig, badkey, badtime, badmode, badname, badalg, badtrunc, badcookie, other | responses/s |

### Per zone

These metrics refer to the DNS zone.

Labels:

| Label      | Description     |
|:-----------|:----------------|
| zone_name | Zone name. |

Metrics:

| Metric | Dimensions | Unit |
|:------|:----------|:----|
| coredns.zone_dns_request_count_total | requests | requests/s |
| coredns.zone_dns_responses_count_total | responses | responses/s |
| coredns.zone_requests_count_total_per_proto | udp, tcp | requests/s |
| coredns.zone_requests_count_total_per_ip_family | v4, v6 | requests/s |
| coredns.zone_requests_count_total_per_per_type | a, aaaa, mx, soa, cname, ptr, txt, ns, ds, dnskey, rrsig, nsec, nsec3, ixfr, any, other | requests/s |
| coredns.zone_responses_count_total_per_rcode | noerror, formerr, servfail, nxdomain, notimp, refused, yxdomain, yxrrset, nxrrset, notauth, notzone, badsig, badkey, badtime, badmode, badname, badalg, badtrunc, badcookie, other | responses/s |



## Alerts

There are no alerts configured by default for this integration.


## Setup

### Prerequisites

No action required.

### Configuration

#### File

The configuration file name for this integration is `go.d/coredns.conf`.


You can edit the configuration file using the `edit-config` script from the
Netdata [config directory](https://github.com/netdata/netdata/blob/master/docs/configure/nodes.md#the-netdata-config-directory).

```bash
cd /etc/netdata 2>/dev/null || cd /opt/netdata/etc/netdata
sudo ./edit-config go.d/coredns.conf
```
#### Options

The following options can be defined globally: update_every, autodetection_retry.


<details><summary>All options</summary>

| Name | Description | Default | Required |
|:----|:-----------|:-------|:--------:|
| update_every | Data collection frequency. | 1 | False |
| autodetection_retry | Recheck interval in seconds. Zero means no recheck will be scheduled. | 0 | False |
| url | Server URL. | http://127.0.0.1:9153/metrics | True |
| per_server_stats | Server filter. |  | False |
| per_zone_stats | Zone filter. |  | False |
| username | Username for basic HTTP authentication. |  | False |
| password | Password for basic HTTP authentication. |  | False |
| proxy_url | Proxy URL. |  | False |
| proxy_username | Username for proxy basic HTTP authentication. |  | False |
| proxy_password | Password for proxy basic HTTP authentication. |  | False |
| timeout | HTTP request timeout. | 2 | False |
| method | HTTP request method. | GET | False |
| body | HTTP request body. |  | False |
| headers | HTTP request headers. |  | False |
| not_follow_redirects | Redirect handling policy. Controls whether the client follows redirects. | False | False |
| tls_skip_verify | Server certificate chain and hostname validation policy. Controls whether the client performs this check. | False | False |
| tls_ca | Certification authority that the client uses when verifying the server's certificates. |  | False |
| tls_cert | Client tls certificate. |  | False |
| tls_key | Client tls key. |  | False |

##### per_server_stats

Metrics of servers matching the selector will be collected.
- Logic: (pattern1 OR pattern2) AND !(pattern3 or pattern4)
- Pattern syntax: [matcher](https://github.com/netdata/go.d.plugin/tree/master/pkg/matcher#supported-format).
- Syntax:

```yaml
per_server_stats:
  includes:
    - pattern1
    - pattern2
  excludes:
    - pattern3
    - pattern4
```


##### per_zone_stats

Metrics of zones matching the selector will be collected.
- Logic: (pattern1 OR pattern2) AND !(pattern3 or pattern4)
- Pattern syntax: [matcher](https://github.com/netdata/go.d.plugin/tree/master/pkg/matcher#supported-format).
- Syntax:

```yaml
per_zone_stats:
  includes:
    - pattern1
    - pattern2
  excludes:
    - pattern3
    - pattern4
```


</details>

#### Examples

##### Basic

An example configuration.

<details><summary>Config</summary>

```yaml
jobs:
  - name: local
    url: http://127.0.0.1:9153/metrics

```
</details>

##### Basic HTTP auth

Local server with basic HTTP authentication.

<details><summary>Config</summary>

```yaml
jobs:
  - name: local
    url: http://127.0.0.1:9153/metrics
    username: foo
    password: bar

```
</details>

##### Multi-instance

> **Note**: When you define multiple jobs, their names must be unique.

Collecting metrics from local and remote instances.


<details><summary>Config</summary>

```yaml
jobs:
  - name: local
    url: http://127.0.0.1:9153/metrics

  - name: remote
    url: http://203.0.113.10:9153/metrics

```
</details>



## Troubleshooting

### Debug Mode

To troubleshoot issues with the `coredns` collector, run the `go.d.plugin` with the debug option enabled. The output
should give you clues as to why the collector isn't working.

- Navigate to the `plugins.d` directory, usually at `/usr/libexec/netdata/plugins.d/`. If that's not the case on
  your system, open `netdata.conf` and look for the `plugins` setting under `[directories]`.

  ```bash
  cd /usr/libexec/netdata/plugins.d/
  ```

- Switch to the `netdata` user.

  ```bash
  sudo -u netdata -s
  ```

- Run the `go.d.plugin` to debug the collector:

  ```bash
  ./go.d.plugin -d -m coredns
  ```


