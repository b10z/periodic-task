# Periodic-Task

### 🛫 How to start the project: 

* clone the repo
* `cd` at `/periodic-task` directory
* execute `make app.start ADDR=DesiredAddressHere PORT=DesiredPortHere`. <br/> For example `make app.start ADDR=0.0.0.0 PORT=8080`
* If some or no arguments are given, the dockerfile is set to provide default values for the program to launch. The **default SERVER_ADDRESS is `0.0.0.0`** and the **default "SERVER_PORT" is `8000`**

The project is using **Makefile**, a **docker-compose.yml** (for expandability) and **three dockerfiles** for **development, testing and production builds**. 

 <br/> 

### 💼 API DOCUMENTATION: 
Using the endpoint bellow we can generate new timestamps. <br/>

```sh
GET - /ptlist 
```

#### **/ptlist - Supported Parameters**

* `period`:  **REQUIRED** <br/> _Is used to set the period (step) between the generated timestamps created from the periodic task. <br/> Currently supporting values: <br/> **"1h"** - for 1 hour, **"1d"** - for 1 day, **"1mo"** - for 1 month, **"1y"**- for 1 year_
* `tz`: **REQUIRED** <br/> _Is used to set the timezone of the timestamps.  <br/> Accepted values: <br/> A valid timezone. For example **Europe/Athens**_
* `t1`:  **REQUIRED**  <br/> _Is used to set the first timestamp (startDate) of the periodic task. <br/> Accepted values: <br/> A valid datetime timestamp. For example **20210214T204603Z**_
* `t2`:  **REQUIRED**  <br/> _Is used to set the last timestamp (endDate) of the periodic task. <br/> Accepted values: <br/> A valid datetime timestamp. For example **20210215T204603Z**_

Please note that the addition of new periods is an extremely easy process. Some comments about it can be found here **/internal/helper/timestamp_generator.go:52**

 <br/> 

### 🔦 EXAMPLES: 



<details>
<summary>Successful GET example </summary>
<br>

 Request:
 
```sh
0.0.0.0:8080/ptlist?period=1mo&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z
```

<br>

 Response:
 
`[
"20210228T220000Z",
"20210331T210000Z",
"20210430T210000Z",
"20210531T210000Z",
"20210630T210000Z",
"20210731T210000Z",
"20210831T210000Z",
"20210930T210000Z",
"20211031T220000Z"
]`
</details>


<details>
<summary> Failed GET example </summary>

<br>
 
 Request:

```sh
0.0.0.0:8080/ptlist?tz=Europe/Athens&t1=20210214T200000Z&t2=20210219T200000Z
```

<br>
 
Response:

`{
"status": "error",
"desc": "invalid period parameter"
}`
</details>




 <br/> 

### 🛠 OTHER MAKEFILE COMMANDS: 

* `make app.stop`: <br/> _This command actually runs **"docker compose stop"**_
* `make tests.generate-mock`: <br/> _This command takes as an input a .go file with an interface and generates a mock file for tests needs._ 
* `make tests.tests-all`: <br/> _This command runs all the test files inside the project and provides a coverage number for each package._  
* `make tests.test-build`: <br/> _This command is used from make tests.tests-all in order to create a test build for the application (using dockerfile.test)_ 

### 📳 TEST INFORMATION

All the test cases given in the assessment can be found tested in [timestamp_generator_test.go](powerfactors-assignment%2Finternal%2Fhelper%2Ftimestamp_generator_test.go) with good and bad paths.

TEST COVERAGE PERCENTAGES:

* `/internal/api	coverage: 100.0% of statements`
* `/internal/service	coverage: 100.0% of statements`
* `/internal/helper	coverage: 100.0% of statements`

