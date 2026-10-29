// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * #SearchResultsMetadata carries the paging token and total authorised match
 * count.
 */
public class SearchResultsMetadata {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("continue")
    public String continueArg;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("totalHits")
    public Long totalHits;
    // Always read totalHits together with totalHitsRelation.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("totalHitsRelation")
    public TotalHitsRelation totalHitsRelation;
    public SearchResultsMetadata() {
    }
    public SearchResultsMetadata(String continueArg,Long totalHits,TotalHitsRelation totalHitsRelation) {
        this.continueArg = continueArg;
        this.totalHits = totalHits;
        this.totalHitsRelation = totalHitsRelation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
