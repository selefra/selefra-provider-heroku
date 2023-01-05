# Table: heroku_team_members

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| identity_provider | json | X | √ |  | 
| role | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| user | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| email | string | X | √ |  | 
| federated | bool | X | √ |  | 
| two_factor_authentication | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


