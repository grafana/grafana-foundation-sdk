// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class V2PanelSpec {
    @JsonProperty("id")
    public Double id;
    @JsonProperty("title")
    public String title;
    // Shown in a info icon tooltip next to panel title
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // Shown in a sub header below the title.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("subtitle")
    public String subtitle;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("links")
    public List<DataLink> links;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("data")
    public V2QueryGroupKind data;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("vizConfig")
    public VizConfigKind vizConfig;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("transparent")
    public Boolean transparent;
    public V2PanelSpec() {
        this.id = 0.0;
        this.title = "";
        this.links = new LinkedList<>();
        this.data = new com.grafana.foundation.dashboardv2beta1.V2QueryGroupKind();
        this.vizConfig = new com.grafana.foundation.dashboardv2beta1.VizConfigKind();
    }
    public V2PanelSpec(Double id,String title,String description,String subtitle,List<DataLink> links,V2QueryGroupKind data,VizConfigKind vizConfig,Boolean transparent) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.subtitle = subtitle;
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
