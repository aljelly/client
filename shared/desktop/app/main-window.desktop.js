// @flow
import URL from 'url-parse'
import AppState from './app-state.desktop'
import Window from './window.desktop'
import hotPath from './hot-path.desktop'
import * as SafeElectron from '../../util/safe-electron.desktop'
import {showDevTools} from '../../local-debug.desktop'
import {hideDockIcon} from './dock-icon.desktop'
import {getRendererHTML} from './dev.desktop'
import {windowStyle} from '../../styles'
import {isWindows} from '../../constants/platform'
import logger from '../../logger'

export default function() {
  // We are not using partitions on webviews, so this essentially disables
  // download for webviews. If we decide to start using partitions for
  // webviews, we should make sure to attach this to those partitions too.
  SafeElectron.getSession().defaultSession.on('will-download', event => event.preventDefault())
  // Disallow any permissions requests except for notifications
  SafeElectron.getSession().defaultSession.setPermissionRequestHandler(
    (webContents, permission, callback) => {
      const ourURL = new URL(getRendererHTML('mainWindow'))
      const requestURL = new URL(webContents.getURL())
      if (
        permission === 'notifications' &&
        requestURL.pathname.toLowerCase() === ourURL.pathname.toLowerCase()
      ) {
        // Allow notifications
        return callback(true)
      }
      return callback(false)
    }
  )

  let appState = new AppState()
  appState.checkOpenAtLogin()

  const mainWindow = new Window(getRendererHTML('mainWindow'), {
    backgroundThrottling: false,
    height: appState.state.height,
    minHeight: windowStyle.minHeight,
    minWidth: windowStyle.minWidth,
    show: false,
    webPreferences: {
      devTools: showDevTools,
      nodeIntegration: true,
      nodeIntegrationInWorker: false,
    },
    width: appState.state.width,
    x: appState.state.x,
    y: appState.state.y,
  })

  const webContents = mainWindow.window.webContents
  webContents.on('did-finish-load', () => {
    webContents.send('load', {
      scripts: [
        {
          async: false,
          src: hotPath('index.bundle.js'),
        },
      ],
    })
  })

  if (showDevTools) {
    webContents.openDevTools({mode: 'detach'})
  }

  appState.manageWindow(mainWindow.window)

  const app = SafeElectron.getApp()

  const openedAtLogin = app.getLoginItemSettings().wasOpenedAtLogin
  // app.getLoginItemSettings().restoreState is Mac only, so consider it always on in Windows
  const isRestore =
    !!process.env['KEYBASE_RESTORE_UI'] || app.getLoginItemSettings().restoreState || isWindows
  const hideWindowOnStart = process.env['KEYBASE_START_UI'] === 'hideWindow'
  const openHidden = app.getLoginItemSettings().wasOpenedAsHidden
  logger.info('Opened at login:', openedAtLogin)
  logger.info('Is restore:', isRestore)
  logger.info('Open hidden:', openHidden)
  if (
    isWindows &&
    appState &&
    appState.state &&
    typeof appState.state.x === 'number' &&
    typeof appState.state.y === 'number' &&
    typeof appState.state.width === 'number' &&
    typeof appState.state.height === 'number'
  ) {
    // DPI scaling issues
    // https://github.com/electron/electron/issues/10862
    mainWindow.window.setBounds({
      x: appState.state.x,
      y: appState.state.y,
      width: appState.state.width,
      height: appState.state.height,
    })
  }

  // Don't show main window:
  // - If we are set to open hidden,
  // - or, if we hide window on start,
  // - or, if we are restoring and window was hidden
  // - or, if we were opened from login (but not restoring)
  const hideMainWindow =
    openHidden ||
    hideWindowOnStart ||
    (isRestore && appState.state.windowHidden) ||
    (openedAtLogin && !isRestore)

  logger.info('Hide main window:', hideMainWindow)
  if (!hideMainWindow) {
    // On Windows we can try showing before Windows is ready
    // This will result in a dropped .show request
    // We add a listener to `did-finish-load` so we can show it when
    // Windows is ready.
    mainWindow.show()
    mainWindow.window.webContents.once('did-finish-load', () => {
      mainWindow.show()
    })
  }

  // Don't show dock:
  // - If we are set to open hidden,
  // - or, if we are restoring and dock was hidden
  // - or, if we were opened from login (but not restoring)
  const shouldHideDockIcon =
    openHidden || (isRestore && appState.state.dockHidden) || (openedAtLogin && !isRestore)
  logger.info('Hide dock icon:', shouldHideDockIcon)
  if (shouldHideDockIcon) {
    hideDockIcon()
  }

  SafeElectron.getIpcMain().on('showMain', () => {
    logger.info('Show main window (requested)')
    mainWindow.show()
    const window = mainWindow.window
    if (window) {
      window.focus()
      logger.info('...showMain: visible=', window.isVisible(), window.getBounds())
    } else {
      logger.info('...showMain: no mainWindow!')
    }
  })

  return mainWindow
}
