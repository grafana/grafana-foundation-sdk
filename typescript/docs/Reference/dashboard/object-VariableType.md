---
title: <span class="badge object-type-enum"></span> VariableType
---
# <span class="badge object-type-enum"></span> VariableType

Dashboard variable type

`query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.

`adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).

`constant`: 	Define a hidden constant.

`datasource`: Quickly change the data source for an entire dashboard.

`interval`: Interval variables represent time spans.

`textbox`: Display a free text input field with an optional default value.

`custom`: Define the variable options manually using a comma-separated list.

`system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables

## Definition

```typescript
export enum VariableType {
	Query = "query",
	Adhoc = "adhoc",
	Constant = "constant",
	Datasource = "datasource",
	Interval = "interval",
	Textbox = "textbox",
	Custom = "custom",
	System = "system",
}

```
