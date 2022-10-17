# 日志收集

> 本项目主要使用 EFK 进行日志收集

- Elasticsearch
- Filebeat
- Kibana

> 安装方法

- [Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)
- [Filebeat](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-installation-configuration.html)
- [Kibana](https://www.elastic.co/guide/en/kibana/current/docker.html)

> 测试环境快速安装方法 \
> Docker

```shell
# es
docker run --name es01 --net elastic -p 9200:9200 -p 9300:9300 -e ES_JAVA_OPTS="-Xms1g -Xmx1g" -m 3g -it docker.elastic.co/elasticsearch/elasticsearch:8.4.3

# kibana
docker run --name kib-01 --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:8.4.3
```

> Filebeat

修改 filebeat-deploy.yaml， 位置 simple-admin-core/deploy/k8s/log-collection/filebeat/
> 可以添加 log 来源位置，默认有 core

```yaml
 - type: log
      paths:
        - /home/data/logs/core/*/*.log
```

> 配置环境变量

```yaml
          env:
            - name: ELASTICSEARCH_HOST   
              value: "192.168.50.216"  # ES的地址
            - name: ELASTICSEARCH_PORT
              value: "9200"  # ES的端口
            - name: ELASTICSEARCH_USERNAME
              value: elastic # ES 的用户名
            - name: ELASTICSEARCH_PASSWORD
              value: UQ==CXXjw47bK_I13*f1 # 密码
            - name: ELASTICSEARCH_CA_FINGERPRINT
              value: 8d6aed6bba745f2f0aaa46f628e3124c82ae6727c1f5e207e3d821ffeefb5e5e # 信任的CA指纹
            - name: ELASTIC_CLOUD_ID 
              value: # 云 ID， 可选
            - name: ELASTIC_CLOUD_AUTH 
              value: # 云 Token， 可选
```

> 然后使用脚本部署 filebeat

```shell
# 进入 simple-admin-core/deploy/k8s/log-collection/filebeat
kubectl apply -f filebeat-deploy.yaml
```

> 效果展示

![Pic](../../assets/kibana.png)