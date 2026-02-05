"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.NestedBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class NestedBuilder {
    constructor() {
        this.internal = elasticsearch.defaultNested();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    field(field) {
        this.internal.field = field;
        return this;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    settings(settings) {
        this.internal.settings = settings;
        return this;
    }
}
exports.NestedBuilder = NestedBuilder;
//# sourceMappingURL=nestedBuilder.gen.js.map