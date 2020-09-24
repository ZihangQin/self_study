//require引用包-->引用包加速引用自己的一个特殊功能
var http = require('http');//加载http核心模块
var os = require('os')
var sever = http.createServer();//创建服务器

//request请求时间处理函数，需要两个函数
//	request请求对象
//		请求对象可以用来获取客户端的一些请求信息
//	response响应对象
//		响应对象可以用来给客户端发送相应消息
sever.on('request',function(request,response){
	console.log('收到了请求，请求路径是：'+ request.url);
	console.log('请求我的客户端地址是：',request.socket.remoteAddress, request.socket.remotePort);
// Content用来告知服务器我给你发送的数据内容是什么类型
// text/plain 就是普通文本  text/html是HTML文本
	response.writeHeader(200,{"Content-Type":"text/html;charset=utf-8"});

	var url = request.url;
	if (url === '/') {
		response.end('index page');
	}else if (url === '/login') {
 	var cpus = os.cpus();
 	// console.log(os.cpus());
 	var tm = os.totalmem();
 	// console.log();
		response.end (JSON.stringify(cpus)+(Number.parseInt(tm/1024/1024/1024+1)));
	}else if (url === '/products'){
		var products = [
			{
				name: '笔记本',   
				price: 9000
			},
			{
				name: '手机',
				price: 5000
			},
			{
				name: '憨憨',
				price: 1999
			}
		]

		response.end(JSON.stringify(products));
	}else{
		response.end('<h1>404</h1> Not Found.');
	}
	
	 

});

sever.listen(3300,function(){//绑定端口号
	console.log('Server Is Running ......');
});
//运行服务器，监听3300端口