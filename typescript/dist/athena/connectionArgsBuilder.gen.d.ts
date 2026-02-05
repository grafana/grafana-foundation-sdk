import * as cog from '../cog';
import * as athena from '../athena';
export declare class ConnectionArgsBuilder implements cog.Builder<athena.ConnectionArgs> {
    protected readonly internal: athena.ConnectionArgs;
    constructor();
    /**
     * Builds the object.
     */
    build(): athena.ConnectionArgs;
    region(region: string): this;
    catalog(catalog: string): this;
    database(database: string): this;
    resultReuseEnabled(resultReuseEnabled: boolean): this;
    resultReuseMaxAgeInMinutes(resultReuseMaxAgeInMinutes: number): this;
}
