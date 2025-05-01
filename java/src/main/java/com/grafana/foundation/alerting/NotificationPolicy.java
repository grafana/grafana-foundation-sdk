// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
public class NotificationPolicy {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("continue")
    public Boolean continueArg;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_by")
    public List<String> groupBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_interval")
    public String groupInterval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_wait")
    public String groupWait;
    // Deprecated. Remove before v1.0 release.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("match")
    public Map<String, String> match;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("match_re")
    public Map<String, String> matchRe;
    // Matchers is a slice of Matchers that is sortable, implements Stringer, and
    // provides a Matches method to match a LabelSet against all Matchers in the
    // slice. Note that some users of Matchers might require it to be sorted.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("matchers")
    public List<Matcher> matchers;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mute_time_intervals")
    public List<String> muteTimeIntervals;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("object_matchers")
    public List<List<String>> objectMatchers;
    @JsonProperty("provenance")
    public String provenance;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("receiver")
    public String receiver;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat_interval")
    public String repeatInterval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("routes")
    public List<NotificationPolicy> routes;
    public NotificationPolicy() {
    }
    public NotificationPolicy(Boolean continueArg,List<String> groupBy,String groupInterval,String groupWait,Map<String, String> match,Map<String, String> matchRe,List<Matcher> matchers,List<String> muteTimeIntervals,List<List<String>> objectMatchers,String provenance,String receiver,String repeatInterval,List<NotificationPolicy> routes) {
        this.continueArg = continueArg;
        this.groupBy = groupBy;
        this.groupInterval = groupInterval;
        this.groupWait = groupWait;
        this.match = match;
        this.matchRe = matchRe;
        this.matchers = matchers;
        this.muteTimeIntervals = muteTimeIntervals;
        this.objectMatchers = objectMatchers;
        this.provenance = provenance;
        this.receiver = receiver;
        this.repeatInterval = repeatInterval;
        this.routes = routes;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
