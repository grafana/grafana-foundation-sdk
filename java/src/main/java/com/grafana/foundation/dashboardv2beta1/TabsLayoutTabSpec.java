// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class TabsLayoutTabSpec {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("layout")
    public GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind layout;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("conditionalRendering")
    public ConditionalRenderingGroupKind conditionalRendering;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat")
    public TabRepeatOptions repeat;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("variables")
    public List<VariableKind> variables;
    public TabsLayoutTabSpec() {
        this.layout = new com.grafana.foundation.dashboardv2beta1.GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
    }
    public TabsLayoutTabSpec(String title,GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind layout,ConditionalRenderingGroupKind conditionalRendering,TabRepeatOptions repeat,List<VariableKind> variables) {
        this.title = title;
        this.layout = layout;
        this.conditionalRendering = conditionalRendering;
        this.repeat = repeat;
        this.variables = variables;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
