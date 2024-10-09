<?php

namespace Grafana\Foundation\Cloudwatch;

final class CloudWatchLogsQueryConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\CloudWatchLogsQueryBuilder())',
        ];
            
    
        {
    $buffer = 'queryMode(';
        $arg0 ='\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::fromValue("'.$input->queryMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->region !== "") {
    
        
    $buffer = 'region(';
        $arg0 =\var_export($input->region, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->expression !== null && $input->expression !== "") {
    
        
    $buffer = 'expression(';
        $arg0 =\var_export($input->expression, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->statsGroups !== null && count($input->statsGroups) >= 1) {
    
        
    $buffer = 'statsGroups(';
        $tmparg0 = [];
        foreach ($input->statsGroups as $arg1) {
        $tmpstatsGroupsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpstatsGroupsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->logGroups !== null && count($input->logGroups) >= 1) {
    
        
    $buffer = 'logGroups(';
        $tmparg0 = [];
        foreach ($input->logGroups as $arg1) {
        $tmplogGroupsarg1 = \Grafana\Foundation\Cloudwatch\LogGroupConverter::convert($arg1);
        $tmparg0[] = $tmplogGroupsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->logGroupNames !== null && count($input->logGroupNames) >= 1) {
    
        
    $buffer = 'logGroupNames(';
        $tmparg0 = [];
        foreach ($input->logGroupNames as $arg1) {
        $tmplogGroupNamesarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmplogGroupNamesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

