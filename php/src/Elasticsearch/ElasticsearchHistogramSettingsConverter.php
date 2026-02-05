<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchHistogramSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchHistogramSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchHistogramSettingsBuilder())',
        ];
            if ($input->interval !== null && $input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->minDocCount !== null && $input->minDocCount !== "") {
    
        
    $buffer = 'minDocCount(';
        $arg0 =\var_export($input->minDocCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

