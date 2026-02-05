"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.PanelModelBuilder = void 0;
const tslib_1 = require("tslib");
const librarypanel = tslib_1.__importStar(require("../librarypanel"));
// Dashboard panels are the basic visualization building blocks.
class PanelModelBuilder {
    constructor() {
        this.internal = librarypanel.defaultPanelModel();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // The panel plugin type id. This is used to find the plugin to display the panel.
    type(type) {
        if (!(type.length >= 1)) {
            throw new Error("type.length must be >= 1");
        }
        this.internal.type = type;
        return this;
    }
    // The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
    pluginVersion(pluginVersion) {
        this.internal.pluginVersion = pluginVersion;
        return this;
    }
    // Depends on the panel plugin. See the plugin documentation for details.
    targets(targets) {
        const targetsResources = targets.map(builder1 => builder1.build());
        this.internal.targets = targetsResources;
        return this;
    }
    // Panel title.
    title(title) {
        this.internal.title = title;
        return this;
    }
    // Panel description.
    description(description) {
        this.internal.description = description;
        return this;
    }
    // Whether to display the panel without a background.
    transparent(transparent) {
        this.internal.transparent = transparent;
        return this;
    }
    // The datasource used in all targets.
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    // Panel links.
    links(links) {
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.links = linksResources;
        return this;
    }
    // Name of template variable to repeat for.
    repeat(repeat) {
        this.internal.repeat = repeat;
        return this;
    }
    // Direction to repeat in if 'repeat' is set.
    // `h` for horizontal, `v` for vertical.
    repeatDirection(repeatDirection) {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }
    // Option for repeated panels that controls max items per row
    // Only relevant for horizontally repeated panels
    maxPerRow(maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
    // The maximum number of data points that the panel queries are retrieving.
    maxDataPoints(maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    // List of transformations that are applied to the panel data before rendering.
    // When there are multiple transformations, Grafana applies them in the order they are listed.
    // Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    transformations(transformations) {
        this.internal.transformations = transformations;
        return this;
    }
    // The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
    // This value must be formatted as a number followed by a valid time
    // identifier like: "40s", "3d", etc.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    interval(interval) {
        this.internal.interval = interval;
        return this;
    }
    // Overrides the relative time range for individual panels,
    // which causes them to be different than what is selected in
    // the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
    // time periods or days on the same dashboard.
    // The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
    // `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    timeFrom(timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    // Overrides the time range for individual panels by shifting its start and end relative to the time picker.
    // For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    timeShift(timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    // Controls if the timeFrom or timeShift overrides are shown in the panel header
    hideTimeOverride(hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    // Sets panel queries cache timeout.
    cacheTimeout(cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }
    // Overrides the data source configured time-to-live for a query cache item in milliseconds
    queryCachingTTL(queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }
    // It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
    options(options) {
        this.internal.options = options;
        return this;
    }
    // Field options allow you to change how the data is displayed in your visualizations.
    fieldConfig(fieldConfig) {
        this.internal.fieldConfig = fieldConfig;
        return this;
    }
}
exports.PanelModelBuilder = PanelModelBuilder;
//# sourceMappingURL=panelModelBuilder.gen.js.map