# How to deploy CoC on AWS with k8s

## Service
+ fsvc.yaml -- LoadBalancer -- expose front-end
+ bwsvc.yaml -- NodePort -- expose back-end-weather
+ bcsvc.yaml -- ClusterIP -- expose back-end-chatbox
+ msvc.yaml -- ClusterIP -- expose mysql

## Database
+ mpv.yaml -- persistent volume
+ mpvc.yaml -- persistent volume claim
+ mconfig.yaml -- ConfigMap
+ msecret.yaml -- Secret -- password
+ mysql.yaml -- database
+ exec m-dep pod
  - execute `mysql -u root -p`
  - pwd =  `12345678`
  - execute `create database weather;`

## Backend & Database Initialize
+ bwdep.yaml -- deployment -- backend-weather
+ exec bw-dep pods
  - execute `python manage.py makemigrations`
  - execute `python manage.py migrate`
+ bcdep.yaml -- deployment -- backend-chatbox

## Jobs
+ warning_cj.yaml -- job -- provide warning
+ weather_cj.yaml -- job -- provide weather
+ predict_cj.yaml -- job -- provide predict

## Frontend
+ fdep.yaml -- deployment -- front-end
