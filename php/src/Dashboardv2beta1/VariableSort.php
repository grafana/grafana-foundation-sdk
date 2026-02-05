<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Sort variable options
 * Accepted values are:
 * `disabled`: No sorting
 * `alphabeticalAsc`: Alphabetical ASC
 * `alphabeticalDesc`: Alphabetical DESC
 * `numericalAsc`: Numerical ASC
 * `numericalDesc`: Numerical DESC
 * `alphabeticalCaseInsensitiveAsc`: Alphabetical Case Insensitive ASC
 * `alphabeticalCaseInsensitiveDesc`: Alphabetical Case Insensitive DESC
 * `naturalAsc`: Natural ASC
 * `naturalDesc`: Natural DESC
 * VariableSort enum with default value
 */
final class VariableSort implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VariableSort>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function disabled(): self
    {
        if (!isset(self::$instances["disabled"])) {
            self::$instances["disabled"] = new self("disabled");
        }

        return self::$instances["disabled"];
    }

    public static function alphabeticalAsc(): self
    {
        if (!isset(self::$instances["alphabeticalAsc"])) {
            self::$instances["alphabeticalAsc"] = new self("alphabeticalAsc");
        }

        return self::$instances["alphabeticalAsc"];
    }

    public static function alphabeticalDesc(): self
    {
        if (!isset(self::$instances["alphabeticalDesc"])) {
            self::$instances["alphabeticalDesc"] = new self("alphabeticalDesc");
        }

        return self::$instances["alphabeticalDesc"];
    }

    public static function numericalAsc(): self
    {
        if (!isset(self::$instances["numericalAsc"])) {
            self::$instances["numericalAsc"] = new self("numericalAsc");
        }

        return self::$instances["numericalAsc"];
    }

    public static function numericalDesc(): self
    {
        if (!isset(self::$instances["numericalDesc"])) {
            self::$instances["numericalDesc"] = new self("numericalDesc");
        }

        return self::$instances["numericalDesc"];
    }

    public static function alphabeticalCaseInsensitiveAsc(): self
    {
        if (!isset(self::$instances["alphabeticalCaseInsensitiveAsc"])) {
            self::$instances["alphabeticalCaseInsensitiveAsc"] = new self("alphabeticalCaseInsensitiveAsc");
        }

        return self::$instances["alphabeticalCaseInsensitiveAsc"];
    }

    public static function alphabeticalCaseInsensitiveDesc(): self
    {
        if (!isset(self::$instances["alphabeticalCaseInsensitiveDesc"])) {
            self::$instances["alphabeticalCaseInsensitiveDesc"] = new self("alphabeticalCaseInsensitiveDesc");
        }

        return self::$instances["alphabeticalCaseInsensitiveDesc"];
    }

    public static function naturalAsc(): self
    {
        if (!isset(self::$instances["naturalAsc"])) {
            self::$instances["naturalAsc"] = new self("naturalAsc");
        }

        return self::$instances["naturalAsc"];
    }

    public static function naturalDesc(): self
    {
        if (!isset(self::$instances["naturalDesc"])) {
            self::$instances["naturalDesc"] = new self("naturalDesc");
        }

        return self::$instances["naturalDesc"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "disabled") {
            return self::disabled();
        }

        if ($value === "alphabeticalAsc") {
            return self::alphabeticalAsc();
        }

        if ($value === "alphabeticalDesc") {
            return self::alphabeticalDesc();
        }

        if ($value === "numericalAsc") {
            return self::numericalAsc();
        }

        if ($value === "numericalDesc") {
            return self::numericalDesc();
        }

        if ($value === "alphabeticalCaseInsensitiveAsc") {
            return self::alphabeticalCaseInsensitiveAsc();
        }

        if ($value === "alphabeticalCaseInsensitiveDesc") {
            return self::alphabeticalCaseInsensitiveDesc();
        }

        if ($value === "naturalAsc") {
            return self::naturalAsc();
        }

        if ($value === "naturalDesc") {
            return self::naturalDesc();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableSort");
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

