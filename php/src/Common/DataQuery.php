<?php

namespace Grafana\Foundation\Common;

/**
 * These are the common properties available to all queries in all datasources.
 * Specific implementations will *extend* this interface, adding the required
 * properties for the given context.
 */
class DataQuery implements \JsonSerializable
{
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

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
     * @var mixed|null
     */
    public $datasource;

    /**
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param mixed|null $datasource
     */
    public function __construct(?string $refId = null, ?bool $hide = null, ?string $queryType = null,  $datasource = null)
    {
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
        /** @var array{refId?: string, hide?: bool, queryType?: string, datasource?: mixed} $inputData */
        $data = $inputData;
        return new self(
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            datasource: $data["datasource"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
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
}
