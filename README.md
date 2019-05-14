# deploy-cli

[ops-srv proto](https://github.com/weeon/proto/blob/master/ops/ops.proto)

## Usage

env need:

* OPS_SRV_ADDR
* OPS_TOKEN
* OPS_WORKLOAD_ID



### gitlab-ci.yaml example

```
deploy:
  image: orvice/deploy-cli
  stage: deploy
  retry: 2
  tags:
    - docker
  script:
    - /app/bin/deploy-cli -a deploy
```

### Telegram notify

set env to enable telegra notify:

* TELEGRAM_TOKEN
* TELEGRAM_CHATID