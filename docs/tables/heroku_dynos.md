# Table: heroku_dynos

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| release | json | X | √ |  | 
| size | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| command | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| type | string | X | √ |  | 
| app | json | X | √ |  | 
| attach_url | string | X | √ |  | 


