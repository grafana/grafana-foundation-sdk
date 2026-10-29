// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * #SearchResultHit is a single match: its identity, an optional relevance
 * score (present only when the query included free text), and the requested
 * fields.
 */
public class SearchResultHit {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("resource")
    public SearchResultResource resource;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("score")
    public Double score;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("fields")
    public RuleSearchHitFields fields;
    public SearchResultHit() {
        this.resource = new com.grafana.foundation.alertingv0alpha1.SearchResultResource();
        this.fields = new com.grafana.foundation.alertingv0alpha1.RuleSearchHitFields();
    }
    public SearchResultHit(SearchResultResource resource,Double score,RuleSearchHitFields fields) {
        this.resource = resource;
        this.score = score;
        this.fields = fields;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
