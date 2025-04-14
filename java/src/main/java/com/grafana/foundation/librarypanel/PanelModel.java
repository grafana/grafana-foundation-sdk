// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.dashboard.DataSourceRef;
import com.grafana.foundation.dashboard.DashboardLink;
import com.grafana.foundation.dashboard.DataTransformerConfig;
import com.grafana.foundation.dashboard.FieldConfigSource;

// Dashboard panels are the basic visualization building blocks.
@JsonDeserialize(using = PanelModelDeserializer.class)
public class PanelModel {
    // The panel plugin type id. This is used to find the plugin to display the panel.
    @JsonProperty("type")
    public String type;
    // The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pluginVersion")
    public String pluginVersion;
    // Depends on the panel plugin. See the plugin documentation for details.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("targets")
    public List<Dataquery> targets;
    // Panel title.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    // Panel description.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // Whether to display the panel without a background.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("transparent")
    public Boolean transparent;
    // The datasource used in all targets.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Panel links.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("links")
    public List<DashboardLink> links;
    // Name of template variable to repeat for.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat")
    public String repeat;
    // Direction to repeat in if 'repeat' is set.
    // `h` for horizontal, `v` for vertical.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeatDirection")
    public PanelModelRepeatDirection repeatDirection;
    // Option for repeated panels that controls max items per row
    // Only relevant for horizontally repeated panels
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxPerRow")
    public Double maxPerRow;
    // The maximum number of data points that the panel queries are retrieving.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxDataPoints")
    public Double maxDataPoints;
    // List of transformations that are applied to the panel data before rendering.
    // When there are multiple transformations, Grafana applies them in the order they are listed.
    // Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("transformations")
    public List<DataTransformerConfig> transformations;
    // The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
    // This value must be formatted as a number followed by a valid time
    // identifier like: "40s", "3d", etc.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    // Overrides the relative time range for individual panels,
    // which causes them to be different than what is selected in
    // the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
    // time periods or days on the same dashboard.
    // The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
    // `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeFrom")
    public String timeFrom;
    // Overrides the time range for individual panels by shifting its start and end relative to the time picker.
    // For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeShift")
    public String timeShift;
    // Controls if the timeFrom or timeShift overrides are shown in the panel header
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideTimeOverride")
    public Boolean hideTimeOverride;
    // Sets panel queries cache timeout.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cacheTimeout")
    public String cacheTimeout;
    // Overrides the data source configured time-to-live for a query cache item in milliseconds
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryCachingTTL")
    public Double queryCachingTTL;
    // It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("options")
    public Object options;
    // Field options allow you to change how the data is displayed in your visualizations.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fieldConfig")
    public FieldConfigSource fieldConfig;
    public PanelModel() {
        this.transparent = false;
    }
    public PanelModel(String type,String pluginVersion,List<Dataquery> targets,String title,String description,Boolean transparent,DataSourceRef datasource,List<DashboardLink> links,String repeat,PanelModelRepeatDirection repeatDirection,Double maxPerRow,Double maxDataPoints,List<DataTransformerConfig> transformations,String interval,String timeFrom,String timeShift,Boolean hideTimeOverride,String cacheTimeout,Double queryCachingTTL,Object options,FieldConfigSource fieldConfig) {
        this.type = type;
        this.pluginVersion = pluginVersion;
        this.targets = targets;
        this.title = title;
        this.description = description;
        this.transparent = transparent;
        this.datasource = datasource;
        this.links = links;
        this.repeat = repeat;
        this.repeatDirection = repeatDirection;
        this.maxPerRow = maxPerRow;
        this.maxDataPoints = maxDataPoints;
        this.transformations = transformations;
        this.interval = interval;
        this.timeFrom = timeFrom;
        this.timeShift = timeShift;
        this.hideTimeOverride = hideTimeOverride;
        this.cacheTimeout = cacheTimeout;
        this.queryCachingTTL = queryCachingTTL;
        this.options = options;
        this.fieldConfig = fieldConfig;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
