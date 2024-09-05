// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Ref to a DataSource instance
public class DataSourceRef {
    // The plugin type-id
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("type")
    public String type;
    // Specific datasource instance
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
