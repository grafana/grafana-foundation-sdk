<?php

namespace Grafana\Foundation\Azuremonitor;

final class AzureTracesFilterConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AzureTracesFilter $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AzureTracesFilterBuilder())',
        ];
            if ($input->property !== "") {
    
        
    $buffer = 'property(';
        $arg0 =\var_export($input->property, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->operation !== "") {
    
        
    $buffer = 'operation(';
        $arg0 =\var_export($input->operation, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->filters) >= 1) {
    
        
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

        return \implode("\n\t->", $calls);
    }
}

