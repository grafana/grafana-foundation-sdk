<?php

namespace Grafana\Foundation\Elasticsearch;

final class BasePipelineMetricAggregationConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\BasePipelineMetricAggregation $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\BasePipelineMetricAggregationBuilder())',
        ];
            if ($input->pipelineAgg !== null && $input->pipelineAgg !== "") {
    
        
    $buffer = 'pipelineAgg(';
        $arg0 =\var_export($input->pipelineAgg, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
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

