<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class SearchResultResourceConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\SearchResultResource $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\SearchResultResourceBuilder())',
        ];
            if ($input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resource !== "") {
    
        
    $buffer = 'resource(';
        $arg0 =\var_export($input->resource, true);
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
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

