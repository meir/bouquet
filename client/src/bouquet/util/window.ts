
declare let global: any;
// @ts-ignore
import electron from 'electron'


export default () => {
  return electron.BrowserWindow.fromId(global.mainWindowId)
}
