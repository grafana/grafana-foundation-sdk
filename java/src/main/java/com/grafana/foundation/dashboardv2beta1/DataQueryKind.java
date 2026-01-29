// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = DataQueryKindDeserializer.class)
public class DataQueryKind {
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("group")
    public String group;
    @JsonProperty("version")
    public String version;
    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public Dashboardv2beta1DataQueryKindDatasource datasource;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spec")
    public Object spec;
    public DataQueryKind() {
        this.kind = "";
        this.group = "";
        this.version = "v0";
    }
    public DataQueryKind(String kind,String group,String version,Dashboardv2beta1DataQueryKindDatasource datasource,Object spec) {
        this.kind = kind;
        this.group = group;
        this.version = version;
        this.datasource = datasource;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
