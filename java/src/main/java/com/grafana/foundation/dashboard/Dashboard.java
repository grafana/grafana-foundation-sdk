// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class Dashboard {
    // Unique numeric identifier for the dashboard.
    // `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("id")
    public Long id;
    // Unique dashboard identifier that can be generated by anyone. string (8-40)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    // Title of dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    // Description of dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // This property should only be used in dashboards defined by plugins.  It is a quick check
    // to see if the version has changed since the last time.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("revision")
    public Long revision;
    // ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("gnetId")
    public String gnetId;
    // Tags associated with dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("tags")
    public List<String> tags;
    // Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public String timezone;
    // Whether a dashboard is editable or not.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("editable")
    public Boolean editable;
    // Configuration of dashboard cursor sync behavior.
    // Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("graphTooltip")
    public DashboardCursorSync graphTooltip;
    // Time range for dashboard.
    // Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("time")
    public DashboardDashboardTime time;
    // Configuration of the time picker shown at the top of a dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timepicker")
    public TimePickerConfig timepicker;
    // The month that the fiscal year starts on.  0 = January, 11 = December
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fiscalYearStartMonth")
    public Integer fiscalYearStartMonth;
    // When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    // This will keep data "moving left" regardless of the query refresh rate. This setting helps
    // avoid dashboards presenting stale live data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("liveNow")
    public Boolean liveNow;
    // Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("weekStart")
    public String weekStart;
    // Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("refresh")
    public String refresh;
    // Version of the JSON schema, incremented each time a Grafana update brings
    // changes to said schema.
    @JsonProperty("schemaVersion")
    public Short schemaVersion;
    // Version of the dashboard, incremented each time the dashboard is updated.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("version")
    public Integer version;
    // List of dashboard panels
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("panels")
    public List<PanelOrRowPanel> panels;
    // Configured template variables
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("templating")
    public DashboardDashboardTemplating templating;
    // Contains the list of annotations that are associated with the dashboard.
    // Annotations are used to overlay event markers and overlay event tags on graphs.
    // Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
    // See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("annotations")
    public AnnotationContainer annotations;
    // Links with references to other dashboards or external websites.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("links")
    public List<DashboardLink> links;
    // Snapshot options. They are present only if the dashboard is a snapshot.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("snapshot")
    public Snapshot snapshot;
    public Dashboard() {
        this.timezone = "browser";
        this.editable = true;
        this.graphTooltip = DashboardCursorSync.OFF;
        this.fiscalYearStartMonth = 0;
        this.schemaVersion = 36;
    }
    
    public Dashboard(Long id,String uid,String title,String description,Long revision,String gnetId,List<String> tags,String timezone,Boolean editable,DashboardCursorSync graphTooltip,DashboardDashboardTime time,TimePickerConfig timepicker,Integer fiscalYearStartMonth,Boolean liveNow,String weekStart,String refresh,Short schemaVersion,Integer version,List<PanelOrRowPanel> panels,DashboardDashboardTemplating templating,AnnotationContainer annotations,List<DashboardLink> links,Snapshot snapshot) {
        this.id = id;
        this.uid = uid;
        this.title = title;
        this.description = description;
        this.revision = revision;
        this.gnetId = gnetId;
        this.tags = tags;
        this.timezone = timezone;
        this.editable = editable;
        this.graphTooltip = graphTooltip;
        this.time = time;
        this.timepicker = timepicker;
        this.fiscalYearStartMonth = fiscalYearStartMonth;
        this.liveNow = liveNow;
        this.weekStart = weekStart;
        this.refresh = refresh;
        this.schemaVersion = schemaVersion;
        this.version = version;
        this.panels = panels;
        this.templating = templating;
        this.annotations = annotations;
        this.links = links;
        this.snapshot = snapshot;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
