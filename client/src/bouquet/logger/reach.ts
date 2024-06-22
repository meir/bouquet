
import { window_log } from './methods.ts'

type LogFn = (msg: string) => void

interface IReach {
  [key: string]: LogFn[]
}

enum Reach {
  main = 0,
  renderer = 1,
  both = 2,
}

const getReaches = (reach: Reach): LogFn[] => {
  const reaches: IReach = {
    [Reach.main]: [console.log],
    [Reach.renderer]: [window_log],
    [Reach.both]: [console.log, window_log],
  }
  return reaches[reach]
}

export {
  Reach,
  getReaches,
}
