const { BrowserWindow, ipcMain } = require('electron')
const path = require('path')

const isDevelopment = process.env.NODE_ENV === 'development'

let mainWindow = null

function createMainWindow() {
    mainWindow = new BrowserWindow({
        width: 1160,
        height: 752,
        minHeight: 632,
        minWidth: 960,
        frame: true,
        title: 'tiktok',
        webPreferences: {
            nodeIntegration: true,
            // preload: path.resolve(__dirname, '../utils/contextBridge.js')
        },
        // icon: path.resolve(__dirname, '../assets/logo.png')
    })

    mainWindow.setMenu(null)

    if (isDevelopment) {
        mainWindow.loadURL('http://localhost:3000/')
        // mainWindow.loadFile("../../src/index.html")
    } else {
        const entryPath = path.resolve(__dirname, '../../../build/index.html')
        mainWindow.loadFile(entryPath)
    }

    mainWindow.once('ready-to-show', () => {
        mainWindow.show()
    })
}

module.exports = { createMainWindow }
