<?php

namespace Grafana\Foundation\Testdata;

final class StreamingQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, StreamingQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function fetch(): self
    {
        if (!isset(self::$instances["Fetch"])) {
            self::$instances["Fetch"] = new self("fetch");
        }

        return self::$instances["Fetch"];
    }

    public static function logs(): self
    {
        if (!isset(self::$instances["Logs"])) {
            self::$instances["Logs"] = new self("logs");
        }

        return self::$instances["Logs"];
    }

    public static function signal(): self
    {
        if (!isset(self::$instances["Signal"])) {
            self::$instances["Signal"] = new self("signal");
        }

        return self::$instances["Signal"];
    }

    public static function traces(): self
    {
        if (!isset(self::$instances["Traces"])) {
            self::$instances["Traces"] = new self("traces");
        }

        return self::$instances["Traces"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "fetch") {
            return self::fetch();
        }

        if ($value === "logs") {
            return self::logs();
        }

        if ($value === "signal") {
            return self::signal();
        }

        if ($value === "traces") {
            return self::traces();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum StreamingQueryType");
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

