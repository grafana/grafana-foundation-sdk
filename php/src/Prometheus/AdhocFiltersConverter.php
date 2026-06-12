<?php

namespace Grafana\Foundation\Prometheus;

final class AdhocFiltersConverter
{
    public static function convert(\Grafana\Foundation\Prometheus\AdhocFilters $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Prometheus\AdhocFiltersBuilder())',
        ];
            if ($input->key !== "") {
    
        
    $buffer = 'key(';
        $arg0 =\var_export($input->key, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->operator !== "") {
    
        
    $buffer = 'operator(';
        $arg0 =\var_export($input->operator, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->values !== null && count($input->values) >= 1) {
    
        
    $buffer = 'values(';
        $tmparg0 = [];
        foreach ($input->values as $arg1) {
        $tmpvaluesarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpvaluesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

