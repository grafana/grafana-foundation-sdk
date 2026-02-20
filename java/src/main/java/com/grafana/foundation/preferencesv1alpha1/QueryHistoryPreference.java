// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferencesv1alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryHistoryPreference {
    // one of: '' | 'query' | 'starred';
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("homeTab")
    public String homeTab;
    public QueryHistoryPreference() {
    }
    public QueryHistoryPreference(String homeTab) {
        this.homeTab = homeTab;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
