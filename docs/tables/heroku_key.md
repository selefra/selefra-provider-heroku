# Table: heroku_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| comment | string | X | √ | Comment on the key. | 
| created_at | timestamp | X | √ | When key was created. | 
| email | string | X | √ | Deprecated. Please refer to 'comment' instead. | 
| fingerprint | string | X | √ | A unique identifying string based on contents. | 
| id | string | X | √ | Unique identifier of this key. | 
| public_key | string | X | √ | Full public_key as uploaded. | 
| updated_at | timestamp | X | √ | When key was updated. | 


