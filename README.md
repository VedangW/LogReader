# LogReader
A tool for retrieving data from JSON log files and obfuscating crucial data in the process.

### Starting the Server

To generate a random JSON file, you can use the ```generate_json.py``` file as follows:
```bash
python generate_json.py --num_entries 100
```
The ```num_entries``` flag denotes the number of JSON objects which will be generated in the file. A sample JSON file is available as ```sample.json```.

To start the server, build the package and run as follows:
```bash
sh build.sh
```

To connect to the SQL server, enter the ```openLink``` parameter in ```main.go```. 

### Using the API

Requests can be sent to the API endpoints using Postman. 

To use the API as an admin, start the URIs as ```localhost:3000/admin```.

Store a JSON entry into the DB:
```
(POST) localhost:3000/user
```

```json
{
    "username": "John Doe", 
    "order_valid": "No", 
    "user_id": 10, 
    "price": 175, 
    "phone_no": "9998887770", 
    "order_placed": "No", 
    "password": "eg1e3131ff", 
    "order_number": 1, 
    "product_weight": 10, 
    "num_installments": 6
}
```

Retrieve a JSON entry from DB based on ID:
```
(GET) localhost:3000/user/findid/<eid>
```

Retrieve JSON entries from DB based on User ID:
```
(GET) localhost:3000/user/finduser/<uid>
```

Retrieve JSON entries from DB based on validity of order (```ov``` can be ```yes``` or ```no```):
```
(GET) localhost:3000/user/findvalid/<ov>
```

Store a file into the DB:
```
(POST) localhost:3000/user
```

```json
{
	"file_path": "Go/src/github.com/user/LogReader/sample.json"
}
```
