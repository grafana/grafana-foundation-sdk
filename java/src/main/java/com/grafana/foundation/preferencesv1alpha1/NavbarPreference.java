// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferencesv1alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class NavbarPreference {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("bookmarkUrls")
    public List<String> bookmarkUrls;
    public NavbarPreference() {
        this.bookmarkUrls = new LinkedList<>();
    }
    public NavbarPreference(List<String> bookmarkUrls) {
        this.bookmarkUrls = bookmarkUrls;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
