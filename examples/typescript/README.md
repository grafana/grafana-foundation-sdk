# Examples

This folder contains a collection of Grafana dashboards written in Typescript.

Each example showcases different aspects of building dashboards as code:

* [`custom-panel`](./custom-panel): definition and usage of a _custom_ Panel type
* [`custom-query`](./custom-query): definition and usage of a _custom_ Query type
* [`grafana-agent-overview`](./grafana-agent-overview):
    * reproduction of the "Grafana Agent Overview" dashboard from
      the [Grafana Agent integration](https://grafana.com/docs/grafana-cloud/monitor-infrastructure/integrations/integration-reference/integration-grafana-agent/)
      available in Grafana Cloud.
    * dashboard variables
    * `table` panels
    * `timeseries` panels
    * `prometheus` queries
* [`linux-node-overview`](./linux-node-overview):
    * reproduction of the "Grafana Agent Overview" dashboard from
      the [Linux Server integration](https://grafana.com/docs/grafana-cloud/monitor-infrastructure/integrations/integration-reference/integration-linux-node/#dashboards)
      available in Grafana Cloud.
    * dashboard variables
    * dashboard links
    * `stat` panels
    * `table` panels
    * `timeseries` panels
    * `prometheus` queries

## Running the examples

From an example's folder:

```console
$ yarn install
$ yarn dev
```
