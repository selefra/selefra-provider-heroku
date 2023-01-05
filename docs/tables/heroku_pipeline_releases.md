# Table: heroku_pipeline_releases

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| version | int | X | √ |  | 
| description | string | X | √ |  | 
| id | string | √ | √ |  | 
| status | string | X | √ |  | 
| user | json | X | √ |  | 
| output_stream_url | string | X | √ |  | 
| slug | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| addon_plan_names | string_array | X | √ |  | 
| app | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| current | bool | X | √ |  | 


