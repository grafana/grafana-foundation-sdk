# Examples

This folder contains a collection of Grafana dashboards written in PHP.

Each example showcases different aspects of building dashboards as code:

* [`custom-panel`](./custom-panel): definition and usage of a _custom_ Panel type
* [`custom-query`](./custom-query): definition and usage of a _custom_ Query type
* [`red-method`](./red-method):
    * example of a dashboard following
      the [RED method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method)

## Running the examples

From an example's folder:

```console
$ composer install
$ php index.php
```
