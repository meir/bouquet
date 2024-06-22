
import window from '../util/window.ts'
import { Level } from './level.ts'
import { Reach, getReaches } from './reach.ts'

const prefix: string = '[Bouquet]'


const window_log = (log: string) => {
  window().webContents.send('bouquet.log', log)
}

const format = (msg: string, type: Level = Level.info): string => {
  return ` ${type} ${prefix} '${msg}'`
}

const log = (msg: string, type: Level, reach: Reach = 2) => {
  let log = format(msg, type)
  getReaches(reach).forEach((fn: Function) => fn(log))
}

const info = (msg: string, reach: Reach) => log(msg, Level.info, reach)
const warn = (msg: string, reach: Reach) => log(msg, Level.warn, reach)
const fatal = (msg: string, reach: Reach) => log(msg, Level.fatal, reach)

export {
  window_log,
  format,
  log,

  info,
  warn,
  fatal,
}
