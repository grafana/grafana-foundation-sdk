// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

public class Dashboard {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("annotations")
    public List<AnnotationQueryKind> annotations;
    // Configuration of dashboard cursor sync behavior.
    // "Off" for no shared crosshair or tooltip (default).
    // "Crosshair" for shared crosshair.
    // "Tooltip" for shared crosshair AND shared tooltip.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("cursorSync")
    public DashboardCursorSync cursorSync;
    // Description of dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // Whether a dashboard is editable or not.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("editable")
    public Boolean editable;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("elements")
    public Map<String, Element> elements;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("layout")
    public GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind layout;
    // Links with references to other dashboards or external websites.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("links")
    public List<DashboardLink> links;
    // When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    // This will keep data "moving left" regardless of the query refresh rate. This setting helps
    // avoid dashboards presenting stale live data.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("liveNow")
    public Boolean liveNow;
    // When set to true, the dashboard will load all panels in the dashboard when it's loaded.
    @JsonProperty("preload")
    public Boolean preload;
    // Plugins only. The version of the dashboard installed together with the plugin.
    // This is used to determine if the dashboard should be updated when the plugin is updated.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("revision")
    public Short revision;
    // Tags associated with dashboard.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("tags")
    public List<String> tags;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("timeSettings")
    public TimeSettingsSpec timeSettings;
    // Title of dashboard.
    @JsonProperty("title")
    public String title;
    // Configured template variables.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("variables")
    public List<VariableKind> variables;
    public Dashboard() {
        this.annotations = new LinkedList<>();
        this.cursorSync = DashboardCursorSync.OFF;
        this.editable = true;
        this.elements = new HashMap<>();
        this.layout = new com.grafana.foundation.dashboardv2beta1.GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        this.links = new LinkedList<>();
        this.preload = false;
        this.tags = new LinkedList<>();
        this.timeSettings = new com.grafana.foundation.dashboardv2beta1.TimeSettingsSpec();
        this.title = "";
        this.variables = new LinkedList<>();
    }
    public Dashboard(List<AnnotationQueryKind> annotations,DashboardCursorSync cursorSync,String description,Boolean editable,Map<String, Element> elements,GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind layout,List<DashboardLink> links,Boolean liveNow,Boolean preload,Short revision,List<String> tags,TimeSettingsSpec timeSettings,String title,List<VariableKind> variables) {
        this.annotations = annotations;
        this.cursorSync = cursorSync;
        this.description = description;
        this.editable = editable;
        this.elements = elements;
        this.layout = layout;
        this.links = links;
        this.liveNow = liveNow;
        this.preload = preload;
        this.revision = revision;
        this.tags = tags;
        this.timeSettings = timeSettings;
        this.title = title;
        this.variables = variables;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
