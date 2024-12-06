<?php

namespace Grafana\Foundation\Cloudwatch;

final class LogsQueryLanguage implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LogsQueryLanguage>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function cWLI(): self
    {
        if (!isset(self::$instances["CWLI"])) {
            self::$instances["CWLI"] = new self("CWLI");
        }

        return self::$instances["CWLI"];
    }

    public static function sQL(): self
    {
        if (!isset(self::$instances["SQL"])) {
            self::$instances["SQL"] = new self("SQL");
        }

        return self::$instances["SQL"];
    }

    public static function pPL(): self
    {
        if (!isset(self::$instances["PPL"])) {
            self::$instances["PPL"] = new self("PPL");
        }

        return self::$instances["PPL"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "CWLI") {
            return self::cWLI();
        }

        if ($value === "SQL") {
            return self::sQL();
        }

        if ($value === "PPL") {
            return self::pPL();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LogsQueryLanguage");
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

