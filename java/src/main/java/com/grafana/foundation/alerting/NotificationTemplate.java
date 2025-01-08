// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class NotificationTemplate {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provenance")
    public String provenance;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("template")
    public String template;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("version")
    public String version;
    public NotificationTemplate() {
    }
    
    public NotificationTemplate(String name,String provenance,String template,String version) {
        this.name = name;
        this.provenance = provenance;
        this.template = template;
        this.version = version;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
