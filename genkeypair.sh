# 密钥生成流程,存在第四步的原因是go使用的库需要读取pkcs#8格式私钥pem文件：

#0 check sm2 in openssl
openssl ecparam -list_curves | grep SM2

#1 生成sm2私钥: 
openssl ecparam -genkey -name SM2 -out sm2PriKey.pem

#2 sm2私钥导出公钥: 
openssl ec -in sm2PriKey.pem -pubout -out sm2PubKey.pem

#3 查看私钥:
openssl ec -in sm2PriKey.pem -text

#4 私钥pkcs#1转pkcs#8: 
openssl pkcs8 -topk8 -inform PEM -in sm2PriKey.pem -outform pem -nocrypt -out sm2PriKeyPkcs8.pem
