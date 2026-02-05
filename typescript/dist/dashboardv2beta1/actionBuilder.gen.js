"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ActionBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ActionBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultAction();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    type(type) {
        this.internal.type = type;
        return this;
    }
    title(title) {
        this.internal.title = title;
        return this;
    }
    fetch(fetch) {
        const fetchResource = fetch.build();
        this.internal.fetch = fetchResource;
        return this;
    }
    infinity(infinity) {
        const infinityResource = infinity.build();
        this.internal.infinity = infinityResource;
        return this;
    }
    confirmation(confirmation) {
        this.internal.confirmation = confirmation;
        return this;
    }
    oneClick(oneClick) {
        this.internal.oneClick = oneClick;
        return this;
    }
    variables(variables) {
        const variablesResources = variables.map(builder1 => builder1.build());
        this.internal.variables = variablesResources;
        return this;
    }
    style(style) {
        this.internal.style = style;
        return this;
    }
}
exports.ActionBuilder = ActionBuilder;
//# sourceMappingURL=actionBuilder.gen.js.map