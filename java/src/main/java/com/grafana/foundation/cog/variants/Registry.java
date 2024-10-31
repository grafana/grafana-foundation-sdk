// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cog.variants;

import java.util.HashMap;
import java.util.Map;

public class Registry {
    private static final Map<String, PanelConfig> panelRegistry = new HashMap<>();
    private static final Map<String, DataqueryConfig> dataqueryRegistry = new HashMap<>();
    
    static {
        registerPanel("alertgroups", com.grafana.foundation.alertgroups.Options.class, null);
        registerPanel("annolist", com.grafana.foundation.annotationslist.Options.class, null);
        registerPanel("barchart", com.grafana.foundation.barchart.Options.class, com.grafana.foundation.barchart.FieldConfig.class);
        registerPanel("bargauge", com.grafana.foundation.bargauge.Options.class, null);
        registerPanel("candlestick", com.grafana.foundation.candlestick.Options.class, com.grafana.foundation.candlestick.FieldConfig.class);
        registerPanel("canvas", com.grafana.foundation.canvas.Options.class, null);
        registerPanel("dashlist", com.grafana.foundation.dashboardlist.Options.class, null);
        registerPanel("datagrid", com.grafana.foundation.datagrid.Options.class, null);
        registerPanel("debug", com.grafana.foundation.debug.Options.class, null);
        registerPanel("gauge", com.grafana.foundation.gauge.Options.class, null);
        registerPanel("geomap", com.grafana.foundation.geomap.Options.class, null);
        registerPanel("heatmap", com.grafana.foundation.heatmap.Options.class, com.grafana.foundation.heatmap.FieldConfig.class);
        registerPanel("histogram", com.grafana.foundation.histogram.Options.class, com.grafana.foundation.histogram.FieldConfig.class);
        registerPanel("logs", com.grafana.foundation.logs.Options.class, null);
        registerPanel("news", com.grafana.foundation.news.Options.class, null);
        registerPanel("nodegraph", com.grafana.foundation.nodegraph.Options.class, null);
        registerPanel("piechart", com.grafana.foundation.piechart.Options.class, com.grafana.foundation.piechart.FieldConfig.class);
        registerPanel("stat", com.grafana.foundation.stat.Options.class, null);
        registerPanel("state-timeline", com.grafana.foundation.statetimeline.Options.class, com.grafana.foundation.statetimeline.FieldConfig.class);
        registerPanel("status-history", com.grafana.foundation.statushistory.Options.class, com.grafana.foundation.statushistory.FieldConfig.class);
        registerPanel("table", com.grafana.foundation.table.Options.class, null);
        registerPanel("text", com.grafana.foundation.text.Options.class, null);
        registerPanel("timeseries", com.grafana.foundation.timeseries.Options.class, com.grafana.foundation.timeseries.FieldConfig.class);
        registerPanel("trend", com.grafana.foundation.trend.Options.class, com.grafana.foundation.trend.FieldConfig.class);
        registerPanel("xychart", com.grafana.foundation.xychart.Options.class, com.grafana.foundation.xychart.FieldConfig.class);
        registerDataquery("__expr__", com.grafana.foundation.expr.Expr.class);
        registerDataquery("cloud-monitoring", com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQuery.class);
        registerDataquery("cloudwatch", com.grafana.foundation.cloudwatch.CloudWatchQuery.class);
        registerDataquery("elasticsearch", com.grafana.foundation.elasticsearch.Dataquery.class);
        registerDataquery("grafana-azure-monitor-datasource", com.grafana.foundation.azuremonitor.AzureMonitorQuery.class);
        registerDataquery("grafanapyroscope", com.grafana.foundation.grafanapyroscope.Dataquery.class);
        registerDataquery("loki", com.grafana.foundation.loki.Dataquery.class);
        registerDataquery("parca", com.grafana.foundation.parca.Dataquery.class);
        registerDataquery("prometheus", com.grafana.foundation.prometheus.Dataquery.class);
        registerDataquery("tempo", com.grafana.foundation.tempo.TempoQuery.class);
        registerDataquery("testdata", com.grafana.foundation.testdata.Dataquery.class);
    }

    public static void registerDataquery(String type, Class<? extends Dataquery> clazz) {
        dataqueryRegistry.put(type, new DataqueryConfig(clazz));
    }

    public static Class<? extends Dataquery> getDataquery(String type) {
        return dataqueryRegistry.get(type).getDataquery();
    }
    
    public static void registerPanel(String type, Class<?> options, Class<?> fieldConfig) {
        panelRegistry.put(type, new PanelConfig(options, fieldConfig));
    }

    public static PanelConfig getPanel(String type) {
        return panelRegistry.get(type);
    }
}
