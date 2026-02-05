"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataLinkBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class DataLinkBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataLink();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    title(title) {
        this.internal.title = title;
        return this;
    }
    url(url) {
        this.internal.url = url;
        return this;
    }
    targetBlank(targetBlank) {
        this.internal.targetBlank = targetBlank;
        return this;
    }
}
exports.DataLinkBuilder = DataLinkBuilder;
//# sourceMappingURL=dataLinkBuilder.gen.js.map