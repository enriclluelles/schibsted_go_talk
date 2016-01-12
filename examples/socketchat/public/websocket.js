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
