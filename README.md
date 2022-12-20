* On the FE side, I use this repo to implement the Web Interface by ReactJS: [https://github.com/phyohtetarkar/react-ecommerce-template](https://github.com/phyohtetarkar/react-ecommerce-template)
# a. Testing
* I use Python to test my entire API. Follow this link to catch it. [unit_test.ipynb](./unit_test.ipynb)

# b. Structure of project
* I implement the project by using the following structure:<br>
  ![](./img/01.png)
* The source code of BE site was implemented in the **be** directory.<br>
  ![](./img/02.png)
* I use Clean Architecture to implement the project.<br>
  ![](./img/03.png)

# c. Frameworks, and Tools
* Golang 1.18.
* Gin Gonic: [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
  * Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance -- up to 40 times faster.
* GORM: [https://gorm.io](https://gorm.io/)
  * The fantastic ORM library for Golang, aims to be developer friendly. You can easly build your structure of database without direactly interacting with SQL.

* Docker, and docker-compose: I use it in my project because I want you to run the project easily without installing any dependencies.