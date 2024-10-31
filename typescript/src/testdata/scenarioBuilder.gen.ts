// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

// TODO: Should this live here given it's not used in the dataquery?
export class ScenarioBuilder implements cog.Builder<testdata.Scenario> {
    protected readonly internal: testdata.Scenario;

    constructor() {
        this.internal = testdata.defaultScenario();
    }

    /**
     * Builds the object.
     */
    build(): testdata.Scenario {
        return this.internal;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    stringInput(stringInput: string): this {
        this.internal.stringInput = stringInput;
        return this;
    }

    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    hideAliasField(hideAliasField: boolean): this {
        this.internal.hideAliasField = hideAliasField;
        return this;
    }
}
