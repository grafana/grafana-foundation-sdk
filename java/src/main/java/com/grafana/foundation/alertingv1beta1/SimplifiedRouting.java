// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class SimplifiedRouting {
    @JsonProperty("type")
    public NotificationSettingsType type;
    @JsonProperty("receiver")
    public String receiver;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupBy")
    public List<String> groupBy;
    @JsonProperty("groupWait")
    public String groupWait;
    @JsonProperty("groupInterval")
    public String groupInterval;
    @JsonProperty("repeatInterval")
    public String repeatInterval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("muteTimeIntervals")
    public List<String> muteTimeIntervals;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("activeTimeIntervals")
    public List<String> activeTimeIntervals;
    public SimplifiedRouting() {
        this.type = NotificationSettingsType.SIMPLIFIED_ROUTING;
        this.receiver = "";
    }
    public SimplifiedRouting(String receiver,List<String> groupBy,String groupWait,String groupInterval,String repeatInterval,List<String> muteTimeIntervals,List<String> activeTimeIntervals) {
        this.type = NotificationSettingsType.SIMPLIFIED_ROUTING;
        this.receiver = receiver;
        this.groupBy = groupBy;
        this.groupWait = groupWait;
        this.groupInterval = groupInterval;
        this.repeatInterval = repeatInterval;
        this.muteTimeIntervals = muteTimeIntervals;
        this.activeTimeIntervals = activeTimeIntervals;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
