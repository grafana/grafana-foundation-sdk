"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.PanelBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class PanelBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultPanelKind();
        this.internal.kind = "Panel";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    id(id) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.id = id;
        return this;
    }
    title(title) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.title = title;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
    links(links) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.spec.links = linksResources;
        return this;
    }
    data(data) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        const dataResource = data.build();
        this.internal.spec.data = dataResource;
        return this;
    }
    visualization(vizConfig) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        const vizConfigResource = vizConfig.build();
        this.internal.spec.vizConfig = vizConfigResource;
        return this;
    }
    transparent(transparent) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.transparent = transparent;
        return this;
    }
}
exports.PanelBuilder = PanelBuilder;
//# sourceMappingURL=panelBuilder.gen.js.map