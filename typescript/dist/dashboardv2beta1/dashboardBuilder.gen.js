"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DashboardBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class DashboardBuilder {
    constructor(title) {
        this.internal = dashboardv2beta1.defaultDashboard();
        this.internal.title = title;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    annotations(annotations) {
        const annotationsResources = annotations.map(builder1 => builder1.build());
        this.internal.annotations = annotationsResources;
        return this;
    }
    // Configuration of dashboard cursor sync behavior.
    // "Off" for no shared crosshair or tooltip (default).
    // "Crosshair" for shared crosshair.
    // "Tooltip" for shared crosshair AND shared tooltip.
    cursorSync(cursorSync) {
        this.internal.cursorSync = cursorSync;
        return this;
    }
    // Description of dashboard.
    description(description) {
        this.internal.description = description;
        return this;
    }
    // Whether a dashboard is editable or not.
    editable(editable) {
        this.internal.editable = editable;
        return this;
    }
    elements(elements) {
        const elementsResource = (function () {
            let results1 = {};
            for (const key1 in elements) {
                const val1 = elements[key1];
                results1[key1] = val1.build();
            }
            return results1;
        }());
        this.internal.elements = elementsResource;
        return this;
    }
    element(key, element) {
        if (!this.internal.elements) {
            this.internal.elements = {};
        }
        const elementResource = element.build();
        this.internal.elements[key] = elementResource;
        return this;
    }
    gridLayout(gridLayoutKind) {
        const gridLayoutKindResource = gridLayoutKind.build();
        this.internal.layout = gridLayoutKindResource;
        return this;
    }
    rowsLayout(rowsLayoutKind) {
        const rowsLayoutKindResource = rowsLayoutKind.build();
        this.internal.layout = rowsLayoutKindResource;
        return this;
    }
    autoGridLayout(autoGridLayoutKind) {
        const autoGridLayoutKindResource = autoGridLayoutKind.build();
        this.internal.layout = autoGridLayoutKindResource;
        return this;
    }
    tabsLayout(tabsLayoutKind) {
        const tabsLayoutKindResource = tabsLayoutKind.build();
        this.internal.layout = tabsLayoutKindResource;
        return this;
    }
    // Links with references to other dashboards or external websites.
    links(links) {
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.links = linksResources;
        return this;
    }
    // When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    // This will keep data "moving left" regardless of the query refresh rate. This setting helps
    // avoid dashboards presenting stale live data.
    liveNow(liveNow) {
        this.internal.liveNow = liveNow;
        return this;
    }
    // When set to true, the dashboard will load all panels in the dashboard when it's loaded.
    preload(preload) {
        this.internal.preload = preload;
        return this;
    }
    // Plugins only. The version of the dashboard installed together with the plugin.
    // This is used to determine if the dashboard should be updated when the plugin is updated.
    revision(revision) {
        this.internal.revision = revision;
        return this;
    }
    // Tags associated with dashboard.
    tags(tags) {
        this.internal.tags = tags;
        return this;
    }
    timeSettings(timeSettings) {
        const timeSettingsResource = timeSettings.build();
        this.internal.timeSettings = timeSettingsResource;
        return this;
    }
    // Title of dashboard.
    title(title) {
        this.internal.title = title;
        return this;
    }
    // Configured template variables.
    variables(variables) {
        const variablesResources = variables.map(builder1 => builder1.build());
        this.internal.variables = variablesResources;
        return this;
    }
    // Configured template variables.
    variable(variable) {
        if (!this.internal.variables) {
            this.internal.variables = [];
        }
        const variableResource = variable.build();
        this.internal.variables.push(variableResource);
        return this;
    }
}
exports.DashboardBuilder = DashboardBuilder;
//# sourceMappingURL=dashboardBuilder.gen.js.map