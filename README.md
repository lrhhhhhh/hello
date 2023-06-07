
## intro
v0.0.4
- 基础组件的安装：
- 项目代码与CI流水线
- CD repo： 


对项目代码的修改，触发流水线，流水线进行镜像构建与推送，然后修改CD repo，触发ArgoCD 


文件说明：
- 用到一些Tekton hub上的Task，这些yaml文件已经存放在基础组件repo中, 如果想显得更加合理，可以把task变成项目级的
- 


要求基础组件安装了tekton trigger和Interceptor

我们将trigger的配置，放在pipeline中（项目级的）


连通性：
tekton执行结束后，如果需要知道ArgoCD的状态，并且要保证GitOps的理念，那么需要ArgoCD返回一个信息。

如果tekton登录到argocd中，那么这个过程将不够符合GitOps的理念。 

## 生成密钥
```
apiVersion: v1
kind: Secret
metadata:
  name: ssh-secret
data:
  id_rsa: <YOUR_ID_RSA>
  known_hosts: ''
  config: ''

---
apiVersion: v1
kind: Secret
metadata:
  name: webhook-secret
type: Opaque
stringData:
  secretToken: ""
```
cat ~/.ssh/id_rsa | base64 -w0  



