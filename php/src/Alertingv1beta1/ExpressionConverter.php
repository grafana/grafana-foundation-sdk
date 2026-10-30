<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class ExpressionConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\Expression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\ExpressionBuilder())',
        ];
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->relativeTimeRange !== null) {
    
        
    $buffer = 'relativeTimeRange(';
        $arg0 = \Grafana\Foundation\Alertingv1beta1\RelativeTimeRangeConverter::convert($input->relativeTimeRange);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasourceUID !== "") {
    
        
    $buffer = 'datasourceUID(';
        $arg0 =\var_export($input->datasourceUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->model !== null) {
    
        
    $buffer = 'model(';
        $arg0 =\var_export($input->model, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->source !== null) {
    
        
    $buffer = 'source(';
        $arg0 =\var_export($input->source, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

