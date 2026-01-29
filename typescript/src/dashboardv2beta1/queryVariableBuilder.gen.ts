// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Query variable kind
export class QueryVariableBuilder implements cog.Builder<dashboardv2beta1.QueryVariableKind> {
    protected readonly internal: dashboardv2beta1.QueryVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultQueryVariableKind();
        this.internal.kind = "QueryVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.QueryVariableKind {
        return this.internal;
    }

    spec(spec: dashboardv2beta1.QueryVariableSpec): this {
        this.internal.spec = spec;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    current(current: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    refresh(refresh: dashboardv2beta1.VariableRefresh): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    query(query: cog.Builder<dashboardv2beta1.DataQueryKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }

    regex(regex: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.regex = regex;
        return this;
    }

    regexApplyTo(regexApplyTo: dashboardv2beta1.VariableRegexApplyTo): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.regexApplyTo = regexApplyTo;
        return this;
    }

    sort(sort: dashboardv2beta1.VariableSort): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.sort = sort;
        return this;
    }

    definition(definition: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.definition = definition;
        return this;
    }

    options(options: dashboardv2beta1.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    multi(multi: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }

    includeAll(includeAll: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }

    allValue(allValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }

    placeholder(placeholder: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.placeholder = placeholder;
        return this;
    }

    allowCustomValue(allowCustomValue: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }

    staticOptions(staticOptions: dashboardv2beta1.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.staticOptions = staticOptions;
        return this;
    }

    staticOptionsOrder(staticOptionsOrder: "before" | "after" | "sorted"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryVariableSpec();
        }
        this.internal.spec.staticOptionsOrder = staticOptionsOrder;
        return this;
    }
}

