// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Query variable kind
export class QueryVariableBuilder implements cog.Builder<dashboardv2.QueryVariableKind> {
    protected readonly internal: dashboardv2.QueryVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2.defaultQueryVariableKind();
        this.internal.kind = "QueryVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.QueryVariableKind {
        return this.internal;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    current(current: dashboardv2.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    refresh(refresh: dashboardv2.VariableRefresh): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    query(query: cog.Builder<dashboardv2.DataQueryKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }

    regex(regex: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.regex = regex;
        return this;
    }

    regexApplyTo(regexApplyTo: dashboardv2.VariableRegexApplyTo): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.regexApplyTo = regexApplyTo;
        return this;
    }

    sort(sort: dashboardv2.VariableSort): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.sort = sort;
        return this;
    }

    definition(definition: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.definition = definition;
        return this;
    }

    options(options: dashboardv2.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    multi(multi: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }

    includeAll(includeAll: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }

    allValue(allValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }

    placeholder(placeholder: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.placeholder = placeholder;
        return this;
    }

    allowCustomValue(allowCustomValue: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }

    staticOptions(staticOptions: dashboardv2.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.staticOptions = staticOptions;
        return this;
    }

    staticOptionsOrder(staticOptionsOrder: "before" | "after" | "sorted"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        this.internal.spec.staticOptionsOrder = staticOptionsOrder;
        return this;
    }

    origin(origin: cog.Builder<dashboardv2.ControlSourceRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryVariableSpec();
        }
        const originResource = origin.build();
        this.internal.spec.origin = originResource;
        return this;
    }
}

