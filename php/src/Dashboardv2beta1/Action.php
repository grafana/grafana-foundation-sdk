<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Action implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ActionType $type;

    public string $title;

    public ?\Grafana\Foundation\Dashboardv2beta1\FetchOptions $fetch;

    public ?\Grafana\Foundation\Dashboardv2beta1\InfinityOptions $infinity;

    public ?string $confirmation;

    public ?bool $oneClick;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\ActionVariable>|null
     */
    public ?array $variables;

    public ?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle $style;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\ActionType|null $type
     * @param string|null $title
     * @param \Grafana\Foundation\Dashboardv2beta1\FetchOptions|null $fetch
     * @param \Grafana\Foundation\Dashboardv2beta1\InfinityOptions|null $infinity
     * @param string|null $confirmation
     * @param bool|null $oneClick
     * @param array<\Grafana\Foundation\Dashboardv2beta1\ActionVariable>|null $variables
     * @param \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle|null $style
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\ActionType $type = null, ?string $title = null, ?\Grafana\Foundation\Dashboardv2beta1\FetchOptions $fetch = null, ?\Grafana\Foundation\Dashboardv2beta1\InfinityOptions $infinity = null, ?string $confirmation = null, ?bool $oneClick = null, ?array $variables = null, ?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle $style = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Dashboardv2beta1\ActionType::Fetch();
        $this->title = $title ?: "";
        $this->fetch = $fetch;
        $this->infinity = $infinity;
        $this->confirmation = $confirmation;
        $this->oneClick = $oneClick;
        $this->variables = $variables;
        $this->style = $style;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, title?: string, fetch?: mixed, infinity?: mixed, confirmation?: string, oneClick?: bool, variables?: array<mixed>, style?: mixed} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\ActionType::fromValue($input); })($data["type"]) : null,
            title: $data["title"] ?? null,
            fetch: isset($data["fetch"]) ? (function($input) {
    	/** @var array{method?: string, url?: string, body?: string, queryParams?: array<array<string>>, headers?: array<array<string>>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\FetchOptions::fromArray($val);
    })($data["fetch"]) : null,
            infinity: isset($data["infinity"]) ? (function($input) {
    	/** @var array{method?: string, url?: string, body?: string, queryParams?: array<array<string>>, datasourceUid?: string, headers?: array<array<string>>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\InfinityOptions::fromArray($val);
    })($data["infinity"]) : null,
            confirmation: $data["confirmation"] ?? null,
            oneClick: $data["oneClick"] ?? null,
            variables: array_filter(array_map((function($input) {
    	/** @var array{key?: string, name?: string, type?: "string"} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ActionVariable::fromArray($val);
    }), $data["variables"] ?? [])),
            style: isset($data["style"]) ? (function($input) {
    	/** @var array{backgroundColor?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle::fromArray($val);
    })($data["style"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->title = $this->title;
        if (isset($this->fetch)) {
            $data->fetch = $this->fetch;
        }
        if (isset($this->infinity)) {
            $data->infinity = $this->infinity;
        }
        if (isset($this->confirmation)) {
            $data->confirmation = $this->confirmation;
        }
        if (isset($this->oneClick)) {
            $data->oneClick = $this->oneClick;
        }
        if (isset($this->variables)) {
            $data->variables = $this->variables;
        }
        if (isset($this->style)) {
            $data->style = $this->style;
        }
        return $data;
    }
}
