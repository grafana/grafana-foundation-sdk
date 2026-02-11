---
weight: 10
---
# Examples

A collection of sample Grafana dashboards written in Python is available in the [grafana/grafana-foundation-sdk](https://github.com/grafana/grafana-foundation-sdk/) repository.

Each example showcases different aspects of building dashboards as code:

* [`custom-panel`](https://github.com/grafana/grafana-foundation-sdk/blob/main/examples/python/custom-panel): definition and usage of a _custom_ Panel type
* [`custom-query`](https://github.com/grafana/grafana-foundation-sdk/blob/main/examples/python/custom-query): definition and usage of a _custom_ Query type
* [`grafana-agent-overview`](https://github.com/grafana/grafana-foundation-sdk/blob/main/examples/python/grafana-agent-overview):
    * reproduction of the "Grafana Agent Overview" dashboard from
      the [Grafana Agent integration](https://grafana.com/docs/grafana-cloud/monitor-infrastructure/integrations/integration-reference/integration-grafana-agent/)
      available in Grafana Cloud.
    * dashboard variables
    * `table` panels
    * `timeseries` panels
    * `prometheus` queries
* [`linux-node-overview`](https://github.com/grafana/grafana-foundation-sdk/blob/main/examples/python/linux-node-overview):
    * reproduction of the "Grafana Agent Overview" dashboard from
      the [Linux Server integration](https://grafana.com/docs/grafana-cloud/monitor-infrastructure/integrations/integration-reference/integration-linux-node/#dashboards)
      available in Grafana Cloud.
    * dashboard variables
    * dashboard links
    * `stat` panels
    * `table` panels
    * `timeseries` panels
    * `prometheus` queries
* [`red-method`](https://github.com/grafana/grafana-foundation-sdk/blob/main/examples/python/red-method):
    * example of a dashboard following
      the [RED method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method)
