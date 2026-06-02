<?php

namespace Grafana\Foundation\Tempo;

final class TempoQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TempoQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function traceql(): self
    {
        if (!isset(self::$instances["Traceql"])) {
            self::$instances["Traceql"] = new self("traceql");
        }

        return self::$instances["Traceql"];
    }

    public static function traceqlSearch(): self
    {
        if (!isset(self::$instances["TraceqlSearch"])) {
            self::$instances["TraceqlSearch"] = new self("traceqlSearch");
        }

        return self::$instances["TraceqlSearch"];
    }

    public static function serviceMap(): self
    {
        if (!isset(self::$instances["ServiceMap"])) {
            self::$instances["ServiceMap"] = new self("serviceMap");
        }

        return self::$instances["ServiceMap"];
    }

    public static function upload(): self
    {
        if (!isset(self::$instances["Upload"])) {
            self::$instances["Upload"] = new self("upload");
        }

        return self::$instances["Upload"];
    }

    public static function nativeSearch(): self
    {
        if (!isset(self::$instances["NativeSearch"])) {
            self::$instances["NativeSearch"] = new self("nativeSearch");
        }

        return self::$instances["NativeSearch"];
    }

    public static function traceId(): self
    {
        if (!isset(self::$instances["TraceId"])) {
            self::$instances["TraceId"] = new self("traceId");
        }

        return self::$instances["TraceId"];
    }

    public static function clear(): self
    {
        if (!isset(self::$instances["Clear"])) {
            self::$instances["Clear"] = new self("clear");
        }

        return self::$instances["Clear"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "traceql") {
            return self::traceql();
        }

        if ($value === "traceqlSearch") {
            return self::traceqlSearch();
        }

        if ($value === "serviceMap") {
            return self::serviceMap();
        }

        if ($value === "upload") {
            return self::upload();
        }

        if ($value === "nativeSearch") {
            return self::nativeSearch();
        }

        if ($value === "traceId") {
            return self::traceId();
        }

        if ($value === "clear") {
            return self::clear();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TempoQueryType");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

