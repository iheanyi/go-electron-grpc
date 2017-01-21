const $ = require('jquery');
var PROTO_PATH = __dirname + '/../demo/demo.proto';
const grpc = require('grpc');
const demoProto = grpc.load(PROTO_PATH).demo;

let client = new demoProto.Greeter('localhost:50051', grpc.credentials.createInsecure());
let $name = $('#name');
let $nameMessage = $('.response');

$('#name-submit').click(function(evt) {
  let name = $name.val().trim();
  let user;

  if (name.length > 0) {
   user = name; 
  } else {
    user = 'world';
  }

  client.sayHello({name: user}, function(err, response) {
    window.alert("Greeting: " + response.message);
    $nameMessage.text(response.message);
  });
});
