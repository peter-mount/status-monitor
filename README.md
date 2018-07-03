# status-monitor

A simple container that act's as a bridge between [Grafana](https://grafana.com/) and [LambStatus](https://lambstatus.github.io)
allowing for alerts to be automatically reported as an Indcident on the status page and maintain those Indcidents so that when
an alert is cleared then the incident is automatically resolved.

Rules define how these are performed, if an alert is repeated in quick succession to reopen an incident rather than creating multiple ones.

## Real World Examples

You can see a real world example here:

* [NROD-TD API Status](https://grafana.area51.onl/d/HeoVU5Hmk/nrod-td-api-status?refresh=1m&orgId=1) is a live Grafana dashboard with alerts configured for Message Rate & Latency
* [Status Page](https://status.area51.onl/) the LambStatus installation showing the alerts being shown as incidents.
