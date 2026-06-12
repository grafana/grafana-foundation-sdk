<?php

namespace Grafana\Foundation\Prometheus;

final class ScopesConverter
{
    public static function convert(\Grafana\Foundation\Prometheus\Scopes $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Prometheus\ScopesBuilder())',
        ];
            if ($input->defaultPath !== null && count($input->defaultPath) >= 1) {
    
        
    $buffer = 'defaultPath(';
        $tmparg0 = [];
        foreach ($input->defaultPath as $arg1) {
        $tmpdefaultPatharg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpdefaultPatharg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filters !== null && count($input->filters) >= 1) {
    
        
    $buffer = 'filters(';
        $tmparg0 = [];
        foreach ($input->filters as $arg1) {
        $tmpfiltersarg1 = \Grafana\Foundation\Prometheus\ScopesFiltersConverter::convert($arg1);
        $tmparg0[] = $tmpfiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

