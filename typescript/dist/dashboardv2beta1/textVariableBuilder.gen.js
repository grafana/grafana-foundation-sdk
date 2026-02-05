"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TextVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Text variable kind
class TextVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultTextVariableKind();
        this.internal.kind = "TextVariable";
        this.internal.spec.name = name;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    spec(spec) {
        this.internal.spec = spec;
        return this;
    }
    name(name) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    current(current) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}
exports.TextVariableBuilder = TextVariableBuilder;
//# sourceMappingURL=textVariableBuilder.gen.js.map