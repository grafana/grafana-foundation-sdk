<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchSerialDiffSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchSerialDiffSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchSerialDiffSettingsBuilder())',
        ];
            if ($input->lag !== null && $input->lag !== "") {
    
        
    $buffer = 'lag(';
        $arg0 =\var_export($input->lag, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

