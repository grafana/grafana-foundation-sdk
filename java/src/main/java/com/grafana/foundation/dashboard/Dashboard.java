// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;
import java.util.LinkedList;

public class Dashboard {
    // Unique numeric identifier for the dashboard.
    // `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("id")
    public Long id;
    // Unique dashboard identifier that can be generated by anyone. string (8-40)
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    // Title of dashboard.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("title")
    public String title;
    // Description of dashboard.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("description")
    public String description;
    // This property should only be used in dashboards defined by plugins.  It is a quick check
    // to see if the version has changed since the last time.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("revision")
    public Long revision;
    // ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("gnetId")
    public String gnetId;
    // Tags associated with dashboard.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("tags")
    public List<String> tags;
    // Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("timezone")
    public String timezone;
    // Whether a dashboard is editable or not.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("editable")
    public Boolean editable;
    // Configuration of dashboard cursor sync behavior.
    // Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("graphTooltip")
    public DashboardCursorSync graphTooltip;
    // Time range for dashboard.
    // Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("time")
    public DashboardDashboardTime time;
    // Configuration of the time picker shown at the top of a dashboard.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("timepicker")
    public TimePickerConfig timepicker;
    // The month that the fiscal year starts on.  0 = January, 11 = December
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("fiscalYearStartMonth")
    public Integer fiscalYearStartMonth;
    // When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    // This will keep data "moving left" regardless of the query refresh rate. This setting helps
    // avoid dashboards presenting stale live data
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("liveNow")
    public Boolean liveNow;
    // Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("weekStart")
    public String weekStart;
    // Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("refresh")
    public String refresh;
    // Version of the JSON schema, incremented each time a Grafana update brings
    // changes to said schema.
    @JsonProperty("schemaVersion")
    public Short schemaVersion;
    // Version of the dashboard, incremented each time the dashboard is updated.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("version")
    public Integer version;
    // List of dashboard panels
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("panels")
    public List<PanelOrRowPanel> panels;
    // Configured template variables
    @JsonProperty("templating")
    public DashboardDashboardTemplating templating;
    // Contains the list of annotations that are associated with the dashboard.
    // Annotations are used to overlay event markers and overlay event tags on graphs.
    // Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
    // See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
    @JsonProperty("annotations")
    public AnnotationContainer annotations;
    // Links with references to other dashboards or external websites.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("links")
    public List<DashboardLink> links;
    // Snapshot options. They are present only if the dashboard is a snapshot.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("snapshot")
    public Snapshot snapshot;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Dashboard> {
        private final Dashboard internal;
        private Integer currentY;
        private Integer currentX;
        private Integer lastPanelHeight;
        
        public Builder(String title) {
            this.internal = new Dashboard();
    this.internal.title = title;
    this.internal.schemaVersion = 36;
    this.internal.editable = true;
        this.currentY = 0;
        this.currentX = 0;
        this.lastPanelHeight = 0;
        this.timezone("browser");
        this.tooltip(DashboardCursorSync.OFF);
        this.fiscalYearStartMonth(0);
        }
    public Builder id(Long id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder uid(String uid) {
    this.internal.uid = uid;
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
    
    public Builder revision(Long revision) {
    this.internal.revision = revision;
        return this;
    }
    
    public Builder gnetId(String gnetId) {
    this.internal.gnetId = gnetId;
        return this;
    }
    
    public Builder tags(List<String> tags) {
    this.internal.tags = tags;
        return this;
    }
    
    public Builder timezone(String timezone) {
    this.internal.timezone = timezone;
        return this;
    }
    
    public Builder editable() {
    this.internal.editable = true;
        return this;
    }
    
    public Builder readonly() {
    this.internal.editable = false;
        return this;
    }
    
    public Builder tooltip(DashboardCursorSync graphTooltip) {
    this.internal.graphTooltip = graphTooltip;
        return this;
    }
    
    public Builder time(com.grafana.foundation.cog.Builder<DashboardDashboardTime> time) {
    this.internal.time = time.build();
        return this;
    }
    
    public Builder timepicker(com.grafana.foundation.cog.Builder<TimePickerConfig> timepicker) {
    this.internal.timepicker = timepicker.build();
        return this;
    }
    
    public Builder fiscalYearStartMonth(Integer fiscalYearStartMonth) {
        if (!(fiscalYearStartMonth < 12)) {
            throw new IllegalArgumentException("fiscalYearStartMonth must be < 12");
        }
    this.internal.fiscalYearStartMonth = fiscalYearStartMonth;
        return this;
    }
    
    public Builder liveNow(Boolean liveNow) {
    this.internal.liveNow = liveNow;
        return this;
    }
    
    public Builder weekStart(String weekStart) {
    this.internal.weekStart = weekStart;
        return this;
    }
    
    public Builder refresh(String refresh) {
    this.internal.refresh = refresh;
        return this;
    }
    
    public Builder version(Integer version) {
    this.internal.version = version;
        return this;
    }
    
    public Builder withPanel(com.grafana.foundation.cog.Builder<Panel> panel) {
		if (this.internal.panels == null) {
			this.internal.panels = new LinkedList<>();
		}
    PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.panel = panel.build();

    if (panelOrRowPanel.panel.gridPos == null) {
        panelOrRowPanel.panel.gridPos = new GridPos();
    }
    if (panelOrRowPanel.panel.gridPos.x == null) {
        panelOrRowPanel.panel.gridPos.x = 0;
    }
    if (panelOrRowPanel.panel.gridPos.y == null) {
        panelOrRowPanel.panel.gridPos.y = 0;
    }
    if (panelOrRowPanel.panel.gridPos.w == null) {
        panelOrRowPanel.panel.gridPos.w = 0;
    }
    if (panelOrRowPanel.panel.gridPos.h == null) {
        panelOrRowPanel.panel.gridPos.h = 0;
    }
    // The panel either has no position set, or it is the first panel of the dashboard.
    // In that case, we position it on the grid
    if (panelOrRowPanel.panel.gridPos.x == 0 && panelOrRowPanel.panel.gridPos.y == 0) {
        panelOrRowPanel.panel.gridPos.x = this.currentX;
        panelOrRowPanel.panel.gridPos.y = this.currentY;
    }
    this.internal.panels.add(panelOrRowPanel);

	// Prepare the coordinates for the next panel
	this.currentX += panelOrRowPanel.panel.gridPos.w;
	this.lastPanelHeight = java.lang.Math.max(this.lastPanelHeight, panelOrRowPanel.panel.gridPos.h);

	// Check for grid width overflow?
	if (this.currentX >= 24) {
		this.currentX = 0;
		this.currentY += this.lastPanelHeight;
		this.lastPanelHeight = 0;
	}
        return this;
    }
    
    public Builder withRow(com.grafana.foundation.cog.Builder<RowPanel> rowPanel) {
		if (this.internal.panels == null) {
			this.internal.panels = new LinkedList<>();
		}
    PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        panelOrRowPanel.rowPanel = rowPanel.build();

    // Position the row on the grid
    if (panelOrRowPanel.rowPanel.gridPos == null || (panelOrRowPanel.rowPanel.gridPos.x == 0 && panelOrRowPanel.rowPanel.gridPos.y == 0)) {
        GridPos gridPos = new GridPos();
        gridPos.x = 0; // beginning of the line
        gridPos.y = this.currentY;
        gridPos.h = 1;
        gridPos.w = 24; // full width
        panelOrRowPanel.rowPanel.gridPos = gridPos;
    }
    this.internal.panels.add(panelOrRowPanel);

    // Reset the state for the next row
	this.currentX = 0;
	this.currentY += panelOrRowPanel.rowPanel.gridPos.h;
	this.lastPanelHeight = 0;
        return this;
    }
    
    public Builder variables(com.grafana.foundation.cog.Builder<List<VariableModel>> list) {
		if (this.internal.templating == null) {
			this.internal.templating = new com.grafana.foundation.dashboard.DashboardDashboardTemplating.Builder().build();
		}
    this.internal.templating.list = list.build();
        return this;
    }
    
    public Builder withVariable(com.grafana.foundation.cog.Builder<VariableModel> list) {
		if (this.internal.templating == null) {
			this.internal.templating = new com.grafana.foundation.dashboard.DashboardDashboardTemplating.Builder().build();
		}
		if (this.internal.templating.list == null) {
			this.internal.templating.list = new LinkedList<>();
		}
    this.internal.templating.list.add(list.build());
        return this;
    }
    
    public Builder annotations(com.grafana.foundation.cog.Builder<List<AnnotationQuery>> list) {
		if (this.internal.annotations == null) {
			this.internal.annotations = new com.grafana.foundation.dashboard.AnnotationContainer();
		}
    this.internal.annotations.list = list.build();
        return this;
    }
    
    public Builder annotation(com.grafana.foundation.cog.Builder<AnnotationQuery> list) {
		if (this.internal.annotations == null) {
			this.internal.annotations = new com.grafana.foundation.dashboard.AnnotationContainer();
		}
		if (this.internal.annotations.list == null) {
			this.internal.annotations.list = new LinkedList<>();
		}
    this.internal.annotations.list.add(list.build());
        return this;
    }
    
    public Builder links(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
    this.internal.links = links.build();
        return this;
    }
    
    public Builder link(com.grafana.foundation.cog.Builder<DashboardLink> links) {
		if (this.internal.links == null) {
			this.internal.links = new LinkedList<>();
		}
    this.internal.links.add(links.build());
        return this;
    }
    
    public Builder snapshot(com.grafana.foundation.cog.Builder<Snapshot> snapshot) {
    this.internal.snapshot = snapshot.build();
        return this;
    }
    public Dashboard build() {
            return this.internal;
        }
    }
}
