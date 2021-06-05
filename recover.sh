pouch network create --driver bridge --subnet 172.16.1.0/24 GLOBAL_RESOURCE
pouch network create --driver bridge --subnet 10.0.0.0/24 d2fd841db19011e
pouch run -d --name GLOBAL_CONSUL --net GLOBAL_RESOURCE ConsulImage top -b
pouch run -d --name GLOBAL_ZIPKIN --net GLOBAL_RESOURCE ZipkinImage top -b
pouch run -d --name 0edef0abb19f11eb9274_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
pouch run -d --name 24391683b27111eba6c5_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
pouch run -d --name 28d20b0fbad111eb943d_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
pouch run -d --name ad7fe219bb8611eb9c0b_interpidtjuniversity.miniselfop.dev --net d2fd841db19011e JavaImage top -b
pouch run -d --name ad7fe219bb8611ebgrev_interpidtjuniversity.miniselfop.pre --net d2fd841db19011e JavaImage top -b
pouch run -d --name ad7xf679bb8611ebgrev_interpidtjuniversity.miniselfop.pre --net d2fd841db19011e JavaImage top -b
pouch run -d --name ad7xf679berge11ebgrev_interpidtjuniversity.miniselfop.prod --net d2fd841db19011e JavaImage top -b
pouch run -d --name g5d3gp79berge11ebgrev_interpidtjuniversity.miniselfop.prod --net d2fd841db19011e JavaImage top -b


#
#pouch exec GLOBAL_CONSUL sh
#/bin/chmod 777 standalone-consul.sh
#./standalone-consul.sh GLOBAL_CONSUL 172.16.1.2 GLOBAL_CONSUL
#exit

#pouch exec GLOBAL_ZIPKIN sh
#/bin/chmod 777 standalone-zipkin.sh
#./standalone-zipkin.sh
#exit


#pouch stop GLOBAL_CONSUL
#pouch stop GLOBAL_ZIPKIN
#pouch stop 0edef0abb19f11eb9274_interpidtjuniversity.miniselfop.dev
#pouch stop 24391683b27111eba6c5_interpidtjuniversity.miniselfop.dev
#pouch stop 28d20b0fbad111eb943d_interpidtjuniversity.miniselfop.dev
#pouch stop ad7fe219bb8611eb9c0b_interpidtjuniversity.miniselfop.dev
#pouch stop ad7fe219bb8611ebgrev_interpidtjuniversity.miniselfop.pre
#pouch stop ad7xf679bb8611ebgrev_interpidtjuniversity.miniselfop.pre
#pouch stop ad7xf679berge11ebgrev_interpidtjuniversity.miniselfop.prod
#pouch stop g5d3gp79berge11ebgrev_interpidtjuniversity.miniselfop.prod

#pouch rm GLOBAL_CONSUL
#pouch rm GLOBAL_ZIPKIN
#pouch rm 0edef0abb19f11eb9274_interpidtjuniversity.miniselfop.dev
#pouch rm 24391683b27111eba6c5_interpidtjuniversity.miniselfop.dev
#pouch rm 28d20b0fbad111eb943d_interpidtjuniversity.miniselfop.dev
#pouch rm ad7fe219bb8611eb9c0b_interpidtjuniversity.miniselfop.dev
#pouch rm ad7fe219bb8611ebgrev_interpidtjuniversity.miniselfop.pre
#pouch rm ad7xf679bb8611ebgrev_interpidtjuniversity.miniselfop.pre
#pouch rm ad7xf679berge11ebgrev_interpidtjuniversity.miniselfop.prod
#pouch rm g5d3gp79berge11ebgrev_interpidtjuniversity.miniselfop.prod
#
#pouch network remove GLOBAL_RESOURCE
#pouch network remove d2fd841db19011e
