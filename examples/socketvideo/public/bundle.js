/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};

/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {

/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;

/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			exports: {},
/******/ 			id: moduleId,
/******/ 			loaded: false
/******/ 		};

/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;

/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}


/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;

/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;

/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";

/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var Recorder = __webpack_require__(1);
	var Receiver = __webpack_require__(3);
	window.rec = new Recorder();
	window.rcv = new Receiver();

	var buttonRecord = document.querySelector("#record");
	buttonRecord.onclick = function() {
		rec.record();
	};

	var buttonReceive = document.querySelector("#receive");
	buttonReceive.onclick = function() {
		rcv.receive();
	};


/***/ },
/* 1 */
/***/ function(module, exports, __webpack_require__) {

	var Client = __webpack_require__(2);

	var Recorder = function Recorder(element, onFrame, client) {
	  this.video = element || document.createElement("video");
	  this.onFrame = onFrame || function(){};
		this.client = client || (new Client("ws"));
		this.animationFrame = function(f, delay) {
			console.log(delay);
			setTimeout(f, 1000/10 - delay);
		};
	};

	Recorder.prototype.getUserMedia = function(cb) {
	  var gum = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;
	  gum.call(navigator, {audio: false, video: true}, cb, function(){});
	};

	Recorder.prototype.record = function() {
	  this.getUserMedia(function(stream) {
	    this.video.src = window.URL.createObjectURL(stream);
	    this.video.addEventListener("loadedmetadata", this.snapshot.bind(this));
	  }.bind(this));
	};

	Recorder.prototype.snapshot = function snapshot () {
		this.client.listen();
	  this.createCanvas();
		var lastTime = new Date().getTime();
		var context = this.canvas.getContext("2d");
	  this.animationFrame(function step() {
			var t = new Date().getTime();
			var duration = new Date().getTime() - lastTime;
	    context.drawImage(this.video, 0, 0, this.canvas.width, this.canvas.height);
			var imgData = this.canvas.toDataURL('image/jpeg');
			lastTime = new Date().getTime();
			var frame = {duration: duration, imgData: imgData};
			this.client.send(frame);
	    this.onFrame(frame);
			var t2 = new Date().getTime() - t;
	    this.animationFrame(step.bind(this), t2);
	  }.bind(this));
	};

	Recorder.prototype.createCanvas = function() {
	  if (!this.canvas) {
	    this.canvas = document.createElement("canvas");
	    this.canvas.width = this.video.videoWidth;
	    this.canvas.height = this.video.videoHeight;
	  }
	};

	module.exports = Recorder;


/***/ },
/* 2 */
/***/ function(module, exports) {

	function Client(path, onMessage) {
	  var location = window.location;
	  this.wsurl = "wss://" + location.hostname + ":" + location.port + "/" + path;
		this.onMessage = onMessage || function(){};
	};

	Client.prototype.listen = function listen() {
		this.ws = new WebSocket(this.wsurl);
	  this.rebind();
	};

	Client.prototype.rebind = function rebind() {
	  this.ws.onmessage = function (evt) {
	    this.onMessage(evt.data);
		}.bind(this);
	  this.ws.onclose = function() {
			console.log(arguments);
			setTimeout(function() {
			console.log("closed connection");
	    this.ws = new WebSocket(this.wsurl);
	    this.rebind();
			}.bind(this), 1000);
	  }.bind(this);
	};

	Client.prototype.send = function send(msg) {
	  var open = this.ws.readyState == this.ws.OPEN;
	  if (open) {
			if (msg instanceof Object) {
				msg = JSON.stringify(msg);
			}
	    this.ws.send(msg);
	  }
	}

	module.exports = Client;


/***/ },
/* 3 */
/***/ function(module, exports, __webpack_require__) {

	var Client = __webpack_require__(2);

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


/***/ }
/******/ ]);