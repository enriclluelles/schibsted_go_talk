var Client = require("./websocket");

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
