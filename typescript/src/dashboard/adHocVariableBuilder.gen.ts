// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';
import * as common from '../common';

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
export class AdHocVariableBuilder implements cog.Builder<dashboard.VariableModel> {
    protected readonly internal: dashboard.VariableModel;

    constructor(name: string) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Adhoc;
    }

    /**
     * Builds the object.
     */
    build(): dashboard.VariableModel {
        return this.internal;
    }

    // Name of variable
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Optional display name
    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    // Visibility configuration for the variable
    hide(hide: dashboard.VariableHide): this {
        this.internal.hide = hide;
        return this;
    }

    // Description of variable. It can be defined but `null`.
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // Data source used to fetch values for a variable. It can be defined but `null`.
    datasource(datasource: common.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // Allow custom values to be entered in the variable
    allowCustomValue(allowCustomValue: boolean): this {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }
}

