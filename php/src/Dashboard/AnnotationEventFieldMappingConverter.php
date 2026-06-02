<?php

namespace Grafana\Foundation\Dashboard;

final class AnnotationEventFieldMappingConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\AnnotationEventFieldMapping $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\AnnotationEventFieldMappingBuilder())',
        ];
            if ($input->source !== null) {
    
        
    $buffer = 'source(';
        $arg0 ='\Grafana\Foundation\Dashboard\AnnotationEventFieldSource::fromValue("'.$input->source.'")';
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

