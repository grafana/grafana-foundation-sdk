<?php

namespace Grafana\Foundation\Alerting;

class RecordRule implements \JsonSerializable
{
    /**
     * Which expression node should be used as the input for the recorded metric.
     */
    public string $from;

    /**
     * Name of the recorded metric.
     */
    public string $metric;

    /**
     * Which data source should be used to write the output of the recording rule, specified by UID.
     */
    public ?string $targetDatasourceUid;

    /**
     * @param string|null $from
     * @param string|null $metric
     * @param string|null $targetDatasourceUid
     */
    public function __construct(?string $from = null, ?string $metric = null, ?string $targetDatasourceUid = null)
    {
        $this->from = $from ?: "";
        $this->metric = $metric ?: "";
        $this->targetDatasourceUid = $targetDatasourceUid;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: string, metric?: string, target_datasource_uid?: string} $inputData */
        $data = $inputData;
        return new self(
            from: $data["from"] ?? null,
            metric: $data["metric"] ?? null,
            targetDatasourceUid: $data["target_datasource_uid"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->from = $this->from;
        $data->metric = $this->metric;
        if (isset($this->targetDatasourceUid)) {
            $data->target_datasource_uid = $this->targetDatasourceUid;
        }
        return $data;
    }
}
