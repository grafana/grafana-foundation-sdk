// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.starsv1alpha1;

import java.util.List;
import java.util.LinkedList;

public class StarsBuilder implements com.grafana.foundation.cog.Builder<Stars> {
    protected final Stars internal;
    
    public StarsBuilder() {
        this.internal = new Stars();
    }
    public StarsBuilder resources(List<com.grafana.foundation.cog.Builder<Resource>> resource) {
        List<Resource> resourceResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Resource> r1 : resource) {
                Resource resourceDepth1 = r1.build();
                resourceResources.add(resourceDepth1); 
        }
        this.internal.resource = resourceResources;
        return this;
    }
    
    public StarsBuilder resource(com.grafana.foundation.cog.Builder<Resource> resource) {
		if (this.internal.resource == null) {
			this.internal.resource = new LinkedList<>();
		}
    Resource resourceResource = resource.build();
        this.internal.resource.add(resourceResource);
        return this;
    }
    public Stars build() {
        return this.internal;
    }
}
