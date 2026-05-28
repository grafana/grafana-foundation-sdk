<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * Dashboard specific preferences (applied per dashboard = all users using the dashboard)
 */
class Preferences implements \JsonSerializable
{
    /**
     * default layout template to be used when new containers are created
     * @var \Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\GridLayoutKind|null
     */
    public $layout;

    /**
     * @param \Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\GridLayoutKind|null $layout
     */
    public function __construct( $layout = null)
    {
        $this->layout = $layout;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{layout?: mixed|mixed} $inputData */
        $data = $inputData;
        return new self(
            layout: isset($data["layout"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "AutoGridLayout":
            return AutoGridLayoutKind::fromArray($input);
        case "GridLayout":
            return GridLayoutKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["layout"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->layout)) {
            $data->layout = $this->layout;
        }
        return $data;
    }
}
