"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HideSeriesConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class HideSeriesConfigBuilder {
    constructor() {
        this.internal = common.defaultHideSeriesConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    tooltip(tooltip) {
        this.internal.tooltip = tooltip;
        return this;
    }
    legend(legend) {
        this.internal.legend = legend;
        return this;
    }
    viz(viz) {
        this.internal.viz = viz;
        return this;
    }
}
exports.HideSeriesConfigBuilder = HideSeriesConfigBuilder;
//# sourceMappingURL=hideSeriesConfigBuilder.gen.js.map