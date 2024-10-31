// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';

export class KeyBuilder implements cog.Builder<testdata.Key> {
    protected readonly internal: testdata.Key;

    constructor() {
        this.internal = testdata.defaultKey();
    }

    /**
     * Builds the object.
     */
    build(): testdata.Key {
        return this.internal;
    }

    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    tick(tick: number): this {
        this.internal.tick = tick;
        return this;
    }

    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }
}
