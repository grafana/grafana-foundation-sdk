---
title: <span class="badge object-type-class"></span> AzureMonitorQuery
---
# <span class="badge object-type-class"></span> AzureMonitorQuery

## Definition

```java
public class AzureMonitorQuery extends com.grafana.foundation.cog.variants.Dataquery {
  public String refId;
  public Boolean hide;
  public String queryType;
  public String subscription;
  public List<String> subscriptions;
  public AzureMetricQuery azureMonitor;
  public AzureLogsQuery azureLogAnalytics;
  public AzureResourceGraphQuery azureResourceGraph;
  public AzureTracesQuery azureTraces;
  public GrafanaTemplateVariableQuery grafanaTemplateVariableFn;
  public String resourceGroup;
  public String namespace;
  public String resource;
  public String region;
  public DataSourceRef datasource;
  public String query;
}
```
## See also

 * <span class="badge builder"></span> [AzureMonitorQueryBuilder](./builder-AzureMonitorQueryBuilder.md)
