// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import java.util.Map;

/**
 * #SearchResults is the response envelope, mirroring
 * search.grafana.app SearchResults.
 */
public class SearchResults {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("metadata")
    public SearchResultsMetadata metadata;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("items")
    public List<SearchResultHit> items;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("facets")
    public Map<String, List<FacetValue>> facets;
    public SearchResults() {
        this.metadata = new com.grafana.foundation.alertingv0alpha1.SearchResultsMetadata();
        this.items = new LinkedList<>();
    }
    public SearchResults(SearchResultsMetadata metadata,List<SearchResultHit> items,Map<String, List<FacetValue>> facets) {
        this.metadata = metadata;
        this.items = items;
        this.facets = facets;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
