# Table: heroku_app_transfers

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| app | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| id | string | √ | √ |  | 
| owner | json | X | √ |  | 
| recipient | json | X | √ |  | 
| state | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


