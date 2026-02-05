<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchTopMetricsSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchTopMetricsSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchTopMetricsSettingsBuilder())',
        ];
            if ($input->order !== null && $input->order !== "") {
    
        
    $buffer = 'order(';
        $arg0 =\var_export($input->order, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->orderBy !== null && $input->orderBy !== "") {
    
        
    $buffer = 'orderBy(';
        $arg0 =\var_export($input->orderBy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metrics !== null && count($input->metrics) >= 1) {
    
        
    $buffer = 'metrics(';
        $tmparg0 = [];
        foreach ($input->metrics as $arg1) {
        $tmpmetricsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpmetricsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

