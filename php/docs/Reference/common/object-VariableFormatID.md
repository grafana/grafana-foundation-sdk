---
title: <span class="badge object-type-enum"></span> VariableFormatID
---
# <span class="badge object-type-enum"></span> VariableFormatID

Optional formats for the template variable replace functions

See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options

## Definition

```php
final class VariableFormatID implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, VariableFormatID>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function lucene(): self
    {
        if (!isset(self::$instances["Lucene"])) {
            self::$instances["Lucene"] = new self("lucene");
        }

        return self::$instances["Lucene"];
    }

    public static function raw(): self
    {
        if (!isset(self::$instances["Raw"])) {
            self::$instances["Raw"] = new self("raw");
        }

        return self::$instances["Raw"];
    }

    public static function regex(): self
    {
        if (!isset(self::$instances["Regex"])) {
            self::$instances["Regex"] = new self("regex");
        }

        return self::$instances["Regex"];
    }

    public static function pipe(): self
    {
        if (!isset(self::$instances["Pipe"])) {
            self::$instances["Pipe"] = new self("pipe");
        }

        return self::$instances["Pipe"];
    }

    public static function distributed(): self
    {
        if (!isset(self::$instances["Distributed"])) {
            self::$instances["Distributed"] = new self("distributed");
        }

        return self::$instances["Distributed"];
    }

    public static function cSV(): self
    {
        if (!isset(self::$instances["CSV"])) {
            self::$instances["CSV"] = new self("csv");
        }

        return self::$instances["CSV"];
    }

    public static function hTML(): self
    {
        if (!isset(self::$instances["HTML"])) {
            self::$instances["HTML"] = new self("html");
        }

        return self::$instances["HTML"];
    }

    public static function jSON(): self
    {
        if (!isset(self::$instances["JSON"])) {
            self::$instances["JSON"] = new self("json");
        }

        return self::$instances["JSON"];
    }

    public static function percentEncode(): self
    {
        if (!isset(self::$instances["PercentEncode"])) {
            self::$instances["PercentEncode"] = new self("percentencode");
        }

        return self::$instances["PercentEncode"];
    }

    public static function uriEncode(): self
    {
        if (!isset(self::$instances["UriEncode"])) {
            self::$instances["UriEncode"] = new self("uriencode");
        }

        return self::$instances["UriEncode"];
    }

    public static function singleQuote(): self
    {
        if (!isset(self::$instances["SingleQuote"])) {
            self::$instances["SingleQuote"] = new self("singlequote");
        }

        return self::$instances["SingleQuote"];
    }

    public static function doubleQuote(): self
    {
        if (!isset(self::$instances["DoubleQuote"])) {
            self::$instances["DoubleQuote"] = new self("doublequote");
        }

        return self::$instances["DoubleQuote"];
    }

    public static function sQLString(): self
    {
        if (!isset(self::$instances["SQLString"])) {
            self::$instances["SQLString"] = new self("sqlstring");
        }

        return self::$instances["SQLString"];
    }

    public static function date(): self
    {
        if (!isset(self::$instances["Date"])) {
            self::$instances["Date"] = new self("date");
        }

        return self::$instances["Date"];
    }

    public static function glob(): self
    {
        if (!isset(self::$instances["Glob"])) {
            self::$instances["Glob"] = new self("glob");
        }

        return self::$instances["Glob"];
    }

    public static function text(): self
    {
        if (!isset(self::$instances["Text"])) {
            self::$instances["Text"] = new self("text");
        }

        return self::$instances["Text"];
    }

    public static function queryParam(): self
    {
        if (!isset(self::$instances["QueryParam"])) {
            self::$instances["QueryParam"] = new self("queryparam");
        }

        return self::$instances["QueryParam"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "lucene") {
            return self::lucene();
        }

        if ($value === "raw") {
            return self::raw();
        }

        if ($value === "regex") {
            return self::regex();
        }

        if ($value === "pipe") {
            return self::pipe();
        }

        if ($value === "distributed") {
            return self::distributed();
        }

        if ($value === "csv") {
            return self::cSV();
        }

        if ($value === "html") {
            return self::hTML();
        }

        if ($value === "json") {
            return self::jSON();
        }

        if ($value === "percentencode") {
            return self::percentEncode();
        }

        if ($value === "uriencode") {
            return self::uriEncode();
        }

        if ($value === "singlequote") {
            return self::singleQuote();
        }

        if ($value === "doublequote") {
            return self::doubleQuote();
        }

        if ($value === "sqlstring") {
            return self::sQLString();
        }

        if ($value === "date") {
            return self::date();
        }

        if ($value === "glob") {
            return self::glob();
        }

        if ($value === "text") {
            return self::text();
        }

        if ($value === "queryparam") {
            return self::queryParam();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableFormatID");
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
