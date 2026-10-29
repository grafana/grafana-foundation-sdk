<?php

namespace Grafana\Foundation\Alertingv0alpha1;

class PanelRef implements \JsonSerializable
{
    public string $dashboardUID;

    public int $panelID;

    /**
     * @param string|null $dashboardUID
     * @param int|null $panelID
     */
    public function __construct(?string $dashboardUID = null, ?int $panelID = null)
    {
        $this->dashboardUID = $dashboardUID ?: "";
        $this->panelID = $panelID ?: 0;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{dashboardUID?: string, panelID?: int} $inputData */
        $data = $inputData;
        return new self(
            dashboardUID: $data["dashboardUID"] ?? null,
            panelID: $data["panelID"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->dashboardUID = $this->dashboardUID;
        $data->panelID = $this->panelID;
        return $data;
    }
}
