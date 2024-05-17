const { app } = require('electron')
const { createMainWindow } = require('./windows/mainWindow')

app.on('ready', () => {
    createMainWindow()
})
