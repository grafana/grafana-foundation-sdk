<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchGeoHashGridSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchGeoHashGridSettingsBuilder())',
        ];
            if ($input->precision !== null && $input->precision !== "") {
    
        
    $buffer = 'precision(';
        $arg0 =\var_export($input->precision, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

