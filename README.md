# Prometheus Network UPS Exporter

This tool exposes Network UPS (nuts) statistics as Prometheus metrics.

## Exposed Metrics

| Metric                    | Definition                                               |
| ------------------------- | -------------------------------------------------------- |
| nut_battery_charge        | Battery charge (percent)                                 |
| nut_ups_load              | Load on UPS (percent)                                    |
| nut_ups_realpower_nominal | Nominal value of real power (Watts)                      |
| nut_ups_on_line_power     | Displays whether or not the ups is running on line power |

## Configuration

The following environment variables are required to start the service:

| Environment Variable | Definition                                                  | Example       |
| -------------------- | ----------------------------------------------------------- | ------------- |
| UPS_SERVER           | IP address of the UPS server                                | 192.168.1.100 |
| LISTENING_ADDR       | The listening address to which the service binds itself     | 0.0.0.0:9055  |
| INTERVAL             | Polling interval to get the latest information from the UPS | 15            |

## Development

### Run locally

```sh
$ cp .env.example .env # adjust the values to yorur needs!
$ make dev
```
