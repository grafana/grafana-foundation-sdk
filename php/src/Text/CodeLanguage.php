<?php

namespace Grafana\Foundation\Text;

final class CodeLanguage implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, CodeLanguage>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function json(): self
    {
        if (!isset(self::$instances["Json"])) {
            self::$instances["Json"] = new self("json");
        }

        return self::$instances["Json"];
    }

    public static function yaml(): self
    {
        if (!isset(self::$instances["Yaml"])) {
            self::$instances["Yaml"] = new self("yaml");
        }

        return self::$instances["Yaml"];
    }

    public static function xml(): self
    {
        if (!isset(self::$instances["Xml"])) {
            self::$instances["Xml"] = new self("xml");
        }

        return self::$instances["Xml"];
    }

    public static function typescript(): self
    {
        if (!isset(self::$instances["Typescript"])) {
            self::$instances["Typescript"] = new self("typescript");
        }

        return self::$instances["Typescript"];
    }

    public static function sql(): self
    {
        if (!isset(self::$instances["Sql"])) {
            self::$instances["Sql"] = new self("sql");
        }

        return self::$instances["Sql"];
    }

    public static function go(): self
    {
        if (!isset(self::$instances["Go"])) {
            self::$instances["Go"] = new self("go");
        }

        return self::$instances["Go"];
    }

    public static function markdown(): self
    {
        if (!isset(self::$instances["Markdown"])) {
            self::$instances["Markdown"] = new self("markdown");
        }

        return self::$instances["Markdown"];
    }

    public static function html(): self
    {
        if (!isset(self::$instances["Html"])) {
            self::$instances["Html"] = new self("html");
        }

        return self::$instances["Html"];
    }

    public static function plaintext(): self
    {
        if (!isset(self::$instances["Plaintext"])) {
            self::$instances["Plaintext"] = new self("plaintext");
        }

        return self::$instances["Plaintext"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "json") {
            return self::json();
        }

        if ($value === "yaml") {
            return self::yaml();
        }

        if ($value === "xml") {
            return self::xml();
        }

        if ($value === "typescript") {
            return self::typescript();
        }

        if ($value === "sql") {
            return self::sql();
        }

        if ($value === "go") {
            return self::go();
        }

        if ($value === "markdown") {
            return self::markdown();
        }

        if ($value === "html") {
            return self::html();
        }

        if ($value === "plaintext") {
            return self::plaintext();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum CodeLanguage");
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

