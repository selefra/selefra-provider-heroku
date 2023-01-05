# Table: heroku_pipeline_builds

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| slug | json | X | √ |  | 
| source_blob | json | X | √ |  | 
| status | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| buildpacks | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| output_stream_url | string | X | √ |  | 
| stack | string | X | √ |  | 
| user | json | X | √ |  | 
| app | json | X | √ |  | 
| id | string | √ | √ |  | 
| release | json | X | √ |  | 


