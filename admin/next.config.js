module.exports = function() {
  return {
    env: {
      J_ENV: process.env.J_ENV,
      J_CICD: process.env.J_CICD,
    }
  }
}