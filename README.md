# periodic-task

### 🛫 How to start the project: 

* clone the repo
* **cd** at **/periodic-task** directory
* execute **make app.start ADDR=DesiredAddressHere PORT=DesiredPortHere**. <br/> For example **make app.start ADDR=0.0.0.0 PORT=8080**
* If some or no arguments are given, the dockerfile is set to provide default values for the program to launch. The **default SERVER_ADDRESS is "0.0.0.0"** and the **default "SERVER_PORT" is "8000"**

The project is using **Makefile**, a **docker-compose.yml** (for expandability) and **three dockerfiles** for **development, testing and production builds**. 

 <br/> 

### 💼 API DOCUMENTATION: 
Using the endpoint bellow we can generate new timestamps. <br/>

```sh
GET - /plist 
```

#### **/plist - Supported Parameters**

* **period**:  **REQUIRED** <br/> Is used to set the period (step) between the generated timestamps created from the periodic task. <br/> Currently supporting values: <br/> **"1h"** - for 1 hour, **"1d"** - for 1 day, **"1mo"** - for 1 month, **"1y"**- for 1 year
* **tz**: **REQUIRED** <br/> Is used to set the timezone of the timestamps.  <br/> Accepted values: <br/> A valid timezone. For example **Europe/Athens**
* **t1**:  **REQUIRED**  <br/> Is used to set the first timestamp (startDate) of the periodic task. <br/> Accepted values: <br/> A valid datetime timestamp. For example **20210214T204603Z**
* **t2**:  **REQUIRED**  <br/> Is used to set the last timestamp (endDate) of the periodic task. <br/> Accepted values: <br/> A valid datetime timestamp. For example **20210215T204603Z**

Please note that the addition of new periods is an extremely easy process. Some comments about it can be found here **/internal/task.go:48**

 <br/> 

### 🔦 EXAMPLES: 



<details>
<summary>Successful GET example </summary>
<br>
0.0.0.0:8080/ptlist?tz=Europe/Athens&t1=20210214T200000Z&t2=20210219T200000Z&period=1d

`[
"20210214T200000Z",
"20210215T200000Z",
"20210216T200000Z",
"20210217T200000Z",
"20210218T200000Z"
]`
</details>


<details>
<summary> Failed GET example </summary>
<br>
0.0.0.0:8080/ptlist?tz=Europe/Athens&t1=20210214T200000Z&t2=20210219T200000Z&period=INVALID_PERIOD

`{
"status": "error",
"desc": "error while generating the timestamps"
}`
</details>




 <br/> 

### 🛠 OTHER MAKEFILES COMMANDS: 

* **make app.stop**: <br/> This command actually runs **"docker compose stop"**
* **make tests.generate-mock**: <br/> This command takes as an input a .go file with an interface and generates a mock file for tests needs. 
* **make tests.tests-all**: <br/> This command runs all the test files inside the project and provides a coverage number for each package.  
* **make tests.test-build**: <br/> This command is used from make tests.tests-all in order to create a test build for the application (using dockerfile.test) 



