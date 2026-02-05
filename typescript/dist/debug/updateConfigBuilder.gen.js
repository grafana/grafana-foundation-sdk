"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.UpdateConfigBuilder = void 0;
const tslib_1 = require("tslib");
const debug = tslib_1.__importStar(require("../debug"));
class UpdateConfigBuilder {
    constructor() {
        this.internal = debug.defaultUpdateConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    render(render) {
        this.internal.render = render;
        return this;
    }
    dataChanged(dataChanged) {
        this.internal.dataChanged = dataChanged;
        return this;
    }
    schemaChanged(schemaChanged) {
        this.internal.schemaChanged = schemaChanged;
        return this;
    }
}
exports.UpdateConfigBuilder = UpdateConfigBuilder;
//# sourceMappingURL=updateConfigBuilder.gen.js.map