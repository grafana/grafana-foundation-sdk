// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import java.util.List;

public class NavbarPreferenceBuilder implements com.grafana.foundation.cog.Builder<NavbarPreference> {
    protected final NavbarPreference internal;
    
    public NavbarPreferenceBuilder() {
        this.internal = new NavbarPreference();
    }
    public NavbarPreferenceBuilder bookmarkUrls(List<String> bookmarkUrls) {
        this.internal.bookmarkUrls = bookmarkUrls;
        return this;
    }
    public NavbarPreference build() {
        return this.internal;
    }
}
