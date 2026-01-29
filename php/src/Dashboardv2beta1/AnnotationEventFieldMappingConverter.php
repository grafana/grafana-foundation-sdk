<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AnnotationEventFieldMappingConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMappingBuilder())',
        ];
            if ($input->source !== null && $input->source !== "" && $input->source !== "field") {
    
        
    $buffer = 'source(';
        $arg0 =\var_export($input->source, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== null && $input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->regex !== null && $input->regex !== "") {
    
        
    $buffer = 'regex(';
        $arg0 =\var_export($input->regex, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

