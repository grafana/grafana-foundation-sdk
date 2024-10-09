<?php

namespace Grafana\Foundation\Elasticsearch;

final class MetricAggregationWithFieldConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\MetricAggregationWithField $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\MetricAggregationWithFieldBuilder())',
        ];
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'type(';
        switch (true) {
            case is_string($input->type):
                $disjunctiontype =\var_export($input->type, true);
                $arg0 = $disjunctiontype;
                break;
            case $input->type instanceof \Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType:
                $disjunctiontype ='\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationType::fromValue("'.$input->type.'")';
                $arg0 = $disjunctiontype;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
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

