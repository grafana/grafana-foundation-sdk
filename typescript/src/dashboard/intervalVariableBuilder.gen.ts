// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
export class IntervalVariableBuilder implements cog.Builder<dashboard.VariableModel> {
    protected readonly internal: dashboard.VariableModel;

    constructor(name: string) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Interval;
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

    // Query used to fetch values for a variable
    values(query: string | Record<string, any>): this {
        this.internal.query = query;
        return this;
    }

    // Shows current selected variable text/value on the dashboard
    current(current: dashboard.VariableOption): this {
        this.internal.current = current;
        return this;
    }

    // Allow custom values to be entered in the variable
    allowCustomValue(allowCustomValue: boolean): this {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }

    // Options that can be selected for a variable.
    options(options: dashboard.VariableOption[]): this {
        this.internal.options = options;
        return this;
    }

    // Dynamically calculates interval by dividing time range by the count specified.
    auto(auto: boolean): this {
        this.internal.auto = auto;
        return this;
    }

    // The minimum threshold below which the step count intervals will not divide the time.
    minInterval(autoMin: string): this {
        this.internal.auto_min = autoMin;
        return this;
    }

    // How many times the current time range should be divided to calculate the value, similar to the Max data points query option.
    // For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.
    stepCount(autoCount: number): this {
        if (!(autoCount > 0)) {
            throw new Error("autoCount must be > 0");
        }
        this.internal.auto_count = autoCount;
        return this;
    }
}
