<?php

namespace Grafana\Foundation\Starsv1alpha1;

final class ResourceConverter
{
    public static function convert(\Grafana\Foundation\Starsv1alpha1\Resource $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Starsv1alpha1\ResourceBuilder())',
        ];
            if ($input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->kind !== "") {
    
        
    $buffer = 'kind(';
        $arg0 =\var_export($input->kind, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->names) >= 1) {
    
        
    $buffer = 'names(';
        $tmparg0 = [];
        foreach ($input->names as $arg1) {
        $tmpnamesarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpnamesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

