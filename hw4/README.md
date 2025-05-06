# HW4

使用 Dockerfile 建置一個 flask app

## 本機建置

```bash
docker login
docker build -t rulerchen/2025cloud:r1 .
docker push rulerchen/2025cloud:r1
```

## 本機執行

```bash
docker build -t rulerchen/2025cloud:r1 .
docker run -d -p 8080:80 rulerchen/2025cloud:r1
```

See the result at [http://localhost:8080](http://localhost:8080)

## 設計流程

當 push 到 main branch 時，會自動建置 docker image 並 push 到 docker hub 上，image tag 選擇為 github.sha 的前 8 碼，這樣可以確保每次 push 都會有一個新的 image tag，並且可以追蹤到是哪一次的 commit。

## Link List

- github repo: https://github.com/RulerChen/NTUCS-CNAD/tree/main/hw4
- docker hub: https://hub.docker.com/repository/docker/rulerchen/2025cloud/general
- success actions : https://github.com/RulerChen/NTUCS-CNAD/actions/runs/14853228255
- failure actions : https://github.com/RulerChen/NTUCS-CNAD/actions/runs/14853261888
- failure pr : https://github.com/RulerChen/NTUCS-CNAD/pull/7