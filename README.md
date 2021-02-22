个推v2 golang sdk  

目前实现了 鉴权、推送 接口
  
1、无需处理token，过期自动刷新 (务必单例使用client，否则自动刷新机制无效)  
2、request_id不传时默认使用16位shortuuid
 
