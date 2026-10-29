<?php

namespace Grafana\Foundation\Alertingv1beta1;

class NamedRoutingTree implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv1beta1\NotificationSettingsType $type;

    public string $routingTree;

    /**
     * @param string|null $routingTree
     */
    public function __construct(?string $routingTree = null)
    {
        $this->type = \Grafana\Foundation\Alertingv1beta1\NotificationSettingsType::simplifiedRouting();
        $this->routingTree = $routingTree ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "NamedRoutingTree", routingTree?: string} $inputData */
        $data = $inputData;
        return new self(
            routingTree: $data["routingTree"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->routingTree = $this->routingTree;
        return $data;
    }
}
