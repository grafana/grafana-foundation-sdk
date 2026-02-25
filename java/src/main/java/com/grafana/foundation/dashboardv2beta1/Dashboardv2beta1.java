// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.grafana.foundation.resource.ManifestBuilder;
import com.grafana.foundation.resource.Resource;
public class Dashboardv2beta1 {    
    /**
     * Creates a resource manifest from a Dashboard.
     */
    public static ManifestBuilder manifest(String name,com.grafana.foundation.cog.Builder<Dashboard> dashboard) {
        ManifestBuilder builder = new ManifestBuilder();
        builder.apiVersion("dashboard.grafana.app/v2beta1");
        builder.kind("Dashboard");
        builder.metadata(Resource.named(name));
        builder.spec(dashboard.build());
        return builder;
    }
    
    public static RowsBuilder rows() {
        RowsBuilder builder = new RowsBuilder();
        return builder;
    }
    
    public static RowBuilder row(String title) {
        RowBuilder builder = new RowBuilder();
        builder.title(title);
        return builder;
    }
    
    public static AutoGridBuilder autoGrid() {
        AutoGridBuilder builder = new AutoGridBuilder();
        return builder;
    }
    
    public static AutoGridItemBuilder autoGridItem(String name) {
        AutoGridItemBuilder builder = new AutoGridItemBuilder();
        builder.name(name);
        return builder;
    }
    
    public static TabsBuilder tabs() {
        TabsBuilder builder = new TabsBuilder();
        return builder;
    }
    
    public static TabBuilder tab(String title) {
        TabBuilder builder = new TabBuilder();
        builder.title(title);
        return builder;
    }
    
    public static GridBuilder grid() {
        GridBuilder builder = new GridBuilder();
        return builder;
    }
    
    public static GridItemBuilder gridItem(String name) {
        GridItemBuilder builder = new GridItemBuilder();
        builder.name(name);
        return builder;
    }
}
