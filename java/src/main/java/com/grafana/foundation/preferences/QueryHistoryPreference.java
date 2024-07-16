// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class QueryHistoryPreference {
    // one of: '' | 'query' | 'starred'; 
    @JsonProperty("homeTab")
    public String homeTab;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryHistoryPreference> {
        private final QueryHistoryPreference internal;
        
        public Builder() {
            this.internal = new QueryHistoryPreference();
        }
    public Builder homeTab(String homeTab) {
    this.internal.homeTab = homeTab;
        return this;
    }
    public QueryHistoryPreference build() {
            return this.internal;
        }
    }
}
