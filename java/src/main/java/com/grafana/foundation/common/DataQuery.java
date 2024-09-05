// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
public class DataQuery {
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("datasource")
    public Object datasource;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DataQuery> {
        private final DataQuery internal;
        
        public Builder() {
            this.internal = new DataQuery();
        }
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder datasource(Object datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    public DataQuery build() {
            return this.internal;
        }
    }
}
