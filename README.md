# Prometheus Network UPS Exporter

This tool exposes Network UPS (nuts) statistics as Prometheus metrics.

This exporter works similar to the [Blackbox Exporter](https://github.com/prometheus/blackbox_exporter) where you can specify different targets in Prometheus to scrape Mikrotik devices through this exporter.

## Exposed Metrics

| Metric                    | Definition                                               |
| ------------------------- | -------------------------------------------------------- |
| nut_battery_charge        | Battery charge (percent)                                 |
| nut_ups_load              | Load on UPS (percent)                                    |
| nut_ups_realpower_nominal | Nominal value of real power (Watts)                      |
| nut_ups_on_line_power     | Displays whether or not the ups is running on line power |

## Prometheus Configuration

This exporter implements the multi-target exporter pattern, therefore we recommend reading the guide [Understanding and using the multi-target exporter pattern](https://prometheus.io/docs/guides/multi-target-exporter/) to get an overview of the configuration.

The target must be passed to the nut-exporter as a parameter, this can be done with relabelling.

Example config:

```yaml
scrape_configs:
  - job_name: "nut"
    metrics_path: /probe
    static_configs:
      - targets:
          - 10.0.1.1
          - 10.0.20.1
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 127.0.0.1:9055 # The nut-exporter's real hostname:port.
```
