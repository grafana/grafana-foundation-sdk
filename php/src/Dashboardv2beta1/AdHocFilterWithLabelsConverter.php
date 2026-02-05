<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AdHocFilterWithLabelsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabels $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabelsBuilder())',
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
            if ($input->keyLabel !== null && $input->keyLabel !== "") {
    
        
    $buffer = 'keyLabel(';
        $arg0 =\var_export($input->keyLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->valueLabels !== null && count($input->valueLabels) >= 1) {
    
        
    $buffer = 'valueLabels(';
        $tmparg0 = [];
        foreach ($input->valueLabels as $arg1) {
        $tmpvalueLabelsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpvalueLabelsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->forceEdit !== null) {
    
        
    $buffer = 'forceEdit(';
        $arg0 =\var_export($input->forceEdit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->origin !== null) {
    
        
    $buffer = 'origin(';
        $arg0 =\var_export($input->origin, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->condition !== null && $input->condition !== "") {
    
        
    $buffer = 'condition(';
        $arg0 =\var_export($input->condition, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

