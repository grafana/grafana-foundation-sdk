---
title: <span class="badge object-type-enum"></span> CodeLanguage
---
# <span class="badge object-type-enum"></span> CodeLanguage

## Definition

```php
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
        if (!isset(self::$instances["json"])) {
            self::$instances["json"] = new self("json");
        }

        return self::$instances["json"];
    }

    public static function yaml(): self
    {
        if (!isset(self::$instances["yaml"])) {
            self::$instances["yaml"] = new self("yaml");
        }

        return self::$instances["yaml"];
    }

    public static function xml(): self
    {
        if (!isset(self::$instances["xml"])) {
            self::$instances["xml"] = new self("xml");
        }

        return self::$instances["xml"];
    }

    public static function typescript(): self
    {
        if (!isset(self::$instances["typescript"])) {
            self::$instances["typescript"] = new self("typescript");
        }

        return self::$instances["typescript"];
    }

    public static function sql(): self
    {
        if (!isset(self::$instances["sql"])) {
            self::$instances["sql"] = new self("sql");
        }

        return self::$instances["sql"];
    }

    public static function go(): self
    {
        if (!isset(self::$instances["go"])) {
            self::$instances["go"] = new self("go");
        }

        return self::$instances["go"];
    }

    public static function markdown(): self
    {
        if (!isset(self::$instances["markdown"])) {
            self::$instances["markdown"] = new self("markdown");
        }

        return self::$instances["markdown"];
    }

    public static function html(): self
    {
        if (!isset(self::$instances["html"])) {
            self::$instances["html"] = new self("html");
        }

        return self::$instances["html"];
    }

    public static function plaintext(): self
    {
        if (!isset(self::$instances["plaintext"])) {
            self::$instances["plaintext"] = new self("plaintext");
        }

        return self::$instances["plaintext"];
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

```
