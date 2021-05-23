#pouch network create --driver bridge --subnet 172.16.1.0/24 GLOBAL_RESOURCE
#pouch network create --driver bridge --subnet 10.0.0.0/24 d2fd841db19011e
#pouch run -d --name GLOBAL_CONSUL --net GLOBAL_RESOURCE ConsulImage top -b
#pouch run -d --name GLOBAL_ZIPKIN --net GLOBAL_RESOURCE ZipkinImage top -b
#pouch run -d --name 0edef0abb19f11eb9274_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
#pouch run -d --name 24391683b27111eba6c5_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
#pouch run -d --name 28d20b0fbad111eb943d_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
#pouch run -d --name ad7fe219bb8611eb9c0b_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
#pouch run -d --name ad7fe219bb8611ebgrev_interpidtjuniversity.miniselfop.pre --net d2fd841db19011e JavaImage top -b
#pouch run -d --name ad7xf679bb8611ebgrev_interpidtjuniversity.miniselfop.pre --net d2fd841db19011e JavaImage top -b


#
#pouch exec GLOBAL_CONSUL sh
#/bin/chmod 777 standalone-consul.sh
#./standalone-consul.sh GLOBAL_CONSUL 172.16.1.2 GLOBAL_CONSUL
#exit

#pouch exec GLOBAL_ZIPKIN sh
#/bin/chmod 777 standalone-zipkin.sh
#./standalone-zipkin.sh
#exit