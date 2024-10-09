<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureMetricDimensionConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AzureMetricDimension $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureMetricDimensionBuilder())',
        ];
            if ($input->dimension !== null && $input->dimension !== "") {
    
        
    $buffer = 'dimension(';
        $arg0 =\var_export($input->dimension, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->operator !== null && $input->operator !== "") {
    
        
    $buffer = 'operator(';
        $arg0 =\var_export($input->operator, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filters !== null && count($input->filters) >= 1) {
    
        
    $buffer = 'filters(';
        $tmparg0 = [];
        foreach ($input->filters as $arg1) {
        $tmpfiltersarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpfiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filter !== null && $input->filter !== "") {
    
        
    $buffer = 'filter(';
        $arg0 =\var_export($input->filter, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

