
import window from '../util/window.ts';

function window_log(log: string) {
  window().webContents.send('bouquet.log', log)
}

const prefix: string = '[Bouquet]'

enum Level {
  info = 'ℹ',
  warn = '⚠',
  fatal = '✖',
}

type LogFn = (msg: string) => void
interface IReach {
  [key: string]: LogFn[]
}

enum Reach {
  main = 0,
  renderer = 1,
  both = 2,
}

const reach: IReach = {
  [Reach.main]: [console.log],
  [Reach.renderer]: [window_log],
  [Reach.both]: [console.log, window_log],
}

function format(msg: string, type: Level = Level.info): string {
  return ` ${type} ${prefix} '${msg}'`
}

function log(msg: string, type: Level, r: Reach = 2) {
  let log = format(msg, type)
  reach[r].forEach((fn: Function) => fn(log))
}

const info = (msg: string, reach: Reach) => log(msg, Level.info, reach)
const warn = (msg: string, reach: Reach) => log(msg, Level.warn, reach)
const fatal = (msg: string, reach: Reach) => log(msg, Level.fatal, reach)

export {
  info,
  warn,
  fatal,
}
