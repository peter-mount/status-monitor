# status-monitor

A simple container that act's as a bridge between [Grafana](https://grafana.com/) and [LambStatus](https://lambstatus.github.io)
allowing for alerts to be automatically reported as an Indcident on the status page and maintain those Indcidents so that when
an alert is cleared then the incident is automatically resolved.

Rules define how these are performed, if an alert is repeated in quick succession to reopen an incident rather than creating multiple ones.
