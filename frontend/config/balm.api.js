const publish = require('./balm.publish');

module.exports = (mix) => {
  if (mix.env.isProd) {
    publish(mix);
  }
};
