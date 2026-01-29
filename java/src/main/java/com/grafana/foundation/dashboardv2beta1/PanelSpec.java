// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

public class PanelSpec {
    @JsonProperty("id")
    public Double id;
    @JsonProperty("title")
    public String title;
    @JsonProperty("description")
    public String description;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("links")
    public List<DataLink> links;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("data")
    public QueryGroupKind data;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("vizConfig")
    public VizConfigKind vizConfig;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("transparent")
    public Boolean transparent;
    public PanelSpec() {
        this.id = 0.0;
        this.title = "";
        this.description = "";
        this.links = new LinkedList<>();
        this.data = new com.grafana.foundation.dashboardv2beta1.QueryGroupKind();
        this.vizConfig = new com.grafana.foundation.dashboardv2beta1.VizConfigKind();
    }
    public PanelSpec(Double id,String title,String description,List<DataLink> links,QueryGroupKind data,VizConfigKind vizConfig,Boolean transparent) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.links = links;
        this.data = data;
        this.vizConfig = vizConfig;
        this.transparent = transparent;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
