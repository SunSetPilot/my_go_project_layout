## router
路由注册目录，路由注册需要分组，总入口放在在router.go，里面只有一个函数RegisterRouter，在API服务被初始化时调用。RegisterRouter中一般情况下不注册具体的路由，只创建路由组，每个路由组对应本包下的一个文件，在该文件中注册具体的路由信息