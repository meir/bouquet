import process from "node:process";

function setup() {
  process.env.PULSE_LATENCY_MSEC = process.env.PULSE_LATENCY_MSEC || "30";
}

export default {
  setup,
  disabled: [],
  switch: {},
};
