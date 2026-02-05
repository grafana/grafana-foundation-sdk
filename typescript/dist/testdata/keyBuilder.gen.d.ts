import * as cog from '../cog';
import * as testdata from '../testdata';
export declare class KeyBuilder implements cog.Builder<testdata.Key> {
    protected readonly internal: testdata.Key;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.Key;
    tick(tick: number): this;
    type(type: string): this;
    uid(uid: string): this;
}
