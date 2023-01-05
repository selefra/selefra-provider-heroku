# Table: heroku_oauth_clients

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| id | string | √ | √ |  | 
| ignores_delinquent | bool | X | √ |  | 
| name | string | X | √ |  | 
| redirect_uri | string | X | √ |  | 
| secret | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


