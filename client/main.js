

const {app, BrowserWindow} = require('electron');
const path = require('path');
const url = require('url');
const $ = require('jquery');

var PROTO_PATH = __dirname + '/../demo/demo.proto';
const grpc = require('grpc');
const demoProto = grpc.load(PROTO_PATH).demo;

// Global reference of window object so it won't be closed automatically during
// JS garbage collection.
let win

function createWindow() {
  // Create Browser Window
  win = new BrowserWindow({width: 800, height: 600});

  // and load the index.html of the app
  win.loadURL(url.format({
    pathname: path.join(__dirname, 'index.html'),
    protocol: 'file',
    slashes: true,
  }));

  // Open the DevTools
  win.webContents.openDevTools()

  // Emitted when the window is closed
  win.on('closed', () => {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    win = null;
  });
}

app.on('ready', createWindow)

app.on('window-all-closed', () => {
  // On macOS it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  //if (process.platform !== 'darwin') {
    app.quit()
  //}
}); 
