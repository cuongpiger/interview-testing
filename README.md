# 1. Standard questions:
## 1.1. Describe your experience level with Python
* Web
  * Using `django` to build simple server-side websites, such as a **Blog website**, **Naive version of Google Classroom website** _(people can post their topic and everyone can leave their comments to discuss that topic)_.
  * Using `djangorestframework` to build simple CRUD API.
  * Using `ibm-db` to connect to IBM DB2 database.
* Data Science
  * Using `numpy`, `pandas`,... for data cleaning and data analysis. I used to use them to count the number of records from data that I read from the CSV files, find min and max values on quantitative and qualitative data,...
  * Using `matplotlib`, `seaborn`, `plotly`,... to draw plots, visualize data, such as **box plot**, **bar chart**, **scatter plot**, **line plot**, **confusion matrix**,...
  * Applying traditional machine learning algorithms, such as **Linear Regression**, **Decision Tree**, **SVM**,... to build predictive and classification models, like **sentiment models**, **house price models**,...
* Deep Learning
  * Applying `tensorflow` to solve deep learning problems. I used to build a **sentiment model** to predict whether customers' comments are positive or negative in **LSTM network**, image generator of furniture, animals, vehicles in **DC-GANs network**.
  * Setting up **NVIDIA cuDNN** on Ubuntu 20.04 distro to optimize time complexity of Deep Neural Network Models of `tensorflow`.
* Big Data
  * Setting up a small cluster consists mainly of two physical machines:
    * The first machine runs **HDFS service**, **one master node** and **one worker node**.
    * The second machine runs **one worker node**.
  * Using `spark`, `graphx` basically to analyze data, apply traditional machine learning algorithms like **logistic regression**, **k-mean**, **BFS**,... to perform distributed computing on data which reading from HDFS.
  * Reading Twitter document to build **spark-streaming** service, counting number of tweets based on **hashtags**. We need to register a Twitter Developer Account to do this problem.
* Data crawling
  * Using `seleninum` and `scrapy` to crawl data from dynamic websites using AJAX, or modern front-end frameworks of ReactJS, VueJS,... such as [https://shopee.vn](https://shopee.vn/), [https://vietnamnews.vn](https://vietnamnews.vn/),...
* Others
  * `pygame`, `PyQT5`,...

## 1.2. Describe your experience level with Git
* I use Git on projects of multiple members. I create a branch named `development`. When any member of the team is given a task, they have to create a new branch named \<name of member\>_\<task name\>, they code all functions, and screens on their own branch, then **commit** and **push** to host, create a **pull request**. The leader reads the **pull request** and decides whether or not to merge this branch with the `development` branch. The `master` branch is the production stage, when a large function or a screen is complete at the `development` branch, it will be merged into the `master` branch. This process is repeated until the project is completed.
* When one member needs to code a new function or a new screen, they have to **pull** from the `development` branch and create a new branch then switch to this new branch.

# 2. Build a simple CRUD API in Python Flask/falcon.
## 2.1. Setting up the environment
* **Operating system**: Ubuntu 20.04 distro.
* **Docker 20.10.12** and **docker-compose 1.26.2**.
* You can check whether or not **docker** and **docker-compose** are installed on your machine by these two followed commands.<br>
  ![](images/01.png)

## 2.2. Download the source code
### 2.2.1. Git
* If Git is installed on your machine, you can run this command to clone the entire source code to your local machine.
  ```bash
  git clone https://github.com/cuongpiger/thuocsi.git
  ```

### 2.2.2. Zip
* You can also download the source code to your local machine as a zip file, and then unzip the source code directory.

## 2.3. Running
* Go to inside the `thuocsi/code/` directory.
* I am currently using the image of Python 3.6 to run the Django web server. The dependencies are listed in the `requirements.txt` file.
* Open the `Dockerfile` file to see the details.
  ![](images/02.png)

* The `docker-compose.yml` is used to tell docker engine to know how to run the entire application accurately.
  ![](images/03.png)
* Following the `docker-compose.yml` file. I am currently using **Postgres 11** as my main database.
* You can see three services: **web**, **db** and **migration**. There are some constraints here. The **web** service only runs if the **db** and **migration** services were run successfully. But the **migration** service only runs if the **db** service was run successfully. So the running order of the services is **db** -> **migration** -> **web**.
* Because Django uses ORM architecture to map the models to the under database. So to make sure that the entire application runs correctly. We need to perform `migrate` command before `runserver` command to surely the models of Django and datatables of Postgres are matched to each other.
* Open the terminal inside the `thuocsi/code/` directory and then run the below command. It will take a few minutes to complete. _(Surely that you installed Docker and Docker-compose on your machine)_
  ```bash
  docker-compose up -d --build
  ```
  ![](images/04.png)
  ![](images/05.png)
* To make sure that the **Django web server container** and **Posrgres database container** are running successfully. Running `docker ps` command to check.
  ![](images/06.png)
* You can see, there are two running containers. The first one is our Django web server and the second one is our Postgres database.

## 2.4. Demo CRUD on `Customer` model
* Firstly, we need to register a new user account. We will send a POST request to the URL [http://localhost:8000/api/v1/auth/users/create/](http://localhost:8000/api/v1/auth/users/create/) with two fields `username` and `password`.
* Remembering that you have to set the `Content-type=application/json` at the **Headers** tab.
  ![](images/08.png)
* Following the below image. I create an `admin` user with the password `qaswedfr`.
  ![](images/07.png)

* Now, we need to get the **access token** and **refresh token** in logging the `admin` user. Go to the URL [localhost:8000/api/v1/auth/token/](localhost:8000/api/v1/auth/token/) and send a POST request in the `admin` user.
  ![](images/09.png)

* To get all `Customer` in the Posrgres database. Send a GET request to the URL [localhost:8000/api/v1/customer/](localhost:8000/api/v1/customer/). But before sending the request, we need to attach the `access token` at the **Authorization** tab.
  ![](images/10.png)
  ![](images/11.png)

* In the case, you get this error response. It means that your `access token` is expired and you need to use your `refresh token`, send a POST request to the URL [localhost:8000/api/v1/auth/token/refresh/](localhost:8000/api/v1/auth/token/refresh/) to generate a new `access token` (Remembering you have to set `Content-Type=application/json` at **Header** tab).
  ![](images/12.png)
  ![](images/13.png)

* To get detail of a `Customer` based on her/his `id`, send a GET request to the URL [localhost:8000/api/v1/customer/\<id\>/](localhost:8000/api/v1/customer/<id>)
  ![](images/14.png)

* To create a new `Customer` object, send a POST to the URL [localhost:8000/api/v1/customer/](localhost:8000/api/v1/customer/) _(Surely that you set up Bearer Token at **Authorization** tab and Content-Type=application/json at **Header** tab)_.
  ![](images/15.png)

* Send a GET request to the URL [localhost:8000/api/v1/customer/4/](localhost:8000/api/v1/customer/4/) to confirm.
  ![](images/16.png)

* Send a PUT request to the [localhost:8000/api/v1/customer/4/](localhost:8000/api/v1/customer/4/) to update the `Customer` who has `id` 4.
  ![](images/17.png)
* Check the `Customer` who has `id` 4 has been updated successfully.
  ![](images/18.png)

* Finally, I will delete the `Customer` who has `id` 4. Send a DELETE request to the URL [localhost:8000/api/v1/customer/4/](localhost:8000/api/v1/customer/4/).
  ![](images/19.png)

* Send a GET request to the URL [localhost:8000/api/v1/customer/4/](localhost:8000/api/v1/customer/4/) to confirm that this `Customer` has been deleted successfully.
  ![](images/20.png)

* Using Docker and Postgres Bash to discover our database.
  ![](images/21.png)

* Shutdown all Docker containers.
  ```bash
  docker-compose down
  ```
  ![](images/22.png)