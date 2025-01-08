// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.athena;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ConnectionArgs {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("region")
    public String region;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("catalog")
    public String catalog;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("database")
    public String database;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultReuseEnabled")
    public Boolean resultReuseEnabled;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultReuseMaxAgeInMinutes")
    public Double resultReuseMaxAgeInMinutes;
    public ConnectionArgs() {
        this.region = "__default";
        this.catalog = "__default";
        this.database = "__default";
        this.resultReuseEnabled = false;
        this.resultReuseMaxAgeInMinutes = 60.0;
    }
    
    public ConnectionArgs(String region,String catalog,String database,Boolean resultReuseEnabled,Double resultReuseMaxAgeInMinutes) {
        this.region = region;
        this.catalog = catalog;
        this.database = database;
        this.resultReuseEnabled = resultReuseEnabled;
        this.resultReuseMaxAgeInMinutes = resultReuseMaxAgeInMinutes;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
