<?php

namespace Grafana\Foundation\Elasticsearch;

final class BaseMetricAggregationConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\BaseMetricAggregation $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\BaseMetricAggregationBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Elasticsearch\MetricAggregationType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

