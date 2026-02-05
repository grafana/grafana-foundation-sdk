<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchMetricAggregationWithMissingSupportSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder())',
        ];
            if ($input->missing !== null && $input->missing !== "") {
    
        
    $buffer = 'missing(';
        $arg0 =\var_export($input->missing, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

