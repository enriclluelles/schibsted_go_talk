var Client = require("./websocket");

var Receiver = function Receiver(element, onMessage, client) {
	if (element) {
		this.canvas = element;
	} else {
		this.canvas = document.createElement("canvas");
		this.canvas.width = 1000;
		this.canvas.height = 1000;
		document.body.appendChild(this.canvas);
	}

	this.queue = [];
  this.onMessage = onMessage || function(){};
	this.client = client || (new Client("ws"));
	this.ratio = 1;
};

Receiver.prototype.stop = function stop() {
	clearTimeout(this.timer);
}

Receiver.prototype.show = function show() {
	// console.log(this.queue.length);
	var time, img;
	data = this.queue.shift();
	if (!data) {
		setTimeout(show.bind(this), 50);
		return;
	}
	time = data.duration / this.ratio;
	// console.log("timeout" + time);
	// console.log("queuelength" + this.queue.length);
	img = new Image;
	img.onload = function() {
		this.context.drawImage(img, 0, 0);
		this.timer = setTimeout(show.bind(this), data.duration / this.ratio);
	}.bind(this);
	img.src = data.imgData;
};

Receiver.prototype.receive = function receive() {
	this.client.listen();
	this.show();
	this.context = this.canvas.getContext("2d");
	var img;
	var lastMessage = new Date().getTime();
	this.client.onMessage = function(msg) {
		console.log("interal" + (new Date().getTime() - lastMessage));
		lastMessage = new Date().getTime();
		var data = JSON.parse(msg);
		this.queue.push(data);
		this.onMessage(data);
	}.bind(this);
};

module.exports = Receiver;
