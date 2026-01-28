<?php

namespace Grafana\Foundation\Bigquery;

/**
 * Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
 */
class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    public ?string $dataset;

    public ?string $table;

    public ?string $project;

    public \Grafana\Foundation\Bigquery\QueryFormat $format;

    public ?bool $rawQuery;

    public string $rawSql;

    public ?string $location;

    public ?bool $partitioned;

    public ?string $partitionedField;

    public ?bool $convertToUTC;

    public ?bool $sharded;

    public ?\Grafana\Foundation\Bigquery\QueryPriority $queryPriority;

    public ?string $timeShift;

    public ?\Grafana\Foundation\Bigquery\EditorMode $editorMode;

    public ?\Grafana\Foundation\Bigquery\SQLExpression $sql;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public ?string $refId;

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Common\DataSourceRef $datasource;

    /**
     * @param string|null $dataset
     * @param string|null $table
     * @param string|null $project
     * @param \Grafana\Foundation\Bigquery\QueryFormat|null $format
     * @param bool|null $rawQuery
     * @param string|null $rawSql
     * @param string|null $location
     * @param bool|null $partitioned
     * @param string|null $partitionedField
     * @param bool|null $convertToUTC
     * @param bool|null $sharded
     * @param \Grafana\Foundation\Bigquery\QueryPriority|null $queryPriority
     * @param string|null $timeShift
     * @param \Grafana\Foundation\Bigquery\EditorMode|null $editorMode
     * @param \Grafana\Foundation\Bigquery\SQLExpression|null $sql
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param \Grafana\Foundation\Common\DataSourceRef|null $datasource
     */
    public function __construct(?string $dataset = null, ?string $table = null, ?string $project = null, ?\Grafana\Foundation\Bigquery\QueryFormat $format = null, ?bool $rawQuery = null, ?string $rawSql = null, ?string $location = null, ?bool $partitioned = null, ?string $partitionedField = null, ?bool $convertToUTC = null, ?bool $sharded = null, ?\Grafana\Foundation\Bigquery\QueryPriority $queryPriority = null, ?string $timeShift = null, ?\Grafana\Foundation\Bigquery\EditorMode $editorMode = null, ?\Grafana\Foundation\Bigquery\SQLExpression $sql = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?\Grafana\Foundation\Common\DataSourceRef $datasource = null)
    {
        $this->dataset = $dataset;
        $this->table = $table;
        $this->project = $project;
        $this->format = $format ?: \Grafana\Foundation\Bigquery\QueryFormat::Timeseries();
        $this->rawQuery = $rawQuery;
        $this->rawSql = $rawSql ?: "";
        $this->location = $location;
        $this->partitioned = $partitioned;
        $this->partitionedField = $partitionedField;
        $this->convertToUTC = $convertToUTC;
        $this->sharded = $sharded;
        $this->queryPriority = $queryPriority;
        $this->timeShift = $timeShift;
        $this->editorMode = $editorMode;
        $this->sql = $sql;
        $this->refId = $refId;
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->datasource = $datasource;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{dataset?: string, table?: string, project?: string, format?: int, rawQuery?: bool, rawSql?: string, location?: string, partitioned?: bool, partitionedField?: string, convertToUTC?: bool, sharded?: bool, queryPriority?: string, timeShift?: string, editorMode?: string, sql?: mixed, refId?: string, hide?: bool, queryType?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            dataset: $data["dataset"] ?? null,
            table: $data["table"] ?? null,
            project: $data["project"] ?? null,
            format: isset($data["format"]) ? (function($input) { return \Grafana\Foundation\Bigquery\QueryFormat::fromValue($input); })($data["format"]) : null,
            rawQuery: $data["rawQuery"] ?? null,
            rawSql: $data["rawSql"] ?? null,
            location: $data["location"] ?? null,
            partitioned: $data["partitioned"] ?? null,
            partitionedField: $data["partitionedField"] ?? null,
            convertToUTC: $data["convertToUTC"] ?? null,
            sharded: $data["sharded"] ?? null,
            queryPriority: isset($data["queryPriority"]) ? (function($input) { return \Grafana\Foundation\Bigquery\QueryPriority::fromValue($input); })($data["queryPriority"]) : null,
            timeShift: $data["timeShift"] ?? null,
            editorMode: isset($data["editorMode"]) ? (function($input) { return \Grafana\Foundation\Bigquery\EditorMode::fromValue($input); })($data["editorMode"]) : null,
            sql: isset($data["sql"]) ? (function($input) {
    	/** @var array{columns?: array<mixed>, from?: string, whereString?: string, groupBy?: array<mixed>, orderBy?: mixed, orderByDirection?: string, limit?: int, offset?: int} */
    $val = $input;
    	return \Grafana\Foundation\Bigquery\SQLExpression::fromArray($val);
    })($data["sql"]) : null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->format = $this->format;
        $data->rawSql = $this->rawSql;
        if (isset($this->dataset)) {
            $data->dataset = $this->dataset;
        }
        if (isset($this->table)) {
            $data->table = $this->table;
        }
        if (isset($this->project)) {
            $data->project = $this->project;
        }
        if (isset($this->rawQuery)) {
            $data->rawQuery = $this->rawQuery;
        }
        if (isset($this->location)) {
            $data->location = $this->location;
        }
        if (isset($this->partitioned)) {
            $data->partitioned = $this->partitioned;
        }
        if (isset($this->partitionedField)) {
            $data->partitionedField = $this->partitionedField;
        }
        if (isset($this->convertToUTC)) {
            $data->convertToUTC = $this->convertToUTC;
        }
        if (isset($this->sharded)) {
            $data->sharded = $this->sharded;
        }
        if (isset($this->queryPriority)) {
            $data->queryPriority = $this->queryPriority;
        }
        if (isset($this->timeShift)) {
            $data->timeShift = $this->timeShift;
        }
        if (isset($this->editorMode)) {
            $data->editorMode = $this->editorMode;
        }
        if (isset($this->sql)) {
            $data->sql = $this->sql;
        }
        if (isset($this->refId)) {
            $data->refId = $this->refId;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        if (isset($this->queryType)) {
            $data->queryType = $this->queryType;
        }
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return 'grafana-bigquery-datasource';
    }
}
