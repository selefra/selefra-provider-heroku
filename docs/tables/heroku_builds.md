# Table: heroku_builds

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| output_stream_url | string | X | √ |  | 
| slug | json | X | √ |  | 
| source_blob | json | X | √ |  | 
| stack | string | X | √ |  | 
| status | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| app | json | X | √ |  | 
| buildpacks | json | X | √ |  | 
| id | string | √ | √ |  | 
| release | json | X | √ |  | 
| user | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


