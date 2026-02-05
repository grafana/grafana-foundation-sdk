"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Query variable kind
class QueryVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultQueryVariableKind();
        this.internal.kind = "QueryVariable";
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
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    current(current) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    refresh(refresh) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }
    regex(regex) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.regex = regex;
        return this;
    }
    sort(sort) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.sort = sort;
        return this;
    }
    definition(definition) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.definition = definition;
        return this;
    }
    options(options) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }
    multi(multi) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }
    includeAll(includeAll) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }
    allValue(allValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }
    placeholder(placeholder) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.placeholder = placeholder;
        return this;
    }
    allowCustomValue(allowCustomValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
    staticOptions(staticOptions) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.staticOptions = staticOptions;
        return this;
    }
    staticOptionsOrder(staticOptionsOrder) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.staticOptionsOrder = staticOptionsOrder;
        return this;
    }
}
exports.QueryVariableBuilder = QueryVariableBuilder;
//# sourceMappingURL=queryVariableBuilder.gen.js.map