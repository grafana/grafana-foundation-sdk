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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ConnectionArgs> {
        protected final ConnectionArgs internal;
        
        public Builder() {
            this.internal = new ConnectionArgs();
        this.region("__default");
        this.catalog("__default");
        this.database("__default");
        this.resultReuseEnabled(false);
        this.resultReuseMaxAgeInMinutes(60.0);
        }
    public Builder region(String region) {
    this.internal.region = region;
        return this;
    }
    
    public Builder catalog(String catalog) {
    this.internal.catalog = catalog;
        return this;
    }
    
    public Builder database(String database) {
    this.internal.database = database;
        return this;
    }
    
    public Builder resultReuseEnabled(Boolean resultReuseEnabled) {
    this.internal.resultReuseEnabled = resultReuseEnabled;
        return this;
    }
    
    public Builder resultReuseMaxAgeInMinutes(Double resultReuseMaxAgeInMinutes) {
    this.internal.resultReuseMaxAgeInMinutes = resultReuseMaxAgeInMinutes;
        return this;
    }
    public ConnectionArgs build() {
            return this.internal;
        }
    }
}
