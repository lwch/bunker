const config = require('./config/balmrc');
const api = require('./config/balm.api');

module.exports = () => {
  return {
    config,
    api
  };
};
