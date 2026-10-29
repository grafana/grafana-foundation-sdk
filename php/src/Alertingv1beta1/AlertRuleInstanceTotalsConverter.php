<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class AlertRuleInstanceTotalsConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotals $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\AlertRuleInstanceTotalsBuilder())',
        ];
            if ($input->healthy !== null) {
    
        
    $buffer = 'healthy(';
        $arg0 =\var_export($input->healthy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->firing !== null) {
    
        
    $buffer = 'firing(';
        $arg0 =\var_export($input->firing, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pending !== null) {
    
        
    $buffer = 'pending(';
        $arg0 =\var_export($input->pending, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->recovering !== null) {
    
        
    $buffer = 'recovering(';
        $arg0 =\var_export($input->recovering, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->nodata !== null) {
    
        
    $buffer = 'nodata(';
        $arg0 =\var_export($input->nodata, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->error !== null) {
    
        
    $buffer = 'error(';
        $arg0 =\var_export($input->error, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

