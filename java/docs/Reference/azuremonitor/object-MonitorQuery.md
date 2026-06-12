---
title: <span class="badge object-type-class"></span> MonitorQuery
---
# <span class="badge object-type-class"></span> MonitorQuery

## Definition

```java
public class MonitorQuery extends com.grafana.foundation.cog.variants.Dataquery {
  public String refId;
  public Boolean hide;
  public String queryType;
  public String subscription;
  public List<String> subscriptions;
  public MetricQuery azureMonitor;
  public LogsQuery azureLogAnalytics;
  public ResourceGraphQuery azureResourceGraph;
  public TracesQuery azureTraces;
  public GrafanaTemplateVariableQuery grafanaTemplateVariableFn;
  public String resourceGroup;
  public String namespace;
  public String resource;
  public String region;
  public String customNamespace;
  public DataSourceRef datasource;
  public String query;
}
```
## See also

 * <span class="badge builder"></span> [MonitorQueryBuilder](./builder-MonitorQueryBuilder.md)
