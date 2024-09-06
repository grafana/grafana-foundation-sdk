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
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import com.grafana.foundation.dashboard.DataTransformerConfig;
import com.grafana.foundation.dashboard.FieldConfigSource;

@JsonDeserialize(using = LibrarypanelLibraryPanelModelDeserializer.class)
public class LibrarypanelLibraryPanelModel {
    // The panel plugin type id. This is used to find the plugin to display the panel.
    @JsonProperty("type")
    public String type;
    // The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pluginVersion")
    public String pluginVersion;
    // Tags for the panel.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("tags")
    public List<String> tags;
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
    public LibraryPanelRepeatDirection repeatDirection;
    // Id of the repeating panel.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeatPanelId")
    public Long repeatPanelId;
    // The maximum number of data points that the panel queries are retrieving.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxDataPoints")
    public Double maxDataPoints;
    // List of transformations that are applied to the panel data before rendering.
    // When there are multiple transformations, Grafana applies them in the order they are listed.
    // Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
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
    // It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("options")
    public Object options;
    // Field options allow you to change how the data is displayed in your visualizations.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("fieldConfig")
    public FieldConfigSource fieldConfig;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LibrarypanelLibraryPanelModel> {
        private final LibrarypanelLibraryPanelModel internal;
        
        public Builder() {
            this.internal = new LibrarypanelLibraryPanelModel();
        this.transparent(false);
        }
    public Builder type(String type) {
        if (!(type.length() >= 1)) {
            throw new IllegalArgumentException("type.length() must be >= 1");
        }
    this.internal.type = type;
        return this;
    }
    
    public Builder pluginVersion(String pluginVersion) {
    this.internal.pluginVersion = pluginVersion;
        return this;
    }
    
    public Builder tags(List<String> tags) {
    this.internal.tags = tags;
        return this;
    }
    
    public Builder targets(com.grafana.foundation.cog.Builder<List<Dataquery>> targets) {
    this.internal.targets = targets.build();
        return this;
    }
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public Builder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public Builder transparent(Boolean transparent) {
    this.internal.transparent = transparent;
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder links(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
    this.internal.links = links.build();
        return this;
    }
    
    public Builder repeat(String repeat) {
    this.internal.repeat = repeat;
        return this;
    }
    
    public Builder repeatDirection(LibraryPanelRepeatDirection repeatDirection) {
    this.internal.repeatDirection = repeatDirection;
        return this;
    }
    
    public Builder repeatPanelId(Long repeatPanelId) {
    this.internal.repeatPanelId = repeatPanelId;
        return this;
    }
    
    public Builder maxDataPoints(Double maxDataPoints) {
    this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public Builder transformations(List<DataTransformerConfig> transformations) {
    this.internal.transformations = transformations;
        return this;
    }
    
    public Builder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public Builder timeFrom(String timeFrom) {
    this.internal.timeFrom = timeFrom;
        return this;
    }
    
    public Builder timeShift(String timeShift) {
    this.internal.timeShift = timeShift;
        return this;
    }
    
    public Builder options(Object options) {
    this.internal.options = options;
        return this;
    }
    
    public Builder fieldConfig(FieldConfigSource fieldConfig) {
    this.internal.fieldConfig = fieldConfig;
        return this;
    }
    public LibrarypanelLibraryPanelModel build() {
            return this.internal;
        }
    }
}
