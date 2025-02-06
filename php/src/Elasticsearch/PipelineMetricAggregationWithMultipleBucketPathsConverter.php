<?php

namespace Grafana\Foundation\Elasticsearch;

final class PipelineMetricAggregationWithMultipleBucketPathsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\PipelineMetricAggregationWithMultipleBucketPaths $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\PipelineMetricAggregationWithMultipleBucketPathsBuilder())',
        ];
            if ($input->pipelineVariables !== null && count($input->pipelineVariables) >= 1) {
    
        
    $buffer = 'pipelineVariables(';
        $tmparg0 = [];
        foreach ($input->pipelineVariables as $arg1) {
        $tmppipelineVariablesarg1 = \Grafana\Foundation\Elasticsearch\PipelineVariableConverter::convert($arg1);
        $tmparg0[] = $tmppipelineVariablesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
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

