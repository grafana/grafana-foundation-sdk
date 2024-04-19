# Examples

This folder contains a collection of Grafana dashboards written in Python.

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
* [`red-method`](./red-method):
    * example of a dashboard following
      the [RED method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method)

## Running the examples

From an example's folder:

```console
$ python -m pip install -r requirements.txt
$ python main.py
```
