// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.grafana.foundation.cog.variants.Dataquery;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = QueryDeserializer.class)
public class Query { 
    @JsonProperty("datasourceUid")
    public String datasourceUid; 
    @JsonProperty("model")
    public Dataquery model; 
    @JsonProperty("queryType")
    public String queryType; 
    @JsonProperty("refId")
    public String refId; 
    @JsonProperty("relativeTimeRange")
    public RelativeTimeRange relativeTimeRange;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Query> {
        private final Query internal;
        
        public Builder(String refId) {
            this.internal = new Query();
    this.internal.refId = refId;
        }
    public Builder datasourceUid(String datasourceUid) {
    this.internal.datasourceUid = datasourceUid;
        return this;
    }
    
    public Builder model(com.grafana.foundation.cog.Builder<Dataquery> model) {
    this.internal.model = model.build();
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder relativeTimeRange(RelativeTimeRange relativeTimeRange) {
    this.internal.relativeTimeRange = relativeTimeRange;
        return this;
    }
    public Query build() {
            return this.internal;
        }
    }
}
