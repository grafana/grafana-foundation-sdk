// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.grafana.foundation.common.MapLayerOptions;
import java.util.List;
import java.util.LinkedList;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder view(com.grafana.foundation.cog.Builder<MapViewConfig> view) {
    MapViewConfig viewResource = view.build();
        this.internal.view = viewResource;
        return this;
    }
    
    public OptionsBuilder controls(com.grafana.foundation.cog.Builder<ControlsOptions> controls) {
    ControlsOptions controlsResource = controls.build();
        this.internal.controls = controlsResource;
        return this;
    }
    
    public OptionsBuilder basemap(com.grafana.foundation.cog.Builder<MapLayerOptions> basemap) {
    MapLayerOptions basemapResource = basemap.build();
        this.internal.basemap = basemapResource;
        return this;
    }
    
    public OptionsBuilder layers(List<com.grafana.foundation.cog.Builder<MapLayerOptions>> layers) {
        List<MapLayerOptions> layersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<MapLayerOptions> r1 : layers) {
                MapLayerOptions layersDepth1 = r1.build();
                layersResources.add(layersDepth1); 
        }
        this.internal.layers = layersResources;
        return this;
    }
    
    public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<TooltipOptions> tooltip) {
    TooltipOptions tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
