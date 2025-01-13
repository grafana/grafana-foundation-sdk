<?php

namespace Grafana\Foundation\Athena;

/**
 * Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
 */
class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    public \Grafana\Foundation\Athena\FormatOptions $format;

    public \Grafana\Foundation\Athena\ConnectionArgs $connectionArgs;

    public ?string $table;

    public ?string $column;

    public ?string $queryID;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    public string $rawSQL;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * @param \Grafana\Foundation\Athena\FormatOptions|null $format
     * @param \Grafana\Foundation\Athena\ConnectionArgs|null $connectionArgs
     * @param string|null $table
     * @param string|null $column
     * @param string|null $queryID
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param string|null $rawSQL
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     */
    public function __construct(?\Grafana\Foundation\Athena\FormatOptions $format = null, ?\Grafana\Foundation\Athena\ConnectionArgs $connectionArgs = null, ?string $table = null, ?string $column = null, ?string $queryID = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?string $rawSQL = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null)
    {
        $this->format = $format ?: \Grafana\Foundation\Athena\FormatOptions::TimeSeries();
        $this->connectionArgs = $connectionArgs ?: new \Grafana\Foundation\Athena\ConnectionArgs();
        $this->table = $table;
        $this->column = $column;
        $this->queryID = $queryID;
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->rawSQL = $rawSQL ?: "";
        $this->datasource = $datasource;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{format?: int, connectionArgs?: mixed, table?: string, column?: string, queryID?: string, refId?: string, hide?: bool, queryType?: string, rawSQL?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            format: isset($data["format"]) ? (function($input) { return \Grafana\Foundation\Athena\FormatOptions::fromValue($input); })($data["format"]) : null,
            connectionArgs: isset($data["connectionArgs"]) ? (function($input) {
    	/** @var array{region?: string, catalog?: string, database?: string, resultReuseEnabled?: bool, resultReuseMaxAgeInMinutes?: float} */
    $val = $input;
    	return \Grafana\Foundation\Athena\ConnectionArgs::fromArray($val);
    })($data["connectionArgs"]) : null,
            table: $data["table"] ?? null,
            column: $data["column"] ?? null,
            queryID: $data["queryID"] ?? null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            rawSQL: $data["rawSQL"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "format" => $this->format,
            "connectionArgs" => $this->connectionArgs,
            "refId" => $this->refId,
            "rawSQL" => $this->rawSQL,
        ];
        if (isset($this->table)) {
            $data["table"] = $this->table;
        }
        if (isset($this->column)) {
            $data["column"] = $this->column;
        }
        if (isset($this->queryID)) {
            $data["queryID"] = $this->queryID;
        }
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return 'grafana-athena-datasource';
    }
}
