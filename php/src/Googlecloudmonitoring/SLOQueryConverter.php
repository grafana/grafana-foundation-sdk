<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class SLOQueryConverter
{
    public static function convert(\Grafana\Foundation\Googlecloudmonitoring\SLOQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\SLOQueryBuilder())',
        ];
            if ($input->projectName !== "") {
    
        
    $buffer = 'projectName(';
        $arg0 =\var_export($input->projectName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->perSeriesAligner !== null && $input->perSeriesAligner !== "") {
    
        
    $buffer = 'perSeriesAligner(';
        $arg0 =\var_export($input->perSeriesAligner, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alignmentPeriod !== null && $input->alignmentPeriod !== "") {
    
        
    $buffer = 'alignmentPeriod(';
        $arg0 =\var_export($input->alignmentPeriod, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->selectorName !== "") {
    
        
    $buffer = 'selectorName(';
        $arg0 =\var_export($input->selectorName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->serviceId !== "") {
    
        
    $buffer = 'serviceId(';
        $arg0 =\var_export($input->serviceId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->serviceName !== "") {
    
        
    $buffer = 'serviceName(';
        $arg0 =\var_export($input->serviceName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sloId !== "") {
    
        
    $buffer = 'sloId(';
        $arg0 =\var_export($input->sloId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sloName !== "") {
    
        
    $buffer = 'sloName(';
        $arg0 =\var_export($input->sloName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->goal !== null) {
    
        
    $buffer = 'goal(';
        $arg0 =\var_export($input->goal, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lookbackPeriod !== null && $input->lookbackPeriod !== "") {
    
        
    $buffer = 'lookbackPeriod(';
        $arg0 =\var_export($input->lookbackPeriod, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

