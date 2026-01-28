// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RowsLayoutRowSpec {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("collapse")
    public Boolean collapse;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideHeader")
    public Boolean hideHeader;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillScreen")
    public Boolean fillScreen;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("conditionalRendering")
    public ConditionalRenderingGroupKind conditionalRendering;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat")
    public RowRepeatOptions repeat;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("layout")
    public GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind layout;
    public RowsLayoutRowSpec() {
        this.layout = new com.grafana.foundation.dashboardv2beta1.GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
    }
    public RowsLayoutRowSpec(String title,Boolean collapse,Boolean hideHeader,Boolean fillScreen,ConditionalRenderingGroupKind conditionalRendering,RowRepeatOptions repeat,GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind layout) {
        this.title = title;
        this.collapse = collapse;
        this.hideHeader = hideHeader;
        this.fillScreen = fillScreen;
        this.conditionalRendering = conditionalRendering;
        this.repeat = repeat;
        this.layout = layout;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
