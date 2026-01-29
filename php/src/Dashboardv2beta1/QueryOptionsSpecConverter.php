<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class QueryOptionsSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpecBuilder())',
        ];
            if ($input->timeFrom !== null && $input->timeFrom !== "") {
    
        
    $buffer = 'timeFrom(';
        $arg0 =\var_export($input->timeFrom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxDataPoints !== null) {
    
        
    $buffer = 'maxDataPoints(';
        $arg0 =\var_export($input->maxDataPoints, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeShift !== null && $input->timeShift !== "") {
    
        
    $buffer = 'timeShift(';
        $arg0 =\var_export($input->timeShift, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryCachingTTL !== null) {
    
        
    $buffer = 'queryCachingTTL(';
        $arg0 =\var_export($input->queryCachingTTL, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->interval !== null && $input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->cacheTimeout !== null && $input->cacheTimeout !== "") {
    
        
    $buffer = 'cacheTimeout(';
        $arg0 =\var_export($input->cacheTimeout, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hideTimeOverride !== null) {
    
        
    $buffer = 'hideTimeOverride(';
        $arg0 =\var_export($input->hideTimeOverride, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeCompare !== null && $input->timeCompare !== "") {
    
        
    $buffer = 'timeCompare(';
        $arg0 =\var_export($input->timeCompare, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

