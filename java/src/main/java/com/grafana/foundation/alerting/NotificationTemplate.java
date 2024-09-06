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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NotificationTemplate> {
        private final NotificationTemplate internal;
        
        public Builder() {
            this.internal = new NotificationTemplate();
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder provenance(String provenance) {
    this.internal.provenance = provenance;
        return this;
    }
    
    public Builder template(String template) {
    this.internal.template = template;
        return this;
    }
    
    public Builder version(String version) {
    this.internal.version = version;
        return this;
    }
    public NotificationTemplate build() {
            return this.internal;
        }
    }
}
