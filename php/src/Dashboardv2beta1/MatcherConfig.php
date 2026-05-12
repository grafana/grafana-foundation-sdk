<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
 * It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
 */
class MatcherConfig implements \JsonSerializable
{
    /**
     * The matcher id. This is used to find the matcher implementation from registry.
     */
    public string $id;

    /**
     * If set, limits this matcher to fields of that type. If not set, "series" mode is used.
     */
    public ?\Grafana\Foundation\Dashboardv2beta1\MatcherScope $scope;

    /**
     * The matcher options. This is specific to the matcher implementation.
     * @var mixed|null
     */
    public $options;

    /**
     * @param string|null $id
     * @param \Grafana\Foundation\Dashboardv2beta1\MatcherScope|null $scope
     * @param mixed|null $options
     */
    public function __construct(?string $id = null, ?\Grafana\Foundation\Dashboardv2beta1\MatcherScope $scope = null,  $options = null)
    {
        $this->id = $id ?: "";
        $this->scope = $scope;
        $this->options = $options;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, scope?: string, options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            scope: isset($data["scope"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\MatcherScope::fromValue($input); })($data["scope"]) : null,
            options: $data["options"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->id = $this->id;
        if (isset($this->scope)) {
            $data->scope = $this->scope;
        }
        if (isset($this->options)) {
            $data->options = $this->options;
        }
        return $data;
    }
}
