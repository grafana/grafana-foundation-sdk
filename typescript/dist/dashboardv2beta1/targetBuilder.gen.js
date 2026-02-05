"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TargetBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class TargetBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultPanelQueryKind();
        this.internal.kind = "PanelQuery";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelQuerySpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }
    refId(refId) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelQuerySpec();
        }
        this.internal.spec.refId = refId;
        return this;
    }
    hidden(hidden) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelQuerySpec();
        }
        this.internal.spec.hidden = hidden;
        return this;
    }
}
exports.TargetBuilder = TargetBuilder;
//# sourceMappingURL=targetBuilder.gen.js.map