// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.DataSourceRef;
import com.grafana.foundation.cog.variants.Dataquery;

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
@JsonDeserialize(using = AnnotationQueryDeserializer.class)
public class AnnotationQuery {
    // Name of annotation.
    @JsonProperty("name")
    public String name;
    // Datasource where the annotations data is
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // When enabled the annotation query is issued with every dashboard refresh
    @JsonProperty("enable")
    public Boolean enable;
    // Annotation queries can be toggled on or off at the top of the dashboard.
    // When hide is true, the toggle is not shown in the dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Color to use for the annotation event markers
    @JsonProperty("iconColor")
    public String iconColor;
    // Filters to apply when fetching annotations
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filter")
    public AnnotationPanelFilter filter;
    // TODO.. this should just be a normal query target
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("target")
    public Dataquery target;
    // TODO -- this should not exist here, it is based on the --grafana-- datasource
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public String type;
    // Set to 1 for the standard annotation query all dashboards have by default.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("builtIn")
    public Double builtIn;
    // Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("placement")
    public AnnotationQueryPlacement placement;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("expr")
    public String expr;
    public AnnotationQuery() {
        this.name = "";
        this.enable = true;
        this.hide = false;
        this.iconColor = "";
        this.builtIn = 0.0;
        this.placement = AnnotationQueryPlacement.IN_CONTROLS_MENU;
    }
    public AnnotationQuery(String name,DataSourceRef datasource,Boolean enable,Boolean hide,String iconColor,AnnotationPanelFilter filter,Dataquery target,String type,Double builtIn,String expr) {
        this.name = name;
        this.datasource = datasource;
        this.enable = enable;
        this.hide = hide;
        this.iconColor = iconColor;
        this.filter = filter;
        this.target = target;
        this.type = type;
        this.builtIn = builtIn;
        this.placement = AnnotationQueryPlacement.IN_CONTROLS_MENU;
        this.expr = expr;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
