var Recorder = require("./recorder.js");
var Receiver = require("./receiver.js");
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
