
declare let global: any;

import version from './version.ts'
import { info, warn, fatal } from './logger/index.ts'
import window from './util/window.ts'
import preload from './preload.ts'

function init() {
  global.bouquet = {
    info, warn, fatal,
    version,
    window,
  }

  global.bouquet.info(`Version ${version}`)
}

export {
  preload,
  init,
}
