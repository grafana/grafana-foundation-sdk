<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * SLO sub-query properties.
 */
class SLOQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public ?string $perSeriesAligner;

    /**
     * Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public ?string $alignmentPeriod;

    /**
     * SLO selector.
     */
    public string $selectorName;

    /**
     * ID for the service the SLO is in.
     */
    public string $serviceId;

    /**
     * Name for the service the SLO is in.
     */
    public string $serviceName;

    /**
     * ID for the SLO.
     */
    public string $sloId;

    /**
     * Name of the SLO.
     */
    public string $sloName;

    /**
     * SLO goal value.
     */
    public ?float $goal;

    /**
     * Specific lookback period for the SLO.
     */
    public ?string $lookbackPeriod;

    /**
     * @param string|null $projectName
     * @param string|null $perSeriesAligner
     * @param string|null $alignmentPeriod
     * @param string|null $selectorName
     * @param string|null $serviceId
     * @param string|null $serviceName
     * @param string|null $sloId
     * @param string|null $sloName
     * @param float|null $goal
     * @param string|null $lookbackPeriod
     */
    public function __construct(?string $projectName = null, ?string $perSeriesAligner = null, ?string $alignmentPeriod = null, ?string $selectorName = null, ?string $serviceId = null, ?string $serviceName = null, ?string $sloId = null, ?string $sloName = null, ?float $goal = null, ?string $lookbackPeriod = null)
    {
        $this->projectName = $projectName ?: "";
        $this->perSeriesAligner = $perSeriesAligner;
        $this->alignmentPeriod = $alignmentPeriod;
        $this->selectorName = $selectorName ?: "";
        $this->serviceId = $serviceId ?: "";
        $this->serviceName = $serviceName ?: "";
        $this->sloId = $sloId ?: "";
        $this->sloName = $sloName ?: "";
        $this->goal = $goal;
        $this->lookbackPeriod = $lookbackPeriod;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{projectName?: string, perSeriesAligner?: string, alignmentPeriod?: string, selectorName?: string, serviceId?: string, serviceName?: string, sloId?: string, sloName?: string, goal?: float, lookbackPeriod?: string} $inputData */
        $data = $inputData;
        return new self(
            projectName: $data["projectName"] ?? null,
            perSeriesAligner: $data["perSeriesAligner"] ?? null,
            alignmentPeriod: $data["alignmentPeriod"] ?? null,
            selectorName: $data["selectorName"] ?? null,
            serviceId: $data["serviceId"] ?? null,
            serviceName: $data["serviceName"] ?? null,
            sloId: $data["sloId"] ?? null,
            sloName: $data["sloName"] ?? null,
            goal: $data["goal"] ?? null,
            lookbackPeriod: $data["lookbackPeriod"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "projectName" => $this->projectName,
            "selectorName" => $this->selectorName,
            "serviceId" => $this->serviceId,
            "serviceName" => $this->serviceName,
            "sloId" => $this->sloId,
            "sloName" => $this->sloName,
        ];
        if (isset($this->perSeriesAligner)) {
            $data["perSeriesAligner"] = $this->perSeriesAligner;
        }
        if (isset($this->alignmentPeriod)) {
            $data["alignmentPeriod"] = $this->alignmentPeriod;
        }
        if (isset($this->goal)) {
            $data["goal"] = $this->goal;
        }
        if (isset($this->lookbackPeriod)) {
            $data["lookbackPeriod"] = $this->lookbackPeriod;
        }
        return $data;
    }
}
