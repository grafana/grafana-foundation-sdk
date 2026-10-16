<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class CodeCellContentSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpecBuilder())',
        ];
            if ($input->language !== "") {
    
        
    $buffer = 'language(';
        $arg0 =\var_export($input->language, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->code !== "") {
    
        
    $buffer = 'code(';
        $arg0 =\var_export($input->code, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->highlight !== null && count($input->highlight) >= 1) {
    
        
    $buffer = 'highlight(';
        $tmparg0 = [];
        foreach ($input->highlight as $arg1) {
        $tmphighlightarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmphighlightarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->annotation !== null && $input->annotation !== "") {
    
        
    $buffer = 'annotation(';
        $arg0 =\var_export($input->annotation, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

