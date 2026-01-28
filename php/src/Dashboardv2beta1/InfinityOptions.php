<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class InfinityOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod $method;

    public string $url;

    public ?string $body;

    /**
     * These are 2D arrays of strings, each representing a key-value pair
     * We are defining them this way because we can't generate a go struct that
     * that would have exactly two strings in each sub-array
     * @var array<array<string>>|null
     */
    public ?array $queryParams;

    public string $datasourceUid;

    /**
     * @var array<array<string>>|null
     */
    public ?array $headers;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod|null $method
     * @param string|null $url
     * @param string|null $body
     * @param array<array<string>>|null $queryParams
     * @param string|null $datasourceUid
     * @param array<array<string>>|null $headers
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod $method = null, ?string $url = null, ?string $body = null, ?array $queryParams = null, ?string $datasourceUid = null, ?array $headers = null)
    {
        $this->method = $method ?: \Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod::GET();
        $this->url = $url ?: "";
        $this->body = $body;
        $this->queryParams = $queryParams;
        $this->datasourceUid = $datasourceUid ?: "";
        $this->headers = $headers;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{method?: string, url?: string, body?: string, queryParams?: array<array<string>>, datasourceUid?: string, headers?: array<array<string>>} $inputData */
        $data = $inputData;
        return new self(
            method: isset($data["method"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod::fromValue($input); })($data["method"]) : null,
            url: $data["url"] ?? null,
            body: $data["body"] ?? null,
            queryParams: $data["queryParams"] ?? null,
            datasourceUid: $data["datasourceUid"] ?? null,
            headers: $data["headers"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->method = $this->method;
        $data->url = $this->url;
        $data->datasourceUid = $this->datasourceUid;
        if (isset($this->body)) {
            $data->body = $this->body;
        }
        if (isset($this->queryParams)) {
            $data->queryParams = $this->queryParams;
        }
        if (isset($this->headers)) {
            $data->headers = $this->headers;
        }
        return $data;
    }
}
