<?php

namespace Grafana\Foundation\Bigquery;

final class DataqueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Bigquery\Dataquery);
        $calls = [
            '(new \Grafana\Foundation\Bigquery\DataqueryBuilder())',
        ];
            if ($input->dataset !== null && $input->dataset !== "") {
    
        
    $buffer = 'dataset(';
        $arg0 =\var_export($input->dataset, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->table !== null && $input->table !== "") {
    
        
    $buffer = 'table(';
        $arg0 =\var_export($input->table, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->project !== null && $input->project !== "") {
    
        
    $buffer = 'project(';
        $arg0 =\var_export($input->project, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'format(';
        $arg0 ='\Grafana\Foundation\Bigquery\QueryFormat::fromValue("'.$input->format.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->rawQuery !== null) {
    
        
    $buffer = 'rawQuery(';
        $arg0 =\var_export($input->rawQuery, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->rawSql !== "") {
    
        
    $buffer = 'rawSql(';
        $arg0 =\var_export($input->rawSql, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->location !== null && $input->location !== "") {
    
        
    $buffer = 'location(';
        $arg0 =\var_export($input->location, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->partitioned !== null) {
    
        
    $buffer = 'partitioned(';
        $arg0 =\var_export($input->partitioned, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->partitionedField !== null && $input->partitionedField !== "") {
    
        
    $buffer = 'partitionedField(';
        $arg0 =\var_export($input->partitionedField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->convertToUTC !== null) {
    
        
    $buffer = 'convertToUTC(';
        $arg0 =\var_export($input->convertToUTC, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sharded !== null) {
    
        
    $buffer = 'sharded(';
        $arg0 =\var_export($input->sharded, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryPriority !== null) {
    
        
    $buffer = 'queryPriority(';
        $arg0 ='\Grafana\Foundation\Bigquery\QueryPriority::fromValue("'.$input->queryPriority.'")';
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
            if ($input->editorMode !== null) {
    
        
    $buffer = 'editorMode(';
        $arg0 ='\Grafana\Foundation\Bigquery\EditorMode::fromValue("'.$input->editorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sql !== null) {
    
        
    $buffer = 'sql(';
        $arg0 = \Grafana\Foundation\Bigquery\SQLExpressionConverter::convert($input->sql);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== null && $input->refId !== "") {
    
        
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
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Common\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

