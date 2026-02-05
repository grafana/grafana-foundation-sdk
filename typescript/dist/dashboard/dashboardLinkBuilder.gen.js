"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DashboardLinkBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// Links with references to other dashboards or external resources
class DashboardLinkBuilder {
    constructor(title) {
        this.internal = dashboard.defaultDashboardLink();
        this.internal.title = title;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Title to display with the link
    title(title) {
        this.internal.title = title;
        return this;
    }
    // Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    type(type) {
        this.internal.type = type;
        return this;
    }
    // Icon name to be displayed with the link
    icon(icon) {
        this.internal.icon = icon;
        return this;
    }
    // Tooltip to display when the user hovers their mouse over it
    tooltip(tooltip) {
        this.internal.tooltip = tooltip;
        return this;
    }
    // Link URL. Only required/valid if the type is link
    url(url) {
        this.internal.url = url;
        return this;
    }
    // List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
    tags(tags) {
        this.internal.tags = tags;
        return this;
    }
    // If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
    asDropdown(asDropdown) {
        this.internal.asDropdown = asDropdown;
        return this;
    }
    // If true, the link will be opened in a new tab
    targetBlank(targetBlank) {
        this.internal.targetBlank = targetBlank;
        return this;
    }
    // If true, includes current template variables values in the link as query params
    includeVars(includeVars) {
        this.internal.includeVars = includeVars;
        return this;
    }
    // If true, includes current time range in the link as query params
    keepTime(keepTime) {
        this.internal.keepTime = keepTime;
        return this;
    }
}
exports.DashboardLinkBuilder = DashboardLinkBuilder;
//# sourceMappingURL=dashboardLinkBuilder.gen.js.map