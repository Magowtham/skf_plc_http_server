# SKF HTTP SERVER API

## 1. Database Initialization Process

```bash
https://skfplc.http.vsensetech.in/admin/database/init
```

### HTTP Method → GET

### HTTP Responses

json message with http status

json message format

```json
{"message":"__server_status_message__"}
```

HTTP status codes

400 → Bad Request (User side error)

500 → Internal server error (Server side error)

200 → ok (Operation successful without any error)

## 2. Creating Admin Account

```bash
https://skfplc.http.vsensetech.in/root/create/admin
```

### HTTP Method → POST

Request Body format

```json
{
    "email":"magowtham7@gmail.com",
    "password":"Vsense@2024"
}
```

## 3. Deleting Admin Account

```bash
https://skfplc.http.vsensetech.in/root/delete/admin/__ADMIN_ID__
```

### HTTP Method → DELETE

## 4. Admin Login

```bash
https://skfplc.http.vsensetech.in/login/admin
```

### HTTP Method → POST

Request body format

```json
{
    "email":"magowtham7@gmail.com",
    "password":"Vsense@2024"
}
```

HTTP only cookie will be get setted in the cookie section of the browser

## 5. Create User

```bash
https://skfplc.http.vsensetech.in/admin/create/user
```

### HTTP Method → POST

Request body format

```json
{
    "label":"rice mill",
    "email":"magowtham7@gmail.com",
    "password":"Vsense@2024"
}
```

## 6. Delete User

```bash
https://skfplc.http.vsensetech.in/admin/delete/user/__USER_ID__
```

### HTTP Method → DELETE

## 7. Get Users

```bash
https://skfplc.http.vsensetech.in/admin/users
```

### HTTP Method → GET

HTTP Response format

```json
{
"users":[
			{
			"user_id":"3545e4c1-529f-4d54-83f5-5f76cb6411d1",
			"email":"magowtham7@gmail.com",
			"label":"rice mill"
			},
		]
}
```

## 8. Create PLC

```bash
https://skfplc.http.vsensetech.in/admin/create/plc/__USER_ID__
```

### HTTP Method → POST

HTTP Request body format

```json
{
    "label":"rice mill owner plc",
    "plc_id":"vs24skf010"
}
```

## 9. Delete PLC

```bash
https://skfplc.http.vsensetech.in/admin/delete/plc/__PLC_ID__
```

### HTTP Method → DELETE

## 10. Get Plcs

```bash
https://skfplc.http.vsensetech.in/admin/plcs/__USER_ID__
```

### HTTP Method → GET

HTTP Response format

```json
	{
		"plcs":[
			{
				"plc_id":"vs24skf010",
				"user_id":"3545e4c1-529f-4d54-83f5-5f76cb6411d1",
				"label":"rice mill owner plc"
			}
		]
	}
```

## 11. Create Drier

```bash
https://skfplc.http.vsensetech.in/admin/create/drier/__PLC_ID__
```

### HTTP Method → POST

HTTP Request format

```json
{
    "label":"paddy drier2"
}
```

## 12. Delete Drier

```bash
https://skfplc.http.vsensetech.in/admin/delete/drier/__PLC_ID__/__DRIER_ID__
```

### HTTP Method → DELETE

## 13. Get Driers

```bash
https://skfplc.http.vsensetech.in/admin/driers/__PLC_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
	"driers":[
		{
			"drier_id":"03633bbe-fb8b-4e92-a64d-20d6059dbc0b",
			"plc_id":"vs24skf010",
			"recipe_step_count":"0",
			"label":"paddy drier2"
		}
	]
}

```

## 14. Create Register

```bash
https://skfplc.http.vsensetech.in/admin/create/register/__PLC_ID__/__DRIER_ID__
```

### HTTP Method → POST

HTTP Request format

```json
{
    "reg_address":"101",
    "reg_type":"rt_tp",
    "label":"real time temperature"
}
```

## 15.  Delete Register

```bash
https://skfplc.http.vsensetech.in/admin/delete/register/__PLC_ID__/__DRIER_ID__/__REG_ADDRESS__/__REG_TYPE__
```

### HTTP Method → DELETE

## 16  Get Registers

```bash
https://skfplc.http.vsensetech.in/admin/registers/__PLC_ID__/__DRIER_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
	"registers":[
		{
			"reg_address":"101",
			"reg_type":"rt_tp",
			"label":"real time temperature",
			"drier_id":"03633bbe-fb8b-4e92-a64d-20d6059dbc0b",
			"value":"0",
			"last_update_timestamp":"2024-11-05T08:36:03.780595Z"
		}
	]
}

```

## 17.  Create Register Type

```bash
https://skfplc.http.vsensetech.in/admin/create/register_type
```

### HTTP Method → POST

HTTP Request fromat

```json
{
    "type":"stptm",
    "label":"step time"
}
```

## 18.  Delete Register Type

```bash
https://skfplc.http.vsensetech.in/admin/delete/register_type/__REGISTER_TYPE__
```

### HTTP Method → DELETE

## 19.  Give Access to User

```bash
https://skfplc.http.vsensetech.in/admin/give/user/access
```

### HTTP Method → POST

HTTP Request format

```json
{
    "user_id":"62ae57e7-f67b-42ca-8e3d-8b3d14b17591",
    "password":"Vsense@2024"
}
```

## 20.  User Login

```bash
https://skfplc.http.vsensetech.in/user/login
```

### HTTP Method → POST

HTTP Request format

```json
{
    "email":"test@gmail.com",
    "password":"Vsense@2024"
}
```

HTTP Response format

```json
{
	"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNzYxOTAxNDIyLCJsYWJlbCI6InJpY2UgbWlsbCIsInVzZXJfaWQiOiI1ZDA0YjhiMy1kNWNjLTRkZTEtOThkNS01YTZjODE0ZTdhOGEifQ.96jO59vien5IPWnjm2ELtyGyfN1MDeB3TUnPnL-lU9U"
}

```

## 21.  Get User Driers

```bash
https://skfplc.http.vsensetech.in/user/driers/__USER_ID__
```

### HTTP Method → GET

HTTP Request body format

```json
{
	"driers":[
		{
			"drier_id":"03633bbe-fb8b-4e92-a64d-20d6059dbc0b",
			"plc_id":"vs24skf010","recipe_step_count":"0",
			"label":"paddy drier2"
		},
		{
			"drier_id":"1f81751b-fbad-409c-9a6a-27d20a59a0a8",
			"plc_id":"vs24skf010",
			"recipe_step_count":"0",
			"label":"paddy drier2"
		},
		{
			"drier_id":"ad33ab19-164b-44f6-86ab-d9a3a0936a44",
			"plc_id":"vs24skf010",
			"recipe_step_count":"0",
			"label":"paddy drier2"
		}
	]
}

```

## 22.  Get Register Types For Form

```bash
https://skfplc.http.vsensetech.in/admin/register/types/__PLC_ID__/__DRIER_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
	"reg_types":[
		{
			"type":"tm",
			"label":"step time"
		}
	]
}

```

## 23.  Get Recipe Step Count

```bash
https://skfplc.http.vsensetech.in/user/recipe/step/count/__DRIER_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
	"recipe_step_count":0
}
```

## 24.  Get Drier Prev Statuses

```bash
https://skfplc.http.vsensetech.in/user/drier/statuses/__PLC_ID__/__DRIER_ID__
```

### HTTP Method → GET

HTTP Response format

```json
{
	"statuses":[
		{"reg_type":"st_bl_trp","reg_value":"0"},
		{"reg_type":"st_el_trp","reg_value":"0"},
		{"reg_type":"st_rt_trp","reg_value":"0"},
		{"reg_type":"st_bl_rn","reg_value":"1"},
		{"reg_type":"st_el_rn","reg_value":"1"},
		{"reg_type":"st_rt_rn","reg_value":"1"}
	]
}

```

## 25.  Create User Feedback

```bash
https://skfplc.http.vsensetech.in/user/feedback/__USER_ID__
```

### HTTP Method → POST

HTTP Request format

```json
{
    "feedback":"your app is perfect"
}
```
