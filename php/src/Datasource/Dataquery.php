<?php

namespace Grafana\Foundation\Datasource;

class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * Panel ID from wich the queries will be reused.
     */
    public int $panelId;

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

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * @param int|null $panelId
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     */
    public function __construct(?int $panelId = null, ?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null)
    {
        $this->panelId = $panelId ?: 0;
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->datasource = $datasource;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{panelId?: int, refId?: string, hide?: bool, queryType?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            panelId: $data["panelId"] ?? null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
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
            "panelId" => $this->panelId,
            "refId" => $this->refId,
        ];
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
        return 'datasource';
    }
}
